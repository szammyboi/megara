<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements';
	let { ...rest }: SvelteHTMLElements['div'] = $props();
</script>

<div {...rest} id="logo" >
	<h1>MEGARA</h1>
	<h2>EN > ITA</h2>
</div>

<style>
	#logo {
		font-family: "Alagard";
	}

	h1 {
		font-size: 50px;
		line-height: 0.9;
		color: var(--overlay3);
	}

	h2 {
		font-size: 33px;
		line-height: 0.9;
		color: var(--overlay2);
	}
</style>

