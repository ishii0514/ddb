package db

import (
    "testing"
)

//初期化されてるか確認
func TestNode(t *testing.T) {
    testNode := node{}
    for _,v := range testNode.nodes {
        if v != nil {
            t.Error("not nil.")
        }
    }
}


//線形探索
func TestBtreeLinearSearch(t *testing.T) {
    childNode1 := node{}
    childNode1.dataCount = 1
    childNode1.values[0].key = 10
    childNode1.values[0].rows = []ROWNUM{10}
    
    childNode2 := node{}
    childNode2.dataCount = 1
    childNode2.values[0].key = 30
    childNode2.values[0].rows = []ROWNUM{100,200}

    testNode := node{}
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    testNode.nodes[1] = &childNode1
    testNode.nodes[3] = &childNode2
    
    //18の検索
    rows := testNode.Search(Integer(18))
    if len(rows) != 1 {
        t.Error("illegal search.18")
    }
    if rows[0] != 2 {
        t.Error("illegal search.18")
    }
    
    //30の検索
    rows = testNode.Search(Integer(30))
    if len(rows) != 2 {
        t.Error("illegal search.30")
    }
    if rows[1] != 200 {
        t.Error("illegal search.30")
    }
    
    //nohit 0
    rows = testNode.Search(Integer(0))
    if len(rows) != 0 {
        t.Error("illegal search.0")
    }
    
    //nohit 19
    rows = testNode.Search(Integer(19))
    if len(rows) != 0 {
        t.Error("illegal search.19")
    }
    
    //nohit 40
    rows = testNode.Search(Integer(40))
    if len(rows) != 0 {
        t.Error("illegal search.40")
    }
}
