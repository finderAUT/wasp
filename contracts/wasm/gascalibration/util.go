package gascalibration

import (
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/sortkeys"
	"log"
	"os"
)

func SaveTestResultAsJSON(filepath string, results map[uint32]uint64) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	bytes, err := json.Marshal(results)
	check(err)

	_, err = f.Write(bytes)
	check(err)
}

func SaveTestResultAsCSV(filepath string, results map[uint32]uint64) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	keys := getSortedKeys(results)
	for _, key := range keys {
		_, err = f.WriteString(fmt.Sprint(key) + "," + fmt.Sprint(results[key]) + "\n")
		check(err)
	}
}

func getSortedKeys(results map[uint32]uint64) []uint32 {
	keys := make([]uint32, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}
	sortkeys.Uint32s(keys)
	return keys
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
