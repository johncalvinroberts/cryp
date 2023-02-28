import type { WhoamiState } from "../../types/types";
import BaseStore from "./base";

const initialState: WhoamiState = {
	isAuthenticated: false,
	email: undefined,
};

class WhoamiStore extends BaseStore<WhoamiState> {
	constructor() {
		super(initialState);
	}
}

export const whoami = new WhoamiStore();
