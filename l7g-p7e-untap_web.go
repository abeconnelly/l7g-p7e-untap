package main

import "fmt"
import "io"
import "net/http"
import "io/ioutil"

import "strconv"

func (lpud *LPUD) WebDefault(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  url := req.URL
  fmt.Printf("default:\n")
  fmt.Printf("  method: %s\n", req.Method)
  fmt.Printf("  proto:  %s\n", req.Proto)
  fmt.Printf("  scheme: %s\n", url.Scheme)
  fmt.Printf("  host:   %s\n", url.Host)
  fmt.Printf("  path:   %s\n", url.Path)
  fmt.Printf("  frag:   %s\n", url.Fragment)
  fmt.Printf("  body:   %s\n\n", body)

  io.WriteString(w, `{"value":"ok"}`)
}

func (lpud *LPUD) WebAbout(w http.ResponseWriter, req *http.Request) {
  str,e := ioutil.ReadFile( lpud.HTMLDir + "/about.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (lpud *LPUD) WebInteractive(w http.ResponseWriter, req *http.Request) {
  str,e := ioutil.ReadFile( lpud.HTMLDir + "/index.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (lpud *LPUD) WebExec(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  fmt.Printf("webexec got>>>\n%s\n\n", body)

  rstr,e := lpud.JSVMRun(string(body))
  if e!=nil {
    rerr := strconv.Quote(fmt.Sprintf("%v", e))
    io.WriteString(w, `{"value":"error","error":` + rerr + `}`)
    return
  }

  io.WriteString(w, rstr)
}

func (lpud *LPUD) StartSrv() error {
  http.HandleFunc("/", lpud.WebDefault)
  http.HandleFunc("/exec", lpud.WebExec)
  http.HandleFunc("/exec/", lpud.WebExec)
  http.HandleFunc("/about", lpud.WebAbout)
  http.HandleFunc("/about/", lpud.WebAbout)
  http.HandleFunc("/i", lpud.WebInteractive)
  http.HandleFunc("/i/", lpud.WebInteractive)

  port_str := fmt.Sprintf("%d", lpud.Port)
  e := http.ListenAndServe(":" + port_str, nil)
  return e
}
