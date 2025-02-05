<script lang="ts">
	import Card from '$lib/components/Card.svelte';
    import SearchCard from '$lib/components/SearchCard.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<SearchCard word="ispirare, stimolare, influenzare" type="verb" definition="inspire somebody, motivate somebody"/>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

</style>
