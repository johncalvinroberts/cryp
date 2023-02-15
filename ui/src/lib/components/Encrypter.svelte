<script>
	import Dropzone from "./Dropzone.svelte";
	import DecryptionForm from "./DecryptionForm.svelte";
	import { encrypter } from "../stores/encrypter";
	import EncryptionForm from "./EncryptionForm.svelte";
	import Done from "./Done.svelte";
	import Failure from "./Failure.svelte";
	import Processing from "./Processing.svelte";
	import { STATE } from "../constants";

	const { store } = encrypter;
	const elements = {
		[STATE.INITIAL]: Dropzone,
		[STATE.SHOULD_DECRYPT]: DecryptionForm,
		[STATE.SHOULD_ENCRYPT]: EncryptionForm,
		[STATE.PROCESSING]: Processing,
		[STATE.FAILURE]: Failure,
		[STATE.DONE]: Done,
	};
	$: console.info($store.state);
</script>

<div class="encrypter vertical-center">
	<svelte:component this={elements[$store.state]} />
</div>

<style>
	.encrypter {
		width: 100%;
	}
</style>
