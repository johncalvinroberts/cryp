<script lang="ts">
	import AuthForm from "./AuthForm.svelte";
	import { encrypter } from "../stores/encrypter";
	import { STATE } from "$lib/constants";
	const backText = "<- Back";
	const handleBack = () => {
		encrypter.dispatch({ state: STATE.SHOULD_ENCRYPT });
	};
</script>

<AuthForm />
<button class="back" on:click={handleBack}>{backText}</button>

<style>
	button.back {
		border: none;
		color: var(--gray);
	}
</style>
