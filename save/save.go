package save

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v2"
	// "encoding/json"
	"io/ioutil"
	"os"
)

// import "github.com/kelindar/binary"

//
//
// Available formats
//
//

var Formats = map[string]SaveFormat{
	"yaml":       SaveFormat{"yaml", false, yaml.Marshal, yaml.Unmarshal, emptyMarshalerStreamer, emptyUnMarshalerStreamer},
	"yamlstream": SaveFormat{"yaml", true, emptyMarshaler, emptyUnMarshaler, yamlMarshaler, yamlUnMarshaler},
	"bson":       SaveFormat{"bson", false, bson.Marshal, bson.Unmarshal, emptyMarshalerStreamer, emptyUnMarshalerStreamer},
	// "json": SaveFormat{".json", false, json.Marshal, json.Unmarshal, emptyMarshalerStreamer, emptyUnMarshalerStreamer}, // no numerical key
	// "binary": SaveFormat{"bin", false, binary.Marshal, binary.Unmarshal, emptyMarshalerStreamer, emptyUnMarshalerStreamer}, // fail to export reader
	"gob": SaveFormat{"gob", false, emptyMarshaler, emptyUnMarshaler, gobMarshaler, gobUnMarshaler},
}

var FormatNames = []string{"yaml", "yamlstream", "bson", "gob"}

//
//
// Types
//
//

type Marshaler func(interface{}) ([]byte, error)
type UnMarshaler func([]byte, interface{}) error
type MarshalerStreamer func(string, interface{}) ([]byte, error)
type UnMarshalerStreamer func(string, interface{}) error

type SaveFormat struct {
	Extension           string
	Streamer            bool // returns bytes or write directly to stream
	Marshaler           Marshaler
	UnMarshaler         UnMarshaler
	MarshalerStreamer   MarshalerStreamer
	UnMarshalerStreamer UnMarshalerStreamer
}

func emptyMarshaler(val interface{}) ([]byte, error) {
	return []byte{}, *new(error)
}

func emptyUnMarshaler(data []byte, val interface{}) error {
	return *new(error)
}

func emptyMarshalerStreamer(filename string, val interface{}) ([]byte, error) {
	return []byte{}, *new(error)
}

func emptyUnMarshalerStreamer(filename string, val interface{}) error {
	return *new(error)
}

//
//
// General Functions
//
//

func GenFilename(outPrefix string, extension string) string {
	return outPrefix + "." + extension
}

func GetExtensionAndMarshaler(format string) (string, Marshaler, UnMarshaler) {
	sf, ok := Formats[format]

	if !ok {
		fmt.Println("Unknown format: ", format, ". valid formats are:")
		for k, _ := range Formats {
			fmt.Println(" ", k)
		}
		os.Exit(1)
	}

	return sf.Extension, sf.Marshaler, sf.UnMarshaler
}

func GetExtension(format string) (extension string) {
	extension, _, _ = GetExtensionAndMarshaler(format)
	return extension
}

func GetMarshaler(format string) (marshaler Marshaler) {
	_, marshaler, _ = GetExtensionAndMarshaler(format)
	return marshaler
}

func GetUnMarshaler(format string) (unmarshaler UnMarshaler) {
	_, _, unmarshaler = GetExtensionAndMarshaler(format)
	return unmarshaler
}

//
//
// Save
//
//

func Save(outPrefix string, format string, val interface{}) {
	SaveWithExtension(outPrefix, format, GetExtension(format), val)
}

func SaveWithExtension(outPrefix string, format string, extension string, val interface{}) {
	saveFormat(outPrefix, extension, GetMarshaler(format), val)
}

func saveFormat(outPrefix string, extension string, marshaler Marshaler, val interface{}) {
	d, err := marshaler(val)

	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	outfile := GenFilename(outPrefix, extension)
	fmt.Println("saving to ", outfile)

	err = ioutil.WriteFile(outfile, d, 0644)
	fmt.Println("  done")
}

//
//
// Load
//
//

func Load(outPrefix string, format string, val interface{}) {
	LoadWithExtension(outPrefix, format, GetExtension(format), val)
}

func LoadWithExtension(outPrefix string, format string, extension string, val interface{}) {
	loadFormat(outPrefix, extension, GetUnMarshaler(format), val)
}

func loadFormat(outPrefix string, extension string, unmarshaler UnMarshaler, val interface{}) {
	outfile := GenFilename(outPrefix, extension)

	data, err := ioutil.ReadFile(outfile)

	if err != nil {
		fmt.Printf("dump file. Get err   #%v ", err)
	}

	err = unmarshaler(data, val)

	if err != nil {
		fmt.Printf("cannot unmarshal data: %v", err)
	}
}