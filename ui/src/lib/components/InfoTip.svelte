<script lang="ts">
	import Tooltip from "./Tooltip.svelte";
	export let title = "";
</script>

<div class="tooltip-outer">
	<Tooltip {title}>
		<small>?</small>
	</Tooltip>
</div>

<style>
	.tooltip-outer {
		position: relative;
		display: block;
	}
</style>
