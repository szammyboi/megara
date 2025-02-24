<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
    import Footer from "$lib/components/Footer.svelte";
    import { goto } from "$app/navigation";
	let colors = GetColorStore();

	window.addEventListener('focus', (e) => {
		console.log(e);
	})

	let searching = $state(true);
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />

</svelte:head>

<section>
{#if !searching}
<div id="top" tabindex="-1">
	<!--<SearchCard word="ISPIRARE, STIMOLARE, INFLUENZARE" definition="inspire somebody, motivate somebody" type="verb"/>
	<SearchCard word="MOTIVARE" definition="inspire somebody, motivate somebody" type="verb"/>
	<SearchCard word="MOTIVARE" definition="inspire somebody, motivate somebody" type="verb"/>
	<SearchCard word="MOTIVARE" definition="inspire somebody, motivate somebody" type="verb"/>-->
</div>
{:else}
<div id="search_opts" tabindex="-1">
	
</div>
{/if}

<Footer />
</section>
<style>
	section {
		display: flex;
		flex-direction:column;
		margin-top: 3vw;
		height: calc(100vh - 6vw);
		width: 94vw;
	}
	#top {
		flex: 1;
		display: flex;
		align-items:center;
		flex-direction:row;
		gap: 3vw;
		min-width: 100vw;
		margin-bottom: 3vw;
		overflow-x:scroll;
		scrollbar-width: none;
		padding-left: 3vw;
		padding-right: 3vw;
		background-color:red;
	}
	#search_opts {
		flex: 1;
		background-color:blue;
		margin-right: calc(9vw - 3vw);
		margin-left: 3vw;
	}
</style>
