    askurl = checkForValidUrl(req.body.uri)
    options = 
        host: 'open.pc120.com'
        port: 80
        path: askurl
    http.get options, (pres) ->
        data = ''
        #console.log 'STATUS: ' + res.statusCode
        #console.log 'HEADERS: ' + JSON.stringify(res.headers)
        pres.on 'data', (chunk) ->
            data += chunk.toString()
        pres.on 'end', () ->
            #console.log data
            answer = JSON.parse(data)
            console.log answer.success
            pres.end answer.success
