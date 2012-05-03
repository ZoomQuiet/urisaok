package hollo

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "time"
    "strconv"
//    "bytes"
    "encoding/base64"
    "io"
    "encoding/hex"
    "crypto/md5"
    
    "appengine"
    "appengine/urlfetch"
)

func init() {
    http.HandleFunc("/", help)
    http.HandleFunc("/chk", chk)
    http.HandleFunc("/try", try)
}

func help(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, usageHelp)
}
const usageHelp = `
URIsA ~ KSC 4 GAE powdered by go1
usage:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/chk
`

var APPKEY  = "k-60666"
var SECRET  = "99fc9fdbc6761f7d898ad25762407373"
var APIHOST = "open.pc120.com"
var APITYPE = "/phish/"
func chk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    var now = time.Now()
    c.Infof("%v , %v", now.Unix(), now.UnixNano())
    var nano = strconv.FormatInt(now.UnixNano(),10)
    c.Infof("timestamp ~ %v.%v", nano[0:10],nano[10:13])

    //base64 encode...
    dst := make([]byte, 256)
    base64.URLEncoding.Encode(dst, []byte(url))
    c.Infof("%v\n", string(dst))

    //md5 hash
    h := md5.New()
    io.WriteString(h, url)
    c.Infof("%v", hex.EncodeToString(h.Sum(nil)))
 
    fmt.Fprint(w, "uri=\t"+r.FormValue("uri"))
    }
/*
    h := md5.New()
    io.WriteString(h, "The fog is getting thicker!")
    fmt.Printf("%x", h.Sum(nil))

   //base64 encode...
    //var buf = bytes.NewBufferString(url)
    var encoder = base64.NewEncoder(base64.NewEncoding(url), &buf)
    encoder.Encode()
    //encoder.Close()
    c.Infof("base64: %v", buf.String())

    var URLEncoded = base64.NewEncoding(url)
    enc := URLEncoded.EncodeToString()
    c.Infof("enc: %v %v", len(enc), string(enc)) 

def __genQueryArgs(api_path, url):
    args = "appkey=" + cfg.APPKEY
    args += "&q=" + base64.urlsafe_b64encode(url)
    args += "&timestamp=" + "%.3f" % (time.time())
    sign_base_string = api_path + "?" + args
    args += "&sign=" + md5(sign_base_string + cfg.SECRET).hexdigest()
    return args

def _askCloud(api_path, url):
    args = __genQueryArgs(api_path, url)
    api_url = "http://%s%s?%s"% (cfg.OPEN_HOST, cfg.APITYPE ,args)
    print api_url
    result = eval(urilib.urlopen(api_url).read())
    print result
    if result['success'] == 1:
        return result['phish']
    else:
        return result
*/

func try(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprint(w, "uri=\t"+r.FormValue("uri"))
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get("http://open.pc120.com/phish/")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)    
    c.Infof("HTTP GET returned status %v", resp.Status)
    if resp.StatusCode != 200 {
        http.Error(w, "couldn't get sale data", http.StatusInternalServerError)
        return
    }

    defer resp.Body.Close()

    c.Infof("resp.ContentLength %v", resp.ContentLength)

    var buf []byte
    buf, _ = ioutil.ReadAll(resp.Body)
    c.Infof("resp.Body %v", string(buf)) 

    type KSC struct {
        Success int    //`json:"success"`
        Msg     string //`json:"msg"`
    }
    //var res KSC
    //json.Unmarshal(body, result)
    //json.Unmarshal(body, &res); 
    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        //panic(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    c.Infof("%+v", result.Msg)

    fmt.Fprintf(w, "HTTP GET returned :: %v", resp.Body)
       
}
  
    // The message body.    Body io.ReadCloser
    /**
    var f interface{}
    json.Unmarshal(body, &f)
    m := f.(map[string]interface{})
    for k, v := range m {
        switch vv := v.(type) {
        case string:
            c.Infof("%v is str %v", k, vv)
        case int:
            c.Infof("%v is int %v", k, vv)
        case []interface{}:
            c.Infof("%v is an array:",k )
            for i, u := range vv {
                c.Infof("%v %v ", i, u)
            }
        default:
            c.Infof(k, "is of a type I don't know how to handle")
        }
    }

    **/
