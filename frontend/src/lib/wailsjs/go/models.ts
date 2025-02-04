export namespace main {
	
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

