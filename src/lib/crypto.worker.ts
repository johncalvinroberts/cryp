import Guu from "guu";
import { encrypt, decrypt } from "./crypto";
import { MESSAGE } from "./constants";
import type { MessagePayload } from "./types";

// alias self to ctx and give it our newly created type
const ctx: Worker = self as any;

const log = new Guu("crypto.worker", "#6a11cb");

// main class wrapper for the worker
class CryptoWorker {
  handleMessage = (msg: MessageEvent<MessagePayload>) => {
    switch (msg.data.type) {
      case MESSAGE.ENCRYPT:
        log.info("Encrypt");
        break;
      case MESSAGE.DECRYPT:
        log.info("Encrypt");
        break;
      default:
        throw new Error("Unknown Message Type");
    }
  };
}

// instantiate a worker
const cryptoWorker = new CryptoWorker();

// add listener to the worker global scope
ctx.addEventListener("message", cryptoWorker.handleMessage);
