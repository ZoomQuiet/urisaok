app = require('express').createServer()
#express = require("express")
#app = module.exports = express.createServer()

app.configure ->
    app.use express.methodOverride()
    app.use express.bodyParser()
    #app.use app.router
    app.get '/', (req, res) -> 
        res.send("""Hello World!
            for URIsaok
            {v12.02.13}
            """)
    app.post '/chk', (req, res) ->
        console.log req.body
        res.send req.body

app.listen process.env.PORT || 8001
