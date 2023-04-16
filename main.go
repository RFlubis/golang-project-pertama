package main

import (
	"fmt"
	"net/http"
	"time"
)
//routing outside main function
func handlerIndex(w http.ResponseWriter, r *http.Request) {
    var message = "Welcome"
    w.Write([]byte(message))
}



func main() {
    //routing inside main function
    handlerHello := func(w http.ResponseWriter, r *http.Request) {
        var message = "Hello world!"
        w.Write([]byte(message))
    }
    //anonymous function routing (have to be inside main function)
    http.HandleFunc("/data",func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("hello again"))
    })

    //routing static assets
    http.Handle("/static/",
        //wraping the FileServer
        //refer to https://dasarpemrogramangolang.novalagung.com/B-routing-static-assets.html
        http.StripPrefix("/static/",
            //get asset folder
            http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", handlerIndex)
    http.HandleFunc("/index", handlerIndex)
    http.HandleFunc("/hello", handlerHello)

    var address = "localhost:9000"
    fmt.Printf("server started at %s\n", address)

    server :=new(http.Server)
    server.Addr = address
    server.ReadTimeout = time.Second * 10
    server.WriteTimeout = time.Second * 10

    err := server.ListenAndServe()
    if err != nil {
        fmt.Println(err.Error())
    }
}

//added from mac