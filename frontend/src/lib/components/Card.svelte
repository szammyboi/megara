<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';
	let { children, ...rest }: SvelteHTMLElements['div'] = $props();
</script>

<div {...rest} id="wrapper" >
	<div id="container">
		{@render children?.()}
	</div>
</div>

<style>
	#wrapper {
		background: #ef540a;
		width: 55vw;
		height: 55vh;
		border-radius: 15px;
		user-select:none;
	}

	#container {
		width: calc(55vw - 10vh);
		height: 45vh;
		margin-left: 5vh;
		margin-top: 5vh;
	}
</style>
