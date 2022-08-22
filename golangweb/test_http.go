/* package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func testGet() {
	url := "http://apis.juhe.cn/simpleWeather/query?city=上海&key=e1a357724172671ad9c3bd77cf6215ba"
	// Request
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func testGet2() {
	params := url.Values{}
	u, err2 := url.Parse("http://apis.juhe.cn/simpleWeather/query")
	if err2 != nil {
		return
	}
	params.Set("city", "上海")
	params.Set("key", "e1a357724172671ad9c3bd77cf6215ba")
	u.RawQuery = params.Encode()
	urlPath := u.String()
	fmt.Printf("urlPath: %v\n", urlPath)
	r2, err3 := http.Get(urlPath)
	if err3 != nil {
		log.Fatal(err3)
	}
	defer r2.Body.Close()
	b, _ := ioutil.ReadAll(r2.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func testParseJson() {
	type result struct {
		Args    string            `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	var res result
	json.Unmarshal(b, &res)
	fmt.Printf("res: %v\n", res)
}

func testAddHeader() {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	r.Header.Add("name", "xtbo97")
	r.Header.Add("age", "25")
	r2, _ := client.Do(r)
	b, _ := ioutil.ReadAll(r2.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func testPost() {
	urlstr := "http://apis.juhe.cn/simpleWeather/query"
	values := url.Values{}
	values.Set("key", "e1a357724172671ad9c3bd77cf6215ba")
	values.Set("city", "上海")
	r, err := http.PostForm(urlstr, values)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func testPost2() {
	urlValues := url.Values{
		"name":  {"xtbo97"},
		"value": {"25"},
	}
	reqBody := urlValues.Encode()
	r, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func testClient() {
	url := "http://apis.juhe.cn/simpleWeather/query?city=上海&key=e1a357724172671ad9c3bd77cf6215ba"
	client := http.Client{
		Timeout: time.Second * 5,
	}
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r2, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(r2.Body)
	fmt.Printf("string(b): %v\n", string(b))
}

func myServer() {
	f := func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "hello world")
	}
	http.HandleFunc("/hello", f)
	http.ListenAndServe(":9999", nil)
}

type countHandler struct {
	mu sync.Mutex
	n  int
}

// 固定ServeHTTP名字
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func testHttpServer2() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	// testGet()
	// testGet2()
	// testParseJson()
	// testAddHeader()
	// testPost()
	// testPost2()
	// testClient()
	// myServer()
	testHttpServer2()
}
 */