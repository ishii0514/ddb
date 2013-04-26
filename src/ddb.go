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
    {"db.TestDeleteValue",db.TestDeleteValue},
    {"db.TestCreateNewRoot",db.TestCreateNewRoot},
    {"db.TestShow",db.TestShow},
    {"db.TestInsert",db.TestInsert},
    {"db.TestBtreeInsert",db.TestBtreeInsert},
    {"db.TestBtreeDelete",db.TestBtreeDelete},
    {"db.TestCanMergeChildNode",db.TestCanMergeChildNode},
    {"db.TestInsertValueTInteger",db.TestInsertValueTInteger},
    {"db.TestInsertValueT2Integer",db.TestInsertValueT2Integer},
    {"db.TestDeleteValueTInteger",db.TestDeleteValueTInteger},
    {"db.TestMaxMinValueInteger",db.TestMaxMinValueInteger},
    {"db.TestPopNodeValueInteger",db.TestPopNodeValueInteger},
    {"db.TestMergeFromLeftNodeInteger",db.TestMergeFromLeftNodeInteger},
    {"db.TestMergeFromRightNodeInteger",db.TestMergeFromRightNodeInteger},
    {"db.TestMergeTailInteger",db.TestMergeTailInteger},
    {"db.TestMergeHeadInteger",db.TestMergeHeadInteger},
    {"db.TestClearTInteger",db.TestClearTInteger},
    {"db.TestRotationLLInteger",db.TestRotationLLInteger},
    {"db.TestRotationRRInteger",db.TestRotationRRInteger},
    {"db.TestRotationLRInteger",db.TestRotationLRInteger},
    {"db.TestRotationRLInteger",db.TestRotationRLInteger},
    {"db.TestGetPositionTInteger",db.TestGetPositionTInteger},
    {"db.TestShowTInteger",db.TestShowTInteger},
    {"db.TestDepthInteger",db.TestDepthInteger},
    {"db.TestTreeInsertInteger",db.TestTreeInsertInteger},
    {"db.TestTreeDeleteRRLotationInteger",db.TestTreeDeleteRRLotationInteger},
    {"db.TestTreeDeleteLRLotationInteger",db.TestTreeDeleteLRLotationInteger},
    {"db.TestDeleteValueT",db.TestDeleteValueT},
    {"db.TestMaxMinValue",db.TestMaxMinValue},
    {"db.TestPopNodeValue",db.TestPopNodeValue},
    {"db.TestMergeFromLeftNode",db.TestMergeFromLeftNode},
    {"db.TestMergeFromRightNode",db.TestMergeFromRightNode},
    {"db.TestMergeTail",db.TestMergeTail},
    {"db.TestMergeHead",db.TestMergeHead},
    {"db.TestClearT",db.TestClearT},
    {"db.TestRotationLL",db.TestRotationLL},
    {"db.TestRotationRR",db.TestRotationRR},
    {"db.TestRotationLR",db.TestRotationLR},
    {"db.TestRotationRL",db.TestRotationRL},
    {"db.TestGetPositionT",db.TestGetPositionT},
    {"db.TestShowT",db.TestShowT},
    {"db.TestDepth",db.TestDepth},
    {"db.TestTreeInsert",db.TestTreeInsert},
    {"db.TestTreeDeleteRRLotation",db.TestTreeDeleteRRLotation},
    {"db.TestTreeDeleteLRLotation",db.TestTreeDeleteLRLotation},


    {"db.TestTypeComp",db.TestTypeComp},
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

