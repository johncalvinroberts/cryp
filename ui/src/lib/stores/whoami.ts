import type { Writable } from 'svelte/store';
import type { WhoamiState } from '../types';
import { writable } from 'svelte/store';

const initialState: WhoamiState = {
	isAuthenticated: false,
	email: undefined,
};

class WhoamiStore {
	constructor(public store: Writable<WhoamiState> = writable(initialState)) {}

	public reset = () => this.store.update(() => initialState);
}

export const whoami = new WhoamiStore();
