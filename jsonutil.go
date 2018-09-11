package gocomm

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"fmt"
)

func JsonDecodeS(jstr string, obj interface{}) {
	json.Unmarshal([]byte(jstr), obj)
}

func JsonDecodeB(value []byte,obj interface{}){
	fmt.Printf("jsondecodeB:%s\n",string(value))
	json.Unmarshal(value,obj);
}

func JsonEncodeS(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

func JsonEncodeB(obj interface{}) []byte {
	res, _ := json.Marshal(obj)
	return res
}

func JsonDecodeSimple(jstr string) *simplejson.Json {
	js, err := simplejson.NewJson([]byte(jstr))
	if err == nil {
		return js
	} else {
		return nil
	}
}
