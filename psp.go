package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":1331")
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] Server started at port 1331")
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("[*] Connection accepted....")
		data := make([]byte, 8192)
		_, err = conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		fmt.Println(reflect.TypeOf(conn))

		go connString(conn, data)
	}
}

func connString(conn net.Conn, data []byte) {
	temp := strings.Split(string(data), "\n")
	//fmt.Println("Splitlines : ", len(temp))
	//fmt.Println("Of tyep : ", reflect.TypeOf(temp))
	//fmt.Println("Printing the array...")
	var addr string
	var tt string
	var port int

	for i := 0; i < len(temp); i++ {
		if strings.Contains(temp[i], "Host: ") {
			tt = temp[i][6:(len(temp[i]) - 1)]
			fmt.Println(tt)
			fmt.Println(strings.Contains(tt, ":"))
			break
		}
	}
	if strings.Contains(tt, ":") {
		addr = strings.Split(tt, ":")[0]
		port, _ = strconv.Atoi(strings.Split(tt, ":")[1])
		fmt.Println("PORT : ", port)
		fmt.Println(addr)
	} else {
		addr = tt
		port = 80
	}
	ipAddr, err := net.LookupIP(string(addr))
	fmt.Println("IP Addres is :", ipAddr[0].String())
	if err != nil {
		panic(err)
	}

	proxyThisCunt(ipAddr[0].String(), port, data, conn)
}

func proxyThisCunt(ip string, port int, data []byte, browserConn net.Conn) {
	dialing := ip + ":" + strconv.Itoa(port)
	//fmt.Println("Dialing this : ", dialing)
	conn, err := net.Dial("tcp", dialing)

	if err != nil {
		panic(err)
	}

	conn.Write(data)

	bs := make([]byte, 8192)
	_, _ = conn.Read(bs)
	fmt.Println(string(bs))
	browserConn.Write(bs)
	conn.Close()
	browserConn.Close()
}
