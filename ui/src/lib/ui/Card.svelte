<script lang="ts">
	let clazz: string = '';
	export { clazz as class };
</script>

<div class={`card ${clazz || ''}`}>
	<slot />
</div>

<style>
	.card {
		background-color: var(--light);
		box-shadow: var(--boxy-shadow);
		border: solid 1px var(--dark);
	}
</style>
