<script lang="ts">
	export let title = "";
	export let disabled = false;
</script>

<form class="form" class:disabled on:submit|preventDefault on:reset autocomplete="off" {...$$props}>
	{#if disabled}
		<div class="disabled-overlay" />
	{/if}
	{#if title}
		<h3>{title}</h3>
	{/if}
	<slot />
</form>

<style>
	.form {
		flex: 0 0 300px;
		position: relative;
		margin-bottom: var(--spacing);
	}
	.form.disabled {
		pointer-events: none;
		cursor: not-allowed;
	}
	.disabled-overlay {
		background-color: var(--opaque);
		position: absolute;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		z-index: var(--z-index-popover);
	}
	h3 {
		text-align: center;
	}
</style>
