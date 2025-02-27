<script lang="ts">
	interface Props {
		focused: boolean;
		color: string;
		symbol: string;
		value?: string;
	}

	let {
		focused,
		color,
		symbol,
		value = $bindable(),
	}: Props = $props();

	let input: HTMLInputElement | undefined = $state();

	$effect(() => {
		if (input && focused)
			input.focus();
	});

	const defocus = () => {
		if (input && focused)
			input.focus();
	}
</script>

<section style="--color: {color};"  >
	<input bind:value={value} autocomplete="off" spellcheck="false" disabled={!focused} bind:this={input} onblur={defocus}>
	<div id="symbolarea" >
		<div id="symbol" class={focused ? "" : "disabled"}>
			{focused ? symbol : "X"}
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
		color: var(--text);
	}

	input::selection {
		color: var(--base);
		background-color: var(--primary1);
	}

	input:disabled {
		color: var(--base);
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

	#symbol.disabled {
		background-color: var(--overlay1);
	}
</style>
