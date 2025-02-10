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
		--background-color:red;
	}

	h1 {
		color: #816a45;
		font-size: 50px;
		line-height: 0.9;
	}

	h2 {
		color: #5c4c2b;
		font-size: 33px;
		line-height: 0.9;
	}
</style>

