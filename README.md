
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

```console
go install github.com/secopsbear/sb-portscanner@latest
```

### Build for linux

```console
go build -o sb-portscanner
```


### Build for window

Generate an executable **`sb-portscanner.exe`** for windows environment.   

```console
env GOOS=windows GOARCH=amd64 go build -o sb-portscanner.exe
```

## Example command

```console
sb-portscanner scan 10.10.11.188 -r 20-100 -o ip188.txt
```

```console
sb-portscanner scan -h
Port scanner to probe for all open ports on a target IP.

Usage:
  sb-portscanner scan [IP address] [flags]

Flags:
  -h, --help               help for scan
  -o, --outFile string     Enter the output file name
  -r, --portRange string   Enter "all" to scan all ports or between 1 to 65535 in format 20-4000 (default "1-1000")
  -p, --protocol string    Protocol to scan in tcp/udp (default "tcp")
      --proxy string       Proxy to use for requests [host:port]
  -s, --smartProbe         Sends the pack with random [0-30]milliseconds time interval to the target
  -t, --threads int        Enter the number of concurrent threads running (default 10)
```

## Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above.