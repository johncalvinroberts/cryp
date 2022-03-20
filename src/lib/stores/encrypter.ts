import { writable } from "svelte/store";
import type { Files } from "filedrop-svelte";

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
  ciphertext: unknown | undefined;
  state: State;
  password: string | undefined;
  hint: string | undefined;
};

const initialState = {
  isProcessing: false,
  files: undefined,
  ciphertext: undefined,
  password: undefined,
  hint: undefined,
  state: State.INITIAL,
};

export const encrypter = writable<Encrypter>(initialState);

export const handleFiles = (files: Files) => {
  encrypter.update((prevState) => {
    // TODO: parse file and identify if should encrypt or decrypt
    return { ...prevState, files, state: State.SHOULD_ENCRYPT };
  });
};

export const reset = () => encrypter.update(() => initialState);
