urisaok = require('zappa').app ->
  @get '/': 'hi'
  
urisaok.app.listen 8001


