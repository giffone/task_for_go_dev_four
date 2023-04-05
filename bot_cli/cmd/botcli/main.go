package main

import "flag"

func main() {
	var token string
	flag.StringVar(&token, "token", "", "telegram token")
	flag.Parse()
}
