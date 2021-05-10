<h1 align="center">FTPHoney</h1>
<h3 align="center">Fake FTP - Honeypot</h3>

---
Project of a **simple** fake FTP service made in Golang.

A sqlite file is generated for the insertion of the logs. 

### Configure in the code file.
```
//-------Configure-------//
var banner string = "220 (vsFTPd 3.0.3)"+"\n"
var host string = "0.0.0.0:2121"

//-------Passwords------//
// All tested passwords are correct.
//var ps string = "230 Password ok, continue\n"

// No correct password.
var ps string = "530 Incorrect password, not logged in\n"
//----------------------//
```
### Download
> go get github.com/farinap5/FTPHoney.git
### Run
> go run ftphoney.go
### Compile and run
> go build ftphoney.go
> 
> ./FTP-Honey

!Too many threads may cause memory corruption in SQLite!
