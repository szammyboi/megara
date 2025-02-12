<script lang="ts">
	import Card from '$lib/components/Card.svelte';
    import { GetColorStore } from '$lib/stores/colors.svelte';

	//import { color_schemes } from '$lib/stores/colors.svelte';

	interface Props {
		word: string;
		type: string;
		definition: string;
	}

	let {
		word,
		type,
		definition
	}: Props = $props();

	let color_scheme = GetColorStore();
</script>

<Card color={$color_scheme.color1}>
	<div id="align">
		<div>
			<h1 id="title">{word.toUpperCase()}:</h1>	
			<div id="type">{type.toLowerCase()}</div>	
			<h1 id="def">{definition.toLowerCase()}</h1>	
		</div>
	</div>
</Card>

<style>
	#align {
		display: flex;
		flex-direction:column;
		align-items: left;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	#title {
		font-size: 42px;
		line-height: 0.95;
		color: #221a02;
		padding-bottom: 10px;
		font-family: 'Alagard';
		overflow-wrap: break-word;
	}

	#type {
		font-family: 'Alagard';
		color: #ef540a;
		background: #221a02;
		width: fit-content;
		font-size: 22px;
		padding-left: 5px;
		padding-right: 5px;
		line-height:1.1;
		margin-bottom: 10px;
	}

	#def {
		color: #d9ae80;
		font-size: 30px;
		color: #221a02;
		line-height: 1;
		font-family:"Alagard";
	}
</style>
