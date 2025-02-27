<script lang="ts">
    import { types } from "$gotypes";
    import SearchCard from "$lib/components/SearchCard.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";
    import { onMount } from "svelte";

    interface Props {
        focused: boolean;
        options: types.WordDetails[];
        select?: (selected: types.WordDetails) => void;
	}

    let { 
        focused = false, 
        options,
        select,
    }: Props = $props();

    let container: HTMLDivElement | undefined = $state();

    let colors = GetColorStore();
    let selected_index = $state(0);

    $effect(() => {if (!focused) selected_index = 0; })

    onMount(() => {
		const down = (event: KeyboardEvent) => {
            if (!focused) return;

			if (event.key == "Enter" && options.length > 0 && select) {
				event.preventDefault();
                select(options[selected_index]);
            } else if (event.key == "l" && container) 
            {
				event.preventDefault();
				if (selected_index+1 < options.length)
                    selected_index++;
                let element: HTMLDivElement = container.children[selected_index] as HTMLDivElement;

                element.scrollIntoView({block: "start", behavior: "smooth"});
			} else if (event.key == "h" && container) 
            {
				event.preventDefault();
				if (selected_index-1 >= 0)
                    selected_index--;
                let element: HTMLDivElement = container.children[selected_index] as HTMLDivElement;

                element.scrollIntoView({block: "start", behavior: "smooth"});
			}
		}

		window.addEventListener("keydown", down);

		return ()=>{
		  window.removeEventListener("keydown", down);
		}
	});	
</script>

<div id="cards" tabindex="-1" bind:this={container}>
    {#each options as details, index}
        <SearchCard word={details.fr_word} definition={details.to_word} type={"verb"} color={index == selected_index ? $colors.primary1 : $colors.overlay3}/>
    {/each}
</div>

<style>
	#cards {
		flex: 1;
		display: flex;
		align-items:center;
		flex-direction:row;
		gap: 3vw;
		width: 100vw;
		overflow-x:scroll;
		scrollbar-width: none;
		padding-left: 3vw;
		padding-right: 3vw;
		scroll-padding-right: 4vw;
		scroll-padding-left: 3vw;
	}
</style>