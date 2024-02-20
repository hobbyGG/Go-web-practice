package main

import (
	"log"
	"net/http"
)

// 本例为获取url的Query信息的两种方法
func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query() //返回一个map

		id := query["id"] //直接用map的关键字来找，该方法呼返回所有id对于的value
		log.Println(id)

		name := query.Get("name") //用该类型上的get方法来得到value值，该方法只会返回第一个name对应的值
		log.Println(name)
	})

	http.ListenAndServe(":8080", nil)
}
