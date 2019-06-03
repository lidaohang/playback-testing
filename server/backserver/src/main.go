package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	LOGPATH string = "backserver.log"
)

func log_info(info string) {
	file, err := os.OpenFile(LOGPATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)

	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(info)
}

func backHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	bytes_sent := r.Header.Get("bytes-sent")
	upstream_status := r.Header.Get("upstream-status")
	upstream_response_time := r.Header.Get("upstream-response-time")

	// Read body
	_, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp_time := 1
	if upstream_response_time != "" {
		ms, _ := strconv.ParseFloat(upstream_response_time, 64)
		resp_time = (int)(ms * 1000)
	}
	status := 200
	if upstream_status != "" {
		status, _ = strconv.Atoi(upstream_status)
	}
	body := ""
	if bytes_sent != "" {
		num, _ := strconv.Atoi(bytes_sent)
		body = strings.Repeat("b", num)
	}

	time.Sleep(time.Duration(resp_time) * time.Millisecond)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	status = 200
	w.WriteHeader(status)
	w.Write([]byte(body))

	log_info(fmt.Sprintf("%s\t%s\t%s\t%s", r.URL.Path, upstream_status, upstream_response_time, bytes_sent))
}

func main() {
	http.HandleFunc("/", backHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
