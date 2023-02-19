export type StateKey =
	| "INITIAL"
	| "SHOULD_ENCRYPT"
	| "SHOULD_DECRYPT"
	| "PROCESSING"
	| "DONE"
	| "FAILURE";

export type MessageKey = "ENCRYPT" | "DECRYPT" | "ENCRYPTED" | "DECRYPTED" | "FAILURE";

export type HexEncodedFile = {
	hex: string;
	name: string;
};

export type EncrypterState = {
	isProcessing: boolean;
	filesToEncrypt: File[] | undefined;
	ciphertext: string | undefined;
	state: StateKey;
	password: string | undefined;
	hint: string | undefined;
	error: Error | undefined;
	crypString: string | undefined;
	decryptedFiles: File[] | undefined;
};

export type WhoamiState = {
	isAuthenticated: boolean;
	email: string | undefined;
};

export type MessagePayload = {
	type: MessageKey;
	payload: EncrypterState;
};
