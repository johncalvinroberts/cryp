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

export const handleEncrypt = async ({
  password,
  hint,
}: {
  password: string;
  hint: string | undefined;
}) => {
  encrypter.update((prevState) => ({
    ...prevState,
    password,
    hint,
    state: State.PROCESSING,
  }));
  const { files } = get(encrypter);
  const accepted = await Promise.all(
    files.accepted.map((item) => item.arrayBuffer())
  );
  const fileStrings = accepted.map((item) => hexEncode(item));
  const plaintext = JSON.stringify(fileStrings);
  const ciphertext = await encrypt(password, plaintext);
  console.log({ ciphertext });
};

export const reset = () => encrypter.update(() => initialState);
