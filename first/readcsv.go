package first

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadCsv() []Item {
	var items []Item

	pwd, _ := os.Getwd()
	csvFile, _ := os.Open(pwd + "/database/anorexia.csv")
	//println(pwd + "/../../database/fortnite-weapons.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		id, err := strconv.Atoi(line[0])
		prewt, err := strconv.ParseFloat(line[2], 64)
		postwt, err := strconv.ParseFloat(line[2], 64)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		items = append(items, Item{
			Id:     id,
			Treat:  line[1],
			Prewt:  prewt,
			Postwt: postwt,
		})
	}

	return items
}
