package router

import (
    "time"

    "gopkg.in/gin-gonic/gin.v1"
)

type CreateAd struct {
    ad_platform string `json:"ads.ad_platform" binding:"required"`
    ad_type int64 `json:"ads.ad_type" binding:"required"`
    ad_bid_method string `json:"ads.ad_bid_method" binding:"required"`
    ad_bid_cost float64 `json:"ads.ad_bid_cost" binding:"required"`
    ad_bid_auto bool `json:"ads.ad_bid_auto,omitempty"`
}

type CreateAds struct {
    ads []CreateAd `json:"ads" binding:"required"`
    token string `json:"token" binding:"required"`
}

func (router *RouteHandler) PostCreateAds(c *gin.Context) {
    var postJson CreateAds
    current_time := time.Now().UTC()
    if c.BindJSON(&postJson) == nil {
        // m.logger.Debug(postJson)
        c.JSON(200, gin.H{
            "data": gin.H{
                "ads": []gin.H{{
                    "ad_id": 1,
                    "ad_created_date": current_time,
                }},
            },
        })
    }
}
