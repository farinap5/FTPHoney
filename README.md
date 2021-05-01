<h1 align="center">FTPHoney</h1>
<h3 align="center">Fake FTP </h3>

---
Project of a **simple** fake FTP service made in Golang.

"No log file yet".

configure in the code file.
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
### Rum
> go run ftphoney.go
### Compile and run
> go build ftphoney.go
> 
> ./ftphoney