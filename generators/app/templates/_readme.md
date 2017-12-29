

# <%- projectname%>

The programming language that has been chosen is golang.

## Install and Setup GO
Download and configure your workspace with latest version of Go and correct environment path.

### Last Go version
https://golang.org/dl/

### Windows
http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/

### Linux
http://www.tecmint.com/install-go-in-linux/

## Get Source Code
On windows:
```
cd %GOPATH%/src/github.com/ricardo-ch
```

On Linux:
```
cd $GOPATH/src/github.com/ricardo-ch
```
Create the folders "github.com" and "ricardo-ch" if not already created.

Then:
```
git clone https://github.com/ricardo-ch/<%-projectname %>
cd <%-projectname %>
```

## Run the API
```
go run main.go
```