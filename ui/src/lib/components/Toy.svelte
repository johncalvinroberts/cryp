<script>
	import { Toy } from "@leveluptuts/svelte-toy";
	import { encrypter } from "../stores/encrypter";
	import { display } from "../stores/display";
	import { whoami } from "../stores/whoami";
</script>

<Toy register={{ encrypter: encrypter.store, whoami: whoami.store, display: display.store }} />
