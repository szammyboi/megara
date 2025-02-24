import { types } from '$gotypes';
import { EventsOff, EventsOn } from '$runtime';
import { getContext, setContext } from 'svelte';
import { readable, type Readable } from 'svelte/store';

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
