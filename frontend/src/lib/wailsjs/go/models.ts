export namespace main {
	
	export class ColorScheme {
	    name: string;
	    base: string;
	    overlay1: string;
	    overlay2: string;
	    overlay3: string;
	    text: string;
	    primary1: string;
	    primary2: string;
	    primary3: string;
	    primary4: string;
	
	    static createFrom(source: any = {}) {
	        return new ColorScheme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.base = source["base"];
	        this.overlay1 = source["overlay1"];
	        this.overlay2 = source["overlay2"];
	        this.overlay3 = source["overlay3"];
	        this.text = source["text"];
	        this.primary1 = source["primary1"];
	        this.primary2 = source["primary2"];
	        this.primary3 = source["primary3"];
	        this.primary4 = source["primary4"];
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

