<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';
	//let { color, children, ...rest }: SvelteHTMLElements['div'] = $props();
	let { color, children, ...rest }: any = $props();
</script>

<div {...rest} id="wrapper" style:background={color}>
	<div id="container">
		{@render children?.()}
	</div>
</div>

<style>
	#wrapper {
		width: 100%;
		height: 10vh;
		border-radius: 15px;
		user-select:none;
		display: flex;
		flex-direction:column;
		align-items:center;
		justify-content:center;
		font-size: 30px;
		font-family: "Alagard";
		color: var(--base);
	}

	#container {
		width: calc(100% - 6vh);
		height: calc(100% - 3vh);
		display: flex;
		flex-direction:row;
		align-items: center;
		justify-content:space-between;
	}
</style>
