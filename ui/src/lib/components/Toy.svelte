<script>
	import { display } from "$lib/stores/display";
	import { whoami } from "$lib/stores/whoami";
	import { Toy } from "@leveluptuts/svelte-toy";
	import { encrypter } from "../stores/encrypter";
</script>

<Toy register={{ encrypter: encrypter.store, whoami: whoami.store, display: display.store }} />

<style>
	/* your styles go here */
</style>
