package json_actions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type fileStruct struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type file struct {
	file *os.File
}

var path = "./.cmdexec/cmdexec.json"

func EncodeJson(name string, command string) error {
	checkFileExists(path)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decode the file into a slice of interfaces
	var arr []fileStruct
	err = json.NewDecoder(f).Decode(&arr)
	if err != nil {
		if err == io.EOF {
			arr = []fileStruct{}
		} else {
			return err
		}
	}

	commandFromFile := FindCommand(arr, name)
	if commandFromFile != "" {
		return fmt.Errorf("name %s already exists", name)
	}

	// Append to the slice
	arr = append(arr, fileStruct{
		Name:    name,
		Command: command,
	})

	// Open the file in write mode using os.Create
	f, err = os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	//Encode the modified slice and write it to the file
	return json.NewEncoder(f).Encode(arr)
}

// AddCommand same as EncodeJson, but 153.131% faster
func AddCommand(name string, command string) {
	checkFileExists(path)
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a new JSON encoder
	enc := json.NewEncoder(f)

	// Encode the data as JSON and write it to the file
	fs := fileStruct{Name: name, Command: command}
	// convert struct to array of structs
	value := reflect.ValueOf(fs)
	arrayValue := reflect.MakeSlice(reflect.SliceOf(value.Type()), 0, 1)
	arrayValue = reflect.Append(arrayValue, value)
	fmt.Println(arrayValue)
	err = enc.Encode(fs)
	if err != nil {
		panic(err)
	}
}

func DecodeJson() fileStruct {
	//path := "./.cmdexec/cmdexec.json"
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	file := fileStruct{}
	err = json.Unmarshal(content, &file)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func checkFileExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

		parts := strings.Split(path, "/")
		err := os.Mkdir(parts[0]+"/"+parts[1], os.ModePerm)
		if err != nil {
			// handle error
		}
		// Create the file or directory
		emptyFile, err := os.Create(path)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
		emptyFile.Close()

		// Create an empty array
		var arr []interface{}

		// Encode the array as JSON and write it to the file
		err = json.NewEncoder(emptyFile).Encode(arr)
		if err != nil {
			// Handle error
		}
	}
}

func FindCommand(commands []fileStruct, name string) string {
	for _, elem := range commands {
		//	check  if name exist in array
		if elem.Name == name {
			return elem.Command
		}
	}

	return ""
}

func DeleteCommand(name string) error {
	checkFileExists(path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decode the file into a slice of fileStructs
	var arr []fileStruct
	err = json.NewDecoder(f).Decode(&arr)
	if err != nil {
		return err
	}

	// Find the index of the element with the matching name
	var indexToDelete int
	for i, elem := range arr {
		if elem.Name == name {
			indexToDelete = i
			break
		}
	}

	// Remove the element from the slice
	arr = append(arr[:indexToDelete], arr[indexToDelete+1:]...)

	// Open the file in write mode using os.Create
	f, err = os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode the modified slice and write it to the file
	return json.NewEncoder(f).Encode(arr)
}
