package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var mConfig = map[string]string{}

var quit chan string = make(chan string)

func LoadConfig() map[string]string {
	if _, err := os.Stat("FrpNodeConfig.json"); os.IsNotExist(err) {
		log.Println("FrpNodeConfig.json file is not found!")
	}
	bConfigFile, err := ioutil.ReadFile("FrpNodeConfig.json")
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(bConfigFile, &mConfig); err != nil {
		log.Println(err)
	}
	return mConfig
}

func StartFrp(nodeName string, configFile string) {
	cmd := exec.Command("frps", "-c", configFile)
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		quit <- nodeName
		return
	}
	if err != nil {
		log.Println(err)
		quit <- nodeName
		return
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		quit <- nodeName
		return
	}
	quit <- nodeName
}

func main() {
	configJson := LoadConfig()
	Theadnum := 0
	log.Println("frpManger Server was start !")
	for k, v := range configJson {
		log.Println(k + " was start!")
		go StartFrp(k, v)
		Theadnum++
	}
	for i := 0; i < Theadnum; i++ {
		sc := <-quit
		log.Printf("%s is exit!", sc)
	}
}
