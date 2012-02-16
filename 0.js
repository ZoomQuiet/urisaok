var http = require('http');
 
var server = http.createServer(function (req, res) {
  res.writeHead(200, { "Content-Type": "text/plain" })
  res.end("Hello world\n URIsaok base KSC \n{v12.02.13}");
});
console.log("hollooooo")
server.listen(process.env.PORT || 8001);
