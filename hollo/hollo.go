package hollo

import (
    "fmt"
    "net/http"
//    "time" 
    "appengine"
)

func init() {
    http.HandleFunc("/", help)
}

func help(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    c.Infof("ZQ as::logging for debug as [INFO]")
    fmt.Fprint(w, usageHelp)
}
const usageHelp = `
URIsA ~ KSC 4 GAE powdered by go1
{12.05.10}
`
