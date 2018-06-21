# Command line WebSocket client

Simple tool that connects to WebSocket and pipes stdin/stdout to/from it.<br/>
Text mode messages are separated by `\n`, binary mode receives/sends bytes as-is.

![Example](https://gist.githubusercontent.com/pixelbender/46a699b2198b246fffaf28b2e336f22d/raw/ws.gif)

Install

```sh
go get github.com/pixelbender/ws
```

Example:

```sh
$ ws wss://echo.websocket.org <<< hello
Connected: wss://echo.websocket.org
hello
```

Usage:

```sh
ws [options] url
  -binary
    	Binary mode
  -origin string
    	Origin
  -protocol string
    	Protocol
```
