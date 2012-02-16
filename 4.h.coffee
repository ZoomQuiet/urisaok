crypto = require('crypto')
http = require('http')
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

app.post '/chk', (req, res) ->
    askurl = checkForValidUrl(req.body.uri)
    answer = 'NULL'
    options = 
        host: 'open.pc120.com'
        port: 80
        path: askurl
    http.get options, (pres) ->
        data = ''
        console.log 'STATUS: ' + pres.statusCode
        console.log 'HEADERS: ' + JSON.stringify(pres.headers)
        pres.on 'data', (chunk) ->
            data += chunk.toString()
        pres.on 'end', () ->
            answer = JSON.parse(data)
            console.log answer
            console.log answer.success
            #pres.end answer.success
    res.send "\n\t..."+answer

app.listen 8001
console.log "Express server listening on port %d in %s mode", app.address().port, app.settings.env
