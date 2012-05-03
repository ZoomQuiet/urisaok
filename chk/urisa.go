package hollo

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "time"
    "strconv"
//    "strings"
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
//var HOSTTRY = "127.0.0.1:1978"//"urisaok.no.de"
func chk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    c.Infof("url len~\t %v\n", len(url))
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    //bin := []byte(url)
    maxEncLen := base64.URLEncoding.EncodedLen(len([]byte(url))) 
    c.Infof("maxEncLen~\t %v\n", maxEncLen)
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    dst := make([]byte, maxEncLen) //<~ 整来的代码,不理解,就一定会出问题...
    //var dst []byte
    base64.URLEncoding.Encode(dst, []byte(url))
    //c.Infof("EncodedLen~ %v\n", base64.URLEncoding.EncodedLenlen(len(url)))
    c.Infof("base64~\t %v\n", string(dst))
    //code64 := string(dst)
    //c.Infof("base64~ %v\n", len(code64))
    //c.Infof("base64~ %v\n", len(strings.TrimSpace(code64)))
    //c.Infof("base64~ %v\n", len(strings.Trim(code64, "\0000")))
    //c.Infof("base64~ %v\n", code64[0:len(strings.Trim(code64, "\0000"))])
    args := "q=" + string(dst)
    args += "&appkey=" + APPKEY
    now := time.Now()
    //c.Infof("%v , %v", now.Unix(), now.UnixNano())
    nano := strconv.FormatInt(now.UnixNano(),10)
    //c.Infof("timestamp ~ %v.%v", nano[0:10],nano[10:13])
    args += "&timestamp=" + nano[0:10] + "." + nano[10:13]
    sign_base_string := APITYPE + "?" + args 
    c.Infof("sign_base_string~\t %v\n", sign_base_string)
    //md5 hash
    h := md5.New()
    io.WriteString(h, sign_base_string + SECRET)
    //tmp := "/phish/?appkey=YXNkZmFzZGZqYXM&q=aHR0cDovL3NoZW56aGVuLWd6Yy5pbmZv&timestamp=1295430113.546"
    //secret := "6a204bd89f3c8348afd5c77c717a097a"
    //io.WriteString(h, tmp + secret)
    args += "&sign=" + hex.EncodeToString(h.Sum(nil))
    c.Infof("sign~\t %v\n", hex.EncodeToString(h.Sum(nil)))
    c.Infof("args~\t %v\n", args)
    api_url := "http://"+ APIHOST + APITYPE + "?" + args 
    c.Infof("api_url~ \n%v", api_url)
    //APIHOST HOSTTRY
    client := urlfetch.Client(c)
    resp, err := client.Get(api_url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
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

    fmt.Fprint(w, "[DEBUG]\t"+string(buf))
 
    //fmt.Fprint(w, "uri=\t"+r.FormValue("uri"))
    }
/*
    req, err := http.NewRequest("POST", "http://urisaok.appsp0t.com/chk", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    req.Header.Set("use_intranet", "yes")
    t := &urlfetch.Transport{Context: appengine.NewContext(r)}
    resp, err := t.RoundTrip(req)


    type KSC struct {
        Success int    //`json:"success"`
        Phish   string //`json:"phish"`
        Msg     string //`json:"msg"`
    }
    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    c.Infof("Success:%v \n Phish:%v", result.Success ,result.Phish)

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

    //var res KSC
    type KSC struct {
        Success int    //`json:"success"`
        Phish   string //`json:"phish"`
        Msg     string //`json:"msg"`
    }
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
