<script>
	import Dropzone from "./Dropzone.svelte";
	import EncryptedPreview from "./EncryptedPreview.svelte";
	import { encrypter } from "../stores/encrypter";
	import FilesPreview from "./FilesPreview.svelte";
	import Done from "./Done.svelte";
	import Failure from "./Failure.svelte";
	import Processing from "./Processing.svelte";
	import { STATE } from "../constants";

	const { store } = encrypter;
	const elements = {
		[STATE.INITIAL]: Dropzone,
		[STATE.SHOULD_DECRYPT]: EncryptedPreview,
		[STATE.SHOULD_ENCRYPT]: FilesPreview,
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
