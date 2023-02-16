<script lang="ts">
	import Empty from "./Empty.svelte";
	import { encrypter } from "../stores/encrypter";
	import FileSize from "./FileSize.svelte";
	import Form from "./form/Form.svelte";
	import Input from "./form/Input.svelte";

	const { store, reset, handleEncrypt, handleFiles } = encrypter;
	$: accepted = $store.filesToEncrypt ?? [];
	$: totalFileBytes = accepted.reduce((memo, current) => {
		return memo + current.size;
	}, 0);
	let name = "";
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
		handleEncrypt(name, password, hint);
	};
</script>

<div class="wrapper">
	{#if accepted.length > 1}
		<div class="title">
			<h2>
				{accepted.length} Files -
			</h2>
			<span class="vertical-center">
				<FileSize bytes={totalFileBytes} />
			</span>
		</div>
	{/if}
	<ul>
		{#each accepted as file}
			<li>
				<span class="file-name truncate">{file.name}</span>
				-
				<FileSize bytes={file.size} class="file-size" />
			</li>
		{/each}
	</ul>
	{#if !accepted?.length}
		<Empty />
	{/if}

	<Form on:submit={handleSubmit} on:reset={reset}>
		<Input label="Select File" type="file" name="files" on:change={handleAddFile} />
		<Input label="Name" type="text" name="name" bind:value={name} />
		<Input label="Secret Key" type="text" name="secret" bind:value={password} />
		<Input name="hint" label="Hint (optional)" bind:value={hint} />
		<div class="bottom-box">
			<button type="reset"> Cancel </button>
			<button type="submit" disabled={!password}> Encrypt </button>
		</div>
	</Form>
</div>

<style>
	.wrapper {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
	}
	ul {
		padding: 0;
		flex: 0 0 100%;
		max-width: 300px;
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
		max-width: 300px;
	}
	.title h2 {
		margin: 0;
	}
</style>
