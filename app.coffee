exp = require 'express'
app = exp.createServer()
#app = require('express').createServer();

app.configure () ->
    app.use exp.methodOverride()
    app.use exp.bodyParser()
    #app.use app.router
app.get '/', (req, res) -> 
    res.send("""Hello World!
        for URIsaok
        {v12.02.13}
        """)
app.post '/chk', (req, res) -> 
    params = req.body
    res.send(params)

#app.listen process.env.PORT || 8001
app.listen 8001
