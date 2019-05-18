package main

import (
	"flag"
	"log"
)

var (
	configFile    = flag.String("config.file", "config.yaml", "Path to configuration file.")
	listenAddress = flag.String("port", ":8080", "Address on which to listen for github webhooks. Can also use $PORT")

	port int
)

func initConfig() {
    port, present := os.LookupEnv(key)
	if !present {
		port = *listenAddress
    } else {
        port = value
    }


}
func main() {
	flag.Parse()

    log.Println("Start")
    
    
    
    if 

}

