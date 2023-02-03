<script lang="ts">
	import Check from '../icons/check.svelte';
	import { encrypter } from '../stores/encrypter';
	import { saveAs } from 'file-saver';
	import { getEncryptedFilename } from '../utils';

	const { store, reset } = encrypter;
	const type = $store.decryptedFiles ? 'decrypted' : 'encrypted';

	let isFailToDownload = false;

	const downloadEncrypted = () => {
		isFailToDownload = false;
		try {
			if (!$store.filesToEncrypt || !$store.crypString) {
				throw new Error('Cannot create download');
			}
			const fileName = getEncryptedFilename($store.filesToEncrypt);
			const file = new File([$store.crypString], fileName);
			saveAs(file);
		} catch (error) {
			isFailToDownload = true;
			console.error(error);
		}
	};

	const downloadDecrypted = () => {
		isFailToDownload = false;
		try {
			if (!$store.decryptedFiles) {
				throw new Error('Cannot create download');
			}
			for (const file of $store.decryptedFiles) {
				saveAs(file, file.name);
			}
		} catch (error) {
			isFailToDownload = true;
			console.error(error);
		}
	};

	const handleDownload = () => (type === 'encrypted' ? downloadEncrypted() : downloadDecrypted());
</script>

<div>
	<h3 class="vertical-center">
		<span class="icon vertical-center">
			<Check width={40} />
		</span>
		Successfully {type}
	</h3>
	{#if isFailToDownload}
		<div class="error">Failed to Download File.</div>
	{/if}
	<div class="bottom-box">
		<button class="vertical-center success-button" on:click={handleDownload}> Download </button>
		<button on:click={reset}> Start Over </button>
	</div>
</div>

<style>
	h3 {
		text-align: center;
	}
	.bottom-box {
		margin: 0 auto;
	}

	.bottom-box button {
		max-width: 150px;
		margin: 0 auto;
	}

	.success-button {
		color: var(--light);
		background: -webkit-linear-gradient(var(--purple), var(--info));
		border: solid 1px var(--light);
		animation: floating 20s ease-in-out infinite;
		border-radius: 2px;
		transition: all 0.3s ease;
	}
	.success-button:hover {
		border: solid 2px var(--light);
	}
	.icon {
		margin-right: 0.5rem;
	}

	.error {
		color: var(--error);
		text-align: center;
		padding-top: 1rem;
	}
</style>
