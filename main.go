package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>hello, 欢迎来到goblog!</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到:(</h1>"+
			"<p>如有疑惑，请联系我们</p>")
	}
}

func aboutHandel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用于记录编程笔记，如您有反馈或建议，请联系"+
		"<a href=\"www.baidu.ciom\">sj</a>")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/about", aboutHandel)
	http.ListenAndServe(":3000", nil)
}
