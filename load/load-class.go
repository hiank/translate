package main 

import (
	"hiank.net/translate/dict"
	"hiank.net/translate/format"
    // "hiank.net/translate/tool"
    // "strings"
	"hiank.net/translate/tool"
)

// Rout used to 
type Rout struct {
    
    
}
// Run used to
func (r *Rout) Run(path string) {
    
    t := trans.NewFClass(path)

    trans.Format(t)
    dstPath = dstDir + strings.TrimLeft(path, classDir)
    trans.SaveFile(t, dstPath)

}
// End used to
func (r *Rout) End() {
    
    
}

func main() {


}