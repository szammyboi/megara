<script lang="ts">
	import type { LayoutProps } from './$types';

	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
    import { CreateColorSchemeStore, SetColorStore } from '$lib/stores/colors.svelte';
    import { goto } from '$app/navigation';
	
	let { data, children }: LayoutProps = $props();

	let color_store = CreateColorSchemeStore(data.color_scheme);

	SetColorStore(color_store);

	let color_css = $derived(`
		--base: ${$color_store.base};
		--overlay1: ${$color_store.overlay1};
		--overlay2: ${$color_store.overlay2};
		--overlay3: ${$color_store.overlay3};
		--text: ${$color_store.text};
		--primary1: ${$color_store.primary1};
		--primary2: ${$color_store.primary2};
		--primary3: ${$color_store.primary3};
		--primary4: ${$color_store.primary4};
	`);

	window.addEventListener('keypress', (e) => {
		if (e.key == '1') {
			e.preventDefault()
			goto("/search2");
		}
		if (e.key == '2') {
			e.preventDefault()
			goto("/flashcard");
		}
	})
</script>

<div class="app" style:background={$color_store.base} style={color_css}>
	{@render children()}
</div>

<style>
	.app {
		position: absolute;
		top: 0;
		left: 0;
		min-width: 100vw;
		height: 100vh;
	}


</style>
