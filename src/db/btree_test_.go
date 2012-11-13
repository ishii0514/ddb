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
    childNode1 := createNode(128)
    childNode1.dataCount = 1
    childNode1.values[0].key = 10
    childNode1.values[0].rows = []ROWNUM{10}
    
    childNode2 := createNode(128)
    childNode2.dataCount = 1
    childNode2.values[0].key = 30
    childNode2.values[0].rows = []ROWNUM{100,200}

    testNode := createNode(128)
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    testNode.nodes[1] = childNode1
    testNode.nodes[3] = childNode2
    
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
    testNode := createNode(128)
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    match,pos := testNode.getPosition(18)
    if match != true {
        t.Error("illegal match.18")
    }
    if pos != 1 {
        t.Error("illegal position.18")
    }
    
    match,pos = testNode.getPosition(19)
    if match != false {
        t.Error("illegal match.19")
    }
    if pos != 2 {
        t.Error("illegal position.19")
    }
    
    match,pos = testNode.getPosition(40)
    if match != false {
        t.Error("illegal match.40")
    }
    if pos != 3 {
        t.Error("illegal position.40")
    }
    
    match,pos = testNode.getPosition(3)
    if match != false {
        t.Error("illegal match.3")
    }
    if pos != 0 {
        t.Error("illegal position.3")
    }
    

}

func TestGetBinarySearch(t *testing.T) {
    testNode := createNode(128)
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    
    match,pos := testNode.binarySearch(18,0,testNode.dataCount-1)
    if match != true {
        t.Error("binarySearch illegal match.18")
    }
    if pos != 1 {
        t.Error("binarySearch illegal position.18")
    }
    
    match,pos = testNode.binarySearch(19,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch illegal match.19")
    }
    if pos != 2 {
        t.Error("binarySearch illegal position.19")
    }
    
    match,pos = testNode.binarySearch(40,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch illegal match.40")
    }
    if pos != 3 {
        t.Error("binarySearch illegal position.40 pos",pos)
    }
    
    match,pos = testNode.binarySearch(3,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch illegal match.3")
    }
    if pos != 0 {
        t.Error("binarySearch illegal position.3 pos",pos)
    }
}
func TestGetBinarySearch2(t *testing.T) {
    testNode := createNode(128)
    testNode.dataCount = 5
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 10
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 15
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 20
    testNode.values[3].rows = []ROWNUM{3,5}
    testNode.values[4].key = 25
    testNode.values[4].rows = []ROWNUM{3,5}
    
    match,pos := testNode.binarySearch(18,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch2 illegal match.18")
    }
    if pos != 3 {
        t.Error("binarySearch2 illegal position.18")
    }
    
    match,pos = testNode.binarySearch(25,0,testNode.dataCount-1)
    if match != true {
        t.Error("binarySearch2 illegal match.25")
    }
    if pos != 4 {
        t.Error("binarySearch2 illegal position.25")
    }
    
    match,pos = testNode.binarySearch(40,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch2 illegal match.40")
    }
    if pos != 5 {
        t.Error("binarySearch2 illegal position.40 pos",pos)
    }
    
    match,pos = testNode.binarySearch(3,0,testNode.dataCount-1)
    if match != false {
        t.Error("binarySearch2 illegal match.3")
    }
    if pos != 0 {
        t.Error("binarySearch2 illegal position.3 pos",pos)
    }
}

func TestInsertValue(t *testing.T) {
    testNode := createNode(128)
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
    testNode := createNode(128)
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    cnode0 := createNode(128)
    cnode1 := createNode(128)
    cnode2 := createNode(128)
    cnode3 := createNode(128)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    
    newNode := createNewNode(testNode,1)
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
    testNode := createNode(128)
    testNode.dataCount = 3
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    cnode0 := createNode(128)
    cnode1 := createNode(128)
    cnode2 := createNode(128)
    cnode3 := createNode(128)
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
    testNode := createNode(128)
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    cnode0 := createNode(128)
    cnode1 := createNode(128)
    cnode2 := createNode(128)
    cnode3 := createNode(128)
    cnode4 := createNode(128)
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
    testNode := createNode(128)
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    cnode0 := createNode(128)
    cnode1 := createNode(128)
    cnode2 := createNode(128)
    cnode3 := createNode(128)
    cnode4 := createNode(128)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    testNode.nodes[4] = cnode4
    
    newRightNode := createNode(128)
    newRightNode.dataCount = 1
    newRightNode.values[0].key = 50
    newRightNode.values[0].rows = []ROWNUM{12}
    cnode00 := createNode(128)
    cnode01 := createNode(128)
    newRightNode.nodes[0] = cnode00
    newRightNode.nodes[1] = cnode01
    
    newNodeValue := nodeValue{key:48,rows: []ROWNUM{100}}
    
    newRoot := createNewRoot(newNodeValue,testNode,newRightNode)
    
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
    testNode := createNode(128)
    testNode.dataCount = 4
    testNode.values[0].key = 5
    testNode.values[0].rows = []ROWNUM{1}
    testNode.values[1].key = 18
    testNode.values[1].rows = []ROWNUM{2}
    testNode.values[2].key = 25
    testNode.values[2].rows = []ROWNUM{3,5}
    testNode.values[3].key = 40
    testNode.values[3].rows = []ROWNUM{6,8,10}
    
    cnode0 := createNode(128)
    cnode0.dataCount = 2
    cnode0.values[0].key = 50
    cnode0.values[0].rows = []ROWNUM{12}
    cnode0.values[1].key = 60
    cnode0.values[1].rows = []ROWNUM{13,14}
    
    cnode1 := createNode(128)
    cnode1.dataCount = 2
    cnode1.values[0].key = 70
    cnode1.values[0].rows = []ROWNUM{21}
    cnode1.values[1].key = 80
    cnode1.values[1].rows = []ROWNUM{22,23}
    
    cnode2 := createNode(128)
    cnode3 := createNode(128)
    cnode4 := createNode(128)
    testNode.nodes[0] = cnode0
    testNode.nodes[1] = cnode1
    testNode.nodes[2] = cnode2
    testNode.nodes[3] = cnode3
    testNode.nodes[4] = cnode4
    
    cnode10 := createNode(128)
    cnode10.dataCount = 2
    cnode10.values[0].key = 91
    cnode10.values[0].rows = []ROWNUM{12}
    cnode10.values[1].key = 92
    cnode10.values[1].rows = []ROWNUM{13,14}
    
    cnode11 := createNode(128)
    cnode11.dataCount = 2
    cnode11.values[0].key = 93
    cnode11.values[0].rows = []ROWNUM{21}
    cnode11.values[1].key = 94
    cnode11.values[1].rows = []ROWNUM{22,23}
    
    cnode1.nodes[0] = cnode10
    cnode1.nodes[1] = cnode11
    
    res := testNode.show()
    
    exp := "[5(1),18(1),25(2),40(3),]\n"
    exp += "-[50(1),60(2),]\n"
    exp += "-[70(1),80(2),]\n"
    exp += "--[91(1),92(2),]\n"
    exp += "--[93(1),94(2),]\n"
    exp += "-[]\n"
    exp += "-[]\n"
    exp += "-[]\n"
    //print(res)
    if res != exp {
        t.Error("illegal show")
    }
}

func TestInsert(t *testing.T) {
    testNode := createNode(128)
    
    //0件
    res := testNode.show()
    exp := "[]\n"
    //print(res)
    if res != exp {
        t.Error("illegal insert 0")
    }
    
    //3件
    testNode.Insert(5,1)
    testNode.Insert(15,2)
    testNode.Insert(10,3)
    res = testNode.show()
    exp = "[5(1),10(1),15(1),]\n"
    //print(res)
    if res != exp {
        t.Error("illegal insert 3")
    }
    
    //同じ値
    testNode.Insert(5,4)
    res = testNode.show()
    exp = "[5(2),10(1),15(1),]\n"
    //print(res)
    if res != exp {
        t.Error("illegal insert 4")
    }
}
func TestBtreeInsert(t *testing.T) {

    btree := CreateBtree(128)
    btree.Insert(5)
    btree.Insert(15)
    btree.Insert(10)
    //260件
    for i :=0 ; i<260 ;i++ {
        btree.Insert(Integer(i))
    }
    res := btree.show()
    //print(res)
    exp := "[128(1),]\n"
    exp += "-[0(1),1(1),2(1),3(1),4(1),5(2),6(1),7(1),8(1),9(1),10(2),11(1),12(1),13(1),14(1),15(2),16(1),17(1),18(1),19(1),20(1),21(1),22(1),23(1),24(1),25(1),26(1),27(1),28(1),29(1),30(1),31(1),32(1),33(1),34(1),35(1),36(1),37(1),38(1),39(1),40(1),41(1),42(1),43(1),44(1),45(1),46(1),47(1),48(1),49(1),50(1),51(1),52(1),53(1),54(1),55(1),56(1),57(1),58(1),59(1),60(1),61(1),62(1),63(1),64(1),65(1),66(1),67(1),68(1),69(1),70(1),71(1),72(1),73(1),74(1),75(1),76(1),77(1),78(1),79(1),80(1),81(1),82(1),83(1),84(1),85(1),86(1),87(1),88(1),89(1),90(1),91(1),92(1),93(1),94(1),95(1),96(1),97(1),98(1),99(1),100(1),101(1),102(1),103(1),104(1),105(1),106(1),107(1),108(1),109(1),110(1),111(1),112(1),113(1),114(1),115(1),116(1),117(1),118(1),119(1),120(1),121(1),122(1),123(1),124(1),125(1),126(1),127(1),]\n"
    exp += "-[129(1),130(1),131(1),132(1),133(1),134(1),135(1),136(1),137(1),138(1),139(1),140(1),141(1),142(1),143(1),144(1),145(1),146(1),147(1),148(1),149(1),150(1),151(1),152(1),153(1),154(1),155(1),156(1),157(1),158(1),159(1),160(1),161(1),162(1),163(1),164(1),165(1),166(1),167(1),168(1),169(1),170(1),171(1),172(1),173(1),174(1),175(1),176(1),177(1),178(1),179(1),180(1),181(1),182(1),183(1),184(1),185(1),186(1),187(1),188(1),189(1),190(1),191(1),192(1),193(1),194(1),195(1),196(1),197(1),198(1),199(1),200(1),201(1),202(1),203(1),204(1),205(1),206(1),207(1),208(1),209(1),210(1),211(1),212(1),213(1),214(1),215(1),216(1),217(1),218(1),219(1),220(1),221(1),222(1),223(1),224(1),225(1),226(1),227(1),228(1),229(1),230(1),231(1),232(1),233(1),234(1),235(1),236(1),237(1),238(1),239(1),240(1),241(1),242(1),243(1),244(1),245(1),246(1),247(1),248(1),249(1),250(1),251(1),252(1),253(1),254(1),255(1),256(1),257(1),258(1),259(1),]\n"
    if res != exp {
        t.Error("illegal insert 4")
    }
    
}

func TestBtreeDelete(t *testing.T) {

    btree := CreateBtree(5)
    btree.Insert(5)
    btree.Insert(15)
    btree.Insert(10)
    //260件
    for i :=0 ; i<153;i++ {
        btree.Insert(Integer(i))
    }
    if btree.dataCount != 156 {
        t.Error("illegal data count.")
    }
    res := btree.show()
    //print(res)
    exp := "[35(1),71(1),107(1),]\n"
    exp += "-[5(2),11(1),17(1),23(1),29(1),]\n"
    exp += "--[0(1),1(1),2(1),3(1),4(1),]\n"
    exp += "--[6(1),7(1),8(1),9(1),10(2),]\n"
    exp += "--[12(1),13(1),14(1),15(2),16(1),]\n"
    exp += "--[18(1),19(1),20(1),21(1),22(1),]\n"
    exp += "--[24(1),25(1),26(1),27(1),28(1),]\n"
    exp += "--[30(1),31(1),32(1),33(1),34(1),]\n"
    exp += "-[41(1),47(1),53(1),59(1),65(1),]\n"
    exp += "--[36(1),37(1),38(1),39(1),40(1),]\n"
    exp += "--[42(1),43(1),44(1),45(1),46(1),]\n"
    exp += "--[48(1),49(1),50(1),51(1),52(1),]\n"
    exp += "--[54(1),55(1),56(1),57(1),58(1),]\n"
    exp += "--[60(1),61(1),62(1),63(1),64(1),]\n"
    exp += "--[66(1),67(1),68(1),69(1),70(1),]\n"
    exp += "-[77(1),83(1),89(1),95(1),101(1),]\n"
    exp += "--[72(1),73(1),74(1),75(1),76(1),]\n"
    exp += "--[78(1),79(1),80(1),81(1),82(1),]\n"
    exp += "--[84(1),85(1),86(1),87(1),88(1),]\n"
    exp += "--[90(1),91(1),92(1),93(1),94(1),]\n"
    exp += "--[96(1),97(1),98(1),99(1),100(1),]\n"
    exp += "--[102(1),103(1),104(1),105(1),106(1),]\n"
    exp += "-[113(1),119(1),125(1),131(1),137(1),143(1),]\n"
    exp += "--[108(1),109(1),110(1),111(1),112(1),]\n"
    exp += "--[114(1),115(1),116(1),117(1),118(1),]\n"
    exp += "--[120(1),121(1),122(1),123(1),124(1),]\n"
    exp += "--[126(1),127(1),128(1),129(1),130(1),]\n"
    exp += "--[132(1),133(1),134(1),135(1),136(1),]\n"
    exp += "--[138(1),139(1),140(1),141(1),142(1),]\n"
    exp += "--[144(1),145(1),146(1),147(1),148(1),149(1),150(1),151(1),152(1),]\n"
    if res != exp {
        t.Error("illegal insert 4")
    }
    
    delCnt := btree.Delete(144)
    if delCnt != 1{
        t.Error("illegal delete count.")
    }
    if btree.dataCount != 155 {
        t.Error("illegal data count.")
    }
    res = btree.show()
    //print(res)
    
    
    delCnt = btree.Delete(10)
    if delCnt != 2 {
        t.Error("illegal delete count.")
    }
    if btree.dataCount != 153 {
        t.Error("illegal data count.")
    }
    res = btree.show()
    print(res)
    
    delCnt = btree.Delete(7)
    if delCnt != 1 {
        t.Error("illegal delete count.")
    }
    if btree.dataCount != 152 {
        t.Error("illegal data count.")
    }
    res = btree.show()
    print(res)
    
    delCnt = btree.Delete(7)
    if delCnt != 0 {
        t.Error("illegal delete count.")
    }
    if btree.dataCount != 152 {
        t.Error("illegal data count.")
    }
}