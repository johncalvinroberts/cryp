import { ENCRYPT_ALGO, KEY_ALGO, CIPHERTEXT_DELIMITER } from "./constants";

export const getRandomBytes = (size = 16): Uint8Array => {
  return crypto.getRandomValues(new Uint8Array(size));
};

export const hexEncode = (buffer: ArrayBuffer) => {
  let s = "";
  for (const i of new Uint8Array(buffer)) {
    s += i.toString(16).padStart(2, "0");
  }
  return s;
};

export const hexDecode = (raw: string) => {
  const size = raw.length / 2;
  const buffer = new Uint8Array(size);
  for (let i = 0; i < size; i++) {
    const idx = i * 2;
    const segment = raw.slice(idx, idx + 2);
    buffer[i] = parseInt(segment, 16);
  }
  return buffer;
};

const getInitialKey = (password: string): Promise<CryptoKey> => {
  return crypto.subtle.importKey(
    "raw",
    new TextEncoder().encode(password),
    KEY_ALGO,
    false,
    ["deriveKey", "deriveBits"]
  );
};

const getDerivedKey = (
  baseKey: CryptoKey,
  salt: ArrayBuffer
): Promise<CryptoKey> => {
  const iterations = 100_000;
  return crypto.subtle.deriveKey(
    { name: "PBKDF2", iterations, salt, hash: "SHA-512" },
    baseKey,
    { name: ENCRYPT_ALGO, length: 256 },
    false,
    ["encrypt", "decrypt"]
  );
};

export const encrypt = async (
  password: string,
  plaintext: string
): Promise<string> => {
  const initialKey = await getInitialKey(password);
  const salt = getRandomBytes();
  const key = await getDerivedKey(initialKey, salt);
  const iv = getRandomBytes(12);
  const ciphertext: ArrayBuffer = await crypto.subtle.encrypt(
    {
      name: ENCRYPT_ALGO,
      iv,
    },
    key,
    new TextEncoder().encode(plaintext)
  );
  const composed = [salt, iv, ciphertext]
    .map((item) => hexEncode(item))
    .join(CIPHERTEXT_DELIMITER);
  return composed;
};

export const decrypt = async (password: string, composed: string) => {
  const [salt, iv, ciphertext] = composed
    .split(CIPHERTEXT_DELIMITER)
    .map((item) => hexDecode(item));
  const initialKey = await getInitialKey(password);
  const key = await getDerivedKey(initialKey, salt);
  const plaintext = await crypto.subtle.decrypt(
    {
      name: ENCRYPT_ALGO,
      iv,
    },
    key,
    ciphertext
  );
  return new TextDecoder().decode(plaintext);
};
