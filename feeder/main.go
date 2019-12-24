package feeder

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type data interface{}

func Feed(index string, reader io.Reader, writer io.Writer) {
	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)

	_, err := dec.Token()

	if err != nil {
		panic(err)
	}

	for dec.More() {
		var d data

		err := dec.Decode(&d)
		if err != nil {
			panic(err)
		}

		eIndex := fmt.Sprintf(`{"index": { "_index" : "%s" } }`, index)
		fmt.Println(eIndex)

		err = enc.Encode(d)
		if err != nil {
			panic(err)
		}

	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

}
