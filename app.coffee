crypto = require('crypto')
http = require('http')
fetch = require('fetch').fetchUrl
express = require("express")
app = module.exports = express.createServer()

app.configure ->
    app.use express.bodyParser()
    app.use express.methodOverride()
    app.use app.router

app.configure "production", ->
    app.use express.errorHandler()

app.get "/", (req, res) ->
    res.send "hollo..."

app.post '/chk', (req, res) ->
    askurl = checkForValidUrl(req.body.uri)
    answer = ''
    fetch ASKHOST+askurl , (error, meta, body) ->
        if error
            console.log "ERROR", error.message || error
        console.log meta
        answer = JSON.parse(body)   #body.toString()
        #console.log PHISHTYPE(answer.phish)
        res.send "/cnk KSC::\t"+PHISHTYPE(answer.phish)
    #res.send "\n\t..."+answer

PHISHTYPE = (code) ->
    switch code.toString()
      when "-1" then 'UNKNOW'
      when "0"  then 'GOOD'
      when "1"  then 'PHISH'
      when "2"  then 'MAYBE PHISH'

APPKEY = "k-60666"
SECRET = "99fc9fdbc6761f7d898ad25762407373"
ASKHOST = "http://open.pc120.com"
ASKTYPE = "/phish/?"

checkForValidUrl = (uri) ->
    crtURI = Buffer(uri).toString('base64')
    timestamp = Date.parse(new Date())/1000+".512"
    signbase = ASKTYPE+"appkey="+APPKEY+"&q="+crtURI+"&timestamp="+ timestamp
    #console.log signbase
    #console.log signbase+SECRET
    sign = crypto.createHash('md5').update(signbase+SECRET).digest("hex")
    #ASKHOST+signbase+"&sign="+sign
    signbase+"&sign="+sign

app.listen 80

console.log "Express server listening on port %d in %s mode", app.address().port, app.settings.env
