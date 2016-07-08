package trans

import (

    "os"
	"fmt"
)

// Data was the little cell to operate translate
type Data interface {

    Load(f *os.File)
    // GetData() 
    // GetData() []Item
}


// LoadData used to load data in item
func LoadData(d Data, path string) {

    f, err := os.Open(path)
    if err != nil {

        fmt.Printf("when load %v : %v\n", path, err.Error())
        return
    }
    defer f.Close()

    d.Load(f)

}


