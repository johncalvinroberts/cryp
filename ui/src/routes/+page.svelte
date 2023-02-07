<script lang="ts">
	import Encrypter from '../lib/ui/Encrypter.svelte';
	import Card from '../lib/ui/Card.svelte';
	import MacintoshBar from '../lib/ui/MacintoshBar.svelte';
	import Banner from '../lib/ui/Banner.svelte';
	import { encrypter } from '../lib/stores/encrypter';
	const { store: encrypterStore } = encrypter;
	$: state = $encrypterStore.state;
</script>

<Card class="card" isDraggable>
	<MacintoshBar>
		{state}
	</MacintoshBar>
	<div class="inner vertical-center">
		<Encrypter />
	</div>
</Card>
<div class="banner-wrapper">
	<Banner />
</div>

<style>
	.inner {
		min-height: 600px;
		padding-bottom: calc(var(--spacing) * 4);
	}

	:global(.card) {
		min-height: 600px;
		position: absolute;
		max-width: 320px;
		min-width: 320px;
	}
	.banner-wrapper {
		min-height: calc(100vh - var(--nav-height));
		display: flex;
		justify-content: flex-end;
		align-items: center;
		max-width: 1200px;
	}
</style>
