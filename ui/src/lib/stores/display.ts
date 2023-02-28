import { get } from "svelte/store";
import { THEME_LOCAL_STORAGE_KEY } from "../constants";
import type { MaybeError, ThemeState, Theme } from "../../types/types";
import BaseStore from "./base";
import { browser } from "$app/environment";

const initialState: ThemeState = {
	theme: "light",
	isAuthModalOpen: false,
	errors: [],
};

class DisplayStore extends BaseStore<ThemeState> {
	constructor() {
		super(initialState);
	}

	public init() {
		let initialTheme = localStorage.getItem(THEME_LOCAL_STORAGE_KEY);
		if (!initialTheme) {
			initialTheme = "light";
		}
		this.setTheme(<Theme>initialTheme);
	}

	public setTheme(nextTheme: Theme) {
		localStorage.setItem(THEME_LOCAL_STORAGE_KEY, nextTheme);
		this.dispatch({ theme: nextTheme });
		if (browser) {
			const toRemove = nextTheme === "dark" ? "light" : "dark";
			document.documentElement.classList.remove(toRemove);
			document.documentElement.classList.add(nextTheme);
		}
	}

	public toggleTheme() {
		const { theme } = get(this.store);
		const nextTheme = theme === "dark" ? "light" : "dark";
		this.setTheme(nextTheme);
	}

	public toggleAuthModal() {
		const { isAuthModalOpen } = get(this.store);
		this.dispatch({ isAuthModalOpen: !isAuthModalOpen });
	}

	public enqueueError(err: MaybeError) {
		const { errors } = get(this.store);
		this.dispatch({ errors: [...errors, err] });
	}

	public dequeueError(): MaybeError {
		const { errors } = get(this.store);
		const [err, ...rest] = errors;
		this.dispatch({ errors: rest });
		return err;
	}
}

export const display = new DisplayStore();
