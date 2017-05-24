package grgutil

import (
    "strings"

    "gopkg.in/gin-gonic/gin.v1"
    "github.com/tsaikd/go-raml-parser/parser"

    r "nick.cc/go-raml-to-gin/routes"
)

func bindRouter(router gin.IRouter, httpMethod string, path string, resourceHandlerName string) {
    r.SetRouter(router, httpMethod, path, resourceHandlerName)
}

func bindRootDocument(router gin.IRouter, rootdoc parser.RootDocument) {
    for ramlPath, resource := range rootdoc.Resources {

        ginPath := ToGinResource(ramlPath)
        resourceHandlerName := resource.DisplayName

        for name, _ := range resource.Methods {
            method := strings.ToUpper(name)
            bindRouter(router, method, ginPath, resourceHandlerName)
        }
    }
}

func SetUp(rootdoc parser.RootDocument) *gin.Engine {
    router := gin.Default()
    router.Use(gin.ErrorLogger())
    bindRootDocument(router, rootdoc)
    return router
}
