<script lang="ts">
	import type { LayoutProps } from './$types';

	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
    import { CreateColorSchemeStore, SetColorStore } from '$lib/stores/colors.svelte';
	
	let { data, children }: LayoutProps = $props();

	let color_store = CreateColorSchemeStore(data.color_scheme);

	SetColorStore(color_store);
</script>

<div class="app">
	<main>
		{@render children()}
	</main>

	<Footer />
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
		background: #221a02;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}
</style>
