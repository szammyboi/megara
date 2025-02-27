<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import { types } from "$gotypes";
	import { onMount } from "svelte";
    import { FetchDetails, SearchWord } from "$go";

    import Selection from "./Selection.svelte";
    import Browsing from "./Browsing.svelte";

	type states = 'typing' | 'selecting' | 'browsing';

	let active_state: states = $state('typing');

	let colors = GetColorStore();

	let source_size = $state(0);

	let search_term = $state('');

	let selected_index: number = $state(0);

	let search_results: types.SearchResult[] = $state([]);

	let preloaded: types.WordDetails[][] = $state(new Array(5));

	let browse_results = $derived(preloaded[selected_index]);

	$effect(() => {
		if (search_term == '') {
			search_results = [];
			return;
		}

		SearchWord("en", "it", search_term).then((result) => {
			if (result.length > 0)
				search_results = result.slice(0, 5);
		});
	});

	const preload = () => {
		preloaded = new Array(5);
		for (let i = 0; i < search_results.length; i++) {
			FetchDetails("en", "it", search_results[i].word).then((res) => {
				if (res.length > 0)
					preloaded[i] = res;
			});
		}
	};

	onMount(() => {
		const typing = (event: KeyboardEvent) => {
			if (active_state != 'typing') return;

			if (event.key == "Enter" && search_results.length == 1)
			{
				event.stopImmediatePropagation();
				event.preventDefault();
				active_state = 'browsing';
				preload();
				selected_index = 0;
			}
			else if (event.key == "Enter")
			{
				event.stopImmediatePropagation();
				event.preventDefault();
				active_state = 'selecting';
				preload();
			}
		};

		const selecting = (event: KeyboardEvent) => {
			if (active_state != 'selecting') return;

			if (event.key == "Backspace") {
				event.stopImmediatePropagation();
				event.preventDefault();
				active_state = 'typing';
			}
		};

		const browsing = (event: KeyboardEvent) => {
			if (active_state != 'browsing') return;

			if (event.shiftKey && event.key == "Backspace") {
				event.stopImmediatePropagation();
				event.preventDefault();
				search_term = '';
				active_state = 'typing';
			} else if (event.key == "Backspace") {
				event.preventDefault();
				active_state = 'selecting';
			}
		};

		window.addEventListener("keydown", typing);
		window.addEventListener("keydown", selecting);
		window.addEventListener("keydown", browsing);

		return ()=>{
		  	window.removeEventListener("keydown", typing);
			window.removeEventListener("keydown", selecting);
			window.removeEventListener("keydown", browsing);
		}
	});
</script>
<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<top>
		{#if active_state != 'browsing'}
			<Selection 
				focused={active_state == 'selecting'} 
				width={source_size} 
				options={search_results} 
				select={(res: types.SearchResult, i: number) => {
					active_state = 'browsing';
					selected_index = i;
				}} 
			/>
		{:else}
			<Browsing focused={active_state == 'browsing'} options={browse_results} />
		{/if}
	</top>
	<footer>
		<div id="logo" bind:clientWidth={source_size}>
			<Logo />
		</div>
		<SearchBar focused={active_state == 'typing'} symbol="?" color={$colors.primary1} bind:value={search_term} />
	</footer>
</section>

<style>
	section {
		display: flex;
		flex-direction:column;
		height: 100vh;
		width: 100vw;
	}

	#logo {
		padding-right: 3vw;
	}
</style>
