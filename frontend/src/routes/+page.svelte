<script lang="ts">
    import SearchBar from '$lib/components/SearchBar.svelte';
    import SearchCard from '$lib/components/SearchCard.svelte';

	import { SearchWord, FetchDetails } from '$go';
	import type { types } from '$gotypes';

	let search_term: string = $state('');
	let search_results: types.SearchResult[] = $state([]);
	let current_term: types.SearchResult | undefined = $state();
	let current_details: types.WordDetails | undefined = $state();

	$effect(() => {
		console.log(search_term);
		SearchWord("en", "it", search_term).then(result => {
			search_results = result;
		});

		
	});

	$effect(() => {
		if (search_results.length > 0) {
			FetchDetails("en", "it", search_results[0].word).then(result => {
				current_details = result[0];
				//search_results = [];
				//search_term = '';
			});
		}
	});
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div>
		<SearchCard word={search_results[0] ? search_results[0].word : ""} type="verb" definition={current_details ? current_details.to_word : ""}/>
	</div>
	<SearchBar symbol="EN" bind:value={search_term}/>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		align-items: center;
		flex: 1;
	}

	div {
		display: flex;
		flex-direction: column;
		justify-content: space-evenly;
		flex: 1;
	}

</style>
