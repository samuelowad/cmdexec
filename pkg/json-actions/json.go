package json_actions

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type FileStruct struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type file struct {
	file *os.File
}

var path = "./.cmdexec/cmdexec.json"

func EncodeJson(name string, command string) {
	// Create a channel to receive the commands from the go routine
	commandsChan := make(chan []FileStruct)

	// Create a channel to receive the error from the go routine
	errChan := make(chan error)

	// Call LoadCommand function in a goroutine
	go func() {
		arr, err := LoadCommand()
		if err != nil {
			if err == io.EOF {
				arr = []FileStruct{}

			} else {
				errChan <- err
				return
			}
		}
		commandsChan <- arr
	}()

	// Wait for the commands or error to be received from the go routine
	select {
	case arr := <-commandsChan:
		commandFromFile := FindCommand(name)
		if commandFromFile != "" {
			log.Fatalf("Error: name %s already exists", name)
			return
		}

		// Append to the slice
		arr = append(arr, FileStruct{
			Name:    name,
			Command: command,
		})

		err := writeToJSONFile(arr)
		if err != nil {
			log.Println(err)
		}

	case err := <-errChan:
		log.Println(err)
	}
}

//TODO: look into making this function compactable  with changes in codebase and possibly faster
// AddCommand same as EncodeJson, but 153.131% faster
//func AddCommand(name string, command string) {
//	checkFileExists(path)
//	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	// Create a new JSON encoder
//	enc := json.NewEncoder(f)
//
//	// Encode the data as JSON and write it to the file
//	fs := FileStruct{Name: name, Command: command}
//	// convert struct to array of structs
//	value := reflect.ValueOf(fs)
//	arrayValue := reflect.MakeSlice(reflect.SliceOf(value.Type()), 0, 1)
//	arrayValue = reflect.Append(arrayValue, value)
//	fmt.Println(arrayValue)
//	err = enc.Encode(fs)
//	if err != nil {
//		panic(err)
//	}
//}

func DecodeJson() FileStruct {
	//path := "./.cmdexec/cmdexec.json"
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	file := FileStruct{}
	err = json.Unmarshal(content, &file)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func FindCommand(name string) string {
	arr, err := LoadCommand()
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range arr {
		//	check  if name exist in array
		if elem.Name == name {
			return elem.Command
		}
	}

	return ""
}

func DeleteCommand(name string) error {
	arr, err := LoadCommand()
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

	return writeToJSONFile(arr)
}

// private file functions
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

func LoadCommand() ([]FileStruct, error) {
	var arr []FileStruct
	checkFileExists(path)
	f, err := os.Open(path)
	if err != nil {
		return arr, err
	}
	defer f.Close()

	// Decode the file into a slice of FileStructs
	err = json.NewDecoder(f).Decode(&arr)
	if err != nil {
		return arr, err
	}
	return arr, nil
}

func writeToJSONFile(data []FileStruct) error {
	// Open the file in write mode using os.Create
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode the modified slice and write it to the file
	return json.NewEncoder(f).Encode(data)
}
