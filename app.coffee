express = require("express")
app = module.exports = express.createServer()

app.configure ->
  app.use express.bodyParser()
  app.use express.methodOverride()
  app.use app.router

app.configure "development", ->
  app.use express.errorHandler(
    dumpExceptions: true
    showStack: true
  )

app.configure "production", ->
  app.use express.errorHandler()

app.get "/", (req, res) ->
  res.send "hollo..."

app.post '/chk', (req, res) ->
  console.log req.body
  console.log req.body.uri
  res.send req.body.uri

app.listen 8001

console.log "Express server listening on port %d in %s mode", app.address().port, app.settings.env
