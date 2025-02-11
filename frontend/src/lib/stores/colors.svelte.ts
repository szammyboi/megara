import { types } from '$gotypes';
import { EventsOff, EventsOn } from '$runtime';
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


export const color_schemes: Readable<types.ColorScheme[]> = readable([{
	name: types.ColorDefaults.NAME,
	background: types.ColorDefaults.BACKGROUND,
	overlay: types.ColorDefaults.OVERLAY,
	text: types.ColorDefaults.TEXT,
	color1: types.ColorDefaults.COLOR1,
	color2: types.ColorDefaults.COLOR2,
	color3: types.ColorDefaults.COLOR3,
	color4: types.ColorDefaults.COLOR4,
}], function(set) {
	EventsOn("updateColors", (data) => {
		set(data);
	});

	return () => {
		EventsOff("updateColors")
	}
});
