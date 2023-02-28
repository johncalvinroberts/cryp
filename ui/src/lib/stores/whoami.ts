import type { WhoamiState } from "../../types/types";
import type { TryWhoamiChallengeDTO } from "../../types/dtos";
import BaseStore from "./base";
import { apiClient } from "./api";
import { display } from "./display";

const initialState: WhoamiState = {
	isAuthenticated: false,
	email: undefined,
	isChallengeSent: false,
};

class WhoamiStore extends BaseStore<WhoamiState> {
	constructor() {
		super(initialState);
	}

	public async startWhoamiChallenge(email: string) {
		try {
			await apiClient.post("api/whoami/start", { email });
		} catch (error) {
			display.enqueueError(error);
		}
	}

	public async tryWhoamiChallenge(email: string, otp: string): Promise<boolean> {
		try {
			await apiClient.post<TryWhoamiChallengeDTO>("api/whoami/try", { email, otp });
			return true;
		} catch (error) {
			display.enqueueError(error);
			return false;
		}
	}
}

export const whoami = new WhoamiStore();
