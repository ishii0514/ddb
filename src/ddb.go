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
    {"db.TestColumnIntegerSearch", db.TestColumnIntegerSearch},
    {"db.TestColumnIntegerSearchNoMatch", db.TestColumnIntegerSearchNoMatch},
    {"db.TestColumnInsertByString",db.TestColumnInsertByString},
    {"db.TestColumnConvertToInteger",db.TestColumnConvertToInteger},
    {"db.TestColumnCreateColumn",db.TestColumnCreateColumn},
    
    {"db.TestTableName", db.TestTableName},	
    {"db.TestTableInsert0column", db.TestTableInsert0column},
    {"db.TestGetInsertVaue",db.TestGetInsertVaue},
    {"db.TestTableAddcolumns",db.TestTableAddcolumns},
    {"db.TestTableInsert3columns",db.TestTableInsert3columns},
}

func main() {
	testing.Main(func(string, string) (bool, error) { return true, nil },
		tests,
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}

