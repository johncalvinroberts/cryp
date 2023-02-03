<script lang="ts">
	import fileSize from 'pretty-bytes';
	export let bytes: number | undefined = undefined;
	let className = '';
	export { className as class };
</script>

<span class={className}>
	{fileSize(bytes ?? 0)}
</span>

<style>
	span {
		opacity: 0.5;
		font-size: 0.8rem;
		font-weight: normal;
		margin-left: 0.2rem;
	}
</style>
