<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
	let colors = GetColorStore();
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<div id="top">
		
</div>

<div id="bottom">
	<Logo />
	<SearchBar symbol="?" color={$colors.primary1}/>
</div>

<style>
	#bottom {
		float: left;
		width: 100%;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-content:center;
		align-items: center;
	}
	#top {
		flex: 1;
		float: left;
		display: flex;
		align-items:center;
		flex-direction:column-reverse;
		width: 100%;
	}
</style>
