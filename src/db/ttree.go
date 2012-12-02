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
}
//Tnode削除
//TODO test
//TODO リバランス
func(p *tnode) Delete(deleteNodeValue nodeValue) ROWNUM {
	if p.leftNode != nil && deleteNodeValue.key > p.maxValue()  {
		deleteNum := p.leftNode.Delete(deleteNodeValue)
		if p.leftNode.dataCount == 0 {
			p.leftNode = nil
		}
		return deleteNum
	}
	if p.rightNode != nil && deleteNodeValue.key < p.minValue() {
		deleteNum := p.rightNode.Delete(deleteNodeValue)
		if p.rightNode.dataCount == 0 {
			p.rightNode = nil
		}
		return deleteNum
	}
	isMatch,pos := p.getPosition(deleteNodeValue.key)
	if isMatch == false {
		//該当データなし
		return 0
	}
	//削除
	deleteNum := p.deleteValue(pos)
	//underflow時処理
	p.forUnderFlow()
	//TODO リバランス
	return deleteNum
}

//underflow時の処理
func(p *tnode) forUnderFlow() {
	if p.IsUnderFlow() == false{
		return
	}
	if p.IsInternalNode() {
		//GLBから値を持って来て先頭に補填する
		p.insertValue(0,p.leftNode.popMaxNode())
	} else if p.IsLeafNode() {
		//leaf
	} else {
		//half-leaf
		switch p.CanMergeChildNode() {
			case MERGE_TYPE_LEFT:
				p.mergeFromLeftNode()
			case MERGE_TYPE_RIGHT:
				p.mergeFromRightNode()
		}
	}
}
//次に挿入したらOverFlowするか
func(p *tnode) IsOverFlow() bool{
	return p.dataCount  == p.t
}
//under flowしているか
func(p *tnode) IsUnderFlow() bool{
	return p.dataCount  <= p.t-3
}
func(p *tnode) IsInternalNode() bool{
	return p.leftNode != nil && p.rightNode != nil
}
func(p *tnode) IsLeafNode() bool{
	return p.leftNode == nil && p.rightNode == nil
}
//マージできる子ノードの有無
//0なし 1左　2右　3両方
func(p *tnode) CanMergeChildNode() MergeType{
	canMerge := MERGE_TYPE_NONE
	if p.leftNode != nil && p.dataCount + p.leftNode.dataCount  <= p.t {
		canMerge += MERGE_TYPE_LEFT
	}
	if p.rightNode != nil && p.dataCount + p.rightNode.dataCount  <= p.t {
		canMerge += MERGE_TYPE_RIGHT
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
//TODOリバランス
func(p *tnode) popMaxNode() nodeValue{
	if p.rightNode != nil {
		retNode := p.rightNode.popMaxNode()
		if p.rightNode.dataCount == 0 {
			p.rightNode = nil
		}
		return retNode
	}
	retNode := p.popNodeValue(p.dataCount-1)
	p.forUnderFlow()
	return retNode
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

//左子ノードからマージ
func (p *tnode) mergeFromLeftNode(){
	p.mergeHead(p.leftNode,0)
	p.rightNode = p.leftNode.rightNode	
	p.leftNode = p.leftNode.leftNode
	
}
//右子ノードからマージ
func (p *tnode) mergeFromRightNode(){
	p.mergeTail(p.rightNode,0)
	p.leftNode = p.rightNode.leftNode
	p.rightNode = p.rightNode.rightNode	
}
//対象ノードを後ろ側にマージする
//start番目以降を移す
func (p *tnode) mergeTail(srcNode *tnode,start int){
	cnt := p.dataCount
	for i:= start; i < srcNode.dataCount ; i++{
		p.values[cnt+i-start] = srcNode.values[i]
	}
	p.dataCount = p.dataCount + srcNode.dataCount - start
}
//対象ノードを前側にマージする
//start番目以降を移す
func (p *tnode) mergeHead(srcNode *tnode,start int){
	cnt := srcNode.dataCount - start
	for i:= 0; i < p.dataCount ; i++{
		p.values[cnt+i] = p.values[i]
	}
	for i:= start; i < srcNode.dataCount ; i++{
		p.values[i-start] = srcNode.values[i]
	}
	p.dataCount = p.dataCount + srcNode.dataCount - start
}
//start番目以降をクリアする
func (p *tnode) clear(start int){
	for i:= start ;i < p.dataCount;i++ {
		p.values[i] = nodeValue{0,nil}
	}
	p.dataCount =start
}

//LLローテーション
func rotationLL(root *tnode) *tnode{
	//新たにrootになる
	newRoot := root.leftNode
	
	//左子を付け替え
	root.leftNode = newRoot.rightNode
	root.leftNode.parentNode = root
	
	//親を付け替え
	newRoot.parentNode = root.parentNode
	root.parentNode = newRoot
	newRoot.rightNode = root
	
	return newRoot
}
//RRローテーション
func rotationRR(root *tnode) *tnode{
	//新たにrootになる
	newRoot := root.rightNode
	
	//左子を付け替え
	root.rightNode = newRoot.leftNode
	root.rightNode.parentNode = root
	
	//親を付け替え
	newRoot.parentNode = root.parentNode
	root.parentNode = newRoot
	newRoot.rightNode = root
	
	return newRoot
}
//LRローテーション
func rotationLR(root *tnode) *tnode{
	//新たにrootになる
	newRoot := root.leftNode.rightNode
	//新たなleftNode
	newLeft := root.leftNode
	
	//leftNode(B)の付け替え
	newLeft.rightNode = newRoot.leftNode
	newRoot.leftNode.parentNode = newLeft
	
	//rootNode(A)の付け替え
	root.leftNode = newRoot.rightNode
	newRoot.rightNode.parentNode = root
	
	//newRoot(C)の付け替え
	newRoot.parentNode = root.parentNode
	
	newRoot.leftNode = newLeft
	newLeft.parentNode = newRoot
	
	newRoot.rightNode = root
	root.parentNode = newRoot
	
	//special lotation
	if newRoot.dataCount == 1 {
		newRoot.mergeHead(newLeft,1)
		newLeft.clear(1)
	}
	
	return newRoot
}
//RLローテーション
func rotationRL(root *tnode) *tnode{
	//新たにrootになる
	newRoot := root.rightNode.leftNode
	//新たなrightNode
	newRight := root.rightNode
	
	//rightNode(B)の付け替え
	newRight.leftNode = newRoot.rightNode
	newRoot.rightNode.parentNode = newRight
	
	//rootNode(A)の付け替え
	root.rightNode = newRoot.leftNode
	newRoot.leftNode.parentNode = root
	
	//newRoot(C)の付け替え
	newRoot.parentNode = root.parentNode
	
	newRoot.rightNode = newRight
	newRight.parentNode = newRoot
	
	newRoot.leftNode = root
	root.parentNode = newRoot
	
	//special lotation
	if newRoot.dataCount == 1 {
		newRoot.mergeTail(newRight,1)
		newRight.clear(1)
	}
	
	return newRoot
}