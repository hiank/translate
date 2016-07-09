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

// Routine used to operate
type Routine interface {
	Filter() Filter

	Run(path string) RoutineChan
	End(RoutineChan)
}

// RoutineChan used to
type RoutineChan interface {
}

// Filter used to check up if the file name  match condition
type Filter interface {
	Match(s string) bool
}
