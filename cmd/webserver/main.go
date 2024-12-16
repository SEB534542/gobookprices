// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	goodreads "github.com/SEB534542/gobookprices/goodreads"
// )

// func main() {
// 	url := "/review/list/68156753"
// 	fmt.Println("Fetching books...")
// 	books, err := goodreads.GetBooks(url, "to-read")
// 	if err != nil {
// 		fmt.Printf("error getting books from '%s': %v", url, err)
// 	}

	

// 	fmt.Println("Saving to json...")
// 	err = saveToJSON(books, "test.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// // saveToJson takes an interface and stores it into the filename. It returns an error.
// func saveToJSON(i interface{}, fileName string) error {
// 	bs, err := json.Marshal(i)
// 	if err != nil {
// 		return fmt.Errorf("got error marshalling: %w\nfor file:\n%+V", err, i)
// 	}
// 	//  Use below if you want JSON pretty printed
// 	var prettyJSON bytes.Buffer
// 	_ = json.Indent(&prettyJSON, bs, "", "    ")

// 	err = os.WriteFile(fileName, prettyJSON.Bytes(), 0644)
// 	if err != nil {
// 		return fmt.Errorf("error saving JSON to '%s': %w", fileName, err)
// 	}
// 	return nil
// }
