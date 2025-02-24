<script lang="ts">
    import Logo from '$lib/components/Logo.svelte';
    import SearchBar from '$lib/components/SearchBar.svelte';
    import { GetColorStore } from '$lib/stores/colors.svelte';

	let color_scheme = GetColorStore();

	interface Props {
		search?: string;
	}

	let {
		search = $bindable(),
	}: Props = $props();
</script>

<footer>
	<Logo />
	<SearchBar symbol="?" color={$color_scheme.primary1} bind:value={search}/>
</footer>

<style>
	footer {
		display: flex;
		flex-direction: row;
		align-content:center;
		align-items: center;
		gap: 9vw;
		padding-right: calc(9vw - 3vw);
	}
</style>


