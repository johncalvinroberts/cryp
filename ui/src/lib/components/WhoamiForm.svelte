<script lang="ts">
	import { whoami } from "../stores/whoami";
	import Button from "./Button.svelte";
	import Form from "./form/Form.svelte";
	import Input from "./form/Input.svelte";

	let step: "START_WHOAMI" | "TRY_WHOAMI" = "START_WHOAMI";
	let email: string;
	let otp: string;

	const handleStartWhoami = () => {
		whoami.startWhoamiChallenge(email);
	};

	const tryWhoamiChallenge = () => {
		whoami.tryWhoamiChallenge(email, otp);
	};
</script>

<div class="box">
	<Form on:submit={handleStartWhoami}>
		<Input name="email" type="email" label="Email" bind:value={email} />
		<Button type="submit">Send Code</Button>
	</Form>
	<Form on:submit={tryWhoamiChallenge} disabled={step !== "TRY_WHOAMI"} prevent:default>
		<Input name="otp" type="text" label="Code" disabled={step !== "TRY_WHOAMI"} bind:value={otp} />
		<Button type="submit" disabled={step !== "TRY_WHOAMI"}>Submit</Button>
	</Form>
</div>

<style>
	.box {
		margin-bottom: 100px;
	}
</style>
