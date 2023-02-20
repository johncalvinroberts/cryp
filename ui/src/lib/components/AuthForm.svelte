<script lang="ts">
	import Button from "./Button.svelte";
	import Form from "./form/Form.svelte";
	import Input from "./form/Input.svelte";

	let step: "START_WHOAMI" | "TRY_WHOAMI" = "START_WHOAMI";

	const handleStartWhoami = () => {
		//
	};

	const tryWhoamiChallenge = () => {
		//
	};
</script>

<Form on:submit={handleStartWhoami}>
	<Input name="email" type="email" label="Email" />
	<Button type="submit">Send Code</Button>
</Form>
<Form on:submit={tryWhoamiChallenge} disabled={step !== "TRY_WHOAMI"}>
	<Input name="code" type="text" label="Code" />
	<Button type="submit">Submit</Button>
</Form>

<style>
</style>
