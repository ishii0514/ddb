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


func TestGetPositionLinear(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    match,pos := testNode.getPositionLinear(18)
    if match != true {
        t.Error("illegal match.18")
    }
    if pos != 1 {
        t.Error("illegal position.18")
    }
    
    match,pos = testNode.getPositionLinear(19)
    if match != false {
        t.Error("illegal match.19")
    }
    if pos != 2 {
        t.Error("illegal position.19")
    }
    
    match,pos = testNode.getPositionLinear(40)
    if match != false {
        t.Error("illegal match.40")
    }
    if pos != 3 {
        t.Error("illegal position.40")
    }
    

}

func TestInsertValue(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    testNode.insertValue(1,nodeValue{key:10,rows:[]ROWNUM{30}},nil)
    if testNode.dataCount != 4 {
        t.Error("illegal data count.")
    }
    
    if testNode.values[0].key  != 5 {
        t.Error("illegal data key.0")
    }
    if testNode.values[0].rows[0] != 1 {
        t.Error("illegal data rows.0")
    }
    
    if testNode.values[1].key  != 10 {
        t.Error("illegal data key.1")
    }
    if testNode.values[1].rows[0] != 30 {
        t.Error("illegal data rows.1")
    }
    
    if testNode.values[2].key  != 18 {
        t.Error("illegal data key.2")
    }
    if testNode.values[2].rows[0] != 2 {
        t.Error("illegal data rows.2")
    }
    if testNode.values[3].key  != 25 {
        t.Error("illegal data key.3")
    }
    if testNode.values[3].rows[0] != 3 {
        t.Error("illegal data rows.3")
    }
}

func TestCreateNewNode(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    cnode0 := new(node)
    cnode1 := new(node)
    cnode2 := new(node)
    cnode3 := new(node)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    
    newNode := createNewNode(&testNode,1)
    if newNode.dataCount  != 1 {
        t.Error("illegal data count.")
    }
    
    //values 0番目
    if newNode.values[0].key != 25 {
        t.Error("illegal data key[0]")
    }
    if newNode.values[0].rows[0] != 3 {
        t.Error("illegal data rows[0]")
    }
    
    //values 1番目
    if newNode.values[1].key != 0 {
        t.Error("illegal data key[1]")
    }
    if newNode.values[1].rows != nil {
        t.Error("illegal data rows[1]")
    }
    
    //nodes 0番目
    if newNode.nodes[0] == nil {
        t.Error("illegal data nodes[0]")
    }
    //nodes 0番目
    if newNode.nodes[1] == nil {
        t.Error("illegal data nodes[1]")
    }
    //nodes 0番目
    if newNode.nodes[2] != nil {
        t.Error("illegal data nodes[2]")
    }
   
}
func TestClear(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    cnode0 := new(node)
    cnode1 := new(node)
    cnode2 := new(node)
    cnode3 := new(node)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    
    testNode.clear(1)
    if testNode.dataCount  != 1 {
        t.Error("illegal data count.")
    }
    
    if testNode.dataCount  != 1 {
        t.Error("illegal data count.")
    }
    if testNode.values[0].key  != 5 {
        t.Error("illegal data key[0].")
    }
    if testNode.values[1].key  != 0 {
        t.Error("illegal data key[1].")
    }
    if testNode.values[0].rows == nil {
        t.Error("illegal data rows[0]")
    }
    if testNode.values[1].rows != nil {
        t.Error("illegal data rows[1]")
    }
    
    //nodes 0番目
    if testNode.nodes[0] == nil {
        t.Error("illegal data nodes[0]")
    }
    //nodes 0番目
    if testNode.nodes[1] == nil {
        t.Error("illegal data nodes[1]")
    }
    //nodes 0番目
    if testNode.nodes[2] != nil {
        t.Error("illegal data nodes[2]")
    }
    
}

func TestDevideNode(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    cnode0 := new(node)
    cnode1 := new(node)
    cnode2 := new(node)
    cnode3 := new(node)
    cnode4 := new(node)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    testNode.nodes[4] = cnode4
    
    returnNodeValue,newNode := testNode.devideNode(2)
    if returnNodeValue.key != 25 {
        t.Error("illegal returnNodeValue.key")
    }
    if len(returnNodeValue.rows) != 2 {
        t.Error("illegal returnNodeValue.rows")
    }
    
    //旧ノード
    if testNode.dataCount != 2 {
        t.Error("illegal testNode dataCount")
    }
    if testNode.values[1].key != 18 {
        t.Error("illegal testNode [1]")
    }
    if testNode.values[2].key !=  0 {
        t.Error("illegal testNode [2]")
    }
    
    //新規ノード
    if newNode.dataCount != 1 {
        t.Error("illegal newNode dataCount")
    }
    if newNode.values[0].key != 40 {
        t.Error("illegal newtNode [0]")
    }
    if len(newNode.values[0].rows) != 3 {
        t.Error("illegal newtNode rows[0]")
    }
    
    if newNode.values[1].key !=  0 {
        t.Error("illegal newNode [1]")
    }
    if newNode.values[1].rows != nil {
        t.Error("illegal newtNode rows[1]")
    }

}
func TestCreateNewRoot(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    cnode0 := new(node)
    cnode1 := new(node)
    cnode2 := new(node)
    cnode3 := new(node)
    cnode4 := new(node)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    testNode.nodes[4] = cnode4
    
    newRightNode := node{}
    newRightNode.dataCount = 1
    newRightNode.values[0].key = 50
    newRightNode.values[0].rows = []ROWNUM{12}
    cnode00 := new(node)
    cnode01 := new(node)
    newRightNode.nodes[0] = cnode00
    newRightNode.nodes[1] = cnode01
    
    newNodeValue := nodeValue{key:48,rows: []ROWNUM{100}}
    
    newRoot := createNewRoot(newNodeValue,&testNode,&newRightNode)
    
    //rootnode
    if newRoot.dataCount != 1 {
        t.Error("illegal newRoot dataCount")
    }
    if newRoot.values[0].key != 48 {
        t.Error("illegal newtNode [0]")
    }
    if len(newRoot.values[0].rows) != 1 {
        t.Error("illegal newtNode rows[0]")
    }
    if newRoot.nodes[0] == nil {
        t.Error("illegal newtNode nodes[0]")
    }
    if newRoot.nodes[1] == nil {
        t.Error("illegal newtNode nodes[1]")
    }
    if newRoot.nodes[2] != nil {
        t.Error("illegal newtNode nodes[2]")
    }
    
    //leftnode
    if newRoot.nodes[0].dataCount != 4 {
        t.Error("illegal leftNode dataCount")
    }
    if newRoot.nodes[0].values[0].key != 5 {
        t.Error("illegal leftNode [0]")
    }
    if len(newRoot.nodes[0].values[0].rows) != 1 {
        t.Error("illegal leftNode rows[0]")
    }
}

func TestShow(t *testing.T) {
    testNode := node{}
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    
    cnode0 := new(node)
    cnode0.dataCount = 2
    cnode0.values[0].key = 50
    cnode0.values[0].rows = []ROWNUM{12}
    cnode0.values[1].key = 60
    cnode0.values[1].rows = []ROWNUM{13,14}
    
    cnode1 := new(node)
    cnode1.dataCount = 2
    cnode1.values[0].key = 70
    cnode1.values[0].rows = []ROWNUM{21}
    cnode1.values[1].key = 80
    cnode1.values[1].rows = []ROWNUM{22,23}
    
    cnode2 := new(node)
    cnode3 := new(node)
    cnode4 := new(node)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    testNode.nodes[4] = cnode4
    
    cnode10 := new(node)
    cnode10.dataCount = 2
    cnode10.values[0].key = 91
    cnode10.values[0].rows = []ROWNUM{12}
    cnode10.values[1].key = 92
    cnode0.values[1].rows = []ROWNUM{13,14}
    
    cnode11 := new(node)
    cnode11.dataCount = 2
    cnode11.values[0].key = 93
    cnode11.values[0].rows = []ROWNUM{21}
    cnode11.values[1].key = 94
    cnode11.values[1].rows = []ROWNUM{22,23}
    
    cnode1.nodes[0] = cnode10
    cnode1.nodes[1] = cnode11
    
    testNode.show()
}