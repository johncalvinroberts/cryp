import { type Writable, writable } from "svelte/store";

abstract class BaseStore<T> {
	private initialState: T;
	public store: Writable<T>;

	constructor(initialState: T) {
		this.store = writable(initialState);
		this.initialState = initialState;
	}

	public reset = () => this.store.update(() => this.initialState);

	public dispatch(payload: Partial<T>) {
		this.store.update((prevState) => ({ ...prevState, ...payload }));
	}
}

export default BaseStore;
