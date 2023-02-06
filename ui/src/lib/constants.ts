import type { StateKey, MessageKey } from './types';

export const ENCRYPT_ALGO = 'AES-GCM';
export const KEY_ALGO = 'PBKDF2';
export const CRYP_DELIMITER = '@CRYP@';
export const CIPHERTEXT_DELIMITER = '::';
export const CRYP_FILE_EXTENSION = '.cryp';
export const FALLBACK_FILE_NAME = 'file';

export const STATE: Record<string, StateKey> = {
	INITIAL: 'INITIAL',
	SHOULD_ENCRYPT: 'SHOULD_ENCRYPT',
	SHOULD_DECRYPT: 'SHOULD_DECRYPT',
	PROCESSING: 'PROCESSING',
	DONE: 'DONE',
	FAILURE: 'FAILURE',
};

export const MESSAGE: Record<string, MessageKey> = {
	ENCRYPT: 'ENCRYPT',
	DECRYPT: 'DECRYPT',
	ENCRYPTED: 'ENCRYPTED',
	DECRYPTED: 'DECRYPTED',
	FAILURE: 'FAILURE',
};
