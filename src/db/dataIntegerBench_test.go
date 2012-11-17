package db

import (
    "testing"
    "math/rand"
)
//benchmark実行コマンド
//cd src/db
//go test -bench=".*"


//データ件数
var DATA_CNT int = 10000000
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
    data1  := CreateBtree(128)
    for i:=0;i<datanumber;i++ {
        data1.Insert(Integer(i))
    }
    return data1
}

//ランダムなIntegerを生成
func randumValues(roopCnt int,upperNumber int) []Integer{
    var seachValues = []Integer{}
    for i:=0;i<roopCnt;i++ {
        seachValues = append(seachValues,Integer(rand.Intn(upperNumber)))
    }
    return seachValues
}

//Insertの計測

func BenchmarkArrayIntegerInsert(b *testing.B) {
    var data1  = ArrayInteger{}
    for i:=0;i<b.N;i++ {
        data1.Insert(Integer(i))
    }
}

//Searchの計測
func BenchmarkArrayIntegerSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataArray(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//Deleteの計測
func BenchmarkArrayIntegerDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataArray(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}

//Insertの計測
func BenchmarkBtreeIntegerInsert(b *testing.B) {
   data1  := CreateBtree(128)
    for i:=0;i<b.N;i++ {
        data1.Insert(Integer(i))
    }
}

//Searchの計測
func BenchmarkBtreeIntegerSearch(b *testing.B) {
    b.StopTimer()
    var data1  = createDataBtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Search(v)
    }
}

//Deleteの計測
func BenchmarkBtreeIntegerDelete(b *testing.B) {
    b.StopTimer()
    var data1  = createDataBtree(DATA_CNT)
    var seachValues = randumValues(b.N,DATA_CNT)
    b.StartTimer()
    for _, v := range seachValues {
        data1.Delete(v)
    }
}

