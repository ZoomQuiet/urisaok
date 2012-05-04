package urisa

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
    "appengine/datastore"
//    "appengine/user"
)

func init() {
    http.HandleFunc("/", help)
    http.HandleFunc("/chk", chk)
    http.HandleFunc("/qchk", qchk)
}

func help(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, usageHelp)
}
const usageHelp = `
URIsA ~ KSC 4 GAE powdered by go1
{v12.05.4}
usage:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/chk
or with GAE Datastore quick resp. if ahd checked:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/qchk
`

var APPKEY  = "k-60666"
var SECRET  = "99fc9fdbc6761f7d898ad25762407373"
var APIHOST = "open.pc120.com"
var APITYPE = "/phish/"
var PHISHID = map[int] string {
    -1:   "UNKNOW",
    0:    "GOOD",
    1:    "PHISH!",
    2:    "MAYBE...",
}
type KSC struct {
    Success int    //`json:"success"`
    Phish   int //`json:"phish"`
    Msg     string //`json:"msg"`
}
func _genKSCuri(url string) string {
    println("url len~\t ", len(url))
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    //bin := []byte(url)
    maxEncLen := base64.URLEncoding.EncodedLen(len([]byte(url))) 
    println("maxEncLen~\t ", maxEncLen)
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    dst := make([]byte, maxEncLen) //<~ 整来的代码,不理解,就一定会出问题...
    //var dst []byte
    base64.URLEncoding.Encode(dst, []byte(url))
    //c.Infof("EncodedLen~ %v\n", base64.URLEncoding.EncodedLenlen(len(url)))
    println("base64~\t", string(dst))
    //code64 := string(dst)
    //c.Infof("base64~ %v\n", len(code64))
    //c.Infof("base64~ %v\n", len(strings.TrimSpace(code64)))
    //c.Infof("base64~ %v\n", len(strings.Trim(code64, "\0000")))
    //c.Infof("base64~ %v\n", code64[0:len(strings.Trim(code64, "\0000"))])
    args := "appkey=" + APPKEY
    args += "&q=" + string(dst)
    now := time.Now()
    //c.Infof("%v , %v", now.Unix(), now.UnixNano())
    nano := strconv.FormatInt(now.UnixNano(),10)
    //c.Infof("timestamp ~ %v.%v", nano[0:10],nano[10:13])
    args += "&timestamp=" + nano[0:10] + "." + nano[10:13]
    sign_base_string := APITYPE + "?" + args 
    println("sign_base_string~\t ", sign_base_string)
    //md5 hash 严格参数顺序:: appkey -> q -> timestamp
    h := md5.New()
    io.WriteString(h, sign_base_string + SECRET)
    //tmp := "/phish/?appkey=YXNkZmFzZGZqYXM&q=aHR0cDovL3NoZW56aGVuLWd6Yy5pbmZv&timestamp=1295430113.546"
    //secret := "6a204bd89f3c8348afd5c77c717a097a"
    //io.WriteString(h, tmp + secret)
    args += "&sign=" + hex.EncodeToString(h.Sum(nil))
    println("sign~\t ", hex.EncodeToString(h.Sum(nil)))
    println("args~\t ", args)
    api_url := "http://"+ APIHOST + APITYPE + "?" + args 
    println("api_url~ ", api_url)

    return api_url
}

func _asKSC(uri string, r *http.Request) (int, int) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get(uri)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
    }
    c.Infof("HTTP GET returned status %v", resp.Status)
    if resp.StatusCode != 200 {
        panic(err)
        c.Infof("couldn't get sale data %v", http.StatusInternalServerError)
        //http.Error(w, "couldn't get sale data", http.StatusInternalServerError)
        //return
    }
    defer resp.Body.Close()
    c.Infof("resp.ContentLength %v", resp.ContentLength)
    var buf []byte
    buf, _ = ioutil.ReadAll(resp.Body)
    c.Infof("resp.Body %v", string(buf))

    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
    }
    
    return result.Success, result.Phish
}


type Chked struct {
    Uri     string
    Phish    int
    Tstamp  time.Time
    Cip     string
}
func qchk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    //c.Infof("r.RemoteAddr %v", r.RemoteAddr)
    //key := url
    ukey := datastore.NewKey(c, "Uri", url, 0, nil)
    var e2 Chked
    if err := datastore.Get(c, ukey, &e2); err != nil {
        fmt.Fprint(w, "~ ", err.Error() , "\n")

        //panic(err)
        c.Infof("Get Err.~\n\t !!! %v", err.Error())
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
        api_url := _genKSCuri(url)
        _,p := _asKSC(api_url, r)
        //c.Infof("Success:%v \t Phish:%s", s ,PHISHID[p])
        e1 := Chked{
            Uri:    url,
            Phish:  p,
            Tstamp: time.Now(),
            Cip:    r.RemoteAddr,
        }
        /*
        if ukey, err = datastore.DecodeKey(url); err != nil {
            fmt.Fprint(w, "datastore.DecodeKey~ ", err.Error() , "\n")
        }
        */
        //key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "chked", nil), &e1)
        ukey, err = datastore.Put(c, ukey, &e1)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprint(w, "ask KCS API srv.\n")
        fmt.Fprint(w, "/qchk(KCS):\t" + PHISHID[p])

    }else{
        c.Infof("datastore ukey=%v", ukey.String())

        fmt.Fprint(w, "datastore Get OK;-) \n")
        c.Infof("co4 datastore:%v \t Phish:%s", e2.Phish ,PHISHID[e2.Phish])

        fmt.Fprint(w, "/qchk(GAE):\t" + PHISHID[e2.Phish])
    }
    //c.Infof("c2.Pish[4datastore] %v", c2.Pish)
    //c.Infof("genKSCuri(url) %v", _genKSCuri(url))

    //fmt.Fprint(w, "/qchk(KCS):\t" + url)
}

func chk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    
    api_url := _genKSCuri(url)

    s,p := _asKSC(api_url, r)
    c.Infof("Success:%v \t Phish:%s", s ,PHISHID[p])

    fmt.Fprint(w, "/chk(KCS):\t" + PHISHID[p])
}

/*
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

    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //c.Infof("PHISHID:\t %v %v", result.Phish, PHISHID[0])
    pishmsg, _ := PHISHID[result.Phish]
    
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