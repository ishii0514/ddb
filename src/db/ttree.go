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
//TODO リファクタ
//TODO リバランス
func(p *tnode) Insert(insertNodeValue nodeValue) {
	if p.leftNode != nil && insertNodeValue.key > p.maxValue()  {
		p.leftNode.Insert(insertNodeValue)
		return
	}
	if p.rightNode != nil && insertNodeValue.key < p.minValue() {
		p.rightNode.Insert(insertNodeValue)
		return
	}
	isMatch,pos := p.getPosition(insertNodeValue.key)
	if isMatch == true {
		p.values[pos].rows = append(p.values[pos].rows,insertNodeValue.rows...)
		return
	}
	
	//新規データ
	if p.IsOverFlow() == false {
		//オーバーフローなし
		p.insertValue(pos,insertNodeValue)
		return
	}

	//オーバフローする
	if pos == 0 {
		p.createLeftNode()
		p.leftNode.Insert(insertNodeValue)
	}else if pos == p.dataCount {
		p.createRightNode()
		p.rightNode.Insert(insertNodeValue)
	} else {
		//minimumを取得して左ノードに再帰的にインサート
		minNode := p.popNodeValue(0)
		p.insertValue(pos,insertNodeValue)
		if p.leftNode == nil {
			p.createLeftNode()
		}
		p.leftNode.Insert(minNode)
	}
	//TODO リバランス
}
//Tnode削除
//TODO test
//TODO リバランス
func(p *tnode) Delete(deleteNodeValue nodeValue) ROWNUM {
	if p.leftNode != nil && deleteNodeValue.key > p.maxValue()  {
		return p.leftNode.Delete(deleteNodeValue)
	}
	if p.rightNode != nil && deleteNodeValue.key < p.minValue() {
		return p.rightNode.Delete(deleteNodeValue)
	}
	isMatch,pos := p.getPosition(deleteNodeValue.key)
	if isMatch == false {
		//該当データなし
		return 0
	}
	//削除
	deleteNum := p.deleteValue(pos)
	if p.IsUnderFlow()  == false{
		return deleteNum
	}
	
	//under flow
	if p.IsInternalNode() {
		//GLBから値を持って来て先頭に補填する
		p.insertValue(0,p.leftNode.popMaxNode())
		return deleteNum
	} else if p.IsLeafNode() {
		//leaf
		//TODO 0件なら本ノード削除して(5)へ
		return deleteNum
	} else {
		//half-leaf
		//TODO 子ノードとマージ可能ならマージして(5)へ
		return deleteNum
	}
	//TODO リバランス
	return 0
}
func(p *tnode) IsOverFlow() bool{
	return p.dataCount  == p.t
}
func(p *tnode) IsUnderFlow() bool{
	return p.dataCount  < p.t-3
}
func(p *tnode) IsInternalNode() bool{
	return p.leftNode != nil && p.rightNode != nil
}
func(p *tnode) IsLeafNode() bool{
	return p.leftNode == nil && p.rightNode == nil
}
//マージできる子ノードの有無
//0なし 1左　2右　3両方
func(p *tnode) CanMergeChildNode() int{
	canMerge := 0
	if p.leftNode != nil && p.dataCount + p.leftNode.dataCount  < p.t {
		canMerge += 1
	}
	if p.rightNode != nil && p.dataCount + p.rightNode.dataCount  < p.t {
		canMerge += 2
	}
	return canMerge
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
//指定ポジションをpop
func(p *tnode) popNodeValue(pos int) nodeValue{
	minNodeValue := p.values[pos]
	p.deleteValue(pos)
	return minNodeValue
}
//最大値をpop
func(p *tnode) popMaxNode() nodeValue{
	if p.rightNode != nil {
		return p.rightNode.popMaxNode()
	}
	return p.popNodeValue(p.dataCount-1)
}
//ノード内に値を挿入する
//TODO test　挿入位置がずれないか
func(p *tnode) insertValue(insertPos int,insertNodeValue nodeValue) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
    }
    p.values[insertPos] = insertNodeValue
    p.dataCount += 1
}
//ノード内の値を削除する
//TODO test　位置が正しくスライドするか
func(p *tnode) deleteValue(deletePos int) ROWNUM {
    rows := len(p.values[deletePos].rows)
    for i:= deletePos ; i < p.dataCount-1;i++ {
        p.values[i] = p.values[i+1]
    }
    //初期化
    p.values[p.dataCount-1] = nodeValue{}   
    p.dataCount -= 1
    return ROWNUM(rows)
}
//ノード内の操作対象箇所を検索する
func(p *tnode) getPosition(searchValue Integer) (bool,int) {
    return binarySearch(p.values,searchValue,0,p.dataCount-1)
}
//左ノードを作る
func(p *tnode) createLeftNode(){
    p.leftNode = createTnode(p.t)
    p.leftNode.parentNode = p
}
//右ノードを作る
func(p *tnode) createRightNode(){
    p.rightNode = createTnode(p.t)
    p.rightNode.parentNode = p
}