import { writable, get } from "svelte/store";
import type { Files } from "filedrop-svelte";
import { decrypt, encrypt, hexDecode, hexEncode } from "../crypto";
import { formatCrypString, parseCrypString } from "../utils";
import { CRYP_FILE_EXTENSION } from "../constants";

export enum State {
  INITIAL = "INITIAL",
  SHOULD_ENCRYPT = "SHOULD_ENCRYPT",
  SHOULD_DECRYPT = "SHOULD_DECRYPT",
  PROCESSING = "PROCESSING",
  DONE = "DONE",
  FAILURE = "FAILURE",
}

type HexEncodedFile = {
  hex: string;
  name: string;
};

type Encrypter = {
  isProcessing: boolean;
  filesToEncrypt: Files | undefined;
  ciphertext: string | undefined;
  state: State;
  password: string | undefined;
  hint: string | undefined;
  error: Error | undefined;
  crypString: string | undefined;
  decryptedFiles: File[] | undefined;
};

const initialState = {
  isProcessing: false,
  filesToEncrypt: undefined,
  ciphertext: undefined,
  password: undefined,
  hint: undefined,
  state: State.INITIAL,
  error: undefined,
  crypString: undefined,
  decryptedFiles: undefined,
};

export const encrypter = writable<Encrypter>(initialState);

export const dispatch = (payload: Partial<Encrypter>) =>
  encrypter.update((prevState) => ({ ...prevState, ...payload }));

export const handleFiles = async (filesToEncrypt: Files) => {
  const isCrypFile = filesToEncrypt?.accepted?.[0]?.name
    ?.trim()
    ?.endsWith(CRYP_FILE_EXTENSION);
  if (!isCrypFile) {
    dispatch({ filesToEncrypt, state: State.SHOULD_ENCRYPT });
  } else {
    const arrayBuffer = await filesToEncrypt?.accepted?.[0].arrayBuffer();
    const crypString = new TextDecoder().decode(arrayBuffer);
    const { ciphertext, hint } = parseCrypString(crypString);
    dispatch({ ciphertext, hint, crypString, state: State.SHOULD_DECRYPT });
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
    state: State.PROCESSING,
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
    dispatch({ ciphertext, crypString, state: State.DONE });
  } catch (error) {
    console.error(error);
    dispatch({
      state: State.FAILURE,
      error,
    });
  }
};

export const handleDecrypt = async (password: string) => {
  dispatch({
    password,
    state: State.PROCESSING,
  });
  try {
    const { ciphertext } = get(encrypter);
    const plaintext = await decrypt(password, ciphertext);
    const hexEncodedFiles: HexEncodedFile[] = JSON.parse(plaintext);
    const decryptedFiles = hexEncodedFiles.map((item) => {
      const blob = new Blob([hexDecode(item.hex)]);
      return new File([blob], item.name);
    });
    dispatch({ decryptedFiles, state: State.DONE });
  } catch (error) {
    console.error(error);
    dispatch({
      state: State.FAILURE,
      error,
    });
  }
};

export const reset = () => encrypter.update(() => initialState);
