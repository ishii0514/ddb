package db

import (
    "testing"
)

func BenchmarkArrayIntegerInsert(b *testing.B) {
    var data1  = ArrayInteger{}
    for i:=0;i<b.N;i++ {
        data1.Insert(1)
    }
}
