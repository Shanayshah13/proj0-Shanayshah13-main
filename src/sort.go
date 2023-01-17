package main
import "fmt"
import (
	"log"
	"os"
	"io"
	"sort"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}
	readPath := os.Args[1]					//  first argument is for input file
	
	readFile, err := os.Open(readPath)
	if err != nil {
		log.Println("Error opening file: ", err)
	}
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])

	recordStream := [][]byte{}
	
	for {
		// record := make(map[byte] byte) 
		record := make([]byte, 100)  // making a slice
		n, err := readFile.Read(record);
		if err != nil {
			if err == io.EOF{
				break
			}
			log.Println("Error reading file: ", err)
		}
		record = record[:n] 
		recordStream = append(recordStream, record)
	}
	
	readFile.Close()
	sort.Slice(recordStream, func(i, j int) bool {return string(recordStream[i][:10]) < string(recordStream[j][:10])}) 
	
	fmt.Println("Sort key according to their names:")
    fmt.Println(recordStream)

	// writing into file 
	writePath := os.Args[2]						// second argument is for outputfile
	writeFile, err := os.Create(writePath)
	if err != nil {
		log.Println("Error opening writefile: ", err)
	}

	for _, record := range recordStream {
		writeFile.Write(record)
	}
	writeFile.Close()



}
