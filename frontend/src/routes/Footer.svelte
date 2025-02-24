<script lang="ts">
	import { page } from '$app/state';
    import Logo from '$lib/components/Logo.svelte';
    import SearchBar from '$lib/components/SearchBar.svelte';
    import { GetColorStore } from '$lib/stores/colors.svelte';

	let color_scheme = GetColorStore();
</script>
<!--
<footer>
	<Logo />
	<SearchBar symbol="?" color={$color_scheme.primary1}/>
</footer>
-->
<style>
	footer {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-content:center;
		align-items: center;
		width: 94vw;
		margin-left: 3vw;
		margin-bottom: 3vw;
		color: #755c3b;
		--background-color: blue;
	}

	footer div {
		cursor: pointer; 
	}

	#star {
		display: flex;
		justify-content: left;
		width: 85vw;
		margin-left: 7.5vw;
		color: #755c3b;
		font-family: 'Daydream';
		font-size: 25px;
		margin-bottom: -1.25vw;
	}

	#active1 {
		color: #da1d4c;
	}

	#active2 {
		color: #ff9500;
	}

	#active3 {
		color: #488b99;
	}

	#active4 {
		color: #ff5813;
	}
</style>


