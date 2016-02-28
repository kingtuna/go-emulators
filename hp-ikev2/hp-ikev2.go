package main
import (
    "fmt"
    "net"
    "reflect"
    "time"
)




func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, spi_val []byte, slp_time int) {
    if slp_time != 0 {
        slp_time = slp_time*1000
        time.Sleep(time.Duration(slp_time) * time.Millisecond)
    }
    spi_and_ciphers := []byte("\xee\x1b\x7a\x88\x27\x22\xe5\xfa\x01\x10\x02\x00\x00\x00\x00\x00\x00\x00\x01\x08\x0d\x00\x00\x40\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x34\x01\x01\x08\x01\xee\x1b\x7a\x88\x27\x22\xe5\xfa\x00\x00\x00\x24\x02\x01\x00\x00\x80\x01\x00\x05\x80\x02\x00\x01\x80\x03\x00\x01\x80\x04\x00\x02\x80\x0b\x00\x01\x00\x0c\x00\x04\x00\x00\x70\x80\x0d\x00\x00\x14\x00\x48\xe2\x27\x0b\xea\x83\x95\xed\x77\x8d\x34\x3c\xc2\xa0\x76\x0d\x00\x00\x14\x5c\xbe\xb3\x99\xeb\x83\x5a\x7d\x7a\x2e\xb4\x95\x90\x5d\xb0\x61\x0d\x00\x00\x14\x81\x0f\xa5\x65\xf8\xab\x14\x36\x91\x05\xd7\x06\xfb\xd5\x72\x79\x0d\x00\x00\x14\x7d\x94\x19\xa6\x53\x10\xca\x6f\x2c\x17\x9d\x92\x15\x52\x9d\x56\x0d\x00\x00\x14\x90\xcb\x80\x91\x3e\xbb\x69\x6e\x08\x63\x81\xb5\xec\x42\x7b\x1f\x0d\x00\x00\x14\xcd\x60\x46\x43\x35\xdf\x21\xf8\x7c\xfd\xb2\xfc\x68\xb6\xa4\x48\x0d\x00\x00\x14\x44\x85\x15\x2d\x18\xb6\xbb\xcd\x0b\xe8\xa8\x46\x95\x79\xdd\xcc\x0d\x00\x00\x0c\x09\x00\x26\x89\xdf\xd6\xb7\x12\x00\x00\x00\x14\xaf\xca\xd7\x13\x68\xa1\xf1\xc9\x6b\x86\x96\xfc\x77\x57\x01\x00")
    // include inbound SPI in reply, trick ike-scan
    response := append(spi_val,spi_and_ciphers...)
    _,err := conn.WriteToUDP(response, addr)
    // fmt.Printf("Success")
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 500,
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
        h := "\x00\x00\x00\x00\x00\x00\x00\x00\x01\x10"
        m := []byte(h)
        spi_in := []byte(p)[0:8]
        pp := []byte(p)[8:18]
        // fmt.Printf("SPI: %x\n",spi_in)
        // fmt.Printf("HERE: %x\n",pp)
        if reflect.DeepEqual(pp, m) {
            // simulate an IKEv2 exponential backoff
            go sendResponse(ser, remoteaddr, spi_in, 0)
            go sendResponse(ser, remoteaddr, spi_in, 1)
            go sendResponse(ser, remoteaddr, spi_in, 2)
            go sendResponse(ser, remoteaddr, spi_in, 4)
            go sendResponse(ser, remoteaddr, spi_in, 8)
            go sendResponse(ser, remoteaddr, spi_in, 16)
        }
        // fmt.Println(pp,m)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        continue
    }
}
