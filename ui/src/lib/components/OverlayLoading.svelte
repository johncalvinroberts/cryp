<script>
	import { onDestroy, onMount } from "svelte";
	import { delay } from "../utils";
	let progress = 0;
	let keepGoing = true;

	const go = async () => {
		if (!keepGoing) return;
		const ms = progress * 1.25;
		await delay(ms);
		progress++;
		go();
	};

	const stop = () => (keepGoing = false);
	onMount(() => go());
	onDestroy(() => stop());
</script>

<div class="box">
	<div class="overlay vertical-center">
		<div class="bar">
			<div class="progress" style="width: {progress}%" />
		</div>
	</div>
</div>

<style>
	.box {
		position: relative;
	}
	.overlay {
		position: absolute;
		left: 0;
		right: 0;
		top: 0;
		bottom: 0;
		background-color: var(--opaque);
	}
	.bar {
		width: 100%;
		height: calc(var(--nav-height) * 2);
		position: relative;
	}
	.progress {
		height: calc(var(--nav-height) * 2);
		left: 0;
		top: 0;
		background-color: var(--dark);
	}
</style>
