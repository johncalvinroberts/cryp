<script lang="ts">
	import Encrypter from "../Encrypter.svelte";
	import Card from "../Card.svelte";
	import MacintoshBar from "../MacintoshBar.svelte";
	import Banner from "../Banner.svelte";
	import { encrypter } from "../../stores/encrypter";
	const { store: encrypterStore } = encrypter;
	$: state = $encrypterStore.state;
</script>

<Card class="encrypter-card" isDraggable>
	<MacintoshBar>
		{state}
	</MacintoshBar>
	<div class="card-inner vertical-center">
		<Encrypter />
	</div>
</Card>
<div class="banner-wrapper">
	<Banner />
</div>

<style>
	.card-inner {
		min-height: 600px;
		padding-bottom: calc(var(--spacing) * 4);
	}

	:global(.encrypter-card) {
		min-height: 600px;
		width: 380px;
		position: absolute;
	}
	.banner-wrapper {
		min-height: calc(100vh - var(--nav-height));
		display: flex;
		justify-content: flex-end;
		align-items: center;
		max-width: 1200px;
	}
</style>
