package main

import "github.com/labstack/echo/v4"

type EchoServer struct {
	Echo *echo.Echo
}

type NexusApp struct {
	Echo EchoServer
}

func NewNexusApp(e EchoServer) NexusApp {
	return NexusApp{
		Echo: e,
	}
}

func NewEcho() EchoServer {
	e := echo.New()
	return EchoServer{
		Echo: e,
	}
}

func (n NexusApp) Start() {
	n.Echo.Echo.Logger.Fatal(n.Echo.Echo.Start(":8080"))
}

func main() {
	e := InitializeApp()

	e.Start()
}
