package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// "os"
	// "os/signal"
	// "syscall"
)

func main() {

	//sendPostRequest("https://apothemxdcpayrpc.blocksscan.io/",["hello"],)
	start := time.Now()

	readMessages(1)
	// go readMessages(2)
	// go readMessages(3)
	elapsed := time.Since(start)
	log.Printf("time for hits %s", elapsed)

}
func readMessages(n int) {
	//timer := time.NewTicker(15 * time.Second)
	//defer timer.Stop()
	//time.Sleep(1 * time.Second)
	// Poll before waiting for ticker
	//rpc.poll()
	fmt.Println("read messages called  for: ", n)
	//_, err := sendPostRequest("http://18.118.94.131:8989/", []byte(`123`))

	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	//i := 0
	for {
		//time.Sleep(1 * time.Second)
		resp, err := sendPostRequest("http://18.118.94.131:8989/", []byte(`1234`))
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("response:", resp)
		}
		//i++
	}
	// for {
	// 	select {
	// 	case <-timer.C:
	// 		//Added
	// 		time.Sleep(2 * time.Second)
	// 		_, err := sendPostRequest("http://18.118.94.131:8989/", []byte(``))
	// 		if err != nil {
	// 			fmt.Println("Error", err)
	// 		}
	// 	}
	// }
	fmt.Println("read messages ended for: ", n)

}

func sendPostRequest(url string, body []byte) ([]byte, error) {
	//time.Sleep(2 * time.Second)
	// fmt.Println("polling url:", url)
	// fmt.Println("body:", body)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//defer logger.ErrorIfCalling(r.Body.Close)

	if r.StatusCode < 200 || r.StatusCode >= 400 {
		return nil, errors.New("got unexpected status code")
	}

	return ioutil.ReadAll(r.Body)
}
