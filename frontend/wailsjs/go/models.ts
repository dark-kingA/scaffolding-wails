export namespace main {
	
	export class Response {
	    rs: number;
	    code: number;
	    address: string;
	    ip: string;
	    isDomain: number;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rs = source["rs"];
	        this.code = source["code"];
	        this.address = source["address"];
	        this.ip = source["ip"];
	        this.isDomain = source["isDomain"];
	    }
	}

}

