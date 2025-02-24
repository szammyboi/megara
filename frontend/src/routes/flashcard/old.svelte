<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
	import { onMount } from 'svelte';
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
    import Footer from "$lib/components/Footer.svelte";
    import Flashcard from "$lib/components/Flashcard.svelte";
	let colors = GetColorStore();

	let flipped = $state(false);

	onMount(() => {
		const spacebar = (event: KeyboardEvent) => {
			if (event.key == " ") {
				flipped = !flipped;	
				event.preventDefault();
			}
		}
		window.addEventListener("keypress", spacebar);

		return ()=>{
		  // this function is called when the component is destroyed
		  window.removeEventListener("keypress", spacebar);
		}
  });	
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section >
	<div id="left">
		<Logo />
	</div>
	<div id="right" tabindex="-1" >
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={flipped} />
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={false} />
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={false} />
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction:row;
		--margin-top: 3vw;
		height: 100vh;
		min-width: 100vw;
	}

	#left {
		display: flex;
		flex-direction:column-reverse;
		overflow-x:scroll;
		scrollbar-width: none;
		margin-top: 3vw;
		margin-bottom: 3vw;
		margin-left: 3vw;
		margin-right: 9vw;
	}

	#right {
		flex: 1;
		display: flex;
		flex-direction:column;
		align-items:center;
		padding-top: 3vw;
		padding-bottom: 3vw;
		margin-right: 9vw;
		gap: 3vw;
		overflow-y:scroll;
		scrollbar-width: none;
		user-select: none;
        -moz-user-select: none;
        -khtml-user-select: none;
        -webkit-user-select: none;
        -o-user-select: none;
	}
</style>
