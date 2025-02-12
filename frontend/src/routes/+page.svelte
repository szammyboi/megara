<script lang="ts">
    import SearchBar from '$lib/components/SearchBar.svelte';
    import SearchCard from '$lib/components/SearchCard.svelte';

</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div id="searchcontainer">
		<SearchCard word="ISPIRARE, STIMOLARE, INFLUENZARE" definition="inspire somebody, motivate somebody" type="verb"/>
		<SearchCard word="MOTIVARE" definition="inspire somebody, motivate somebody" type="verb"/>
	</div>
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

	#searchcontainer {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: left;
		column-gap: 3vw;
		width: fit-content;
		max-width: 100%;
		margin-left: 3vw;
	}

</style>
