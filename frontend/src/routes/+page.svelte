<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
    import Footer from "$lib/components/Footer.svelte";
    import { goto } from "$app/navigation";
	let colors = GetColorStore();

	goto("/flashcard");
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>



<style>
	
</style>
