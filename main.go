package main

import (
	. "project/parameters"

	"github.com/gin-gonic/gin"
)

var (
	clock Clock
	queue []Process
)

func main() {
	// defining the routers and run the
	router := gin.Default()
	router.POST("/addProcess", addProcess)
	router.GET("/setClockNow", setClockNow)
	router.GET("/setClockMax", setClockMax)
	router.GET("/update", update)
	router.GET("/getView", getView)
	router.Run()
}

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
