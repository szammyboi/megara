<script lang="ts">
    import SearchBar from '$lib/components/SearchBar.svelte';
    import SearchCard from '$lib/components/SearchCard.svelte';

	import { SearchWord, FetchDetails } from '$go';
	import type { types } from '$gotypes';
    import Logo from '$lib/components/Logo.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<SearchCard word="test" definition="test" type="verb"/>
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
