package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Worker interface {
	Task()
	isStopWork() bool
}

type JobInfo struct {
	delay       int
	name        string
	isTerminate bool
}

func (j JobInfo) Task() {
	fmt.Println("Doing work for ", j.name)
	sleep(j.delay + 2)
	fmt.Println("Complete work for ", j.name)
}

func (j JobInfo) isStopWork() bool {
	return j.isTerminate
}

func main() {
	fmt.Println("Begin")
	workChannel := make(chan Worker, 10)
	var waitGroupDone sync.WaitGroup
	for i := 0; i < 3; i++ {
		waitGroupDone.Add(1)
		go doWork(i, workChannel, &waitGroupDone)
	}

	delay := 1
	for j := 0; j < 20; j++ {
		jStr := strconv.Itoa(j)
		v := JobInfo{delay, "name " + jStr, false}
		workChannel <- v
		delay += 1
		if delay > 4 {
			delay = 0
		}
	}

	for j := 20; j < 23; j++ {
		jStr := strconv.Itoa(j)
		v := JobInfo{0, "end name " + jStr, true}
		workChannel <- v
	}

	fmt.Println("Waiting for all jobs to complete")
	waitGroupDone.Wait()

	fmt.Println("End")

}

func doWork(counter int, workQueue <-chan Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("worker: ", counter, "reading queue")
		work := <-workQueue
		if work.isStopWork() {
			break
		}
		work.Task()
	}
}

func sleep(secs int) {
	for {
		if secs <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			secs--
		}
	}
}
