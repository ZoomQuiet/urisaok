-- try openresty easy creat RESTful API srv.
ngx.req.read_body()
local method   = ngx.var.request_method
ngx.say("remote_addr:\t",ngx.var.remote_addr)
--ngx.log(ngx.INFO,"\n\trequest_method: ",method,"\n\t")
-- intra. func.
curl = require "luacurl"
function _fetch_uri(url, c)
    local result = { }
    if c == nil then 
        c = curl.new() 
    end
    c:setopt(curl.OPT_URL, url)
    c:setopt(curl.OPT_WRITEDATA, result)
    c:setopt(curl.OPT_WRITEFUNCTION, function(tab, buffer)
        table.insert(tab, buffer)
        return #buffer
    end)
    local ok = c:perform()
    return ok, table.concat(result)
end
function _split(str, sep)
    fields = {}
    str:gsub("([^"..sep.."]*)"..sep, function(c) table.insert(fields, c) end)
    return fields
end
-- global var
PHISHTYPE = {["-1"]='UNKNOW'
    ,["0"]='GOOD'
    ,["1"]='PHISH'
    ,["2"]='MAYBE PHISH'
    }
--ngx.say(PHISHTYPE["2"])
APPKEY = "k-60666"
SECRET = "99fc9fdbc6761f7d898ad25762407373"
ASKHOST = "http://open.pc120.com"
ASKTYPE = "/phish/?"
function checkForValidUrl(uri)
    crtURI = ngx.encode_base64(uri)
    timestamp = ngx.now()
    signbase = ASKTYPE .. "appkey=" .. APPKEY .. "&q=" .. crtURI .. "&timestamp=" .. timestamp
    sign = ngx.md5(signbase .. SECRET)
    return ASKHOST .. signbase .. "&sign=" .. sign
end

if method ~= 'POST' then
    ngx.say('only POST chk me;-)')
    local readme = ngx.location.capture("/")
    if res.status == 200 then
        ngx.say(readme.body)
    end
else
    local data = ngx.req.get_body_data()
    ngx.say("get_body_data:\t",data)
    local args = ngx.req.get_post_args()
    local uri = args.uri
    ngx.say("get_post_args:\t",args.uri)
    fields = _split(uri,"/")
    ngx.say("str:gsub\t",fields[3])
    local url = fields[3]
    local chkURI = checkForValidUrl(url)
    ngx.say(chkURI)
    ok, html = _fetch_uri(chkURI)
    if ok then
        local cjson = require "cjson"
        json = cjson.decode(html)
        if 1 == json.success then
            ngx.log(ngx.INFO,"\n\tKCS say:  ",html,"\n\t")
            ngx.say("KCS /phish?:\t", PHISHTYPE[tostring(json.phish)])
        else
            ngx.say("KCS say:\t",html)
        end
    end
end