http = require 'http'
port = process.env.PORT || 8001

http.createServer (req,res) -> 
    res.writeHead 200, {'Content-Type': 'text/plain'}
    res.end 'Hello World\n'
.listen port

console.log "Listening on port " + port
