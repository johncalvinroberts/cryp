export type StateKey =
	| "INITIAL"
	| "SHOULD_AUTHENTICATE"
	| "SHOULD_ENCRYPT"
	| "SHOULD_DECRYPT"
	| "PROCESSING"
	| "DONE"
	| "FAILURE";

export type MessageKey = "ENCRYPT" | "DECRYPT" | "ENCRYPTED" | "DECRYPTED" | "FAILURE";

export type MessagePayload = {
	type: MessageKey;
	payload: EncrypterState;
};

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
	totalFileBytes: number | undefined;
};

export type WhoamiState = {
	isAuthenticated: boolean;
	email: string | undefined;
	isChallengeSent: boolean;
	isLoading: boolean;
};

export type APIClientState = {
	isRefreshingToken: boolean;
	tokenExpiresAt: number | undefined;
	token: string | undefined;
};

export type ThemeState = {
	theme: Theme;
	isAuthModalOpen: boolean;
	errors: MaybeError[];
};

export type ApiResponse<T> = {
	success: boolean;
	data: T;
	error: string | undefined;
};

export type HTTPMethod = "GET" | "HEAD" | "POST" | "PUT" | "DELETE" | "PATCH";

export type HTTPRequestBody = Record<string, unknown> | Blob | File;

export type Theme = "dark" | "light";

export type MaybeError = Error | string | unknown;
