<h1 align="center">FTPHoney</h1>
<h3 align="center">Fake FTP - Honeypot</h3>

---
Project of a **simple** fake FTP service made in Golang.

A sqlite file is generated for the insertion of the logs.

!Too many threads may cause memory corruption in SQLite!

#### Help
```
COMMAND  DESCRIPTION                       REQUIRED
-------  -----------                       --------
-l       Local host and port. ip:port      No
-a       All tested password are correct.  No
-v       Show connections in verbose mode.  No
-h       Help menu.
```
### Download
> go get github.com/farinap5/FTPHoney
### Run
> go run ftphoney.go
### Compile and run
> go build ftphoney.go
> 
> ./ftphoney

