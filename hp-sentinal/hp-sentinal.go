package main
import (
    "fmt" 
    "net"  
    "reflect"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    response := "\x7A\x2C\x50\x53\x48\x2C\x27\x41\x7B\x5E\x51\x4F\x48\x70\x65\x5D\x29\x5D\x5C\x5E\x63\x52\x48\x3E\x25\x67\x4E\x51\x58\x62\x48\x3E\x31\x4A\x52\x4E\x5B\x68\x4B\x53\x24\x67\x4D\x50\x66\x6D\x65\x41\x7B\x5A\x5B\x60\x62\x6B\x54\x5B\x29\x68\x3E\x41\x45" 
    _,err := conn.WriteToUDP([]byte(response), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 5093,
        IP: net.ParseIP("0.0.0.0"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := ser.ReadFromUDP(p)
        // fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
	h := "\x7A\x00\x00\x00\x00\x00"
        m := []byte(h)
	pp := []byte(p)[0:6]
	if reflect.DeepEqual(pp, m) {
		go sendResponse(ser, remoteaddr)	
		}
	// fmt.Println(pp,m)
	if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        continue
      }
}
