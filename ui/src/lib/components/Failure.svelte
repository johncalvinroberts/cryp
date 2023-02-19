<script lang="ts">
	import { encrypter } from "../stores/encrypter";
	const { store, reset } = encrypter;
</script>

<div class="box vertical-center">
	<h2>Error</h2>
	<div>An error occured during the encryption or decryption process.</div>
	<div>
		<strong>Error: </strong>
		<em>
			{$store.error}
		</em>
	</div>
	<div class="bottom-box">
		<button on:click={reset}> Start Over </button>
	</div>
</div>

<style>
	.box {
		flex-wrap: wrap;
		width: 100%;
		justify-content: center;
	}
	.box h2,
	.box div {
		width: 100%;
		text-align: center;
		margin-bottom: 0.5rem;
	}
	h2 {
		color: var(--error);
	}
</style>
