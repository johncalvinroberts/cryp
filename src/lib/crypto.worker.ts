import Guu from "guu";
import { encrypt, decrypt, hexEncode, hexDecode } from "./crypto";
import { MESSAGE, STATE } from "./constants";
import { formatCrypString } from "./utils";
import type { MessagePayload, HexEncodedFile, EncrypterState } from "./types";

// alias self to ctx and give it our newly created type
const ctx: Worker = self as any;

const log = new Guu("crypto.worker", "#6a11cb");

// main class wrapper for the worker
class CryptoWorker {
  encrypt = async (encrypterState: EncrypterState) => {
    try {
      log.info("Encrypting");
      const { filesToEncrypt, password, hint } = encrypterState;
      const accepted = await Promise.all(
        filesToEncrypt.accepted.map((item) => item.arrayBuffer())
      );
      const hexEncodedFiles: HexEncodedFile[] = accepted.map((item, index) => ({
        hex: hexEncode(item),
        name: filesToEncrypt?.accepted?.[index].name,
      }));
      // the plaintext is a stringified JSON array of files
      const plaintext = JSON.stringify(hexEncodedFiles);
      const ciphertext = await encrypt(password, plaintext);
      log.info("Encrypted");
      const crypString = formatCrypString(ciphertext, hint);
      const payload = { ciphertext, crypString, state: STATE.DONE };
      ctx.postMessage({ payload, type: MESSAGE.ENCRYPTED });
      log.info("Posted message back to main thread");
    } catch (error) {
      const payload = {
        state: STATE.FAILURE,
        error,
      };
      ctx.postMessage({ payload, type: MESSAGE.FAILURE });
    }
  };

  decrypt = async (encrypterState: EncrypterState) => {
    try {
      log.info("Decrypting");
      const { ciphertext, password } = encrypterState;
      const plaintext = await decrypt(password, ciphertext);
      log.info("Decrypted");
      const hexEncodedFiles: HexEncodedFile[] = JSON.parse(plaintext);
      const decryptedFiles = hexEncodedFiles.map((item) => {
        const blob = new Blob([hexDecode(item.hex)]);
        return new File([blob], item.name);
      });
      const payload = { decryptedFiles, state: STATE.DONE };
      log.info("Posted message back to main thread");
      ctx.postMessage({ payload, type: MESSAGE.DECRYPTED });
    } catch (error) {
      log.error(error);
      ctx.postMessage({
        state: STATE.FAILURE,
        error,
      });
    }
  };

  handleMessage = (msg: MessageEvent<MessagePayload>) => {
    const { type, payload } = msg.data;
    log.info(msg.data);
    switch (type) {
      case MESSAGE.ENCRYPT:
        this.encrypt(payload);
        break;
      case MESSAGE.DECRYPT:
        this.decrypt(payload);
        break;
      default:
        throw new Error("Unknown Message Type");
    }
  };
}

// instantiate a worker
log.info("Instantiating worker");
const cryptoWorker = new CryptoWorker();
// add listener to the worker global scope
ctx.addEventListener("message", cryptoWorker.handleMessage);
