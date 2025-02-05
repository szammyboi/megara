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
		background: #362712;
		width: 450px;
		height: 280px;
		border-radius: 20px;
		border: solid 6px #f18d00;
	}
</style>
