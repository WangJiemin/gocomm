package gocomm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/itkinside/itkconfig"
)

func LoadConfig(filename string, config interface{}) error {
	return itkconfig.LoadConfig(filename, config)
}

func ReadConfig(filepath string) map[string]string {
	res := map[string]string{}
	file, err := os.Open(filepath)
	if err != nil {
		return res
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				return res
			}
			if len(line) == 0 {
				break
			}
		}

		if len(line) == 0 || line == "\r\n" {
			//break
			continue
		}

		if line[0] == '/' {
			continue
		}

		fmt.Println(line)
		i := strings.IndexAny(line, "=")
		value := strings.TrimSpace(line[i+1 : len(line)])
		res[strings.TrimSpace(line[0:i])] = value
	}
	return res
}

/////demo

// func main() {
// 	config := &Config{
// 		ZKServer: []string{},
// 	}
// 	err := LoadConfig("config.conf", config)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println("config load ok")
// 	}
// 	fmt.Printf("zkaddr:%v", config.ZKServer)
// }
