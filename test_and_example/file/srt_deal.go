package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// read file line by line
// Use os.open() function to open the file.
// Use bufio.NewScanner() function to create the file scanner.
// Use bufio.ScanLines() function with the scanner to split the file into lines.
// Then use the scanner Scan() function in a for loop to get each line and process it.
func main() {

	filePath := "./storeConfigFile/_video_id_plsjCyVNZVg.srt"
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	contents := make([]string, 0)

	i := 1
	for fileScanner.Scan() {
		content := fileScanner.Text()
		fmt.Printf("%d:%s \n", i%4, content)
		if i%4 == 3 {
			contents = append(contents, content)
		}
		i++

	}

	readFile.Close()

	// create the file
	f, err := os.Create("./store_config_file/language.srt")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer f.Close()

	for _, c := range contents {
		// write a string
		io.WriteString(f, c+"\n")
	}

}
