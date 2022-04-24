import { Writable, writable, get } from "svelte/store";
import type { Files } from "filedrop-svelte";
import Guu from "guu";
import { decrypt, encrypt, hexDecode, hexEncode } from "../crypto";
import { formatCrypString, parseCrypString } from "../utils";
import { CRYP_FILE_EXTENSION, STATE } from "../constants";
import type { EncrypterState, HexEncodedFile } from "../types";

const log = new Guu("encrypter", "pink");

const initialState: EncrypterState = {
  isProcessing: false,
  filesToEncrypt: undefined,
  ciphertext: undefined,
  password: undefined,
  hint: undefined,
  state: STATE.INITIAL,
  error: undefined,
  crypString: undefined,
  decryptedFiles: undefined,
};

class Encrypter {
  constructor(
    public store: Writable<EncrypterState> = writable(initialState),
    public worker: Worker = new Worker(
      new URL("./crypto.worker.js", import.meta.url)
    )
  ) {
    this.worker.onmessage = this.handleMessage;
    this.worker.onerror = this.handleWorkerError;
  }

  private dispatch(payload: Partial<EncrypterState>) {
    this.store.update((prevState) => ({ ...prevState, ...payload }));
  }

  private handleMessage = (msg: MessageEvent<unknown>) => {
    log.info({ msg });
  };

  private handleWorkerError = (err: ErrorEvent) => {
    log.error(err);
    this.dispatch({
      state: STATE.FAILURE,
      error: err.error,
    });
  };

  public reset = () => this.store.update(() => initialState);

  public handleFiles = async (filesToEncrypt: Files) => {
    log.info("handleFiles", filesToEncrypt);
    const isCrypFile = filesToEncrypt?.accepted?.[0]?.name
      ?.trim()
      ?.endsWith(CRYP_FILE_EXTENSION);
    log.info({ isCrypFile });
    if (!isCrypFile) {
      this.dispatch({ filesToEncrypt, state: STATE.SHOULD_ENCRYPT });
    } else {
      const arrayBuffer = await filesToEncrypt?.accepted?.[0].arrayBuffer();
      const crypString = new TextDecoder().decode(arrayBuffer);
      const { ciphertext, hint } = parseCrypString(crypString);
      this.dispatch({
        ciphertext,
        hint,
        crypString,
        state: STATE.SHOULD_DECRYPT,
      });
    }
  };

  public handleEncrypt = async ({
    password,
    hint,
  }: {
    password: string;
    hint: string | undefined;
  }) => {
    this.dispatch({
      password,
      hint,
      state: STATE.PROCESSING,
    });
    try {
      const { filesToEncrypt } = get(this.store);
      const accepted = await Promise.all(
        filesToEncrypt.accepted.map((item) => item.arrayBuffer())
      );
      const hexEncodedFiles: HexEncodedFile[] = accepted.map((item, index) => ({
        hex: hexEncode(item),
        name: filesToEncrypt?.accepted?.[index].name,
      }));
      // the plaintext is a stringified JSON array of files
      const plaintext = JSON.stringify(hexEncodedFiles);
      // ciphertext is the encrypted array
      const ciphertext = await encrypt(password, plaintext);
      const { hint } = get(this.store);
      const crypString = formatCrypString(ciphertext, hint);
      this.dispatch({ ciphertext, crypString, state: STATE.DONE });
    } catch (error) {
      log.error(error);
      this.dispatch({
        state: STATE.FAILURE,
        error,
      });
    }
  };

  public handleDecrypt = async (password: string) => {
    this.dispatch({
      password,
      state: STATE.PROCESSING,
    });
    try {
      const { ciphertext } = get(this.store);
      const plaintext = await decrypt(password, ciphertext);
      const hexEncodedFiles: HexEncodedFile[] = JSON.parse(plaintext);
      const decryptedFiles = hexEncodedFiles.map((item) => {
        const blob = new Blob([hexDecode(item.hex)]);
        return new File([blob], item.name);
      });
      this.dispatch({ decryptedFiles, state: STATE.DONE });
    } catch (error) {
      log.error(error);
      this.dispatch({
        state: STATE.FAILURE,
        error,
      });
    }
  };
}

export const encrypter = new Encrypter();
