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

app.post '/chk', (req, res) ->
    #console.log req.body
    #console.log req.body.uri
    #console.log checkForValidUrl(req.body.uri)
    askurl = checkForValidUrl(req.body.uri)
    options = 
        host: 'open.pc120.com'
        port: 80
        path: askurl
    
    data = ''
    http.get options, (res) ->
        res.on 'data', (chunk) ->
            data += chunk.toString()
        res.on 'end', () ->
            console.log data
    res.send askurl+"\n\t"+data

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
