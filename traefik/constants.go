package traefik

const (
	// 
	HTTP_ROUTERS     		= "/api/http/routers"		// --> /api/http/routers	Lists all the HTTP routers information.
	//HTTP_ROUTERS_DETAIL     = "/api/http/routers/"		// --> /api/http/routers/{name}	Returns the information of the HTTP router specified by name.
	HTTP_SERVICES 			= "/api/http/services" 		// --> /api/http/services	Lists all the HTTP services information.
	HTTP_MIDDLEWARES 		= "/api/http/middlewares" 	// --> /api/http/middlewares	Lists all the HTTP middlewares information.
	//HTTP_MIDDLEWARES_DETAIL = "/api/http/middlewares/"	// --> /api/http/middlewares/{name}	Returns the information of the HTTP middleware specified by name.
	TCP_ROUTERS 			= "/api/tcp/routers" 		// --> /api/tcp/routers	Lists all the TCP routers information.
	TCP_SERVICES 			= "/api/tcp/services" 		// --> /api/tcp/services	Lists all the TCP services information.
	TCP_MIDDLEWARES 		= "/api/tcp/middlewares" 	// --> /api/tcp/middlewares	Lists all the TCP middlewares information.
	UDP_ROUTERS 			= "/api/udp/routers" 		// --> //api/udp/routers	Lists all the UDP routers information.
	UDP_SERVICES 			= "/api/udp/service" 		// --> /api/udp/services	Lists all the UDP services information.
	ENTRYPOINTS 			= "/api/entrypoints" 		// --> /api/entrypoints	Lists all the entry points information.
	OVERVIEW 				= "/api/overview"			// --> /api/overview Returns statistic information about http and tcp as well as enabled features and providers.
	//
	RAW_DATA 				= "/api/rawdata" 			// --> /api/rawdata	Returns information about dynamic configurations, errors, status and dependency relations.
	API_VERSION 			= "/api/version" 			// --> /api/version	Returns information about Traefik version.
)


var (
	//
	GET_HTTP_ROUTER = func(routerName string) string {
		return HTTP_ROUTERS + "/" + routerName
	}
	//
	GET_HTTP_MIDDLEWARE = func(middlewareName string) string {
		return HTTP_MIDDLEWARES + "/" + middlewareName
	}
	//
	GET_TCP_ROUTER = func(name string ) string {
		return TCP_ROUTERS + "/" + name
	}
	//
	GET_TCP_SERVICE = func(name string ) string {
		return TCP_SERVICES + "/" + name
	}
	//
	GET_TCP_MIDDLEWARE = func(name string ) string {
		return TCP_MIDDLEWARES + "/" + name
	}
	//
	GET_UDP_ROUTER = func(name string ) string {
		return UDP_ROUTERS + "/" + name
	}
	//
	GET_UDP_SERVICE = func(name string ) string {
		return UDP_SERVICES + "/" + name
	}
	//
	GET_ENTRYPOINT = func(name string ) string {
		return ENTRYPOINTS + "/" + name
	}
	//
)