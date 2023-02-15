<script lang="ts">
	import Eye from "../icons/Eye.svelte";
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
	<label for={name}>
		{label}
	</label>
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
	<slot />
</div>

<style>
	input {
		background-color: transparent;
		border: none;
		width: 100%;
		font-size: 0.9rem;
		text-align: center;
		color: var(--dark);
	}

	.input-box {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: var(--spacing);
		position: relative;
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

	.input-box:not(:last-child) {
		border-bottom: solid 1px var(--dark);
	}
</style>
