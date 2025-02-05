<script lang="ts">
	import { page } from '$app/state';

	let chosen = $state(1);
</script>

<div id="star">
	<div id={chosen == 1 ? "active1" : ""}>
		*
	</div>

	<div id={chosen == 2 ? "active2" : ""}>
		*
	</div>

	<div id={chosen == 3 ? "active3" : ""}>
		*
	</div>

	<div id={chosen == 4 ? "active4" : ""}>
		*
	</div>
</div>

<footer>
	<div id={chosen == 1 ? "active1" : ""} on:click={() => chosen = 1}>
		SEARCH	
	</div>

	<div id={chosen == 2 ? "active2" : ""} on:click={() => chosen = 2}>
		LEARN
	</div>

	<div id={chosen == 3 ? "active3" : ""} on:click={() => chosen = 3}>
		CARDS
	</div>

	<div id={chosen == 4 ? "active4" : ""} on:click={() => chosen = 4}>
		REMEMBER
	</div>
</footer>

<style>
	footer {
		display: flex;
		justify-content: space-between;
		width: 85vw;
		margin-left: 7.5vw;
		margin-bottom: 2vw;
		color: #755c3b;
		font-family: 'Jersey';
		font-size: 45px;
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
		font-family: 'Jersey';
		font-size: 37px;
		margin-bottom: -3vw;
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


