package core

// Value used to operate Dict's value
type Value interface {

    ToString() string
}


// Dict used to manager dictionary
type Dict interface {

    AddDict(d Dict)
    AddValue(key string, v Value)

    GetData() map[string]Value
    GetValue(key string) Value
}


// Trans used to 
type Trans interface {

    // SetConfig(tool.Config)
    // Format return nil map int file
    Format(dst string, src string, d Dict) map[string]int

    AddNil(map[string]int)
    GetNilArr(ignoreArr []string) []string
}