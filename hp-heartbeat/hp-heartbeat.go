package main
import (
    "fmt" 
    "net"  
    "reflect"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    response := "\x43\x7D\xB9\xDF\x90\x23\xD6\x1D\x00\x00\x00\x00\x00\x00\x00\x00\x01\x10\x02\x00\x00\x00\x00\x00\x00\x00\x04\xA0\x0D\x00\x04\x44\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x04\x38\x01\x01\x00\x14\x03\x00\x00\x24\x01\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x02\x80\x04\x00\x02\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x02\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x04\x00\x02\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x03\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x02\x80\x04\x00\x01\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x04\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x01\x80\x04\x00\x01\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x05\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x02\x80\x04\x00\x02\x80\x03\x00\x03\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x06\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x04\x00\x02\x80\x03\x00\x03\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x07\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x02\x80\x04\x00\x01\x80\x03\x00\x03\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x08\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x01\x80\x04\x00\x01\x80\x03\x00\x03\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x09\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x02\x80\x04\x00\x02\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x0A\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x04\x00\x02\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x0B\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x02\x80\x04\x00\x01\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x24\x0C\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x01\x80\x04\x00\x01\x80\x03\x00\x01\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x03\x00\x00\x50\x0D\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x02\x80\x04\x00\x02\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x7D\x01\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x0E\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x02\x80\x04\x00\x02\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x40\x00\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x0F\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x04\x00\x02\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x7D\x01\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x10\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x04\x00\x02\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x40\x00\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x11\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x02\x80\x04\x00\x01\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x7D\x01\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x12\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x02\x80\x04\x00\x01\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x40\x00\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x03\x00\x00\x50\x13\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x01\x80\x04\x00\x01\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x7D\x01\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x00\x00\x00\x50\x14\x01\x00\x00\x80\x01\x00\x01\x80\x02\x00\x01\x80\x04\x00\x01\x80\x03\xFD\xE9\x80\x0B\x00\x01\x00\x0C\x00\x04\x00\x00\x70\x80\x40\x00\x00\x28\x6E\x00\x6F\x00\x64\x00\x65\x00\x33\x00\x30\x00\x36\x00\x24\x00\x40\x00\x45\x00\x43\x00\x47\x00\x52\x00\x49\x00\x44\x00\x2E\x00\x4E\x00\x45\x00\x54\x00\x00\x00\x0D\x00\x00\x18\x1E\x2B\x51\x69\x05\x99\x1C\x7D\x7C\x96\xFC\xBF\xB5\x87\xE4\x61\x00\x00\x00\x02\x0D\x00\x00\x14\x40\x48\xB7\xD5\x6E\xBC\xE8\x85\x25\xE7\xDE\x7F\x00\xD6\xC2\xD3\x00\x00\x00\x14\x90\xCB\x80\x91\x3E\xBB\x69\x6E\x08\x63\x81\xB5\xEC\x42\x7B\x1F"
    _,err := conn.WriteToUDP([]byte(response), addr)
    // fmt.Printf("Success")
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 7778,
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
	h := "\x5c\x73\x74\x61\x74\x75\x73\x5c"
        m := []byte(h)
	pp := []byte(p)[0:8]
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
