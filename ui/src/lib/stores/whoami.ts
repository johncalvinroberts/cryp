import type { WhoamiState } from "../../types/types";
import type { TryWhoamiChallengeDTO } from "../../types/dtos";
import BaseStore from "./base";
import { apiClient } from "./api";
import { display } from "./display";

const initialState: WhoamiState = {
	isAuthenticated: false,
	email: undefined,
	isChallengeSent: false,
	isLoading: false,
};

class WhoamiStore extends BaseStore<WhoamiState> {
	constructor() {
		super(initialState);
	}

	public async startWhoamiChallenge(email: string) {
		try {
			this.dispatch({ isLoading: true });
			await apiClient.post("api/whoami/start", { email });
			this.dispatch({ isChallengeSent: true });
		} catch (error) {
			display.enqueueError(error);
			throw error;
		} finally {
			this.dispatch({ isLoading: false });
		}
	}

	public async tryWhoamiChallenge(email: string, otp: string) {
		try {
			this.dispatch({ isLoading: true });
			const res = await apiClient.post<TryWhoamiChallengeDTO>("api/whoami/try", { email, otp });
			this.dispatch({ email, isAuthenticated: true });
			apiClient.handleToken(res.jwt);
		} catch (error) {
			display.enqueueError(error);
			throw error;
		} finally {
			this.dispatch({ isLoading: false });
		}
	}
}

export const whoami = new WhoamiStore();
