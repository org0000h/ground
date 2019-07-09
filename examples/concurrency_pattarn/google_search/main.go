package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	info string
}

var (
	Web1   = fakeSearchGenerator("web1")
	Web2   = fakeSearchGenerator("web2")
	Image1 = fakeSearchGenerator("image1")
	Image2 = fakeSearchGenerator("image2")
	Video1 = fakeSearchGenerator("video1")
	Video2 = fakeSearchGenerator("video2")
)

type Search func(query string) Result

func fakeSearchGenerator(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan Result)
	go func() { c <- First("query", Web1, Web2) }()
	go func() { c <- First("query", Image1, Image2) }()
	go func() { c <- First("query", Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		results := make([]Result, 10)
		select {
		case result := <-c:
			results = append(results, result)
			fmt.Println(result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}
