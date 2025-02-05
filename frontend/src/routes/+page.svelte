<script lang="ts">
    import SearchBar from '$lib/components/SearchBar.svelte';
    import SearchCard from '$lib/components/SearchCard.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div>
		<SearchCard word="ispirare, stimolare, influenzare" type="verb" definition="inspire somebody, motivate somebody"/>
	</div>
	<SearchBar symbol="EN"/>
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
