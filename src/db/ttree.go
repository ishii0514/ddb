package db

import (

)

//Tteeデータ構造
type TtreeInteger struct{
    rootNode *tnode
    dataCount ROWNUM
    rowid ROWNUM
}
//TtreeIntegerコンストラクタ
func CreateTtree(t int) TtreeInteger {
    return TtreeInteger{rootNode:createTnode(t),dataCount:0,rowid:0}
}

//データ数
func(p *TtreeInteger) DataCount() ROWNUM{
    return p.dataCount
}

//行指定（ダミー）
func(p *TtreeInteger) Get(row ROWNUM) (Integer,error){
    return 0,nil
}
//挿入
func(p *TtreeInteger) Insert(insertValue Integer) ROWNUM{
	//TODO tnode.insert
    p.dataCount += 1
    p.rowid += 1
    return ROWNUM(1)
}


//Tツリーのノード
type tnode struct{
    //ノードサイズ
    t int
    //データ数
    dataCount int    
    //値
    values []nodeValue
    
    //parentノード
    parentNode  *tnode
    //leftノード
    leftNode  *tnode
    //rightノード
    rightNode  *tnode
}

//Tノードコンストラクタ
func createTnode(t int) *tnode{
    newNode := new(tnode)
    newNode.values = make([]nodeValue,t)
    newNode.t = t
    newNode.dataCount = 0
    newNode.parentNode = nil
    newNode.leftNode = nil
    newNode.rightNode = nil
    return newNode
}
//Tnodeインサート
//TODO test
func(p *tnode) Insert(insertValue Integer,row ROWNUM) {
	//TODO インサート後の処理
	//TODO リバランス
	//TODO 確認　子ノードがあればdataCount >0 か？
	if p.leftNode != nil && insertValue > p.maxValue()  {
		p.leftNode.Insert(insertValue,row)
	} else if p.rightNode != nil && insertValue < p.minValue() {
		p.rightNode.Insert(insertValue,row)
	} else{
		isMatch,pos := p.getPosition(insertValue)
		if isMatch == true {
			p.values[pos].rows = append(p.values[pos].rows,row)
		} else {
			//一致データなし
			//データ0件の場合も
			p.insertValue(pos,nodeValue{insertValue,[]ROWNUM{row}})
		}
	}
}
func(p *tnode) maxValue() Integer{
	if p.dataCount == 0 {
		return 0
	}
	return p.values[p.dataCount-1].key
}
func(p *tnode) minValue() Integer{
	if p.dataCount == 0 {
		return 0
	}
	return p.values[0].key
}

//ノード内に値を挿入する
func(p *tnode) insertValue(insertPos int,insertNodeValue nodeValue) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
    }
    p.values[insertPos] = insertNodeValue
    p.dataCount += 1
}

//ノード内の操作対象箇所を検索する
func(p *tnode) getPosition(searchValue Integer) (bool,int) {
    return binarySearch(p.values,searchValue,0,p.dataCount-1)
}
