package main

import "fmt"
import "encoding/csv"
import "os"
import "bufio"
import "io"

type trade struct {
	code string
	price float32
	count int 
	time string
}

func main() {
	fmt.Println("Hello, world")

	csvFile, _ := os.Open("trades.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// reader := csv.NewReader(bufio.NewReader(csvFile))
    // var trades []trade
    for {
		line, error := reader.Read()
		fmt.Println("line: %s",line);
        if error == io.EOF {
            break
		} 
		break /*else if error != nil {
            log.Fatal(error)
        }*/
        // trades = append(trades, trade{
        //     code: line[0],
        //     price:  line[1],
		// 	count: line[2],
		// 	time: line[3],
        // })
    }
    // peopleJson, _ := json.Marshal(people)
    // fmt.Println(string(peopleJson))
}
