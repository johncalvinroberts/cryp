<script lang="ts">
	import { encrypter } from '../stores/encrypter';
	import { STATE } from '../constants';
	const { store: encrypterStore } = encrypter;
	$: state = $encrypterStore.state;
</script>

{#if state === STATE.INITIAL}
	<div class="initial vertical-center">
		<div>
			<h1>furizu.</h1>
			<h2>Encrypted deep-freeze file storage</h2>
		</div>
	</div>
{/if}

<style>
	.initial h1 {
		font-size: var(--font-size-huge);
		margin: 0;
	}

	.initial h2 {
		font-size: var(--font-size-h1);
	}
</style>
