package main 

import(
    "db"
	"testing"
)


var tests =[]testing.InternalTest{
	{"db.TestColumnNumber", db.TestColumnNumber},
	{"db.TestColumnNumberInsert", db.TestColumnNumberInsert},
	{"db.TestColumnNumberGet", db.TestColumnNumberGet},
    {"db.TestColumnNumberDeleteAt", db.TestColumnNumberDeleteAt},
	
}

func main() {
	testing.Main(func(string, string) (bool, error) { return true, nil },
		tests,
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}

