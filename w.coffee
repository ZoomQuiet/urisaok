w = require 'webjs'

getRouter = 
    '/': (req, res) ->
        res.send 'hollo'
    '/try': (req, res) ->
        res.send req.url

postRouter = 
    '/chk': (req, res) ->
        res.send 'Hello ' + req.data.uri

console.log "Hello World"

w.run 8001
    .post postRouter  
    .get getRouter 
    .use w.bodyParser
