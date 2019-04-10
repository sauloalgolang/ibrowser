package interfaces

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

import "github.com/sauloalgolang/introgressionbrowser/save"

const IndexExtension = "ibindex"

type ChromosomeInfo struct {
	ChromosomeName string
	StartPosition  int64
	NumRegisters   int64
}

type ChromosomeNamesType struct {
	Infos          []ChromosomeInfo
	NumChromosomes int64
	StartPosition  int64
	EndPosition    int64
	NumRegisters   int64
}

func NewChromosomeNames(size int, cap int) (cn *ChromosomeNamesType) {
	cn = &ChromosomeNamesType{
		Infos: make([]ChromosomeInfo, size, cap),
	}
	return cn
}

func (cn *ChromosomeNamesType) IndexFileName(outPrefix string) (indexFile string) {
	indexFile = save.GenFilename(outPrefix, IndexExtension)
	return indexFile
}

func (cn *ChromosomeNamesType) Save(outPrefix string) {
	save.SaveWithExtension(outPrefix, "yaml", IndexExtension, cn)
}

func (cn *ChromosomeNamesType) Load(outPrefix string) {
	outfile := cn.IndexFileName(outPrefix)

	data, err := ioutil.ReadFile(outfile)

	if err != nil {
		fmt.Printf("yamlFile. Get err   #%v ", err)
	}

	err = yaml.Unmarshal(data, &cn)

	if err != nil {
		fmt.Printf("cannot unmarshal data: %v", err)
	}
}

func (cn *ChromosomeNamesType) Add(chromosomeName string, startPosition int64) {
	if !(chromosomeName == "") { // valid chromosome name
		cn.Infos = append(cn.Infos, ChromosomeInfo{
			ChromosomeName: chromosomeName,
			StartPosition:  startPosition,
			NumRegisters:   -1,
		})

	} else {
		fmt.Println("got last chromosome", cn)

		cn.NumChromosomes = int64(len(cn.Infos))
		cn.NumRegisters = 0

		for p := int64(0); p < cn.NumChromosomes-1; p++ {
			infoC := &cn.Infos[p]
			infoN := &cn.Infos[p+1]
			infoC.NumRegisters = infoN.StartPosition - infoC.StartPosition
			cn.NumRegisters += infoC.NumRegisters
		}

		cn.Infos[cn.NumChromosomes-1].NumRegisters = startPosition - cn.Infos[cn.NumChromosomes-2].StartPosition

		cn.StartPosition = cn.Infos[0].StartPosition
		cn.EndPosition = cn.Infos[cn.NumChromosomes-1].StartPosition

		fmt.Println("fixed chromosome sizes", cn)
	}
}