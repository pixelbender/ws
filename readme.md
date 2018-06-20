# Command line WebSocket client

Simple tool that connects to WebSocket and pipes stdin/stdout to/from it.<br/>
Text mode messages are separated by `\n`, binary mode receives/sends bytes as-is.

```sh
go install github.com/pixelbender/ws
```

Usage:

```sh
ws -h
Usage: ws [options] url
  -binary
    	Binary mode
  -origin string
    	Origin URL
  -protocol string
    	Protocol name
    	
ws wss://echo.websocket.org <<< hello
Connected: wss://echo.websocket.org
hello
```
