package main

import (
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // https://github.com/gin-gonic/gin
    r := gin.Default()
    r.Static("/assets", "./assets")
    //r.LoadHTMLGlob("templates/**/*")
    r.LoadHTMLGlob("templates/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Map Service",
		})
    })
    //r.GET("/api", func(c *gin.Context) {
    //    ret := apiGetFstdata()
    //    c.JSON(200, ret)
    //})
    port := os.Getenv("PORT")
    if port == "" {
        r.Run()
    } else {
        r.Run(":" + port)
    }
}
