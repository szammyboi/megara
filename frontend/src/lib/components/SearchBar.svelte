<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';

	interface Props {
		symbol: string;
	}

	let {
		symbol,
	}: Props = $props();
</script>


<section>
	<input />
	<div id="symbolarea">
		<div id="symbol">
			{symbol}
		</div>
	</div>
</section>

<style>
	section {
		background: #362712;
		width: 85vw;
		height: 80px;
		display: flex;
		flex-direction:row;
		font-family: 'Jersey';
		border-left: solid 6px #da1d4c;
	}

	input {
		all: unset;
		padding-left: 20px;
		outline: none;
		appearance: none;
		width: 100%;
		font-size: 32px;
		color: #d9ae80;
	}

	#symbolarea {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 20%;
		height: 100%;
	}

	#symbol {
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: #ff5813;
		width: 5vw;
		height: 5vw;
		font-size: 22px;
	}
</style>
