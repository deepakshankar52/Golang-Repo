// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	handler := http.HandleFunc(accountInfo)
// 	http.Handle("/accounts", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func accountInfo(w http.ResponseWriter, r *http.Request) {
	
// 	// id := r.Header.Values("acc_no")
// 	// fmt.Println("Account Number: %d", id)

// 	contentType := r.Header.Get("content-type")
// 	fmt.Println("Content-Type: %s", contentType)

// 	// accept := r.Header.Get("accept")
// 	// fmt.Println("Accept: %s", accept)

// 	headers := r.Header
// 	fmt.Println("Header Details: %s", headers)
// 	fmt.Println("Content-Type: %s", headers["Content-Type"])
// 	// fmt.Println("Accept: %s", headers["Accept"])
// 	// fmt.Println("Account Number: %d", headers["ID"])
// }


// set request headers

// package main
// import (
//     "fmt"
//     "net/http"
//     "time"
// )
// func main() {
//     call("https://google.com", "GET")
// }
// func call(url, method string) error {
//     client := &http.Client{
//         Timeout: time.Second * 10,
//     }
//     req, err := http.NewRequest(method, url, nil)
//     if err != nil {
//         return fmt.Errorf("Got error %s", err.Error())
//     }
//     req.Header.Set("user-agent", "golang application")
//     req.Header.Add("foo", "bar1")
//     req.Header.Add("foo", "bar2")
//     response, err := client.Do(req)
//     if err != nil {
//         return fmt.Errorf("Got error %s", err.Error())
//     }
//     defer response.Body.Close()
//     return nil
// }

// get request headers from an incoming http request

package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/account", handler)
	http.ListenAndServe(":8080", nil)
}
	
func handleRequest(w http.ResponseWriter, r *http.Request) {

	nameValues := r.Header.Values("name")
	fmt.Printf("r.Header.Values(\"name\"):: %s\n\n", nameValues)

	contentType := r.Header.Get("content-type")
	fmt.Printf("r.Header.Get(\"content-type\"):: %s\n\n", contentType)

	headers := r.Header
	fmt.Printf("r.Headers:: %s\n\n", headers)
	fmt.Printf("r.Headers[\"Content-Type\"]:: %s\n\n", headers["Content-Type"])
	fmt.Printf("r.Headers[\"Name\"]:: %s", headers["Name"])
}
