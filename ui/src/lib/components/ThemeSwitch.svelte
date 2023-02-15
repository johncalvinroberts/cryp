<script lang="ts">
	import { display } from "../stores/display";
	import Checkbox from "./form/Checkbox.svelte";

	const { store } = display;
	const back = "<-";
</script>

<div>
	<a href="/">
		<span class="arrow">
			{back}
		</span>
		Home
	</a>
	<Checkbox
		checked={$store.theme === "dark"}
		on:change={() => display.toggleTheme()}
		id="theme-toggle"
		label="Dark Mode"
	/>
</div>

<style>
	.arrow {
		transform: rotate(90deg);
	}

	a {
		margin-bottom: calc(var(--spacing) * 2);
		display: block;
	}
</style>
