<script lang="ts">
	import Eye from "../icons/Eye.svelte";
	import MacintoshHD from "../icons/MacintoshHD.svelte";
	export let type:
		| "color"
		| "date"
		| "datetime-local"
		| "email"
		| "file"
		| "hidden"
		| "image"
		| "month"
		| "number"
		| "password"
		| "reset"
		| "submit"
		| "tel"
		| "text"
		| "time"
		| "url"
		| "week"
		| "search" = "text";
	export let name: string;
	export let label = "";
	export let placeholder = "";
	export let value: string | number = "";
	let showPassword = false;
	$: actualType = type === "password" && showPassword ? "text" : type;
	// you need to this to avoid 2-way binding
	const setType = (node: HTMLInputElement, _type: string) => {
		node.type = _type;
		return {
			update(_type: string) {
				node.type = _type;
			},
		};
	};
</script>

<div class="input-box">
	<label for={name}>
		{label}
	</label>
	{#if type !== "file"}
		<input
			id={name}
			{name}
			{placeholder}
			on:change
			on:input
			bind:value
			autocomplete="off"
			spellcheck="false"
			use:setType={actualType}
			{...$$restProps}
		/>
	{/if}
	{#if type === "password"}
		<button
			type="button"
			class="vertical-center"
			on:click={() => (showPassword = !showPassword)}
			title={showPassword ? "Hide" : "Show"}
		>
			<Eye strikethrough={!showPassword} />
		</button>
	{/if}
	{#if type === "file"}
		<div class="file-input" role="button">
			<input type="file" id={name} {name} {placeholder} on:change />
			<div class="vertical-center">
				<MacintoshHD />
			</div>
		</div>
	{/if}
	<slot />
</div>

<style>
	input {
		background-color: transparent;
		border: solid 1px var(--dark);
		width: 100%;
		font-size: 0.9rem;
		color: var(--dark);
		height: 21px;
		max-width: 200px;
	}

	label {
		flex: 0 0 100%;
	}

	.input-box {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: var(--spacing);
		position: relative;
		max-width: 300px;
		flex-wrap: wrap;
	}
	.input-box button {
		background-color: transparent;
		border: none;
		opacity: 0.5;
		transition: opacity 0.5s ease;
		position: absolute;
		right: 0;
	}
	.input-box button:hover {
		opacity: 1;
	}

	.file-input {
		cursor: pointer;
		height: 21px;
		border: solid 1px var(--dark);
		display: flex;
		max-width: 200px;
		padding-right: 10px;
	}
	.file-input input {
		opacity: 0;
		cursor: pointer;
	}
</style>
