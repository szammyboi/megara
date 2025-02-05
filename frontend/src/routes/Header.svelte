<script lang="ts">
	import { page } from '$app/state';

	interface Props {
		native: string;
		learning: string;
	}

	let {
		native,
		learning,
	}: Props = $props();

</script>

<header>
	<div class="left">
		MEGARA	
	</div>

	<!--<nav>
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z" />
		</svg>
		<ul>
			<li aria-current={page.url.pathname === '/' ? 'page' : undefined}>
				<a href="/">Home</a>
			</li>
			<li aria-current={page.url.pathname === '/about' ? 'page' : undefined}>
				<a href="/about">About</a>
			</li>
			<li aria-current={page.url.pathname.startsWith('/sverdle') ? 'page' : undefined}>
				<a href="/sverdle">Sverdle</a>
			</li>
		</ul>
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
		</svg>
	</nav>-->

	<div class="right">
		<span>{native}</span>
		>
		<span>{learning}</span>
	</div>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
		width: 85vw;
		margin-left: 7.5vw;
		margin-top: 2vw;

		font-family: 'Jersey';
		font-size: 45px;
	}

	.left {
		color: #488b99;
	}
	.right {
		color: #ff5813;
	}
</style>
