<script lang="ts">
	interface Props {
		word: string;
		type: string;
		definition: string;
		color: string;
		flipped: boolean;
	}

	let {
		word,
		type,
		definition,
		color,
		flipped,
		...rest
	}: Props = $props();

	
</script>

<div id="wrapper" style="--color: {color};" class:flipped={flipped} {...rest}>
	<div id="front">
		<div class="align">
			<h1 class="title">{definition.toUpperCase()}</h1>	
			<div class="type">
				{type.toLowerCase()}
			</div>	
		</div>
	</div>
	<div id="back">
		<div class="align">
			<h1 class="title">{word.toUpperCase()}</h1>	
		</div>
	</div>
</div>

<style>
	#wrapper {
		flex-shrink: 0;
		width: 100%;
		height: 85%;
		border-radius: 15px;
		user-select:none;
		background-color: var(--color);
		transform-style: preserve-3d;
		--transition: transform 0.8s cubic-bezier(0.175, 0.885, 0.32, 1.275);
		transition: transform 0.4s cubic-bezier(0.39,0.57,0.56,1);
	}

	.flipped {
		transform: rotateX(180deg);
	}

	.align {
		display: flex;
		flex-direction:column;
		align-items: left;
		justify-content: center;
		width: fit-content;
		color: var(--base);
	}

	.title {
		font-size: 42px;
		line-height: 0.95;
		padding-bottom: 10px;
		font-family: 'Alagard';
		overflow-wrap: break-word;
	}

	.type {
		font-family: 'Alagard';
		width: fit-content;
		font-size: 22px;
		padding-left: 5px;
		padding-right: 5px;
		line-height:1.1;
		color: var(--color);
		background-color: var(--base);
	}

	#front {
		position: absolute;
		width: 100%;
		padding: 5vh;
		height: 100%;
		display:flex;
		align-items:center;
		justify-content:center;

		backface-visibility: hidden;
	}

	#back {
		position: absolute;
		width: 100%;
		height: 100%;
		padding: 5vh;
		transform: rotateX(180deg);
		display:flex;
		align-items:center;
		justify-content:center;
		border-radius:15px;
		backface-visibility: hidden;
	}
</style>
