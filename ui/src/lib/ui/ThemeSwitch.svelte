<script lang="ts">
	import { theme } from "../stores/theme";

	const { store, toggleTheme } = theme;
	const back = "<-";
</script>

<div>
	<a href="/">
		<span class="arrow">
			{back}
		</span>
		Home
	</a>
	<div class="theme-switch">
		<input
			checked={$store.theme === "dark"}
			on:click={toggleTheme}
			type="checkbox"
			id="theme-toggle"
		/>
		<label for="theme-toggle"> Dark Mode </label>
	</div>
</div>

<style>
	.theme-switch {
		display: flex;
	}
	.theme-switch label {
		width: 100%;
	}
	.theme-switch input {
		flex: 0;
	}
	.arrow {
		transform: rotate(90deg);
	}

	a {
		margin-bottom: calc(var(--spacing) * 2);
		display: block;
	}
</style>
