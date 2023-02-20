<script lang="ts">
	import ThemeSwitch from "$lib/components/ThemeSwitch.svelte";
	const back = "<-";
</script>

<div>
	<a href="/">
		<span class="arrow">
			{back}
		</span>
		Home
	</a>
	<ThemeSwitch />
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
