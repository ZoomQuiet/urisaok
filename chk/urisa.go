package urisa

import (
    "fmt"
    "net/http"
/**/
    "time"
    "strconv"
    "io"
    "io/ioutil"
    "encoding/hex"
    "encoding/json"
    "encoding/base64"
    "crypto/md5"
/**/    
    "appengine"
    "appengine/urlfetch"
)

func init() {
    http.HandleFunc("/", help)
    http.HandleFunc("/chk", chk)
}

func help(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, usageHelp)
}
const usageHelp = `
URIsA ~ KSC 4 GAE powdered by go1
{v12.05.3}
usage:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/chk
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
//var HOSTTRY = "127.0.0.1:1978"//"urisaok.no.de"
func chk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    //c.Infof("url~\t %v\n", url)
    /**/
    //c.Infof("url len~\t %v\n", len(url))
    maxEncLen := base64.URLEncoding.EncodedLen(len([]byte(url))) 
    c.Infof("maxEncLen~\t %v\n", maxEncLen)
    dst := make([]byte, maxEncLen) //<~ 整来的代码,不理解,就一定会出问题...
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    
    base64.URLEncoding.Encode(dst, []byte(url))
    //c.Infof("EncodedLen~ %v\n", base64.URLEncoding.EncodedLenlen(len(url)))
    c.Infof("base64~\t %v\n", string(dst))
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
    //nano := string(now.UnixNano())
    c.Infof("timestamp ~ %v %v.%v", nano, nano[0:10],nano[10:13])
    args += "&timestamp=" + nano[0:10] + "." + nano[10:13]
    sign_base_string := APITYPE + "?" + args 
    c.Infof("sign_base_string~\t %v\n", sign_base_string)
    //md5 hash 严格参数顺序:: appkey -> q -> timestamp
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
    //resp, err := client.Get(api_url)
    resp, err := client.Get("http://open.pc120.com/phish/")
    
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

    type KSC struct {
        Success int     //`json:"success"`
        Phish   int     //`json:"phish"`
        Msg     string  //`json:"msg"`
    }
    //fmt.Fprint(w, "/chk(KCS):\t[DEBUG]")
    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        //panic(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //c.Infof("PHISHID:\t %v %v", result.Phish, PHISHID[0])
    pishmsg, _ := PHISHID[result.Phish]
    c.Infof("Success:%v \n Phish:%s", result.Success ,pishmsg)

    fmt.Fprint(w, "/chk(KCS):\t" + pishmsg)
    /**/
}

/*
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
