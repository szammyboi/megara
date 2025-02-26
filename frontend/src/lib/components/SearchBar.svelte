<script lang="ts">
	interface Props {
		color: string;
		text: string;
		symbol: string;
		value?: string;
	}

	let {
		color,
		text,
		symbol,
		value = $bindable(),
	}: Props = $props();
</script>

<!--component

should color be activecolor set by layout?
-->

<section style="--color: {color}; --textcolor: {text};">
	<input bind:value={value} autocomplete="off" spellcheck="false" autofocus={true}>
	<div id="symbolarea">
		<div id="symbol">
			{symbol}
		</div>
	</div>
</section>

<style>
	section {
		width: 100%;
		height: 60px;
		display: flex;
		flex-direction:row;
		background-color: var(--overlay1);
	}

	input, input:focus {
		all: unset;
		padding-left: 20px;
		outline: none;
		appearance: none;
		width: 100%;
		font-size: 30px;
		font-family: "Alagard";
		color: var(--textcolor);
	}

	input::selection {
		color: var(--base);
		background-color: var(--primary1);
	}

	#symbolarea {
		display: flex;
		align-items: center;
		justify-content: center;
		width: fit-content;
		height: 100%;
		margin-right: calc(calc(60px - 3.75vw) / 2);
	}

	#symbol {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 3.75vw;
		height: 3.75vw;
		font-size: 20px;
		font-family: "Alagard";

		color: var(--base);
		background-color: var(--color);
	}
</style>
