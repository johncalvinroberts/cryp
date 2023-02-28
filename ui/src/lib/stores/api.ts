import decodeJwt, { type JwtPayload } from "jwt-decode";
import { get } from "svelte/store";
import type { APIClientState, HTTPMethod, HTTPRequestBody } from "../../types/types";
import type { RefreshTokenDTO } from "../../types/dtos";
import HTTPClient from "../http";
import { JWT_LOCAL_STORAGE_KEY, API_BASE_URL, GET, POST, PUT, DELETE, PATCH } from "../constants";
import { delay } from "../utils";
import BaseStore from "./base";
import { display } from "./display";
import { whoami } from "./whoami";

const MIN_TOKEN_REFRESH_MS = 1000 * 60; // 1 min
const TOKEN_REFRESH_BACKOFF_MS = 700;

const initialState: APIClientState = {
	isRefreshingToken: false,
	tokenExpiresAt: undefined,
};

/**
 * This API client is essentially a wrapper around the HTTPClient,
 * but with stateful handling of auth state.
 * Using the Svelte store to track the state of the
 */

class APIClientStore extends BaseStore<APIClientState> {
	private httpClient: HTTPClient;
	constructor() {
		super(initialState);
		const initialToken = localStorage?.getItem(JWT_LOCAL_STORAGE_KEY) ?? undefined;
		this.httpClient = new HTTPClient(initialToken, API_BASE_URL);
		if (initialToken != null) {
			this.handleToken(initialToken);
		}
	}

	public handleToken(token: string) {
		try {
			localStorage.setItem(JWT_LOCAL_STORAGE_KEY, token);
			const decoded: JwtPayload = decodeJwt(token);
			if (decoded.exp) {
				this.dispatch({ tokenExpiresAt: decoded.exp });
			}
			this.httpClient.setToken(token);
		} catch (error) {
			localStorage.removeItem(JWT_LOCAL_STORAGE_KEY);
			this.httpClient.removeToken();
		}
	}

	public async refreshToken() {
		try {
			const res = await this.httpClient.post<RefreshTokenDTO>("api/whoami/refresh", {});
			this.handleToken(res.jwt);
		} catch (error) {
			display.enqueueError(error);
			whoami.reset();
		}
	}

	private async fetch<T>(path: string, method: HTTPMethod, body?: HTTPRequestBody): Promise<T> {
		const { isRefreshingToken, tokenExpiresAt } = get(this.store);
		if (isRefreshingToken) {
			await delay(TOKEN_REFRESH_BACKOFF_MS);
			return this.fetch<T>(path, method, body);
		}
		if (tokenExpiresAt != null && tokenExpiresAt - Date.now() < MIN_TOKEN_REFRESH_MS) {
			await this.refreshToken();
		}
		return this.httpClient.fetch<T>(path, method, body);
	}

	public get<T>(path: string): Promise<T> {
		return this.fetch(path, GET);
	}

	public post<T>(path: string, body: HTTPRequestBody): Promise<T> {
		return this.fetch(path, POST, body);
	}

	public delete<T>(path: string, body: HTTPRequestBody): Promise<T> {
		return this.fetch(path, DELETE, body);
	}

	public patch<T>(path: string, body: HTTPRequestBody): Promise<T> {
		return this.fetch(path, PATCH, body);
	}

	public put<T>(path: string, body: HTTPRequestBody): Promise<T> {
		return this.fetch(path, PUT, body);
	}
}

export const apiClient = new APIClientStore();
