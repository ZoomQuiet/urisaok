mate = require 'coffeemate'

mate.get '/', ->
    @resp.end 'Hello World'

console.log "Hello World"

mate.listen 8001
