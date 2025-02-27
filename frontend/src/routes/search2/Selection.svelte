<script lang="ts">
    import MiniCard from "$lib/components/MiniCard.svelte";
    import { GetColorStore } from "$lib/stores/colors.svelte";
    import { onMount } from "svelte";
    import { types } from "$gotypes";

    interface Props {
        focused: boolean;
		width: number;
        options: types.SearchResult[];
        select: (selected: types.SearchResult) => void;
	}

    let { 
        focused = true, 
        width, 
        options,
        select,
    }: Props = $props();

    let colors = GetColorStore();
    let selected_index = $state(0);

    $effect(() => {
		if (!focused) selected_index = 0; 
	});

    onMount(() => {
		const down = (event: KeyboardEvent) => {
            if (!focused) return;

			if (event.key == "Enter" && options.length > 0) {
				event.preventDefault();
                select(options[selected_index]);
            } else if (event.key == "k") 
            {
				event.preventDefault();
				if (selected_index+1 < options.length)
                    selected_index++;
			} else if (event.key == "j") 
            {
				event.preventDefault();
				if (selected_index-1 >= 0)
                    selected_index--;
			}
		}

		window.addEventListener("keydown", down);

		return ()=>{
		  window.removeEventListener("keydown", down);
		}
	});	
</script>

<left style="width: {width}px;"></left>
<right>
    {#each options as option, index}
        <MiniCard color={(index == selected_index && focused) || options.length == 1 ? $colors.primary1 : $colors.overlay3}>
            <h1>{option.word.toUpperCase()}</h1>
            <h1>{option.language}</h1>
        </MiniCard>
    {/each}
</right>


<style>
	right {
		display: flex;
		flex-direction:column-reverse;
		gap: 1.5vw;
	}
</style>
