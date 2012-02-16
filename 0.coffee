http = require 'http'

http.createServer (req,res) -> 
    res.writeHead 200, {'Content-Type': 'text/plain'}
    res.end 'Hello World\n'
.listen process.env.PORT || 8001

