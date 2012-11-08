package db

import (
    //"fmt"
    "strings"
    "strconv"
)

//Bteeデータ構造
type BtreeInteger struct{
    rootNode *node
    dataCount ROWNUM
}
//BtreeIntegerの生成
func CreateNewBtree() BtreeInteger {
    return BtreeInteger{rootNode:new(node),dataCount:0}
}

//データ数
func(p *BtreeInteger) DataCount() ROWNUM{
    return p.dataCount
}

//行指定（ダミー）
func(p *BtreeInteger) Get(row ROWNUM) (Integer,error){
    return 0,nil
}

//探索
func(p *BtreeInteger) show() string{
    return p.rootNode.show()
}
//探索
func(p *BtreeInteger) Search(searchValue Integer) []ROWNUM{
    return p.rootNode.Search(searchValue)
}

//挿入
func(p *BtreeInteger) Insert(insertValue Integer) ROWNUM{
    newNodeValue,newChildNode := p.rootNode.Insert(insertValue,p.dataCount)
    if newChildNode != nil {
        //ルートノードの付け替え
        p.rootNode = createNewRoot(newNodeValue,p.rootNode,newChildNode)
    }
    p.dataCount += 1
    return ROWNUM(1)
}
/*
 * 新規ルートノードの生成
 */
func createNewRoot(newNodeValue nodeValue,rootNode *node,newChildNode *node) *node{
	newRootNode := new(node)
    newRootNode.values[0] = newNodeValue
    newRootNode.nodes[0] = rootNode
    newRootNode.nodes[1] = newChildNode
    newRootNode.dataCount = 1
    return newRootNode
}

//TODO ノードサイズを可変にする
const MAX_NODE_NUM int = 255

//Bツリーのノード
type node struct{
    //値
    values [MAX_NODE_NUM]nodeValue
    //子ノード
    nodes  [MAX_NODE_NUM+1]*node
    //データ数
    dataCount int
}

type nodeValue struct{
    key Integer
    rows []ROWNUM
}

//探索
func(p *node) Search(searchValue Integer) []ROWNUM{
    isMatch,searchPos := p.getPositionLinear(searchValue)

    //一致
    if isMatch {
    	return p.values[searchPos].rows
    }
    //子ノード
    if  p.nodes[searchPos] != nil {
        return p.nodes[searchPos].Search(searchValue)
    }
    //不一致
    return []ROWNUM{}
}


//挿入
func(p *node) Insert(insertValue Integer,row ROWNUM) (nodeValue,*node){

  	isMatch,insertPos := p.getPositionLinear(insertValue)
	//一致
    if isMatch {
        p.values[insertPos].rows = append(p.values[insertPos].rows,row)
    	return nodeValue{},nil
    }
    
    var newNodeValue = nodeValue{key : insertValue,rows : []ROWNUM{row}}
    var newChildNode *node =nil
    //子ノード
    if  p.nodes[insertPos] != nil {
        newNodeValue,newChildNode = p.nodes[insertPos].Insert(insertValue,row)
        if newChildNode == nil {
            //子ノードで分割が無ければリターン
            return nodeValue{},nil
        }
    }
   
    //新規データの挿入
    p.insertValue(insertPos,newNodeValue,newChildNode)
    if p.dataCount >= MAX_NODE_NUM {
		//ノード分割
    	return p.devideNode(p.dataCount /2)
    }
    //分割なし
    return nodeValue{},nil
}
//削除
//TODO test
//TODO ノードが0件になった場合
//TODO リファクタ
func(p *node) Delete(deleteValue Integer) bool{

    isMatch,deletePos := p.getPositionLinear(deleteValue)
    
    //一致
    if isMatch {
        //葉の場合
        if p.nodes[deletePos] ==nil{
            //普通に消す
            p.deleteValue(deletePos)
            if p.dataCount > MAX_NODE_NUM/2 {
                //データが十分にある
                return false
            }
            return true
        }
        //葉じゃない場合
        if p.nodes[deletePos].dataCount > MAX_NODE_NUM/2 {
            //左子から値を持ってくる。
            
            //deletePosに左子ノードの最大値を挿入
            p.values[deletePos] = p.nodes[deletePos].getMaxValue()
            //左子ノードの最大値を削除
            p.nodes[deletePos].Delete(p.nodes[deletePos].getMaxValue().key)
        } else if p.nodes[deletePos+1].dataCount > MAX_NODE_NUM/2 {
            //右子から値を持ってくる。
            //deletePosに右子ノードの最小値を挿入
            p.values[deletePos] = p.nodes[deletePos+1].getMinValue()
            //右子ノードの最小値を削除
            p.nodes[deletePos+1].Delete(p.nodes[deletePos+1].getMinValue().key)
        } else {
            //とりあえず左子から最大値を持ってくる
            p.values[deletePos] = p.nodes[deletePos].getMaxValue()
            //左子ノードの最大値を削除
            p.nodes[deletePos].Delete(p.nodes[deletePos].getMaxValue().key)
            
            
            //左子、持ってきた値、右子をマージして右子に入れる
            p.nodes[deletePos+1] = nodeMerge(p.nodes[deletePos],p.values[deletePos],p.nodes[deletePos+1])
            //値（と左子）を削除
            p.deleteValue(deletePos)
        }
        return false
    }
    
    //不一致
    //子ノードあり
    if  p.nodes[deletePos] != nil {
        if p.nodes[deletePos].Delete(deleteValue) {
            //子ノードで要素数が足りない
            
            //右ノードに要素が十分ある
            //自ノードの値を渡す
            //代わりを右から持ってくる
            
            //最右端の場合
            //左ノードに要素が十分ある
            //自ノードの値を渡す
            //代わりを左から持ってくる
            
            //上記以外
            //子、自分、右ノードをマージする
            //最右端の場合
            //子、自分、左ノードをマージする
            //自分を消す
        }
        return false
    }
    //データなし
    return false
}

//左右ノードのマージ
//TODO test
func nodeMerge(leftNode *node,med nodeValue,rightNode *node) *node{
    mergeNode := new(node)
    //左ノード代入
    for i:= 0;i<leftNode.dataCount;i++ {
        mergeNode.values[i] = leftNode.values[i]
        mergeNode.nodes[i] = leftNode.nodes[i]
        mergeNode.dataCount +=1
    }
    //中央値
    mergeNode.values[mergeNode.dataCount] = med
    mergeNode.nodes[mergeNode.dataCount] = leftNode.nodes[leftNode.dataCount]
    mergeNode.dataCount +=1
    
    //右ノード代入
    for i:= 0;i<rightNode.dataCount;i++ {
        mergeNode.values[leftNode.dataCount + i] = rightNode.values[i]
        mergeNode.nodes[leftNode.dataCount + i] = rightNode.nodes[i]
        mergeNode.dataCount +=1
    }
    //最後のノードポインタ
    mergeNode.nodes[mergeNode.dataCount] = rightNode.nodes[rightNode.dataCount]
    return mergeNode
}
/*
 *ノードを分割する
 */
func(p *node) devideNode(devidePosition int) (nodeValue,*node){
    //新規ノードの生成
    newNode := createNewNode(p,devidePosition)
   
    //親ノードに返す値        
    returnNodeValue := p.values[devidePosition] 
    
    //元ノードの初期化
    p.clear(devidePosition)
   
    return returnNodeValue,newNode

}

/*
 * 指定ポジション以降を初期化
 */
func(p *node) clear(devidePosition int) {
    for i:= devidePosition ; i<p.dataCount;i= i+1{            
        //初期化
        p.values[i] = nodeValue{0,nil}
        p.nodes[i+1] = nil
    }
    //データ数
    p.dataCount = devidePosition
}

/*
 *新たなノードの生成
 * srcノードの指定ポジション以降をコピー
 */
func createNewNode(srcNode *node,devidePosition int) *node {
    //木の分割
    newNode := new(node)
    newNode.nodes[0] = srcNode.nodes[devidePosition+1]
    for i,j:= devidePosition+1, 0 ; i<srcNode.dataCount;i,j = i+1,j+1{
        //データを移す
        newNode.values[j] = srcNode.values[i]
        newNode.nodes[j+1] = srcNode.nodes[i+1]
        newNode.dataCount +=1
    }    
    return newNode
}

//ノード内の操作対象箇所を線形検索する
func(p *node) getPositionLinear(searchValue Integer) (bool,int) {
	//線形探索
	var i int =0
	for ; i< p.dataCount;i++ {
        if p.values[i].key >= searchValue {
            break
        }
    }
    if i == p.dataCount {
        return false,p.dataCount
    }
    return p.values[i].key == searchValue,i
}

//ノード内に値を挿入する
func(p *node) insertValue(insertPos int,insertNodeValue nodeValue,newNode *node) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
        p.nodes[i+1] = p.nodes[i]   
    }
    p.values[insertPos] = insertNodeValue
    p.nodes[insertPos+1] = newNode
    p.dataCount += 1
}

//ノード内の値を削除する
//TODO test
func(p *node) deleteValue(deletePos int) {
    for i:= deletePos ; i < p.dataCount;i++ {
        p.values[i] = p.values[i+1]
        p.nodes[i] = p.nodes[i+1]   
    }
    p.nodes[p.dataCount] = p.nodes[p.dataCount+1]
    p.dataCount -= 1
}

//ノード配下の最大値を返す
//TODO test
func(p *node) getMaxValue() nodeValue{
    if p.nodes[p.dataCount] != nil {
        return p.nodes[p.dataCount].getMaxValue()
    }
    return p.values[p.dataCount-1]
}
//ノード配下の最小値を返す
//TODO test
func(p *node) getMinValue() nodeValue{
    if p.nodes[0] != nil {
        return p.nodes[0].getMinValue()
    }
    return p.values[0]
}

//ノード内の状態を出力する
func(p *node) show() string {
    return p.showPadding(0)
}
func(p *node) showPadding(pad int) string {
    res := ""
    padding := strings.Repeat("-", pad)
    
    res += padding + "["
    for i:= 0;i < p.dataCount;i++ {
        res += strconv.Itoa(int(p.values[i].key)) + "("
        res += strconv.Itoa(len(p.values[i].rows))
        res += "),"
    }
    res += "]\n"
    
    //子ノード
    for i:= 0;i < p.dataCount+1;i++ {
        if p.nodes[i] != nil {
            res += p.nodes[i].showPadding(pad+1)
        }
    }
    return res
}
