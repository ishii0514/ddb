package db

import (
    "testing"
    "math/rand"
)
//benchmark実行コマンド
//cd src/db
//go test -bench="BenchmarkBinarySearch"

var TEST_NODE_SIZE int = 256
//ランダムなIntegerを生成
func randumValuesInt(dataCnt int,upperNumber int) []int{
    var seachValues = []int{}
    for i:=0;i<dataCnt;i++ {
        seachValues = append(seachValues,rand.Intn(upperNumber))
    }
    return seachValues
}

func BenchmarkBinarySearch(b *testing.B){
    b.StopTimer()
    nodes  := make([]nodeValue,TEST_NODE_SIZE)
    for i:=0;i<TEST_NODE_SIZE;i++ {
        nodes[i] = nodeValue{key:Integer(i),rows:[]ROWNUM{}}
    }
    var seachValues = randumValues(b.N,TEST_NODE_SIZE)
    b.StartTimer()
    for _, v := range seachValues {
        binarySearch(nodes,v,0,TEST_NODE_SIZE-1)
    }
}

func BenchmarkBinarySearchInteger(b *testing.B){
    b.StopTimer()
    nodes  := make([]nodeValueInteger,TEST_NODE_SIZE)
    for i:=0;i<TEST_NODE_SIZE;i++ {
        nodes[i] = nodeValueInteger{key:Integer(i),rows:[]ROWNUM{}}
    }
    var seachValues = randumValues(b.N,TEST_NODE_SIZE)
    b.StartTimer()
    for _, v := range seachValues {
        binarySearchInteger(nodes,v,0,TEST_NODE_SIZE-1)
    }
}
func BenchmarkBinarySearchInt(b *testing.B){
    b.StopTimer()
    nodes  := make([]nodeValueInt,TEST_NODE_SIZE)
    for i:=0;i<TEST_NODE_SIZE;i++ {
        nodes[i] = nodeValueInt{key:i,rows:[]ROWNUM{}}
    }
    var seachValues = randumValuesInt(b.N,TEST_NODE_SIZE)
    b.StartTimer()
    for _, v := range seachValues {
        binarySearchInt(nodes,v,0,TEST_NODE_SIZE-1)
    }
}
func BenchmarkBinarySearchIntArray(b *testing.B){
    b.StopTimer()
    nodes  := make([]Integer,TEST_NODE_SIZE)
    for i:=0;i<TEST_NODE_SIZE;i++ {
        nodes[i] = Integer(i)
    }
    var seachValues = randumValues(b.N,TEST_NODE_SIZE)
    b.StartTimer()
    for _, v := range seachValues {
        binarySearchIntArray(nodes,v,0,TEST_NODE_SIZE-1)
    }
}