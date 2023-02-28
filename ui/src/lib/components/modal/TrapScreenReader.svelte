<script lang="ts">
	import { onMount } from "svelte";
	type PartialAttribute = {
		inert: string | undefined;
		ariaHidden: string | undefined;
	};
	// Props
	export let enabled: boolean;
	let originalAttributes: PartialAttribute[] = [];

	const hideFromScreenReader = (node: HTMLElement) => {
		originalAttributes.push({
			ariaHidden: node.getAttribute("aria-hidden") ?? undefined,
			inert: node.getAttribute("inert") ?? undefined,
		});
		node.setAttribute("aria-hidden", "true");
		node.setAttribute("inert", "true");
	};

	const exposeToScreenReader = (node: HTMLElement, i: number) => {
		const { ariaHidden, inert } = originalAttributes[i];
		if (!ariaHidden) {
			node.removeAttribute("aria-hidden");
		} else {
			node.setAttribute("aria-hidden", ariaHidden);
		}
		if (!inert) {
			node.removeAttribute("inert");
		}
	};
	onMount(() => {
		if (!enabled) {
			// `DialogContent` has the `aria-modal` attribute. This indicates to screen readers
			// that only content contained within the dialog should be accessible to the user.
			// Modern screen readers respect this attribute. In cases where support is inadequate,
			// this legacy workaround can be enabled.
			return;
		}
		// Grab all children in the `body` except for the dialog portal
		const children = Array.from(
			document.querySelectorAll("body > *:not([data-svelte-dialog-portal])"),
		);
		children.forEach((node) => hideFromScreenReader(node as HTMLElement));
		return () => {
			children.forEach((node, i) => exposeToScreenReader(node as HTMLElement, i));
		};
	});
</script>

<slot />
