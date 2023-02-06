<script lang="ts">
	type Variant = "dropdown" | "default";
	export let disabled: boolean = false;
	export let variant: Variant = "default"
	const additionalClasses: Record<Variant, string> = {
		"default": "",
		"dropdown": "dropdown-button"
	}
</script>

<button class={`button ${additionalClasses[variant] || ''}`} disabled={disabled}>
	<slot></slot>
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
	}	
	.dropdown-button:hover {
		box-shadow: none;
	}
	.dropdown-button:focus,.dropdown-button:active {
		border: none;
		outline: none;
		background-color: var(--dark);
		color: var(--light);
	}
</style>
