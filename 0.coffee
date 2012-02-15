express: require 'coffee-script'
express: require 'express'
app: express.createServer()
eat: (food) -> food + " om nom nom..."

app.get "/", (req, res) ->
    nomnom: eat food for food in ["Banana", "Apple", "Turkey"]
    res.send nomnom                                               
port: 80                                               

console.log "Listening on port " + port

app.listen port
