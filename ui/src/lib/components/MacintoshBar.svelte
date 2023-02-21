<div class="parent">
	<div class="lines">
		<div class="text vertical-center">
			<span>
				<slot />
			</span>
		</div>
	</div>
	<div class="bottom-bar" />
</div>

<style>
	.parent {
		cursor: grab;
	}

	.lines {
		border-bottom: solid 1px var(--dark);
		height: 17px;
		align-items: center;
		background: linear-gradient(#000 21%, transparent 0);
		background-clip: content-box;
		background-size: 6.6666666667% 15.333333%;
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		padding: 2px 5px;
	}
	.text {
		flex: 0;
		padding: 0 var(--spacing);
		background-color: var(--light);
		height: 13px;
		line-height: 0;
	}

	.bottom-bar {
		height: 4px;
		border-bottom: 1px solid var(--dark);
		width: 100%;
	}
</style>
