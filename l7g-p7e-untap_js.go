package main

import "fmt"
//import "bytes"
//import "bufio"
import "strings"
import "strconv"
import "io/ioutil"
import "github.com/robertkrimen/otto"

//import "github.com/abeconnelly/sloppyjson"

func (lpud *LPUD) status_otto(call otto.FunctionCall) otto.Value {
  v,e := otto.ToValue("ok status")
  if e!=nil { return otto.Value{} }
  return v
}

func (lpud *LPUD) sqlexec_otto(call otto.FunctionCall) otto.Value {
  otto_err,err := otto.ToValue("error")
  if err!=nil { return otto.Value{} }

  sqlstr := call.Argument(0).String()
  ofmt := call.Argument(1).String() ; _ = ofmt

  sql_ret,e := lpud.SQLExec(sqlstr)
  if e!=nil {

    //DEBUG
    fmt.Printf(">>>> sql error: %v\n", e)

    errstr := fmt.Sprint("%v", e)
    oerr,e := otto.ToValue(errstr)
    if e!=nil { return otto_err }
    return oerr
  }

  s := _strstr_to_json(sql_ret)

  v,e := otto.ToValue(s)
  if e!=nil { return otto_err }

  return v
}


func _strstr_to_json(ssa [][]string) string {
  x := []string{}
  x = append(x, "{")
  x = append(x, `"result":[`)
  for i:=0; i<len(ssa); i++ {
    if i>0 { x = append(x, `,`) }
    x = append(x, `[`)
    for j:=0; j<len(ssa[i]); j++ {
      if j>0 { x = append(x, `,`) }
      //x = append(x, `"` + ssa[i][j] + `"`)
      x = append(x, strconv.Quote(ssa[i][j]))
    }
    x = append(x, `]`)

  }
  x = append(x, `]`)
  x = append(x, "}")
  //return strings.Join(x, "")
  return strings.Join(x, "")
}

func (lpud *LPUD) JSVMRun(src string) (rstr string, e error) {
  js_vm := otto.New()

  fmt.Printf("JSVM_run:\n\n")

  init_js,err := ioutil.ReadFile("js/init.js")
  if err!=nil { e = err; return }
  js_vm.Run(init_js)

  js_vm.Set("pheno_status", lpud.status_otto)
  js_vm.Set("pheno_sql", lpud.sqlexec_otto)

  v,err := js_vm.Run(src)
  if err!=nil {
    e = err
    return
  }

  rstr,e = v.ToString()
  return
}

