package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	store             map[string]*minHeap
	mutex             sync.Mutex
	getRequestsCount  uint64
	postRequestsCount uint64
)

func main() {
	store = make(map[string]*minHeap)
	router := routing.New()

	router.Get("/pq/<customer_id>/top", pqTop)
	router.Post("/pq/<customer_id>/push/<priority>", pqPush)
	router.Get("/info", getInfo)

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}

func getInfo(c *routing.Context) error {
	fmt.Printf("get requests count: %d\n", getRequestsCount)
	fmt.Printf("post requests count: %d\n", postRequestsCount)
	fmt.Printf("pq map size: %d\n", len(store))

	for key, pq := range store {
		_, top := pq.Top()

		fmt.Printf("pq %s top value is %s\n", key, top.(string))
	}

	return nil
}

func getPqByCustomerId(customerId string) *minHeap {
	mutex.Lock()
	defer mutex.Unlock()

	if _, found := store[customerId]; !found {
		store[customerId] = NewMinHeap()
	}

	return store[customerId]
}

func pqPush(c *routing.Context) error {
	customerId := c.Param("customer_id")
	data := string(c.PostBody())
	pq := getPqByCustomerId(customerId)
	priority, _ := strconv.Atoi(c.Param("priority"))

	pq.Push(priority, data)

	atomic.AddUint64(&postRequestsCount, 1)

	return nil
}

func pqTop(c *routing.Context) error {
	customerId := c.Param("customer_id")
	pq := getPqByCustomerId(customerId)

	if !pq.Empty() {
		_, topVal := pq.Top()
		fmt.Fprintf(c, topVal.(string))
	} else {
		fmt.Fprintf(c, "pq is empty")
	}

	atomic.AddUint64(&getRequestsCount, 1)

	return nil
}
