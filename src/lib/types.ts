import type { Files } from "filedrop-svelte";

export type StateKey =
  | "INITIAL"
  | "SHOULD_ENCRYPT"
  | "SHOULD_DECRYPT"
  | "PROCESSING"
  | "DONE"
  | "FAILURE";

export type MessageKey = "ENCRYPT" | "DECRYPT";

export type MessagePayload = {
  type: MessageKey;
  value: string;
};

export type HexEncodedFile = {
  hex: string;
  name: string;
};

export type EncrypterState = {
  isProcessing: boolean;
  filesToEncrypt: Files | undefined;
  ciphertext: string | undefined;
  state: StateKey;
  password: string | undefined;
  hint: string | undefined;
  error: Error | undefined;
  crypString: string | undefined;
  decryptedFiles: File[] | undefined;
};
