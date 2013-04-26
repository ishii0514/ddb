package db

import (
    "strings"
    "strconv"
)

//Tteeデータ構造
type TtreeInteger struct{
    rootNode *tnodeInteger
    dataCount ROWNUM
    rowid ROWNUM
}
//TtreeIntegerコンストラクタ
func CreateTtreeInteger(t int) TtreeInteger {
    return TtreeInteger{rootNode:createTnodeInteger(t),dataCount:0,rowid:0}
}

//データ数
func(p *TtreeInteger) DataCount() ROWNUM{
    return p.dataCount
}

//行指定（ダミー）
func(p *TtreeInteger) Get(row ROWNUM) (Integer,error){
    return 0,nil
}
//探索
func(p *TtreeInteger) Search(searchValue Integer) []ROWNUM{
    return p.rootNode.Search(searchValue)
}
//挿入
func(p *TtreeInteger) Insert(insertValue Integer) ROWNUM{
	_,p.rootNode = p.rootNode.Insert(nodeValueInteger{insertValue,[]ROWNUM{p.rowid}})
    p.dataCount += 1
    p.rowid += 1
    return ROWNUM(1)
}
func(p *TtreeInteger) Delete(deleteValue Integer) ROWNUM{
	var deleteRows ROWNUM = 0
    deleteRows,_,p.rootNode = p.rootNode.Delete(deleteValue)
    p.dataCount = p.dataCount - deleteRows
    return deleteRows
}
//探索
func(p *TtreeInteger) Show() string{
    return p.rootNode.Show()
}
//Tツリーのノード
type tnodeInteger struct{
    //ノードサイズ
    t int
    //データ数
    dataCount int
    //値
    values []nodeValueInteger

    //parentノード
    parentNode  *tnodeInteger
    //leftノード
    leftNode  *tnodeInteger
    //rightノード
    rightNode  *tnodeInteger
}

//Tノードコンストラクタ
func createTnodeInteger(t int) *tnodeInteger{
    newNode := new(tnodeInteger)
    newNode.values = make([]nodeValueInteger,t)
    newNode.t = t
    newNode.dataCount = 0
    newNode.parentNode = nil
    newNode.leftNode = nil
    newNode.rightNode = nil
    return newNode
}
//探索
//TODO test
func(p *tnodeInteger) Search(searchValue Integer) []ROWNUM{
	//再帰なし版
	node := p
	for ; ; {
		if node.leftNode != nil && searchValue < node.minValue()  {
			node = node.leftNode
			continue
		}
		if node.rightNode != nil && searchValue > node.maxValue() {
			node = node.rightNode
			continue
		}
		break
	}
	isMatch,pos := node.getPosition(searchValue)
	if isMatch == true {
		return node.values[pos].rows
	}
	return []ROWNUM{}
}
func(p *tnodeInteger) Search_(searchValue Integer) []ROWNUM{
    if p.leftNode != nil && searchValue < p.minValue()  {
		return p.leftNode.Search(searchValue)
	}
	if p.rightNode != nil && searchValue > p.maxValue() {
		return p.rightNode.Search(searchValue)
	}
	isMatch,pos := p.getPosition(searchValue)
	if isMatch == true {
		return p.values[pos].rows
	}
	return []ROWNUM{}
}
//Tnodeインサート
//戻り値　ノード追加発生,新たなルートノード
//TODO リファクタ
func(p *tnodeInteger) Insert(insertNodeValue nodeValueInteger) (bool,*tnodeInteger) {
	if p.leftNode != nil && insertNodeValue.key < p.minValue()  {
		add,_ := p.leftNode.Insert(insertNodeValue)
		if add {
			return rebalanceInteger(p)
		}
		return false,p
	}
	if p.rightNode != nil && insertNodeValue.key > p.maxValue() {
		add,_ := p.rightNode.Insert(insertNodeValue)
		if add {
			return rebalanceInteger(p)
		}
		return false,p
	}
	isMatch,pos := p.getPosition(insertNodeValue.key)
	if isMatch == true {
		p.values[pos].rows = append(p.values[pos].rows,insertNodeValue.rows...)
		return false,p
	}

	//新規データ
	if p.IsOverFlow() == false {
		//オーバーフローなし
		p.insertValue(pos,insertNodeValue)
		return false,p
	}

	//オーバフローする
	if pos == 0 {
		p.createLeftNode()
		p.leftNode.Insert(insertNodeValue)
		return true,p
	}else if pos == p.dataCount {
		p.createRightNode()
		p.rightNode.Insert(insertNodeValue)
		return true,p
	} else {
		//minimumを取得して左ノードに再帰的にインサート
		minNode := p.popNodeValue(0)
		p.insertValue(pos-1,insertNodeValue)
		if p.leftNode == nil {
			p.createLeftNode()
			p.leftNode.Insert(minNode)
			return true,p
		}
		add,_ := p.leftNode.Insert(minNode)
		if add {
			return rebalanceInteger(p)
		}
		return false,p
	}
	return false,p
}


//Tnode削除
//TODO test
func(p *tnodeInteger) Delete(deleteValue Integer) (ROWNUM,bool,*tnodeInteger) {
	if p.leftNode != nil && deleteValue < p.minValue()  {
		deleteNum,del,_ := p.leftNode.Delete(deleteValue)
		delflg,newRoot := p.doAfterChildDelete(LEFT,del)
		return deleteNum,delflg,newRoot
	}
	if p.rightNode != nil && deleteValue > p.maxValue() {
		deleteNum,del,_ := p.rightNode.Delete(deleteValue)
		delflg,newRoot := p.doAfterChildDelete(RIGHT,del)
		return deleteNum,delflg,newRoot
	}
	isMatch,pos := p.getPosition(deleteValue)
	if isMatch == false {
		//該当データなし
		return 0,false,p
	}
	//削除
	deleteNum := p.deleteValue(pos)
	if p.IsInternalNode() && p.IsUnderFlow()  {
		//underflow時処理
		glb,del,_ := p.leftNode.popMaxNode()
		p.insertValue(0,glb)
		delflg,newRoot := p.doAfterChildDelete(LEFT,del)
		return deleteNum,delflg,newRoot
	}
	del := p.halfLeafMerge()
	return deleteNum,del,p
}
//最大値をpop
//TODO test
func(p *tnodeInteger) popMaxNode() (nodeValueInteger,bool,*tnodeInteger){
	if p.rightNode != nil {
		retNode,del,_ := p.rightNode.popMaxNode()
		delflg,newRoot := p.doAfterChildDelete(RIGHT,del)
		return retNode,delflg,newRoot
	}
	retNode := p.popNodeValue(p.dataCount-1)
	del := p.halfLeafMerge()
	return retNode,del,p
}

//harfnodeのマージ
//戻り値　ノードの削除発生
func(p *tnodeInteger) halfLeafMerge() bool{
	isDel := false
	if p.IsHalfLeaf() {
		//half-leaf
		switch p.canMergeChildNode() {
			case MERGE_TYPE_LEFT:
				p.mergeFromLeftNode()
				isDel = true
			case MERGE_TYPE_RIGHT:
				p.mergeFromRightNode()
				isDel = true
		}
	}
	return isDel
}
//子ノードのDelete実行後処理
//空になったリーフの削除、リバランス
func (p *tnodeInteger) doAfterChildDelete(child ChildType,del bool) (bool,*tnodeInteger){
	newRoot := p
	delflg := del
	if child == LEFT {
		if p.leftNode.dataCount == 0 {
			p.leftNode = nil
			delflg = true
		}
	} else if child == RIGHT {
		if p.rightNode.dataCount == 0 {
			p.rightNode = nil
			delflg = true
		}
	}
	//リバランス
	if delflg {
		_,newRoot = rebalanceInteger(p)
	}
	return delflg,newRoot
}
//次に挿入したらOverFlowするか
func(p *tnodeInteger) IsOverFlow() bool{
	return p.dataCount  == p.t
}
//under flowしているか
func(p *tnodeInteger) IsUnderFlow() bool{
	return p.dataCount  <= p.t-3
}
func(p *tnodeInteger) IsInternalNode() bool{
	return p.leftNode != nil && p.rightNode != nil
}
func(p *tnodeInteger) IsLeafNode() bool{
	return p.leftNode == nil && p.rightNode == nil
}
func(p *tnodeInteger) IsHalfLeaf() bool{
	return p.IsInternalNode() == false && p.IsLeafNode() == false
}
//マージできる子ノードの有無
//0なし 1左　2右　3両方
func(p *tnodeInteger) canMergeChildNode() MergeType{
	canMerge := MERGE_TYPE_NONE
	if p.leftNode != nil && p.dataCount + p.leftNode.dataCount  <= p.t {
		canMerge += MERGE_TYPE_LEFT
	}
	if p.rightNode != nil && p.dataCount + p.rightNode.dataCount  <= p.t {
		canMerge += MERGE_TYPE_RIGHT
	}
	return canMerge
}
func(p *tnodeInteger) maxValue() Integer{
	return p.values[p.dataCount-1].key
}
func(p *tnodeInteger) minValue() Integer{
	return p.values[0].key
}
//指定ポジションをpop
func(p *tnodeInteger) popNodeValue(pos int) nodeValueInteger{
	minNodeValue := p.values[pos]
	p.deleteValue(pos)
	return minNodeValue
}
//ノード内に値を挿入する
func(p *tnodeInteger) insertValue(insertPos int,insertNodeValue nodeValueInteger) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
    }
    p.values[insertPos] = insertNodeValue
    p.dataCount += 1
}
//ノード内の値を削除する
func(p *tnodeInteger) deleteValue(deletePos int) ROWNUM {
    rows := len(p.values[deletePos].rows)
    for i:= deletePos ; i < p.dataCount-1;i++ {
        p.values[i] = p.values[i+1]
    }
    //初期化
    p.values[p.dataCount-1] = nodeValueInteger{}
    p.dataCount -= 1
    return ROWNUM(rows)
}
//ノード内の操作対象箇所を検索する
func(p *tnodeInteger) getPosition(searchValue Integer) (bool,int) {
    return binarySearch(p.values,searchValue,0,p.dataCount-1)
}
//左ノードを作る
func(p *tnodeInteger) createLeftNode(){
    p.leftNode = createTnodeInteger(p.t)
    p.leftNode.parentNode = p
}
//右ノードを作る
func(p *tnodeInteger) createRightNode(){
    p.rightNode = createTnodeInteger(p.t)
    p.rightNode.parentNode = p
}

//左子ノードからマージ
func (p *tnodeInteger) mergeFromLeftNode(){
	p.mergeHead(p.leftNode,0)
	p.rightNode = p.leftNode.rightNode
	p.leftNode = p.leftNode.leftNode

}
//右子ノードからマージ
func (p *tnodeInteger) mergeFromRightNode(){
	p.mergeTail(p.rightNode,0)
	p.leftNode = p.rightNode.leftNode
	p.rightNode = p.rightNode.rightNode
}
//対象ノードを後ろ側にマージする
//srcNodeのstart番目以降をマージ
func (p *tnodeInteger) mergeTail(srcNode *tnodeInteger,start int){
	cnt := p.dataCount
	for i:= start; i < srcNode.dataCount ; i++{
		p.values[cnt+i-start] = srcNode.values[i]
	}
	p.dataCount = p.dataCount + srcNode.dataCount - start

	srcNode.clear(start)
}
//対象ノードを前側にマージする
//srcNodeのstart番目以降をマージ
func (p *tnodeInteger) mergeHead(srcNode *tnodeInteger,start int){
	cnt := srcNode.dataCount - start
	for i:= 0; i < p.dataCount ; i++{
		p.values[cnt+i] = p.values[i]
	}
	for i:= start; i < srcNode.dataCount ; i++{
		p.values[i-start] = srcNode.values[i]
	}
	p.dataCount = p.dataCount + srcNode.dataCount - start

	srcNode.clear(start)
}
//start番目以降をクリアする
func (p *tnodeInteger) clear(start int){
	for i:= start ;i < p.dataCount;i++ {
		p.values[i] = nodeValueInteger{0,nil}
	}
	p.dataCount =start
}

//LLローテーション
func rotationLLInteger(root *tnodeInteger) *tnodeInteger{
	//新たにrootになる
	newRoot := root.leftNode

	//左子を付け替え
	root.leftNode = newRoot.rightNode
	if root.leftNode != nil {
		root.leftNode.parentNode = root
	}

	//親を付け替え
	newRoot.parentNode = root.parentNode
	root.parentNode = newRoot
	newRoot.rightNode = root

	return newRoot
}
//RRローテーション
func rotationRRInteger(root *tnodeInteger) *tnodeInteger{
	//新たにrootになる
	newRoot := root.rightNode

	//右子を付け替え
	root.rightNode = newRoot.leftNode
	if root.rightNode != nil {
		root.rightNode.parentNode = root
	}

	//親を付け替え
	newRoot.parentNode = root.parentNode
	root.parentNode = newRoot
	newRoot.leftNode = root

	return newRoot
}
//LRローテーション
func rotationLRInteger(root *tnodeInteger) *tnodeInteger{
	//新たにrootになる
	newRoot := root.leftNode.rightNode
	//新たなleftNode
	newLeft := root.leftNode

	//leftNode(B)の付け替え
	newLeft.rightNode = newRoot.leftNode
	if newRoot.leftNode != nil {
		newRoot.leftNode.parentNode = newLeft
	}

	//rootNode(A)の付け替え
	root.leftNode = newRoot.rightNode
	if newRoot.rightNode != nil {
		newRoot.rightNode.parentNode = root
	}

	//newRoot(C)の付け替え
	newRoot.parentNode = root.parentNode

	newRoot.leftNode = newLeft
	newLeft.parentNode = newRoot

	newRoot.rightNode = root
	root.parentNode = newRoot

	//special lotation
	if newRoot.dataCount == 1 {
		newRoot.mergeHead(newLeft,1)
	}

	return newRoot
}
//RLローテーション
func rotationRLInteger(root *tnodeInteger) *tnodeInteger{
	//新たにrootになる
	newRoot := root.rightNode.leftNode
	//新たなrightNode
	newRight := root.rightNode

	//rightNode(B)の付け替え
	newRight.leftNode = newRoot.rightNode
	if newRoot.rightNode != nil {
		newRoot.rightNode.parentNode = newRight
	}
	//rootNode(A)の付け替え
	root.rightNode = newRoot.leftNode
	if newRoot.leftNode != nil {
		newRoot.leftNode.parentNode = root
	}
	//newRoot(C)の付け替え
	newRoot.parentNode = root.parentNode

	newRoot.rightNode = newRight
	newRight.parentNode = newRoot

	newRoot.leftNode = root
	root.parentNode = newRoot

	//special lotation
	if newRoot.dataCount == 1 {
		newRoot.mergeTail(newRight,1)
	}

	return newRoot
}

//リバランスチェックし必要ならリバランス
//戻り値　リバランスしたらfalse,新しいroot
func rebalanceInteger(root *tnodeInteger) (bool,*tnodeInteger){
	newRoot :=root
	def := root.leftDepth() - root.rightDepth()
	if def > -2 && def < 2 {
		//差が2以内
		return true,root
	}
	//親ノード退避
	parent := root.parentNode

	if def > 0 {
		//左が深い
		if root.leftNode.leftDepth() >= root.leftNode.rightDepth() {
			//LL回転
			newRoot = rotationLLInteger(root)
		} else {
			//LR回転
			newRoot = rotationLRInteger(root)
		}
	}else {
		//右が深い
		if root.rightNode.rightDepth() >= root.rightNode.leftDepth() {
			//RR回転
			newRoot = rotationRRInteger(root)
		} else {
			//RL回転
			newRoot = rotationRLInteger(root)
		}
	}

	//親ノードのポインタ付け替え
	if parent != nil {
		if parent.leftNode == root {
			parent.leftNode = newRoot
		} else if parent.rightNode == root {
			parent.rightNode = newRoot
		}
	}
	return false,newRoot
}
func (p *tnodeInteger) leftDepth() int{
	if p.leftNode != nil {
		return p.leftNode.depth() + 1
	}
	return 0
}
func (p *tnodeInteger) rightDepth() int{
	if p.rightNode != nil {
		return p.rightNode.depth() + 1
	}
	return 0
}
//木の深さを取得する
func (p *tnodeInteger) depth() int{
	leftDepth := p.leftDepth()
	rightDepth := p.rightDepth()
	//深い方
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

//ノード内の状態を出力する
func(p *tnodeInteger) Show() string {
    return p.showPadding(0)
}
func(p *tnodeInteger) showPadding(pad int) string {
    res := ""
    padding := strings.Repeat("-", pad)

    res += padding + "["
    for i:= 0;i < p.dataCount;i++ {
        res += strconv.Itoa(int(p.values[i].key)) + "("
        res += strconv.Itoa(len(p.values[i].rows))
        res += "),"
    }
    res += "]\n"
    if p.leftNode !=nil {
    	res += "l:" + p.leftNode.showPadding(pad+1)
    }
    if p.rightNode !=nil {
    	res += "r:" + p.rightNode.showPadding(pad+1)
    }
    return res
}