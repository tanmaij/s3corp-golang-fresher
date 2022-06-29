package service

import (
	"log"
	"os"
)

func ExportCSVFile(cols string, cnt string) []byte {
	f, err := os.Create("tmp.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("tmp.csv")
	defer f.Close()
	_, err2 := f.WriteString(cols + "\n")
	_, err2 = f.WriteString(cnt)

	if err2 != nil {
		log.Fatal(err2)
	}
	b, err := os.ReadFile("tmp.csv")
	if err != nil {
		log.Fatal(err)
	}

	return b
}
