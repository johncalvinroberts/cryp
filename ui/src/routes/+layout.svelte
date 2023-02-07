<script>
	import { whoami } from '../lib/stores/whoami';
	import { encrypter } from '../lib/stores/encrypter';
	import Button from '../lib/ui/Button.svelte';

	const { store: encrypterStore } = encrypter;
	const { store: whoamiStore } = whoami;
	$: state = $encrypterStore.state;
	let title = `furizu. | ${state}`;
	const isAuthenticated = $whoamiStore.isAuthenticated;
	const email = $whoamiStore.email;
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<nav>
	<Button variant="dropdown">
		{isAuthenticated ? email : 'Guest'}
	</Button>
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
		padding: var(--spacing);
	}
</style>
