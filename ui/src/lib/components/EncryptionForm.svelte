<script lang="ts">
	import Empty from "./Empty.svelte";
	import { encrypter } from "../stores/encrypter";
	import FileSize from "./FileSize.svelte";
	import Form from "./form/Form.svelte";
	import Input from "./form/Input.svelte";

	const { store, reset, handleEncrypt } = encrypter;
	const accepted = $store.filesToEncrypt?.accepted || [];
	const rejected = $store.filesToEncrypt?.rejected || [];
	const totalFileBytes = accepted.reduce((memo, current) => {
		return memo + current.size;
	}, 0);
	let password = "";
	let hint = "";

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		handleEncrypt(password, hint);
	};
</script>

<div class="wrapper">
	{#if accepted.length > 1}
		<div class="title">
			<h6>
				{accepted.length} Files -
			</h6>
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

	{#if rejected?.length}
		<h6>Rejected Files</h6>
		<ul>
			{#each rejected as rejected}
				<li>{rejected.file.name} - {rejected.error.message}</li>
			{/each}
		</ul>
	{/if}
	<Form on:submit={handleSubmit} on:reset={reset}>
		<Input placeholder="Password" type="password" name="secret" bind:value={password} />
		<Input name="hint" placeholder="Hint (optional)" bind:value={hint} />
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
	h6 {
		text-align: center;
		margin: 0;
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
		justify-content: center;
		margin-top: 3rem;
	}
</style>
