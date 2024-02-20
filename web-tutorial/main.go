package main

import "net/http"

type helloHandler struct{}
type aboutHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me!"))
}

func welcom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {

	//http.ListenAndServe的底层操作
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, //DefaultServeMux
	}

	//网页处理
	//go内置handler，handler是个有SeverHTTP方法的类型
	//NotFOundHandler、RedirectHandler、StripPrefix、TimeoutHandler、FileSever
	mh := &helloHandler{}
	a := &aboutHandler{}
	http.Handle("/hello", mh)
	http.Handle("/about", a)
	
	//HandleFunc是要传入一个与SeverHTTP参数一样的函数，该函数会被转换为具有SeverHTTP方法的HandlerFunc类型
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home."))
	})
	http.HandleFunc("/welcome", http.HandlerFunc(welcom)) //可以直接传welcome，HandlerFunc有Handler方法，本质还是调用handler
	
	//开始监听
	server.ListenAndServe()

	// http.ListenAndServe("localhost:8080", nil) //defaultServeMux默认路由
}
