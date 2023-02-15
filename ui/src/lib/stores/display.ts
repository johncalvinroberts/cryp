import { get } from "svelte/store";
import { browser } from "$app/environment";
import { THEME_STORAGE_KEY } from "../constants";
import BaseStore from "./base";

type Theme = "dark" | "light";

type ThemeState = {
	theme: Theme;
	isAuthModalOpen: boolean;
};

const initialState: ThemeState = {
	theme: "light",
	isAuthModalOpen: false,
};

class DisplayStore extends BaseStore<ThemeState> {
	constructor() {
		super(initialState);
	}

	public init() {
		let initialTheme = localStorage.getItem(THEME_STORAGE_KEY);
		if (!initialTheme) {
			initialTheme = "light";
		}
		this.setTheme(<Theme>initialTheme);
	}

	public setTheme(nextTheme: Theme) {
		localStorage.setItem(THEME_STORAGE_KEY, nextTheme);
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
}

export const display = new DisplayStore();
