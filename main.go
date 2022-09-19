package main

import (
	usecases "github.com/moohbr/WebMonitor/src/use_cases"
)

func main() {

	usecases.PingTest()

	usecases.SendMail([]string{"matheus.araujo@kukac.com.br"}, usecases.TestMail)

}
