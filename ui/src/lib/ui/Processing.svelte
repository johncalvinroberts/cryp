<script lang="ts">
	import { onDestroy } from 'svelte';
	import { getRandomUnicodeString } from '../utils';
	const CHAR_AMT = 25;

	let randomChars = getRandomUnicodeString(CHAR_AMT);
	let bottom = 0;
	const shiftString = () => {
		const nextStringArr = randomChars.split('');
		for (let i = 0; i < 4; i++) {
			const randomIndex = Math.floor(Math.random() * CHAR_AMT);
			nextStringArr[randomIndex] = getRandomUnicodeString(1);
		}
		randomChars = nextStringArr.join('');
		bottom = bottom + 1;
	};
	const interval = setInterval(shiftString, 100);
	onDestroy(() => clearInterval(interval));
</script>

<div class="container vertical-center">
	<div class="bar truncate" style="bottom:{bottom}px;">
		{randomChars}
	</div>
</div>

<style>
	.container {
		min-height: 100px;
		position: relative;
	}
	.bar {
		background-color: var(--yellow);
		color: var(--dark);
		width: 100%;
		text-align: center;
		transition: all 0.2s ease;
		position: absolute;
		left: 0;
		right: 0;
	}
</style>
