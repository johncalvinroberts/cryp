import { writable, get } from "svelte/store";
import type { Files } from "filedrop-svelte";
import { encrypt, hexEncode } from "../cryp";

export enum State {
  INITIAL = "INITIAL",
  SHOULD_ENCRYPT = "SHOULD_ENCRYPT",
  SHOULD_DECRYPT = "SHOULD_DECRYPT",
  PROCESSING = "PROCESSING",
  DONE = "DONE",
  FAILURE = "FAILURE",
}

type Encrypter = {
  isProcessing: boolean;
  files: Files | undefined;
  ciphertext: string | undefined;
  state: State;
  password: string | undefined;
  hint: string | undefined;
  error: Error | undefined;
};

const initialState = {
  isProcessing: false,
  files: undefined,
  ciphertext: undefined,
  password: undefined,
  hint: undefined,
  state: State.INITIAL,
  error: undefined,
};

export const encrypter = writable<Encrypter>(initialState);

export const dispatch = (payload: Partial<Encrypter>) =>
  encrypter.update((prevState) => ({ ...prevState, ...payload }));

export const handleFiles = (files: Files) => {
  // TODO: parse file and identify if should encrypt or decrypt
  dispatch({ files, state: State.SHOULD_ENCRYPT });
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
    const { files } = get(encrypter);
    const accepted = await Promise.all(
      files.accepted.map((item) => item.arrayBuffer())
    );
    const fileStrings = accepted.map((item) => hexEncode(item));
    const plaintext = JSON.stringify(fileStrings);
    const ciphertext = await encrypt(password, plaintext);
    dispatch({ ciphertext, state: State.DONE });
  } catch (error) {
    console.error(error);
    dispatch({
      state: State.FAILURE,
      error,
    });
  }
};

export const reset = () => encrypter.update(() => initialState);
