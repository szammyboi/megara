import { GetActiveColor } from '$go';
import { CreateColorSchemeStore } from '$lib/stores/colors.svelte';
// @ts-ignore
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async () => {
	let color_scheme = await GetActiveColor();

	return {
		color_scheme: color_scheme 
	};
};

export const prerender = true;
export const ssr = false;

