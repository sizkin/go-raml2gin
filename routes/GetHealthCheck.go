package router

import (

    "gopkg.in/gin-gonic/gin.v1"
)

func (router *RouteHandler) GetHealthCheck(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "ok",
    })
}
