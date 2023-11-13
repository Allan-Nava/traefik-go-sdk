package traefik

const (
	// 
	HTTP_ROUTERS     		= "/api/http/routers"		// --> /api/http/routers	Lists all the HTTP routers information.
	//HTTP_ROUTERS_DETAIL     = "/api/http/routers/"		// --> /api/http/routers/{name}	Returns the information of the HTTP router specified by name.
	HTTP_SERVICES 			= "/api/http/services" 		// --> /api/http/services	Lists all the HTTP services information.
	HTTP_MIDDLEWARES 		= "/api/http/middlewares" 	// --> /api/http/middlewares	Lists all the HTTP middlewares information.
	//HTTP_MIDDLEWARES_DETAIL = "/api/http/middlewares/"	// --> /api/http/middlewares/{name}	Returns the information of the HTTP middleware specified by name.
	TCP_ROUTERS 			= "/api/tcp/routers" 		// --> /api/tcp/routers	Lists all the TCP routers information.
	TCP_SERVICES 			= "/api/tcp/services" // --> /api/tcp/services	Lists all the TCP services information.
	//
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
)