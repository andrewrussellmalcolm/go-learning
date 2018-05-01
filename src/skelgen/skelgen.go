package main

import (
	"os"

	"strings"
	"fmt"
	"flag"
	"io/ioutil"
)

var template = `

package [package-name]

import (
	"fmt"
	"google.golang.org/grpc"
)

func (dialer *Dialer) [client-name]() ([package-name].[client-name], error) {

	err := dialer.DialService("[package-name].[service-name]")
	if err != nil {
		return nil, err
	}

	dialer.RLock()
	defer dialer.RUnlock()

	if dialer.conn == nil {
		return nil, fmt.Errorf("DIALER CONNECTION IS NIL")
	}

	return [package-name].New[client-name](dialer.conn), nil
}`

var (
	protoFile = flag.String("proto", "service.proto", "The protocol file for the service")
)
func main(){
	flag.Parse()

	var packageName string
	var serviceName string
	var clientName string

	dat, err := ioutil.ReadFile(*protoFile)
	bailOnError(err)
	
	lines :=strings.Split(string(dat),"\n")

	for _, line := range lines {

		line = strings.Trim(line,";")
	
		words := strings.Split(line, " ") 

		if words[0] == "package" {
			packageName = words[1]
			
		}

		if words[0] == "service" {
			serviceName = words[1]

		}		
	}

	clientName = serviceName+"Client"
	fmt.Println("package is "+ packageName)
	fmt.Println("service is "+ serviceName)
	fmt.Println("client is "+ clientName)

	template = strings.Replace(template, "[client-name]",clientName,-1)
	template = strings.Replace(template, "[package-name]",packageName,-1)
	template = strings.Replace(template, "[service-name]",serviceName,-1)
	
	fmt.Println(template)
}

func bailOnError(err error) {

	if os.IsNotExist(err){
		 
		fmt.Println("error opening file: file does not exist: " +*protoFile)
		os.Exit(1)
	}
}