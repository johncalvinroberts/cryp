<script lang="ts">
	import Empty from "./Empty.svelte";
	import { encrypter } from "../stores/encrypter";
	import FileSize from "./FileSize.svelte";
	import Form from "./form/Form.svelte";
	import Input from "./form/Input.svelte";
	import Button from "./Button.svelte";

	const { store, reset, handleEncrypt, handleFiles } = encrypter;
	$: files = $store.filesToEncrypt ?? [];
	$: totalFileBytes = files.reduce((memo, current) => {
		return memo + current.size;
	}, 0);
	let password = "";
	let hint = "";

	const handleAddFile = (e: Event) => {
		const files = (<HTMLInputElement>e.target).files;
		if (files) {
			handleFiles(Array.from(files));
		}
	};

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		handleEncrypt(password, hint);
	};
</script>

<div class="box">
	<div class="file-preview">
		<div class="title">
			<h2>
				{files.length} Files -
			</h2>
			<span class="vertical-center">
				<FileSize bytes={totalFileBytes} />
			</span>
		</div>
		<ul>
			{#each files as file}
				<li>
					<span class="file-name truncate">{file.name}</span>
					-
					<FileSize bytes={file.size} class="file-size" />
				</li>
			{/each}
		</ul>
		{#if !files?.length}
			<Empty />
		{/if}
	</div>
	<hr />
	<Form on:submit={handleSubmit} on:reset={reset}>
		<Input
			label="Secret Key"
			type="text"
			name="secret"
			bind:value={password}
			tip="Special password for decrypting the encrypted file. Do not lose this."
		/>
		<Input name="hint" label="Hint" bind:value={hint} tip="Optional secret key hint" />
		<Input
			label="Add more files ->"
			type="file"
			name="files"
			variant="minimal"
			on:change={handleAddFile}
			multiple={true}
		/>
		<div class="bottom-box">
			<Button type="reset">Cancel</Button>
			<Button type="submit" disabled={!password}>Encrypt</Button>
		</div>
	</Form>
</div>

<style>
	.box {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		max-width: 300px;
	}
	ul {
		padding: 0;
		flex: 0 0 100%;
	}
	li {
		display: flex;
		width: 100%;
		justify-content: flex-start;
		align-items: center;
	}
	.file-name {
		display: block;
		max-width: calc(70%);
	}
	.title {
		display: flex;
		width: 100%;
	}
	.title h2 {
		margin: 0;
	}

	.file-preview {
		flex: 0 0 100%;
		min-height: 130px;
	}
</style>
