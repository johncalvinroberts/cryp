import { get } from "svelte/store";
import { parseCrypString } from "../utils";
import { CRYP_FILE_EXTENSION, STATE, MESSAGE } from "../constants";
import type { EncrypterState, MessageKey, MessagePayload } from "../types";
import IsomorphicWorker from "../isomorphic-worker";
import BaseStore from "./base";

const initialState: EncrypterState = {
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

class EncrypterStore extends BaseStore<EncrypterState> {
	constructor(
		private worker: Worker = new IsomorphicWorker(new URL("../crypto.worker.ts", import.meta.url), {
			type: "module",
		}),
	) {
		super(initialState);
		this.worker.onmessage = this.handleMessage;
		this.worker.onerror = this.handleWorkerError;
		this.worker.onmessageerror = this.handleWorkerError;
	}

	private handleMessage = (msg: MessageEvent<MessagePayload>) => {
		const { payload } = msg.data;
		this.dispatch(payload);
	};

	private postMessage = (type: MessageKey) => {
		const payload = get(this.store);
		this.worker.postMessage({ type, payload });
	};

	private handleWorkerError = (err: ErrorEvent | MessageEvent) => {
		if (err instanceof ErrorEvent) {
			this.dispatch({
				state: STATE.FAILURE,
				error: err.error,
			});
		}
	};

	public handleFiles = async (files: File[]) => {
		const isCrypFile = files?.[0]?.name?.trim()?.endsWith(CRYP_FILE_EXTENSION);
		if (!isCrypFile) {
			const { filesToEncrypt: currentFiles = [] } = get(this.store);
			const filesToEncrypt = [...currentFiles, ...files];
			this.dispatch({ filesToEncrypt, state: STATE.SHOULD_ENCRYPT });
		}
		if (isCrypFile) {
			const arrayBuffer = await files?.[0].arrayBuffer();
			const crypString = new TextDecoder().decode(arrayBuffer);
			const { ciphertext, hint } = parseCrypString(crypString);
			this.dispatch({
				ciphertext,
				hint,
				crypString,
				state: STATE.SHOULD_DECRYPT,
			});
		}
	};

	public handleEncrypt = async (password: string, hint: string) => {
		this.dispatch({
			password,
			hint,
			state: STATE.PROCESSING,
		});
		this.postMessage(MESSAGE.ENCRYPT);
	};

	public handleDecrypt = async (password: string) => {
		this.dispatch({
			password,
			state: STATE.PROCESSING,
		});
		this.postMessage(MESSAGE.DECRYPT);
	};
}

export const encrypter = new EncrypterStore();
