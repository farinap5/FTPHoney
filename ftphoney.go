package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"github.com/cheynewallace/tabby"
	_"github.com/mattn/go-sqlite3"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	// All password are correct
	var apc = flag.Bool("a",false,"All tested passwords are correct.")
	// log connectons
	var lc = flag.Bool("v",false,"Show conections in verbose mode.")
	// Host
	var h = flag.String("l","0.0.0.0:2121","Local Host ip:port. ")
	var help = flag.Bool("h",false,"Help menu.")
	flag.Parse()

	if (*help == true) {
		help_m()
	} else {
		database,_ := sql.Open("sqlite3","./logs.db")
		show_options(*apc,*lc)
		setdb(database)
		hp(*lc,*h,*apc,database)
	}
}

// Honeypot starter service.
func hp(lc bool,host string,apc bool,database *sql.DB) {
	var ps string
	var banner string = "220 (vsFTPd 3.0.3)"+"\n"

	if (apc == true) {
		ps = "230 Password ok, continue\n"
	} else {
		ps = "530 Incorrect password, not logged in\n"
	}

	init_server(lc,database,host,banner,ps)
}

// Init server and wait for clients.
func init_server(lc bool,database *sql.DB,host string,banner string,ps string) {
	var id int = 1
	serve, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("- Listenig",host)
	defer serve.Close()
	for {
		client, err := serve.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		h := client.RemoteAddr()

		if (lc == true) {
			println(id,"- New Connection -> ",h.String())
		}

		go handle_conn(database,client, id, banner, ps)
		id++
	}
}
func handle_conn(database *sql.DB,conn net.Conn,id int, banner string, ps string) {
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
		//println(id,"- User -> ",ss0[1])
		conn.Write([]byte("331 User name ok, password required.\n"))
	}
	data0, err := bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	ss1 := strings.Split(data0," ")
	if (ss1[0] != "PASS") {
		conn.Write([]byte("500 Command not found.\n"))
	} else {
		conn.Write([]byte(ps))
		//println(id,"- Password -> ",ss1[1])
	}
	conn.Close()

	d := ctime()
	h := conn.RemoteAddr()
	hs := h.String()
	usr := ss0[1]
	pssk := ss1[1]
	loging(database,hs,d,usr,pssk)

}

// Set database.
// Create a new database and table if not exist,
// columns: id,datetime, ip, username, password.
func setdb(database *sql.DB) {
	println("- Setting up database.")
	state,_ := database.Prepare("CREATE TABLE IF NOT EXISTS frst (id INTEGER PRIMARY KEY,datetime TEXT,IP TEXT, username TEXT,password TEXT)")
	state.Exec()
}
// Current date and time.
func ctime() string {
	date := time.Now()
	dat := date.Format(("02/Jan/2006 15:04:05"))
	return dat
}

// Save information in database.
func loging(database *sql.DB,host string,dt string,user string,passkey string) {
	state,_ := database.Prepare("INSERT INTO frst (datetime,IP,username,password) VALUES (?,?,?,?)")
	state.Exec(dt,host,user,passkey)
}
func show_options(apc bool,lc bool) {
	if (apc == true) {
		println("- All tested pessword are correct:",apc)
	}
	if (lc == true) {
		println("- Show connections in verbose mode:",lc)
	}
}

func help_m() {
	t := tabby.New()
	t.AddHeader("COMMAND","DESCRIPTION","REQUIRED")
	t.AddLine("-l","Local host and port. ip:port","No")
	t.AddLine("-a","All tested pessword are correct.","No")
	t.AddLine("-v","Show conections in verbose mode.","No")
	t.AddLine("-h","Help menu.")
	t.Print()
}
