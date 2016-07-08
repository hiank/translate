package tool

import (
    "encoding/json"
	"io/ioutil"
	"fmt"
)

// DictConfig used to 
type DictConfig struct {
    
    CSVDir  []string    `json:"csv"`
    XMLDir  []string    `json:"xml"`
    EXDir   []string    `json:"ex"`
}

// 
type DictIgnore struct {

    Ignore  bool        `json:"true"`
    Array   []string    `json:"array"`
}

// Config used to
type Config struct {
    
    DictDir DictConfig  `json:"dir.dict"`
    DstDir  string      `json:"dir.dst"`
    SrcDir  string      `json:"dir.src"`
    Type    string      `json:"type"`

    DIgnore DictIgnore  `json:"dict.ignore"`
}

// LoadConfig used to load json config 
func LoadConfig(c *Config, path string) {
    
    content, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("cann't load the config from " + path)
        return
    }
    
    err = json.Unmarshal(content, c)
    if err != nil {
        fmt.Println("load config error : " + err.Error())
        return
    }
    
}