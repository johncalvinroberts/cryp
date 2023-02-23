<script lang="ts">
	import { encrypter } from "../stores/encrypter";
	import { STATE } from "../constants";
	import AuthForm from "./AuthForm.svelte";
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
