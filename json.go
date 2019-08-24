package gogo

import (
	"encoding/json"
	"io"
	"os"
)

func EncodeJSON(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func EncodeJSONFile(file string, v interface{}) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return EncodeJSON(f, v)
}

func DecodeJSON(r io.Reader, v interface{}) error {
	enc := json.NewDecoder(r)
	return enc.Decode(v)
}

func DecodeJSONFile(file string, v interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return DecodeJSON(f, v)
}
