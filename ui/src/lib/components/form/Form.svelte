<script lang="ts">
	export let title = "";
</script>

<form class="form" on:submit on:reset autocomplete="off">
	{#if title}
		<h3>{title}</h3>
	{/if}
	<slot />
</form>

<style>
	.form {
		flex: 0 0 300px;
	}
	h3 {
		text-align: center;
	}
</style>
