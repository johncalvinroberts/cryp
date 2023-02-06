<div class="parent">
	<div class="lines">
		<div class="line-box">
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
		</div>
		<div class="text">
			<slot />
		</div>
		<div class="line-box">
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
			<span class="line" />
		</div>
	</div>
	<div class="bottom-bar" />
</div>

<style>
	.parent {
		cursor: grab;
	}
	.lines {
		padding: 0 3px;
		display: flex;
		height: var(--spacing) * 2;
		border-bottom: solid 1px var(--dark);
	}
	.text {
		flex: 0;
		padding: 0 var(--spacing);
	}
	.line-box {
		flex: 1;
		display: flex;
		flex-wrap: wrap;
		flex-direction: row;
		justify-content: space-around;
		padding: 3px 0;
	}
	.line {
		width: 100%;
		border-bottom: solid 1px var(--dark);
		display: block;
		flex: 0 0 100%;
	}
	.bottom-bar {
		height: 4px;
		border-bottom: 1px solid var(--dark);
		width: 100%;
	}
</style>
