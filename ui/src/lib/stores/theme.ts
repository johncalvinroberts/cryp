import { get, writable, type Writable } from "svelte/store";
import { browser } from "$app/environment";
import { THEME_STORAGE_KEY } from "../constants";

type Theme = "dark" | "light";

type ThemeState = {
	theme: Theme;
};

const initialState: ThemeState = {
	theme: "light",
};

class ThemeStore {
	constructor(public store: Writable<ThemeState> = writable(initialState)) {}

	private dispatch(payload: Partial<ThemeState>) {
		this.store.update((prevState) => ({ ...prevState, ...payload }));
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
}

export const theme = new ThemeStore();
