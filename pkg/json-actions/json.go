package json_actions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type fileStruct struct {
	Name    string
	Command string
	IsSudo  bool
}

func EncodeJson(name string, command string, isSudo bool) {
	path := "./.cmdexec"
	checkFileExists(path)
	path = path + "/cmdexec.json"

	eFile, err := json.Marshal(fileStruct{Name: name, Command: command, IsSudo: isSudo})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, eFile, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func DecodeJson() {
	path := "./.cmdexec/cmdexec.json"
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	file := fileStruct{}
	err = json.Unmarshal(content, &file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
func checkFileExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}
