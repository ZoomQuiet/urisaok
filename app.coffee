#120215 appended sync http get support
fetch = require('fetch').fetchUrl
#120211 appended base web dev support
crypto = require('crypto')
http = require('http')

express = require("express")
app = module.exports = express.createServer()
app.configure ->
    app.use express.bodyParser()
    app.use express.methodOverride()
    #app.use express.logger()
    app.use app.router
app.configure "production", ->
    app.use express.errorHandler()

app.get "/", (req, res) ->
    res.send '''Hollo there!
        this is URIsaOK v12.02.26, usage as:
        $ curl --data "uri=http://douban.com" http://urisaok.cnodejs.net/chk
        or
        $ curl --data "uri=http://douban.com" http://urisaok.cnodejs.net/qchk
            ~ use loc MongoDB quickly return checked uri KSC result ;-)
        doc: https://github.com/ZoomQuiet/urisaok
        '''
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

app.post '/chk', (req, res) ->
    askurl = checkForValidUrl(req.body.uri)
    answer = ''
    fetch ASKHOST+askurl , (error, meta, body) ->
        if error
            console.log "ERROR", error.message || error
        else
            console.log meta
            console.log body.toString()
            answer = JSON.parse(body)   #body.toString()
            console.log PHISHTYPE(answer.phish)
            res.send "/cnk KSC::\t"+PHISHTYPE(answer.phish)
            #res.send "/cnk KSC::\t"+answer.phish
    #res.send "\n\t..."+answer

#120221 appended Mongo support
db = require('mongoskin').db('8b2qv5kfx2imv:7okgyd5s6d2@127.0.0.1:20088/mYoQDTSPcCca?auto_reconnect')
#8b2qv5kfx2imv:7okgyd5s6d2@127.0.0.1:20088/mYoQDTSPcCca
#mongo 127.0.0.1:20088/mYoQDTSPcCca -u 8b2qv5kfx2imv -p 7okgyd5s6d2
#db = require('mongoskin').db('localhost:27017/chaos?auto_reconnect')
chked = db.collection('chked')
### Mongo doc design:
'uri':""
'phishcode':""
'timestamp':""
'clientip':""
###
app.post '/qchk', (req, res) ->
    uri = req.body.uri.split("/" ,3)[2]
    console.log uri
    timestamp = Date.parse(new Date())/1000+".512"
    phishcode = "NULL"
    clientip = req.header('x-forwarded-for') || req.connection.remoteAddress
    chked.find({'uri':uri}).toArray (err, result) ->
        if err
            console.log err
        else
            if result.length is 0
                # not chk ever
                console.log "%s \n\tnever chk,ask KSC now!" ,uri
                askurl = checkForValidUrl   "http://"+uri
                fetch ASKHOST+askurl , (error, meta, body) ->
                    if error
                        console.log "ERROR", error.message || error
                        console.log "ERROR"
                    else
                        #console.log meta
                        answer = JSON.parse(body)   #body.toString()
                        console.log answer
                        phishcode = answer.phish
                        doc =
                            'uri': uri
                            'timestamp': timestamp
                            'clientip': clientip
                            'phishcode': phishcode
                        console.log doc
                        #db.collection('test').insert(doc)
                        chked.insert(doc)
                        res.send "/cnk KSC::\t"+PHISHTYPE phishcode
            else
                # had chk.ed
                console.log "%s \n\thad chk.ed,return from MongoDB ;=)" ,uri
                console.log result
                console.log "/cnk KSC::\t"+PHISHTYPE result[0].phishcode
                res.send "/cnk KSC::\t"+PHISHTYPE result[0].phishcode

    #res.send "\n\tQuickly chk with MongoDB!"

app.listen 8001

console.log "Express server listening on port %d in %s mode", app.address().port, app.settings.env
