express = require 'express'
# about express
app = express.createServer()
app.configure () ->
    app.use express.methodOverride()
    app.use express.bodyParser()
    app.use app.router
#app = require('express').createServer();
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
app.listen 8001
console.log "run server. port #{port}."
