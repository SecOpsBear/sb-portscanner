
# Port scanner in GoLang

  This is simple Port scanner cli tool written in GO.

## A fully functional cli tool to probe a target IP address scanning for open ports!  

This project is a simple implementation of port scanner in GO.  

* Create a simple port scanner  
* Scan for open ports on the target ip address
* Scan for open ports in a given port range or all the ports
* Dynamically change the timing of the packets sent  

## How to build   

### Install in linux sb-portscanner

```bash
go install github.com/secopsbear/sb-portscanner@latest
```

### Build for linux

```bash
go build -o sb-portscanner
```


### Build for window

Generate an executable **`sb-portscanner.exe`** for windows environment.   

```bash
env GOOS=windows GOARCH=amd64 go build -o sb-portscanner.exe
```

## Example command

```bash
sb-portscanner scan 10.10.11.188 -r 20-100 -o ip188.txt
```

```bash
$ sb-portscanner -h
Port scanner to probe for all open ports on a target IP.

Usage:
  sb-portscanner [IP address] [flags]

Flags:
  -h, --help               help for sb-portscanner
  -o, --outFile string     Enter the output file name
  -r, --portRange string   Port range - all or range[1-100] format (default "1-1000")
  -p, --protcol string     Protocol to scan in tcp/udp (default "tcp")
      --proxy string       Proxy to use for requests [host:port]
  -s, --smartProbe         Sends the pack with random [0-30]milliseconds time interval to the target
  -t, --threads int        Enter the number of concurrent threads running. (default 10)

```
