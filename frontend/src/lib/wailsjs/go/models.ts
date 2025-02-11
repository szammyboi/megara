export namespace main {
	
	export enum ColorDefaults {
	    NAME = "Megara",
	    BACKGROUND = "#201a01",
	    OVERLAY = "#342711",
	    TEXT = "#d9ae80",
	    COLOR1 = "#ef540a",
	    COLOR2 = "#cc094d",
	    COLOR3 = "#578b99",
	    COLOR4 = "#e58c00",
	}
	export class ColorScheme {
	    name: string;
	    background: string;
	    overlay: string;
	    text: string;
	    color1: string;
	    color2: string;
	    color3: string;
	    color4: string;
	
	    static createFrom(source: any = {}) {
	        return new ColorScheme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.background = source["background"];
	        this.overlay = source["overlay"];
	        this.text = source["text"];
	        this.color1 = source["color1"];
	        this.color2 = source["color2"];
	        this.color3 = source["color3"];
	        this.color4 = source["color4"];
	    }
	}
	export class SearchResult {
	    language: string;
	    word: string;
	    strength: number;
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.language = source["language"];
	        this.word = source["word"];
	        this.strength = source["strength"];
	    }
	}
	export class WordDetails {
	    fr_word: string;
	    to_word: string;
	
	    static createFrom(source: any = {}) {
	        return new WordDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fr_word = source["fr_word"];
	        this.to_word = source["to_word"];
	    }
	}

}

