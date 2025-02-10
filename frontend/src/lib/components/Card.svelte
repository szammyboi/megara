<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';
	let { children, ...rest }: SvelteHTMLElements['div'] = $props();
</script>

<div {...rest} >
	{@render children?.()}
</div>

<style>
	div {
		--background: #362712;
		background: #ef540a;
		width: 50vw;
		height: 55vh;
		border-radius: 10px;
	}
</style>
