package main

import (
	"fmt"
	"math"
	"net/http"
	. "project/parameters"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addProcess(c *gin.Context) {
	pid := string(c.PostForm("pid"))
	priority, err := strconv.Atoi(c.PostForm("priority"))
	arrive, err1 := strconv.Atoi(c.PostForm("arrive"))
	burst, err2 := strconv.Atoi(c.PostForm("burst"))
	// remain, err3 := strconv.Atoi(c.PostForm("remain"))

	if err != nil {
		fmt.Println("priority conversion wrong")
	}
	if err1 != nil {
		fmt.Println("arrive time conversion wrong")
	}
	if err2 != nil {
		fmt.Println("burst time conversion wrong")
	}
	// if err3 != nil {
	// 	fmt.Println("remain time conversion wrong")
	// }

	queue = append(queue, Process{
		pid, priority, arrive, burst, burst, "ready",
	})

	c.JSON(http.StatusOK, gin.H{
		"msg": "add successfully",
	})
}

func setClockNow(c *gin.Context) {
	now, err := strconv.Atoi(c.DefaultQuery("now", "0"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong input clock format",
		})
		return
	}
	clock.Now = now
	msg := "clock now is " + strconv.Itoa(clock.Now)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func setClockMax(c *gin.Context) {
	for _, v := range queue {
		clock.Max += v.Burst
	}
	msg := "clock max is set to " + strconv.Itoa(clock.Max)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func update(c *gin.Context) {
	// if the clock is over the maximum, only send out message
	if clock.Now > clock.Max {
		c.JSON(http.StatusOK, gin.H{
			"msg": "all process is done",
		})
		return
	}

	// default method is shortest job first, and default is preemptive
	// method := "SJF"
	// preemptive := "true"
	method := c.DefaultQuery("method", "SJF")
	preemptive := c.DefaultQuery("preemptive", "true")

	// if the clock now is 0, means we need to start the first process
	if clock.Now == 0 {
		nextProcessIn(method)
	}
	// based on different preemptive status do the logic
	// basically preemptive mode will put all the process from running to ready
	// then pick a new one to run. if the running process is over, turn it to finished
	if preemptive == "true" {
		for i := 0; i < len(queue); i++ {
			if queue[i].State == "run" {
				queue[i].Remain--
				if queue[i].Remain == 0 {
					queue[i].State = "finished"
				} else {
					queue[i].State = "ready"
				}
			}
		}

		nextProcessIn(method)

		// if it is non-prermptive, we only set the ended process to finished
		// then pick a new process to run, or do nothing
	} else {
		for i := 0; i < len(queue); i++ {
			if queue[i].State == "run" {
				queue[i].Remain--
				if queue[i].Remain == 0 {
					queue[i].State = "finished"
					nextProcessIn(method)
				}
			}
		}
	}
	// increase the clock now
	clock.Now++

	c.JSON(http.StatusOK, gin.H{
		"methdod":    method,
		"preemptive": preemptive,
		"msg":        "clock update successfully",
	})
}

// this function picks the next process from ready to running
func nextProcessIn(method string) {
	switch method {
	case "SJF":
		key, min := -1, math.MaxInt32
		// if there are two or more processes have the same remain time and they are all the shortest,
		// the first iterated one will be run
		for i := 0; i < len(queue); i++ {
			if queue[i].Remain > 0 && queue[i].Remain < min {
				key = i
				min = queue[i].Remain
			}
		}
		if key == -1 {
			fmt.Println("no process remains burst")
			return
		}
		queue[key].State = "run"
	case "FCFS":
		key, min := -1, math.MaxInt32
		// if there are two or more processes have the same arriving time and they are all the earliest,
		// the first iterated one will be run
		for i := 0; i < len(queue); i++ {
			if queue[i].Remain > 0 && queue[i].Arrive < min {
				key = i
				min = queue[i].Arrive
			}
		}
		if key == -1 {
			fmt.Println("no process remains burst")
			return
		}
		queue[key].State = "run"
	case "Priority":
		key, max := -1, -1
		// if there are two or more processes have the same arriving time and they are all the earliest,
		// the first iterated one will be run
		for i := 0; i < len(queue); i++ {
			if queue[i].Remain > 0 && queue[i].Priority > max {
				key = i
				max = queue[i].Priority
			}
		}
		if key == -1 {
			fmt.Println("no process remains burst")
			return
		}
		queue[key].State = "run"
	}
}

func getView(c *gin.Context) {
	var msg []interface{}
	for _, v := range queue {
		msg = append(msg, v)
	}
	c.JSON(http.StatusOK, gin.H{
		"processes": msg,
		"clock now": clock.Now,
	})
}
