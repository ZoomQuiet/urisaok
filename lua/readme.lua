-- reamdme for /=/ export base help!
ngx.req.read_body()
VERTION="URIsAok4openresty v12.03.6"

ngx.say(VERTION
    ,"\n\tusage:"
    ,"$crul -d 'uri=http://sina.com' 127.0.0.1:9090/=/chk"
    )
