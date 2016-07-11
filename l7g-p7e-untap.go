package main

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

//import "net/http"
//import "github.com/robertkrimen/otto"
//import "github.com/abeconnelly/sloppyjson"

type LPUD struct {
  DB *sql.DB
}

func (lpud *LPUD) Init(sql_fn string) error {
  var err error
  lpud.DB, err = sql.Open("sqlite3", sql_fn)
  if err !=nil { panic(err) }
  return nil
}

func (lpud *LPUD) SQLExec(req string) ([][]string, error ) {
  rows,err := lpud.DB.Query(req)
  if err!=nil { return nil, err }
  cols,e := rows.Columns() ; _ = cols
  if e!=nil { return nil, e }

  rawResult := make([][]byte, len(cols))

  res_str_array := [][]string{}

  dest := make([]interface{}, len(cols))
  for i,_ := range rawResult {
    dest[i] = &rawResult[i]
  }

  for rows.Next() {
    err := rows.Scan(dest...)
    if err!=nil { return nil,err }

    result := make([]string, len(cols))

    for i,raw := range rawResult {
      if raw==nil {
        result[i] = "\n"
      } else {
        result[i] = string(raw)

        //DEBUG
        fmt.Printf("raw>>>>\n%v\n", string(raw))

      }
    }

    res_str_array = append(res_str_array, result)

  }

  //DEBUG
  fmt.Printf(">>>>\n%v\n", res_str_array)

  return res_str_array, nil
}

func main() {
  local_debug := true

//  db, err := sql.Open("sqlite3", "./untap.sqlite3")
//  if err !=nil { panic(err) }
//  rows,err := db.Query("select * from demographics limit 10")
//  if err!=nil { panic(err) }
//  for rows.Next() {
//    var id int
//    var human_id string
//    var date_of_birth string
//    var gender string
//    var weight string
//    var height string
//    var blood_type string
//    var race string
//    err = rows.Scan(&id, &human_id, &date_of_birth, &gender, &weight, &height, &blood_type, &race)
//    fmt.Print(id, human_id, date_of_birth, gender, weight, height, blood_type, race, "\n")
//  }

  lpud := LPUD{}

  err := lpud.Init("./untap.sqlite3")
  if err!=nil { panic(err) }

  if local_debug {
    fmt.Printf(">> starting\n")
  }

  err = lpud.StartSrv()
  if err!=nil { panic(err) }

}
