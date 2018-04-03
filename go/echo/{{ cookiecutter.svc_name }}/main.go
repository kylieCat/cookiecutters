package main

import (
	"fmt"

	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/server"
)

func main() {
	fmt.Println("Server started!")
	server.StartServer()
}
