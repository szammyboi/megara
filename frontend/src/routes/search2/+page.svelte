<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
    import MiniCard from "$lib/components/MiniCard.svelte";
	import { types } from "$gotypes";
	import { onMount } from "svelte";
    import { FetchDetails, SearchWord } from "$go";
	let colors = GetColorStore();

	let source_size = $state();

	let searching = $state(true);

	let search_term = $state('');
	let selected = $state(0);
	let selected_word = $state('');
	let search_results: types.SearchResult[] = $state([]);
	let detailed_results: types.WordDetails[] = $state([]);

	$effect(() => {
		searching = true;
		detailed_results = [];

		if (search_term == '') {
			search_results = [];
			return;
		}

		selected = 0;
		SearchWord("en", "it", search_term).then((result) => {
			if (result.length > 0)
				search_results = result.slice(0, 5);
		});
	});

	$effect(() => {
		FetchDetails("en", "it", selected_word).then((result) => {
			if (result.length > 0)
				detailed_results = result;
		});
	});

	onMount(() => {
		const down = (event: KeyboardEvent) => {
			if (event.key == "Enter" && search_results.length > 0) {
				event.preventDefault();
				searching = false;
				selected_word = search_results[selected].word;
			//	selected = 0;
				
			} else if (event.key == "ArrowUp") {
				event.preventDefault();
				if (!searching) searching = true;
				if (selected+1 < search_results.length)
					selected++;
			}
			else if (event.key == "ArrowDown") {
				event.preventDefault();
				if (!searching) searching = true;
				if (selected-1 >= 0)
					selected--;
			}
		}
		window.addEventListener("keydown", down);

		return ()=>{
		  // this function is called when the component is destroyed
		  window.removeEventListener("keydown", down);
		}
	});	
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<top>
		{#if searching}
			<left style="width: {source_size}px;"></left>
			<right>
				{#each search_results as search_res, index}
					<MiniCard color={index == selected ? $colors.primary1 : $colors.overlay3}>
						<h1>{search_res.word.toUpperCase()}</h1>
						<h1>{search_res.language}</h1>
					</MiniCard>
				{/each}
			</right>
		{:else}
			<div id="cards" tabindex="-1">
				{#each detailed_results as details, index}
					<SearchCard word={details.fr_word} definition={details.to_word} type={"verb"}/>	
				{/each}
			</div>
			
		{/if}
	</top>
	<footer>
		<div id="logo" bind:clientWidth={source_size}>
			<Logo />
		</div>
		<SearchBar symbol="?" color={$colors.primary1} bind:value={search_term}/>
	</footer>
</section>

<style>
	section {
		display: flex;
		flex-direction:column;
		height: 100vh;
		min-width: 100vw;
	}

	#logo {
		padding-right: 3vw;
	}

	#cards {
		flex: 1;
		display: flex;
		align-items:center;
		flex-direction:row;
		gap: 3vw;
		min-width: 100vw;
		overflow-x:scroll;
		scrollbar-width: none;
		padding-left: 3vw;
		padding-right: 3vw;
	}

	right {
		display: flex;
		flex-direction:column-reverse;
		gap: 1.5vw;
	}
</style>
