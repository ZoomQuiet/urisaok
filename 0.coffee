http = require 'http'

http.createServer (req,res) -> 
    res.writeHead 200, {'Content-Type': 'text/plain'}
    res.end '''Hello World
        URIsaok base KSC
            {v12.02.13.1}
        '''
.listen process.env.PORT || 8001

