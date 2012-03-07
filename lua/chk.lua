-- try openresty easy creat RESTful API srv.
ngx.req.read_body()
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


local method   = ngx.var.request_method
ngx.say("remote_addr:\t",ngx.var.remote_addr)
--ngx.log(ngx.INFO,"\n\trequest_method: ",method,"\n\t")
if method == 'POST' then
    local data = ngx.req.get_body_data()
    ngx.say("get_body_data:\t",data)

    local args = ngx.req.get_post_args()
    local uri = args.uri
    ngx.say("get_post_args:\t",args.uri)
    ngx.say("md5:\t",ngx.md5(args.uri))
    ngx.say("encode_base64:\t",ngx.encode_base64(args.uri))
    ngx.say("ngx.now():\t",ngx.now())

    urili = string._split(uri,"/",4)
    --print(urili)
    --for k,v in next,  string._split(uri,"/",4) do ngx.say(v) end
    ngx.say("string._split:\t",urili[3])
    --[[
    ok, html = _fetch_uri("http://open.pc120.com/phish/")
    if ok then
        ngx.say("KCS /phish?:\t",html)
    end
    ]]

else
    ngx.say('only POST chk me;-)')
end

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

function string:_split(sSeparator, nMax, bRegexp)
    assert(sSeparator ~= '')
    assert(nMax == nil or nMax >= 1)
    local aRecord = {}
    if self:len() > 0 then
        local bPlain = not bRegexp
        nMax = nMax or -1
        local nField=1 nStart=1
        local nFirst,nLast = self:find(sSeparator, nStart, bPlain)
        while nFirst and nMax ~= 0 do
            aRecord[nField] = self:sub(nStart, nFirst-1)
            nField = nField+1
            nStart = nLast+1
            nFirst,nLast = self:find(sSeparator, nStart, bPlain)
            nMax = nMax-1
        end
        aRecord[nField] = self:sub(nStart)
    end
    return aRecord
end
