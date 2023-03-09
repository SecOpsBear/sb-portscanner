/*
Copyright © 2023 SecOpsBear bear@secopsbear.com
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/proxy"
)

var (
	ipAddress   string
	threadCount int
	portRange   string

	smartProbe bool

	startPort int
	endPort   int

	scanProtocol string
	proxyURL     string
	outFile      string
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [IP address]",
	Short: "Scan the all ports or range of ports",
	Long:  `Scan the all ports or range of ports.`,
	Args:  cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := checkIPAddress(args[0]); err != nil {
			return err
		}
		ipAddress = args[0]
		rangeData := strings.Split(portRange, "-")

		// checking for range format and valid ports
		if len(rangeData) == 2 {
			var err error
			startPort, err = strconv.Atoi(rangeData[0])
			if err != nil {
				return fmt.Errorf("enter range in the format 15-2000")
			}
			endPort, err = strconv.Atoi(rangeData[1])
			if err != nil {
				return fmt.Errorf("enter range in the format 15-2000")
			}
			if !(startPort >= 1 && startPort <= 65535) {
				return fmt.Errorf("%d port is not valid. It should be between 1-65535", startPort)
			}
			if !(endPort >= 1 && endPort <= 65535) {
				return fmt.Errorf("%d port is not valid. It should be between 1-65535", endPort)
			}
		} else {
			if strings.ToLower(portRange) == "all" {
				startPort = 1
				endPort = 65535
			} else {

				return fmt.Errorf("enter range in the format 15-2000")
			}
		}
		proto := [2]string{"tcp", "udp"}

		if !contains(proto, strings.ToLower(scanProtocol)) {
			return fmt.Errorf("protocol: select either tcp or udp")
		}

		if len(proxyURL) == 0 {
			proxyURL = ""
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Scanning %s for open ports between %d-%d\n", ipAddress, startPort, endPort)

		ports := make(chan int, threadCount)
		results := make(chan int)

		var openPorts []int
		for i := 0; i <= cap(ports); i++ {
			go worker(ports, results)
		}

		go func() {
			for i := startPort; i <= endPort; i++ {
				ports <- i
			}
		}()

		for i := startPort; i <= endPort; i++ {

			port := <-results
			if port != 0 {
				fmt.Printf("%d port is open\n", port)
				openPorts = append(openPorts, port)
			}

		}

		close(ports)
		close(results)
		sort.Ints(openPorts)
		op := strings.Join(strings.Fields(fmt.Sprint(openPorts)), ",")
		opData := op[1 : len(op)-1]
		fmt.Printf("List of open ports : %s", opData)
		if outFile != "" {
			outFilefile, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, 0666)
			printFatalError(err)
			defer outFilefile.Close()
			outFilefile.WriteString("Scan IP address: " + ipAddress + "\n")
			outFilefile.WriteString("List of open ports : " + opData + "\n")

		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().IntVarP(&threadCount, "threads", "t", 10, "Enter the number of concurrent threads running.")
	scanCmd.PersistentFlags().StringVarP(&portRange, "portRange", "r", "1-1000", "Port range - all or range[1-100] format")
	scanCmd.PersistentFlags().StringVarP(&scanProtocol, "protcol", "p", "tcp", "Protocol to scan in tcp/udp")
	scanCmd.PersistentFlags().StringVarP(&proxyURL, "proxy", "", "", "Proxy to use for requests [host:port]")
	scanCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "Enter the output file name")
	scanCmd.PersistentFlags().BoolVarP(&smartProbe, "smartProbe", "s", false, "Sends the pack with random [0-30]milliseconds time interval to the target")
}

func worker(ports, results chan int) {

	// startTime := time.Now()
	// Using for loop to loop through the channel input and which ever goroutine is available
	// it access the channels buffer capacity
	for p := range ports {

		if smartProbe {
			rand.Seed(time.Now().UnixNano())
			min := 2
			max := 200
			val := (rand.Intn(max-min+1) + min)
			time.Sleep(time.Millisecond * time.Duration(val))

		}

		if proxyURL != "" {
			dailer, err := proxy.SOCKS5("tcp", proxyURL, nil, proxy.Direct)
			if err != nil {
				fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
				os.Exit(1)
			}
			proxyConns, err := dailer.Dial(scanProtocol, ipAddress+":"+strconv.Itoa(p))
			if err != nil {
				results <- 0
				continue
			}
			proxyConns.Close()
		} else {
			conn, err := net.Dial(scanProtocol, ipAddress+":"+strconv.Itoa(p))

			if err != nil {
				results <- 0
				continue
			}

			conn.Close()
		}
		results <- p

	}
}

// checkIPAddress check for valid ipaddress
func checkIPAddress(ip string) (string, error) {
	if net.ParseIP(ip) == nil {
		return ip, fmt.Errorf("enter valid ip address")
	}
	return ip, nil
}

func contains(arr [2]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func printFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
