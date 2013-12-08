## go-middleware

### Example Usage

    import "net/http"
    import "github.com/shelakel/go-middleware"
    var handler http.HandlerFunc = middleware.Compose(...)

See [example](https://github.com/shelakel/go-middleware/blob/master/example/example.go)

#### Installation

    go get -u github.com/shelakel/go-middleware 

#### Documentation
See [GoDoc on Github](http://godoc.org/github.com/shelakel/go-middleware)

### Copyright and Licensing
See LICENSE

### Middleware

*done* packages may be used  
*work in progress (WIP)* packages should *not* be used - only for inspiration  
*TODO* packages currently lack implementation  

 * context - based on [context](github.com/gorilla/context) - done
 * compression - WIP
 * cors - WIP
 * panic/recover error handling - WIP
 * request logging - WIP
 * static file server - TODO
 * CSRF - [nosurf](https://github.com/justinas/nosurf) - TODO
 * Request throttling - TODO
 * Request/Response buffering - TODO
 * Auth - TODO
 * Sessions - TODO
 * bind - [binding](https://github.com/codegangsta/martini-contrib/tree/master/binding) / [schema](http://www.gorillatoolkit.org/pkg/schema) - TODO

__Middleware/Web Frameworks (for inspiration)__

*Links* indicate interesting implementations  
*\** can be converted or may be useful in middleware

 * [mango](https://github.com/paulbellamy/mango) - context, sessions, error handling, [static file server](https://github.com/paulbellamy/mango/blob/master/static.go), basic auth
 * [martini](https://github.com/codegangsta/martini) - context, request logging, error handling, static file server, buferred response writer
 * [martini-contrib](https://github.com/codegangsta/martini-contrib) - "Accept-Language" header parsing, basic auth, [binding](https://github.com/codegangsta/martini-contrib/tree/master/binding), compression, render (view path provider), sessions, strip, web.go context adapter
 * [beego](https://github.com/astaxie/beego) - context, [error handling](https://github.com/astaxie/beego/blob/master/middleware/error.go), [cache*](https://github.com/astaxie/beego/tree/master/cache), [validation*](https://github.com/astaxie/beego/tree/master/validation)
 * [traffic](https://github.com/pilu/traffic) - [error handling](https://github.com/pilu/traffic/blob/master/show_errors_middleware.go), [request logging](https://github.com/pilu/traffic/blob/master/logger_middleware.go), static file server
 * [handy](https://github.com/go-web-framework/handy) - context
 * [nosurf](https://github.com/justinas/nosurf) - CSRF middleware - can be used but not compatible for chaining as is
 * [revel](https://github.com/robfig/revel) - [data binding](https://github.com/robfig/revel/blob/master/binder.go), [filters](https://github.com/robfig/revel/blob/master/filter.go), [flash cookies](https://github.com/robfig/revel/blob/master/flash.go), [validation*](https://github.com/robfig/revel/blob/master/validation.go) - See [filter](https://github.com/robfig/revel/blob/master/filter.go)
 * [bones](https://github.com/peterskeide/bones) - context, sessions - See [bones.web](https://github.com/peterskeide/bones/tree/master/web)
 * [gorilla](https://github.com/gorilla/) - [context](https://github.com/gorilla/context), [mux](https://github.com/gorilla/mux), [schema*](https://github.com/gorilla/schema), [securecookie*](https://github.com/gorilla/securecookie), [sessions](https://github.com/gorilla/sessions), [handlers*](https://github.com/gorilla/handlers), [websocket*](https://github.com/gorilla/websocket), [feeds*](https://github.com/gorilla/feeds) - See [gorilla](https://github.com/gorilla/)

___Web Frameworks___

* [mango](https://github.com/paulbellamy/mango)
* [web.go](https://github.com/hoisie/web)
* [martini](https://github.com/codegangsta/martini)
* [beego](https://github.com/astaxie/beego)
* [traffic](https://github.com/pilu/traffic)
* [handy](https://github.com/go-web-framework/handy)
* [revel](https://github.com/robfig/revel)
* [bones](https://github.com/peterskeide/bones)

### Author/s

* [Shelakel](https://github.com/shelakel)
* Attribution for inspiration as per **Middleware/Web Frameworks**