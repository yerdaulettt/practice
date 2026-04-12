package app

import (
	v1 "p7/internal/contoller/http/v1"
)

func Run() {
	v1.StartServer()
}
