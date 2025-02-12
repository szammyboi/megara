import { types } from '$gotypes';
import { EventsOff, EventsOn } from '$runtime';
import { getContext, setContext } from 'svelte';
import { readable, type Readable } from 'svelte/store';

/*export const color_schemes: types.ColorScheme = $state({
	name: types.ColorDefaults.NAME,
	background: types.ColorDefaults.BACKGROUND,
	overlay: types.ColorDefaults.OVERLAY,
	text: types.ColorDefaults.TEXT,
	color1: types.ColorDefaults.COLOR1,
	color2: types.ColorDefaults.COLOR2,
	color3: types.ColorDefaults.COLOR3,
	color4: types.ColorDefaults.COLOR4,
});*/

export const CreateColorSchemeStore = (base: types.ColorScheme): Readable<types.ColorScheme> => {
	return readable(base, function(set) {
		EventsOn("updateActiveColor", (data) => {
			set(data);
		});

		return () => {
			EventsOff("updateActiveColor")
		}
	});
}

export const GetColorStore = (): Readable<types.ColorScheme> => {
	return getContext('active_color') as Readable<types.ColorScheme>;
}

export const SetColorStore = (store: Readable<types.ColorScheme>) => {
	setContext('active_color', store);
}

/*export const func CreateColorSchemeStore(arr types.ColorScheme[]): Readable<types.ColorScheme[]> {
	let test =  readable([], function(set) {
		EventsOn("updateColors", (data) => {
			set(data);
		});

		return () => {
			EventsOff("updateColors")
		}
	});
}*/
