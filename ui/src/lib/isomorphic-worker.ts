/* eslint-disable @typescript-eslint/no-unused-vars */

class StubWorker implements Worker {
	onmessage() {
		// noop
	}

	onerror() {
		// noop
	}

	onmessageerror() {
		// noop
	}

	terminate(): void {
		// noop
	}

	postMessage(message: unknown, options?: unknown): void {
		// noop
	}

	addEventListener(
		type: string,
		listener: EventListenerOrEventListenerObject,
		options?: boolean | AddEventListenerOptions | undefined,
	): void {
		// noop
	}

	removeEventListener(type: unknown, listener: unknown, options?: unknown): void {
		// noop
	}

	dispatchEvent(event: Event): boolean {
		return true;
	}
}

const IsomorphicWorker = typeof window == 'undefined' ? StubWorker : Worker;

export default IsomorphicWorker;

/* eslint-enable @typescript-eslint/no-unused-vars */
