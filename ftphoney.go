package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)


func main() {
	//-------Configure-------//
	var banner string = "220 (vsFTPd 3.0.3)"+"\n"
	var host string = "0.0.0.0:2121"
	//-------Passwords------//
	// All tested passwords are correct.
	//var ps string = "230 Password ok, continue\n"

	// No correct password.
	var ps string = "530 Incorrect password, not logged in\n"
	//----------------------//

	init_server(host,banner,ps)
}
func init_server(host string,banner string,ps string) {
	var id int = 1
	serve, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("Listenig",host)
	defer serve.Close()
	for {
		client, err := serve.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		h := client.RemoteAddr()
		println(id,"- New Connection -> ",h.String())
		go handle_conn(client, id, banner, ps)
		id++
	}
}

func handle_conn(conn net.Conn,id int, banner string, ps string) {
	conn.Write([]byte(banner))

	data, err := bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	//println(data)
	ss0 := strings.Split(data, " ")
	if (ss0[0] != "USER") {
		conn.Write([]byte("500 Command not found\n"))
	} else{
		println(id,"- User -> ",ss0[1])
		conn.Write([]byte("331 User name ok, password required\n"))
	}
	data0, err := bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	ss1 := strings.Split(data0," ")
	if (ss1[0] != "PASS") {
		conn.Write([]byte("500 Command not found\n"))
	} else {
		conn.Write([]byte(ps))
		println(id,"- Password -> ",ss1[1])
	}
}
