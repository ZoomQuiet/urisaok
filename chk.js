var crypto = require('crypto')
var http = require('http')
var fetch = require('fetch').fetchUrl

var express = require("express")
var app = module.exports = express.createServer()
app.configure (function() {
    app.use(express.bodyParser())
    app.use(express.methodOverride())
    app.use(app.router)
})

app.get("/", function(req, res) {
    res.send("Hollo there! \n this is URIsaOK v12.02.26, usage as:\n$ curl --data 'uri=http://douban.com' http://urisaok.cnodejs.net/chk\nor\n$ curl --data 'uri=http://douban.com' http://urisaok.cnodejs.net/qchk\n~ use loc MongoDB quickly return checked uri KSC result ;-)\n doc: https://github.com/ZoomQuiet/urisaok")
})

var APPKEY = "k-60666"
var SECRET = "99fc9fdbc6761f7d898ad25762407373"
var ASKHOST = "http://open.pc120.com"
var ASKTYPE = "/phish/?"
PHISHTYPE = (function(code){
    switch(code.toString()){
      case "-1":
        return('UNKNOW')
      case "0":
        return('GOOD')
      case "1":
        return('PHISH')
      case "2":
        return('MAYBE PHISH')
    }
})
//console.log(PHISHTYPE(2))
checkForValidUrl = (function(uri) {
    var crtURI = Buffer(uri).toString('base64')
    var timestamp = Date.parse(new Date())/1000+".512"
    var signbase = ASKTYPE+"appkey="+APPKEY+"&q="+crtURI+"&timestamp="+ timestamp
    var sign = crypto.createHash('md5').update(signbase+SECRET).digest("hex")
    return(signbase+"&sign="+sign)
})
//console.log(checkForValidUrl("http://sina.com"))
app.post('/chk', function(req, res) {
    var askurl = checkForValidUrl(req.body.uri)
    var answer = ''
    fetch(ASKHOST+askurl , function(error, meta, body) {
        if(error){
            console.log("ERROR", error.message || error)
        }else{
            console.log(meta)
            console.log(body.toString())
            answer = JSON.parse(body)
            console.log(PHISHTYPE(answer.phish))
            res.send("/cnk KSC::\t"+PHISHTYPE(answer.phish))
        }
    })
})

var db = require('mongoskin').db('8b2qv5kfx2imv:7okgyd5s6d2@127.0.0.1:20088/mYoQDTSPcCca?auto_reconnect')
var chked = db.collection('chked')
app.post('/qchk', function(req, res) {
    var uri = req.body.uri.split("/" ,3)[2]
    console.log(uri)
    var timestamp = Date.parse(new Date())/1000+".512"
    var phishcode = "NULL"
    var clientip = req.header('x-forwarded-for') || req.connection.remoteAddress
    chked.find({'uri':uri}).toArray(function(err, result) {
        if(err){
            console.log(err)
        }else{
            if(result.length === 0){
                // not chk ever
                console.log("%s \n\tnever chk,ask KSC now!" ,uri)
                var askurl = checkForValidUrl("http://"+uri)
                fetch(ASKHOST+askurl , function(error, meta, body){
                    if(error){
                        console.log("ERROR", error.message || error)
                    }else{
                        var answer = JSON.parse(body)
                        console.log(answer)
                        var phishcode = answer.phish
                        var doc = {'uri': uri
                            ,'timestamp': timestamp
                            ,'clientip': clientip
                            ,'phishcode': phishcode
                            }
                        console.log(doc)
                        chked.insert(doc)
                        res.send("/cnk KSC::\t"+PHISHTYPE(phishcode))
                    } 
                })
            }else{
                // had chk.ed
                console.log("%s \n\thad chk.ed,return from MongoDB ;=)" ,uri)
                console.log(result)
                console.log("/cnk KSC::\t"+PHISHTYPE(result[0].phishcode))
                res.send("/cnk KSC::\t"+PHISHTYPE(result[0].phishcode))
            }
        }
    })
})

app.listen(8001)
console.log("Express server listening on port %d in %s mode",app.address().port, app.settings.env)

