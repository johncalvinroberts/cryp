import type { Files } from "filedrop-svelte";
import { CRYP_DELIMITER, CRYP_FILE_EXTENSION } from "./constants";

const MAX_NAME_LENGTH = 100;

export const getEncryptedFilename = (files: Files) => {
  let concatenatedNames = "";
  for (const file of files?.accepted) {
    const name = file.name?.substring(0, file?.name.lastIndexOf("."));
    concatenatedNames = `${concatenatedNames}${
      concatenatedNames?.length ? "+" : ""
    }${name}`;
    if (concatenatedNames.length >= MAX_NAME_LENGTH) {
      concatenatedNames = `${concatenatedNames.substring(0, 97)}...`;
      break;
    }
  }
  return `${concatenatedNames}${CRYP_FILE_EXTENSION}`;
};

export const formatCrypString = (ciphertext: string, hint: string) =>
  `${CRYP_DELIMITER}${ciphertext}${CRYP_DELIMITER}${hint}`;

export const parseCrypString = (crypString: string) => {
  const [, ciphertext, hint] = crypString.split(CRYP_DELIMITER);
  return { ciphertext, hint };
};

export const getRandomUnicodeString = (length: number): string => {
  const array = new Uint16Array(length);
  window.crypto.getRandomValues(array);
  let str = "";
  for (let i = 0; i < array.length; i++) {
    str += String.fromCharCode(array[i]);
  }
  return str;
};

export const clone = (obj: unknown) => JSON.parse(JSON.stringify(obj));
