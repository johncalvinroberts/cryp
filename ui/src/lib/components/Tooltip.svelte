<script lang="ts">
	export let title = "";
	let isHovered = false;
	let x: number;
	let y: number;

	const handleMouseOver = (event: MouseEvent) => {
		isHovered = true;
		x = event.pageX + 5;
		y = event.pageY + 5;
	};

	const handleMouseMove = (event: MouseEvent) => {
		x = event.pageX + 5;
		y = event.pageY + 5;
	};

	const handleFocus = () => {
		// TODO
	};

	const handleMouseLeave = () => {
		isHovered = false;
	};
</script>

<div
	on:mouseover={handleMouseOver}
	on:mouseleave={handleMouseLeave}
	on:mousemove={handleMouseMove}
	on:focus={handleFocus}
>
	<slot />
	{#if isHovered}
		<div style="top: {y}px; left: {x}px;" class="tooltip">{title}</div>
	{/if}
</div>

<style>
	.tooltip {
		border: 1px solid #ddd;
		box-shadow: 1px 1px 1px #ddd;
		background: white;
		border-radius: 4px;
		padding: 4px;
		position: absolute;
	}
</style>
