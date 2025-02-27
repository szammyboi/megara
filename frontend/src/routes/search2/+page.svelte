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

	let selected_word = $state('');
	let search_results: types.SearchResult[] = $state([]);
	let detailed_results: types.WordDetails[] = $state([]);

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

	$effect(() => {
		detailed_results = [];
		FetchDetails("en", "it", selected_word).then((result) => {
			if (result.length > 0)
				detailed_results = result;
		});
	});

	onMount(() => {
		const typing = (event: KeyboardEvent) => {
			if (active_state != 'typing') return;

			if (event.key == "Enter" && search_results.length == 1)
			{
				event.stopImmediatePropagation();
				event.preventDefault();
				active_state = 'browsing';
				selected_word = search_results[0].word;
			}
			else if (event.key == "Enter")
			{
				event.stopImmediatePropagation();
				event.preventDefault();
				active_state = 'selecting';
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
				select={(e: types.SearchResult) => {
					active_state = 'browsing';
					selected_word = e.word;
				}} 
			/>
		{:else}
			<Browsing focused={active_state == 'browsing'} options={detailed_results} />
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
