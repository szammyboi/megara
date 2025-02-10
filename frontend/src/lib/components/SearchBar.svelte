<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';

	interface Props {
		color?: string;
		symbol: string;
		value?: string;
	}

	let {
		color,
		symbol,
		value = $bindable(),
	}: Props = $props();
</script>


<section>
	<input bind:value={value} autocomplete="off" spellcheck="false">
	<div id="symbolarea">
		<div id="symbol">
			{symbol}
		</div>
	</div>
</section>

<style>
	section {
		background: #362712;
		width: 55vw;
		height: 60px;
		display: flex;
		flex-direction:row;
		--border-left: solid 6px #da1d4c;
		margin-right: 7vw;
	}

	input {
		all: unset;
		padding-left: 20px;
		outline: none;
		appearance: none;
		width: 100%;
		font-size: 30px;
		color: #d9ae80;
		font-family: "Alagard";
	}

	#symbolarea {
		display: flex;
		align-items: center;
		justify-content: center;
		width: fit-content;
		height: 100%;
		margin-right: calc(calc(60px - 4.25vw) / 2);
	}

	#symbol {
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: #ff5813;
		color: #201a01;
		width: 4.25vw;
		height: 4.25vw;
		font-size: 20px;
		font-family: "Alagard";
	}
</style>
