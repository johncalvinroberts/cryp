import type { ApiResponse, HTTPMethod, HTTPRequestBody } from "../types/types";
import { GET, POST, PUT, DELETE, PATCH } from "./constants";

class HTTPClient {
	constructor(private baseURL: string) {}

	public async fetch<T>(
		path: string,
		method: HTTPMethod,
		body?: HTTPRequestBody,
		givenHeaders?: Record<string, string>,
	): Promise<T> {
		const headers: Headers = new Headers(givenHeaders);
		const res = await fetch(`${this.baseURL}/${path}`, {
			headers,
			method,
			body: JSON.stringify(body),
		});
		const duplicate = res.clone();
		const json: ApiResponse<T> = await res.json();
		if (!res.ok || json.error) {
			let message = json.error;
			if (!message) {
				message = await duplicate.text();
			}
			throw new Error(message);
		}
		return json.data;
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

export default HTTPClient;
