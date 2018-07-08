package main

import (
  "fmt"
  "log"
  "net/url"
)

func main() {
  //General form
  u, err := url.Parse("http://user1@bing.com/search?q=foo%2fbar&q2=hoge#fragment")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("URL: %s\n", u.String())
  fmt.Printf("Scheme: %s\n", u.Scheme)
  fmt.Printf("Opaque: %s\n", u.Opaque)
  fmt.Printf("User: %s\n", u.User)
  fmt.Printf("Host: %s\n", u.Host)
  fmt.Printf("Path: %s\n", u.Path)
  fmt.Printf("RawPath: %s\n", u.RawPath)
  fmt.Printf("RawQuery: %s\n", u.RawQuery)
  fmt.Printf("Fragment: %s\n", u.Fragment)

  for key, values := range u.Query() {
    fmt.Printf("Query Key: %s\n", key)    
    for i, v := range values {
      fmt.Printf("Query Value[%d]: %s\n", i, v)          
    } 
  }  
}

/*

$go run parse.go
URL: http://user1@bing.com/search?q=foo%2fbar&q2=hoge#fragment
Scheme: http
Opaque:
User: user1
Host: bing.com
Path: /search
RawPath:
RawQuery: q=foo%2fbar&q2=hoge
Fragment: fragment
Query Key: q
Query Value[0]: foo/bar
Query Key: q2
Query Value[0]: hoge

*/
