package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {

    http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hello)))))
    http.ListenAndServe(":8080", nil)
}

//实现中间件的功能：1）记录请求的url和方法 2）记录请求的网络的地址 3）记录方法的执行时间
func tracing(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("tracing start 记录请求的url和方法: %s, %s", r.URL, r.Method)
        next.ServeHTTP(w, r)
        log.Println("tracing end")
    })
}

func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("logging start 记录请求的网络的地址: %s", r.RemoteAddr)
        next.ServeHTTP(w, r)
        log.Println("logging end")
    })
}

func timeRecording(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("timeRecording start")
        startTime := time.Now()
        next.ServeHTTP(w, r)
        elapsedTime := time.Since(startTime)
        log.Printf("timeRecording end 记录方法的执行时间 %s", elapsedTime)
    })
}

func hello(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "hello")
}
