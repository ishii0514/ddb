package main 

import(
    "db"
	"testing"
)


var tests =[]testing.InternalTest{
	{"db.TestArrayIntegerInsert",db.TestArrayIntegerInsert},
	{"db.TestArrayIntegerGet",db.TestArrayIntegerGet},
	{"db.TestArrayIntegerSearch",db.TestArrayIntegerSearch},
	{"db.TestArrayIntegerSearchNoMatch",db.TestArrayIntegerSearchNoMatch},
	{"db.TestArrayIntegerDelete",db.TestArrayIntegerDelete},
	
	{"db.TestColumnInteger", db.TestColumnInteger},
	{"db.TestColumnIntegerInsert", db.TestColumnIntegerInsert},
	{"db.TestColumnIntegerGet", db.TestColumnIntegerGet},
    {"db.TestColumnIntegerSearch", db.TestColumnIntegerSearch},
    {"db.TestColumnIntegerSearchNoMatch", db.TestColumnIntegerSearchNoMatch},
    {"db.TestColumnIntegerSearchIllegalNum",db.TestColumnIntegerSearchIllegalNum},
    {"db.TestColumnInsertIllegalData",db.TestColumnInsertIllegalData},
    {"db.TestColumnConvertToInteger",db.TestColumnConvertToInteger},
    {"db.TestColumnCreateColumn",db.TestColumnCreateColumn},
    {"db.TestColumnIntegerDelete",db.TestColumnIntegerDelete},
    {"db.TestColumnIntegerDeleteIllegalNum",db.TestColumnIntegerDeleteIllegalNum},
    
    {"db.TestTableName", db.TestTableName},	
    {"db.TestTableInsert0column", db.TestTableInsert0column},
    {"db.TestGetInsertVaue",db.TestGetInsertVaue},
    {"db.TestTableAddcolumns",db.TestTableAddcolumns},
    {"db.TestTableInsert3columns",db.TestTableInsert3columns},
    
    {"db.TestNode",db.TestNode},
    {"db.TestBtreeLinearSearch",db.TestBtreeLinearSearch},
    {"db.TestGetPositionLinear",db.TestGetPositionLinear},
    {"db.TestGetBinarySearch",db.TestGetBinarySearch},
    {"db.TestGetBinarySearch2",db.TestGetBinarySearch2},
    {"db.TestInsertValue",db.TestInsertValue},
    {"db,TestCreateNewNode",db.TestCreateNewNode},
    {"db.TestClear",db.TestClear},
    {"db.TestDevideNode",db.TestDevideNode},
    {"db.TestCreateNewRoot",db.TestCreateNewRoot},
    {"db.TestShow",db.TestShow},
    {"db.TestInsert",db.TestInsert},
    {"db.TestBtreeInsert",db.TestBtreeInsert},
}

var benchmarks =[]testing.InternalBenchmark{
//    {"db.BenchmarkArrayIntegerInsert",db.BenchmarkArrayIntegerInsert},
}
func main() {
	testing.Main(func(string, string) (bool, error) { return true, nil },
		tests,
		benchmarks,
		[]testing.InternalExample{})
}

