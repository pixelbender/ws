package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
	"bufio"
)

var (
	url, protocol, origin string
	binary                bool
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] url\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.BoolVar(&binary, "binary", false, "Binary mode")
	flag.StringVar(&origin, "origin", "", "Origin")
	flag.StringVar(&protocol, "protocol", "", "Protocol")
	flag.Parse()

	url = flag.Arg(0)
	if url == "" {
		flag.Usage()
	}
	if origin == "" {
		origin = url
	}

	conn, err := websocket.Dial(url, protocol, origin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Fprintf(os.Stderr, "Connected: %s\n", url)

	if binary {
		conn.PayloadType = websocket.BinaryFrame
		go io.Copy(conn, os.Stdin)
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			conn.Write([]byte(s.Text()))
		}
	}()
	buf := make([]byte, websocket.DefaultMaxPayloadBytes)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, string(buf[:n]))
	}
}
