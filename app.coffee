
app = require('express').createServer();

app.configure ->
    app.get '/', (req, res) -> 
        res.send('Hello World!')
    app.get '/chk', (req, res) -> 
        res.send('chk World!')

app.listen process.env.PORT || 8001
