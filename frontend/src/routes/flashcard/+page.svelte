<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
	import { onMount } from 'svelte';
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	import SearchCard from '$lib/components/SearchCard.svelte';
    import Footer from "$lib/components/Footer.svelte";
    import Flashcard from "$lib/components/Flashcard.svelte";
	let colors = GetColorStore();

	let source_size = $state();

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

<section>
	<left id="logo">
		<Logo />
	</left>
	<right>
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={flipped} />
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={false} />
		<Flashcard word="alienare" definition="cede, transfer, alienate" type="verb" color={$colors.primary2} flipped={false} />
	</right>
</section>

<style>
	section {
		display: flex;
		flex-direction:row;
		height: 100vh;
		min-width: 100vw;
	}

	#logo {
		padding-right: 3vw;
	}

	left {
		margin-top: 3vw;
		margin-bottom: 3vw;
		height: auto;
		display: flex;
		flex-direction: column-reverse;
	}

	right {
		display: flex;
		flex-direction:column;
		align-items:center;
		padding-top: 3vw;
		padding-bottom: 3vw;
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
