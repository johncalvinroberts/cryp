import { writable, get } from "svelte/store";
import type { Files } from "filedrop-svelte";
import Guu from "guu";
import { decrypt, encrypt, hexDecode, hexEncode } from "../crypto";
import { formatCrypString, parseCrypString } from "../utils";
import { CRYP_FILE_EXTENSION, STATE } from "../constants";
import type { Encrypter, HexEncodedFile } from "../types";

const log = new Guu("encrypter", "pink");

const initialState: Encrypter = {
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

export const encrypter = writable<Encrypter>(initialState);

export const dispatch = (payload: Partial<Encrypter>) =>
  encrypter.update((prevState) => ({ ...prevState, ...payload }));

export const handleFiles = async (filesToEncrypt: Files) => {
  log.info("handleFiles", filesToEncrypt);
  const isCrypFile = filesToEncrypt?.accepted?.[0]?.name
    ?.trim()
    ?.endsWith(CRYP_FILE_EXTENSION);
  log.info({ isCrypFile });
  if (!isCrypFile) {
    dispatch({ filesToEncrypt, state: STATE.SHOULD_ENCRYPT });
  } else {
    const arrayBuffer = await filesToEncrypt?.accepted?.[0].arrayBuffer();
    const crypString = new TextDecoder().decode(arrayBuffer);
    const { ciphertext, hint } = parseCrypString(crypString);
    dispatch({ ciphertext, hint, crypString, state: STATE.SHOULD_DECRYPT });
  }
};

export const handleEncrypt = async ({
  password,
  hint,
}: {
  password: string;
  hint: string | undefined;
}) => {
  dispatch({
    password,
    hint,
    state: STATE.PROCESSING,
  });
  try {
    const { filesToEncrypt } = get(encrypter);
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
    const { hint } = get(encrypter);
    const crypString = formatCrypString(ciphertext, hint);
    dispatch({ ciphertext, crypString, state: STATE.DONE });
  } catch (error) {
    log.error(error);
    dispatch({
      state: STATE.FAILURE,
      error,
    });
  }
};

export const handleDecrypt = async (password: string) => {
  dispatch({
    password,
    state: STATE.PROCESSING,
  });
  try {
    const { ciphertext } = get(encrypter);
    const plaintext = await decrypt(password, ciphertext);
    const hexEncodedFiles: HexEncodedFile[] = JSON.parse(plaintext);
    const decryptedFiles = hexEncodedFiles.map((item) => {
      const blob = new Blob([hexDecode(item.hex)]);
      return new File([blob], item.name);
    });
    dispatch({ decryptedFiles, state: STATE.DONE });
  } catch (error) {
    log.error(error);
    dispatch({
      state: STATE.FAILURE,
      error,
    });
  }
};

export const reset = () => encrypter.update(() => initialState);
