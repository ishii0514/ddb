package main 

import(
    "db"
	"testing"
)


var tests =[]testing.InternalTest{
	{"db.TestColumnInteger", db.TestColumnInteger},
	{"db.TestColumnIntegerInsert", db.TestColumnIntegerInsert},
	{"db.TestColumnIntegerGet", db.TestColumnIntegerGet},
    {"db.TestColumnIntegerDeleteAt", db.TestColumnIntegerDeleteAt},
	
}

func main() {
	testing.Main(func(string, string) (bool, error) { return true, nil },
		tests,
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}

