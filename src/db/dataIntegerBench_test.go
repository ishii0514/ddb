package db

import (
    "testing"
    "math/rand"
)
//benchmark実行コマンド
//cd src/db
//go test -bench=".*"

//BTreeノード数
var BNODE_CNT int = 128
//Treeノード数
var TNODE_CNT int = 128
//データ件数
var DATA_CNT int = 1000000

//ランダムなIntegerを生成
func randumValues(dataCnt int,upperNumber int) []Integer{
    var seachValues = []Integer{}
    for i:=0;i<dataCnt;i++ {
        seachValues = append(seachValues,Integer(rand.Intn(upperNumber)))
    }
    return seachValues
}

//指定したデータ件数のArrayを生成
func createDataArray(datanumber int) ArrayInteger {
    var dataarray  = ArrayInteger{}
    for i:=0;i<datanumber;i++ {
        dataarray.Insert(Integer(i))
    }
    return dataarray
}
//指定したデータ件数のBtreeを生成
func createDataBtree(datanumber int) BtreeInteger {
    data1  := CreateBtree(BNODE_CNT)
    for _, i := range randumValues(datanumber,datanumber) {
        data1.Insert(Integer(i))
    }
    return data1
}
//指定したデータ件数のTtreeを生成
func createDataTtreeInteger(datanumber int) TtreeInteger {
    data1  := CreateTtreeInteger(TNODE_CNT)
    for _, i := range randumValues(datanumber,datanumber) {
        data1.Insert(Integer(i))
    }
    return data1
}
//指定したデータ件数のTtreeを生成
func createDataTtree(datanumber int) Ttree {
    data1  := CreateTtree(TNODE_CNT)
    for _, i := range randumValues(datanumber,datanumber) {
        data1.Insert(Integer(i))
    }
    return data1
}
//Array Insertの計測
func BenchmarkArrayIntegerInsert(b *testing.B) {
    b.StopTimer()
    var data1  = ArrayInteger{}
    var insertValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range insertValues {
        data1.Insert(Integer(v))
    }
}

//Array Searchの計測
func BenchmarkArrayIntegerSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataArray(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//Array Deleteの計測
func BenchmarkArrayIntegerDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataArray(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}

//Btree Insertの計測
func BenchmarkBtreeIntegerInsert(b *testing.B) {
    b.StopTimer()
    var data1  = createDataBtree(0)
    var insertValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range insertValues {
        data1.Insert(Integer(v))
    }
}

//Btree  Searchの計測
func BenchmarkBtreeIntegerSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataBtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//Btree  Deleteの計測
func BenchmarkBtreeIntegerDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataBtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}
//TTreeIngerger Insertの計測
func BenchmarkTtreeIntegerInsert(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtreeInteger(0)
    var insertValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range insertValues {
        data1.Insert(Integer(v))
    }
}

//TTreeIngerger Searchの計測
func BenchmarkTtreeIntegerSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtreeInteger(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//TTreeIngerger Deleteの計測
func BenchmarkTtreeIntegerDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtreeInteger(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}


//TTree Insertの計測
func BenchmarkTtreeInsert(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtree(0)
    var insertValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range insertValues {
        data1.Insert(Integer(v))
    }
}

//TTree Searchの計測
func BenchmarkTtreeSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//TTree Deleteの計測
func BenchmarkTtreeDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataTtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}