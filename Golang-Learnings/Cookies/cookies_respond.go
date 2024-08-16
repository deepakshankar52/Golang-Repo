// Set a cookie while responding to an incoming request

package main
import (
    "net/http"
)
func main() {
    docHandler := http.HandlerFunc(docHandler)
    http.Handle("/doc", docHandler)
    http.ListenAndServe(":8080", nil)
}
func docHandler(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:   "id",
        Value:  "abcd",
        MaxAge: 300,
    }
    http.SetCookie(w, cookie)
    w.WriteHeader(200)
    w.Write([]byte("Doc Get Successful"))
    return
}