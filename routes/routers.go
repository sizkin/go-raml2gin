package router

import (
    "reflect"

    "gopkg.in/gin-gonic/gin.v1"
)

type RouteHandlers struct {
    Id int
    Handlers []*RouteHandler
}

type RouteHandler struct {
    Name string
    // Method Handler
}

// type Handler interface {
//     HandledFunc(c *gin.Context)
// }

var Routes = &RouteHandlers{1, []*RouteHandler{}}

// func SetRouter(h Handler) *RouteHandlers {
//     handlerValue := reflect.ValueOf(h)
//     handlerName := handlerValue.FieldByName("Name").String()
//     Routes.Handlers = append(Routes.Handlers, RouteHandler{handlerName, h})
//
//     //fmt.Println(handlerType)
//     //h.handledFunc(456)
//     return Routes
// }

// TODO:
func execMethodByName(rh *RouteHandler, name string) (func(c *gin.Context)) {
    var out func(c *gin.Context)
    rhReflect := reflect.ValueOf(rh)
    rhReflectFunc := rhReflect.MethodByName(name)
    rhInterface := rhReflectFunc.Interface()
    out = rhInterface.(func(c *gin.Context))
    return out
}

func SetRouter(router gin.IRouter, httpMethod string, path string, resourceHandlerName string) {
    route_ := &RouteHandler{resourceHandlerName}
    routeHandler_ := execMethodByName(route_, resourceHandlerName)
    router.Handle(httpMethod, path, routeHandler_)
    Routes.Handlers = append(Routes.Handlers, route_)
}
