<script lang="ts">
	import Card from '$lib/components/Card.svelte';

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
</script>

<Card>
	<div id="align">
		<div>
			<h1 id="title">{word.toLowerCase()}</h1>	
			<div id="type">{type.toLowerCase()}</div>	
			<h1 class="text" id="def">{definition.toLowerCase()}</h1>	
		</div>
		<div class="text">
			ctrl-a to add
		</div>
	</div>
</Card>

<style>
	#align {
		display: flex;
		flex-direction:column;
		align-items: left;
		justify-content: space-between;
		width: 100%;
		height: 100%;
		--font-family: 'Alagard';
		padding: 20px;
	}

	#title {
		font-size: 45px;
		line-height: 1;
		color: #f18d00;
		padding-bottom: 10px;
	}

	#type {
		background: #f18d00;
		width: fit-content;
		font-size: 19px;
		padding-left: 5px;
		padding-right: 5px;
		line-height:1;
		margin-bottom: 5px;
	}

	#def {
		color: #d9ae80;
	}

	.text {
		font-size: 25px;
		color: #755C3B;
	}
</style>
