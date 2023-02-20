<script lang="ts">
	import { display } from "../stores/display";
	import Checkbox from "./form/Checkbox.svelte";

	const { store } = display;
</script>

<div>
	<Checkbox
		checked={$store.theme === "dark"}
		on:change={() => display.toggleTheme()}
		id="theme-toggle"
		label="Dark Mode"
	/>
</div>

<style>
</style>
