package main

import (
	"fmt"
	"math/rand"
	"time"

	randomUserAgent "github.com/corpix/uarand"
	"github.com/valyala/fasthttp"
)

var (
	client fasthttp.Client

	errors int = 0
	sent   int = 0
	rpm    int = 0
)

func addShare(itemID string) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("POST")
	req.SetRequestURI(generateURL())
	req.SetBody([]byte(fmt.Sprintf("item_id=%v&share_delta=1", itemID)))

	req.Header.Set("User-Agent", randomUserAgent.GetRandom())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err := client.Do(req, res); err != nil {
		errors++
		return
	}

	sent++
}

func rpmCounter() {
	for {
		before := sent
		time.Sleep(6000 * time.Millisecond)
		after := sent

		rpm = (after - before) * 10
	}
}

func statusPrinter() {
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("[+] Sent: %v | Requests per minute: %v | Errors: %v\r", sent, rpm, errors)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Print("[+] github.com/monkeyskid\n\n")

	var threads int
	var itemID string

	fmt.Print("[>] Threads: ")
	fmt.Scanln(&threads)

	fmt.Print("[>] Video ID: ")
	fmt.Scanln(&itemID)

	fmt.Print("\n\n")

	go rpmCounter()
	go statusPrinter()

	for i := 0; i < threads; i++ {
		go func() {
			for {
				addShare(itemID)
			}
		}()
	}

	select {}
}
