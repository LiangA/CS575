package main

import (
	. "project/parameters"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	clock Clock
	queue []Process
)

func main() {
	// defining the routers and run the
	router := gin.Default()
	// deal with the CORS problem. down below have another way to do it
	router.Use(cors.Default())
	router.POST("/addProcess", addProcess)
	router.GET("/setClockNow", setClockNow)
	router.GET("/setClockMax", setClockMax)
	router.GET("/update", update)
	router.GET("/getView", getView)
	router.GET("/reset", reset)
	router.Run()
}

// see https://stackoverflow.com/questions/29418478/go-gin-framework-cors
// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// test of struct
// type test struct {
// 	iii int
// }
// x, y, z := test{123}, test{456}, test{789}
// a := []test{}
// b := []test{}
// a = append(a, x, y, z)
// b = append(b, a[0])
// a = append(a[1:])
// b[0].iii += 100
// a = append(a, b[0])
// b = b[0:0]

// b = append(b, a[0])
// a = append(a[1:])
// b[0].iii += 100
// a = append(a, b[0])
// b = b[0:0]

// fmt.Println(&a)
// fmt.Println(&b)
