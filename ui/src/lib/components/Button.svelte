<script lang="ts">
	type Variant = "dropdown" | "default";
	export let disabled = false;
	export let variant: Variant = "default";
	export let ariaExpanded: boolean | null;
	export let ariaHaspopup:
		| boolean
		| "dialog"
		| "menu"
		| "true"
		| "false"
		| "grid"
		| "listbox"
		| "tree"
		| null
		| undefined;
	const additionalClasses: Record<Variant, string> = {
		default: "",
		dropdown: "dropdown-button",
	};
	const additionalProps = {
		...(disabled ? { disabled } : null),
		...(ariaExpanded != null ? { "aria-expanded": ariaExpanded } : null),
		...(ariaHaspopup ? { "aria-haspopup": ariaHaspopup } : null),
	};
</script>

<button
	class={`button ${additionalClasses[variant] || ""}`}
	{...additionalProps}
	on:click
	on:mouseleave
>
	<slot />
</button>

<style>
	.button {
		padding: 0 var(--spacing);
		min-width: 111px;
		height: var(--nav-height);
		background: var(--light);
		color: var(--dark);
		border: solid 1px var(--dark);
		border-radius: 5px;
		transition: all 0.2s ease;
	}

	.button:hover {
		box-shadow: var(--boxy-shadow);
	}

	.button:active {
		background-color: var(--dark);
		color: var(--light);
		border-color: var(--light);
	}

	.button:focus {
		outline: 2px solid var(--dark);
	}

	.button:disabled {
		opacity: 0.5;
	}

	.button:disabled:active {
		background-color: var(--light);
		color: var(--dark);
	}

	.button:disabled:hover {
		box-shadow: none;
	}

	.dropdown-button {
		border: none;
		border-radius: 0;
		background-color: transparent;
		position: relative;
		padding: 0;
	}
	.dropdown-button:hover {
		box-shadow: none;
	}
	.dropdown-button:focus,
	.dropdown-button:active {
		border: none;
		outline: none;
		background-color: var(--dark);
		color: var(--light);
	}
</style>
