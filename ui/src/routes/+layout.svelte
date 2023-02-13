<script>
	import { onMount } from "svelte";
	import { whoami } from "$lib/stores/whoami";
	import { encrypter } from "$lib/stores/encrypter";
	import Dropdown from "$lib/ui/Dropdown.svelte";
	import Control from "$lib/ui/icons/Control.svelte";
	import Money from "$lib/ui/icons/Money.svelte";
	import Keycaps from "$lib/ui/icons/Keycaps.svelte";
	import { theme } from "$lib/stores/theme";

	const { store: encrypterStore } = encrypter;
	const { store: whoamiStore } = whoami;
	let title = `furizu. | ${$encrypterStore.state}`;
	const isAuthenticated = $whoamiStore.isAuthenticated;
	const email = $whoamiStore.email;
	const dropdownOptions = [
		...(isAuthenticated
			? [
					{
						Icon: Money,
						label: "Authenticate",
						onClick: () => alert("TODO: implement me"),
					},
			  ]
			: [
					{
						Icon: Keycaps,
						label: "Authenticate",
						onClick: () => alert("TODO: implement me"),
					},
			  ]),
		{
			Icon: Control,
			label: "Settings",
			href: "/settings",
		},
	];
	// TODO: make this isomorphic + no FOUC
	onMount(() => theme.init());
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<nav>
	<Dropdown label={isAuthenticated ? email : "Guest"} options={dropdownOptions} />
</nav>

<main>
	<slot />
</main>

<style>
	nav {
		background-color: var(--gray);
		height: var(--nav-height);
		width: 100%;
		display: flex;
		justify-content: flex-end;
		padding: 0 calc(var(--spacing) * 4);
		border-bottom: solid 1px var(--dark);
	}

	main {
		padding: calc(var(--spacing) * 4);
		position: relative;
	}
</style>
