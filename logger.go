package main
 
import (
    "log"
    "net/http"
)

func Logger1(r *http.Request){
        log.Printf(
            "%s\t%s\t%s\t%s\t",
            r.Method,
            r.RequestURI,
            r.Header,
            r.RemoteAddr,
              )
}