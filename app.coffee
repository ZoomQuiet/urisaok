
app = require('express').createServer();

app.configure ->
    app.get '/', (req, res) -> 
        res.send("""Hello World!
            for URIsaok
            {v12.02.13}
            """)
    app.get '/chk', (req, res) -> 
        #params = req.body
        res.send(req.query.content)

app.listen process.env.PORT || 8001
