<script lang="ts">
	export let checked: boolean;
	export let id: string;
	export let label: string;
</script>

<div class="checkbox-box">
	<input {checked} {id} on:change type="checkbox" />
	<label for={id}>
		{label}
		{#if checked}
			<span class="check left" />
			<span class="check right" />
		{/if}
	</label>
</div>

<style>
	.checkbox-box {
		display: flex;
		position: relative;
	}
	.checkbox-box label {
		width: 100%;
	}
	.checkbox-box input {
		flex: 0;
	}

	input[type="checkbox"] {
		opacity: 0;
		position: absolute;
		left: 0;
		top: 0;
		z-index: 9;
	}

	input[type="checkbox"] + label {
		position: relative;
		padding-left: calc(var(--spacing) * 4);
		display: flex;
		align-items: center;
	}

	input[type="checkbox"] + label::before {
		width: calc(var(--spacing) * 3);
		height: calc(var(--spacing) * 3);
		border: solid 1px var(--dark);
		position: absolute;
		left: 0;
		content: " ";
		cursor: pointer;
		z-index: 10;
	}

	.check {
		left: 7px;
		position: absolute;
		height: 18px;
		border-left: solid 1px var(--dark);
		z-index: 10;
	}
	.left {
		transform: rotate(-45deg);
	}

	.right {
		transform: rotate(45deg);
	}
</style>
