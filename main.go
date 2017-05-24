package main

import (
    _ "gopkg.in/gin-gonic/gin.v1"
    "github.com/tsaikd/go-raml-parser/parser"
    _ "github.com/tsaikd/go-raml-parser/parser/parserConfig"

    u "nick.cc/go-raml-to-gin/util"
)

var logger Logger

type RAMLConfig struct{
   RAMLFile string
}

func main() {
    Config := RAMLConfig{"./routes/REST.raml"}
    //fmt.Printf("Config: %u\r", Config)

    ramlParser := parser.NewParser()

    rootdoc, err := ramlParser.ParseFile(Config.RAMLFile)
    if err != nil {
        logger.Error(err)
        return
    }

    // route := engineFromRootDocument(router, rootdoc)
    route := u.SetUp(rootdoc)
    // logger.Debug(route)
    route.Run(":8488")

}
