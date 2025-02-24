<script lang="ts">
    import Logo from "$lib/components/Logo.svelte";
    import SearchBar from "$lib/components/SearchBar.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";

	let colors = GetColorStore();
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<div id="left">
		<Logo />
	</div>
	<div id="right">
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction:row;
		margin-top: 3vw;
		margin-left: 3vw;
		height: calc(100vh - 6vw);
		width: 94vw;
		background-color:yellow;
	}
	#left {
		float: left;
		display: flex;
		flex-direction:column-reverse;
		align-items:start;
		width: fit-content;
		height: calc(100vh - 6vw);
		background-color:red;
	}

	#right {
		flex: 1;
		float: left;
		display: flex;
		align-items:center;
		flex-direction:column-reverse;
		height: calc(100vh - 6vw);
		margin-left: 6vw;
		background-color:green;
	}
<
