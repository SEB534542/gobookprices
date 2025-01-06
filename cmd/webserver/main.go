package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	gr "github.com/SEB534542/gobookprices/goodreads"
)

func main() {
	url := "/review/list/68156753"
	shelf := "to-read"
	log.Println("Creating libary...")
	library := gr.NewLibrary(url, shelf, gr.EbookFormats, []string{"Dutch", "English"})
	log.Println("Libary created")
	log.Println("Fetching books...")
	err := library.GetAll()
	if err != nil {
		log.Printf("error getting books from '%s': %v\n", url, err)
		return
	}
	log.Println("Books fetched.")

	log.Println("Saving to json...")
	err = saveToJSON(library, "library.json")
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Library saved to json.")
}

// saveToJson takes an interface and stores it into the filename. It returns an error.
func saveToJSON(i interface{}, fileName string) error {
	bs, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("got error marshalling: %w\nfor file:\n%+V", err, i)
	}
	//  Use below if you want JSON pretty printed
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, bs, "", "    ")

	err = os.WriteFile(fileName, prettyJSON.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error saving JSON to '%s': %w", fileName, err)
	}
	return nil
}
