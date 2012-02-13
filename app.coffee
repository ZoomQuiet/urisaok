
app = require('express').createServer();

app.configure ->
    app.get '/', (req, res) -> 
        res.send("""Hello World!
            for URIsaok
            {v12.02.13}
            """)
    app.post '/chk', (req, res) -> 
        params = req.body
        res.send('chk World!')

app.listen process.env.PORT || 8001
