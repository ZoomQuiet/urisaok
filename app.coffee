express = require 'express'
#app = require('express').createServer();

# about express
app = express.createServer()
app.configure ()->
    app.use express.methodOverride()
    app.use express.bodyParser()
    app.use app.router

app.configure 'development', ()->
  app.use express.errorHandler({ dumpExceptions: true, showStack: true })

# ====== API ========
app.get '/', (req, res) -> 
    res.send("""Hello World!
        for URIsaok
        {v12.02.13}
        """)

app.post '/chk', (req, res) -> 
    params = req.query.content  #req.body.uri
    res.send(params)


# Bind Application
port = 8001
app.listen port
console.log "run server. port #{port}."
