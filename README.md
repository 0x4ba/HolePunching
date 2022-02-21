# HolePunching
udp HolePunching

## usage
  `main.go
  var (
    s         = host{ipaddr: "xxx.xxx.xxx.xxx", port: "666"}
    localhost = host{ipaddr: "", port: "666"}
  )
  `
  Modify s.ipaddr and port 
  
### server
  go run . -s
  
### client 
  go run . -c
