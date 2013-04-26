package db

import (
    "strings"
    "strconv"
)

//Bteeデータ構造
type BtreeInteger struct{
    rootNode *node
    dataCount ROWNUM
    rowid ROWNUM
}
//BtreeIntegerの生成
func CreateBtree(t int) BtreeInteger {
    return BtreeInteger{rootNode:createNode(t),dataCount:0,rowid:0}
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
func(p *BtreeInteger) Show() string{
    return p.rootNode.Show()
}
//探索
func(p *BtreeInteger) Search(searchValue Integer) []ROWNUM{
    return p.rootNode.Search(searchValue)
}

//挿入
func(p *BtreeInteger) Insert(insertValue Integer) ROWNUM{
    newNodeValue,newChildNode := p.rootNode.Insert(insertValue,p.rowid)
    if newChildNode != nil {
        //ルートノードの付け替え
        p.rootNode = createNewRoot(newNodeValue,p.rootNode,newChildNode)
    }
    p.dataCount += 1
    p.rowid += 1
    return ROWNUM(1)
}
//削除
//test 件数正しいか
func(p *BtreeInteger) Delete(deleteValue Integer) ROWNUM{
    deleteRows := p.rootNode.Delete(deleteValue)
    if p.rootNode.dataCount == 0 && p.rootNode.nodes[0] != nil{
        //ルートノードの付け替え
        p.rootNode = p.rootNode.nodes[0]
    }
    p.dataCount = p.dataCount - deleteRows
    return deleteRows
}
/*
 * 新規ルートノードの生成
 */
func createNewRoot(newNodeValue nodeValueInteger,rootNode *node,newChildNode *node) *node{
	newRootNode := createNode(rootNode.t)
    newRootNode.values[0] = newNodeValue
    newRootNode.nodes[0] = rootNode
    newRootNode.nodes[1] = newChildNode
    newRootNode.dataCount = 1
    return newRootNode
}

//Bツリーのノード
type node struct{
    //データ数
    dataCount int    
    //値
    values []nodeValueInteger
    //子ノード
    nodes  []*node
    //ノードサイズ
    t int
}

func createNode(t int) *node{
    newNode := new(node)
    nodeSize := 2*t-1
    if t <= 0 {
        nodeSize = 0
    }
    //予備で一つ分多く取っておく
    newNode.values = make([]nodeValueInteger,nodeSize+1)
    newNode.nodes = make([]*node,nodeSize+2)
    newNode.t = t
    newNode.dataCount = 0
    return newNode
}

//ノード内の操作対象箇所を二分検索する
func binarySearch(values []nodeValueInteger,searchValue Integer,head int,tail int) (bool,int){
    if head > tail {
        return false,head
    }
    pivot := (head+tail)/2
    
    if values[pivot].key == searchValue {
        return true,pivot
    } else if values[pivot].key > searchValue {
        return binarySearch(values,searchValue,head,pivot-1)
    }
    return binarySearch(values,searchValue,pivot+1,tail)
}



//探索
func(p *node) Search(searchValue Integer) []ROWNUM{
	//再帰なし版
	node := p
	for ;; {
	    isMatch,searchPos := node.getPosition(searchValue)
	    //一致
    	if isMatch {
    		return node.values[searchPos].rows
    	}
    	
    	//子ノード
    	if  node.nodes[searchPos] == nil {
	    	break
    	}
    	node = node.nodes[searchPos]
    }
    //不一致
    return []ROWNUM{}
}
//探索
func(p *node) Search_re(searchValue Integer) []ROWNUM{
    isMatch,searchPos := p.getPosition(searchValue)

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
func(p *node) Insert(insertValue Integer,row ROWNUM) (nodeValueInteger,*node){

  	isMatch,insertPos := p.getPosition(insertValue)
	//一致
    if isMatch {
        p.values[insertPos].rows = append(p.values[insertPos].rows,row)
    	return nodeValueInteger{},nil
    }
    
    var newNodeValue = nodeValueInteger{key : insertValue,rows : []ROWNUM{row}}
    var newChildNode *node =nil
    //子ノード
    if  p.nodes[insertPos] != nil {
        newNodeValue,newChildNode = p.nodes[insertPos].Insert(insertValue,row)
        if newChildNode == nil {
            //子ノードで分割が無ければリターン
            return nodeValueInteger{},nil
        }
    }
   
    //新規データの挿入
    p.insertValue(insertPos,newNodeValue,newChildNode)
    if p.dataCount > p.t*2-1 {
		//ノード分割
    	return p.devideNode(p.t)
    }
    //分割なし
    return nodeValueInteger{},nil
}
//削除
//TODO test
//TODO リファクタ
func(p *node) Delete(deleteValue Integer) ROWNUM{

    isMatch,deletePos := p.getPosition(deleteValue)
    
    //一致
    if isMatch {
        //葉の場合
        if p.nodes[deletePos] ==nil{
            //消す
            return p.deleteValue(deletePos)
        }
        //内部接点の場合
        rows := len(p.values[deletePos].rows)
        if p.nodes[deletePos].dataCount >= p.t {
            //左子から値を持ってくる。
            //deletePosに左子ノードの最大値を挿入
            p.values[deletePos] = p.nodes[deletePos].getMaxValue()
            //左子ノードの最大値を削除
            p.nodes[deletePos].Delete(p.values[deletePos].key)
        } else if p.nodes[deletePos+1].dataCount >= p.t {
            //右子から値を持ってくる。
            //deletePosに右子ノードの最小値を挿入
            p.values[deletePos] = p.nodes[deletePos+1].getMinValue()
            //右子ノードの最小値を削除
            p.nodes[deletePos+1].Delete(p.values[deletePos].key)

        } else {
            //左子、現在の値、右子をマージして右子に入れる
            p.nodes[deletePos+1] = mergeNodes(p.nodes[deletePos],p.values[deletePos],p.nodes[deletePos+1])
            //値（と左子）を削除
            p.deleteValue(deletePos)
            //再帰的に右子から削除
            //根ノードの時、valuesが無くなりnodes[0]だけ残る場合がある。
            p.nodes[deletePos].Delete(deleteValue)
            
        }
        return ROWNUM(rows)
    }
    
    //不一致
    if  p.nodes[deletePos] == nil {
        //葉
        return ROWNUM(0)
    }
    //内部接点
    if p.nodes[deletePos].dataCount <= p.t-1 {
        //対象子ノードに要素が十分に無い場合
        if deletePos < p.dataCount && p.nodes[deletePos+1].dataCount >= p.t {
            //右兄弟ノードに要素が十分ある
            //対象子ノード末尾に要素を挿入
            p.nodes[deletePos].addTail(
                p.values[deletePos],
                p.nodes[deletePos+1].nodes[0])
                    
            //現ノードに右兄弟から値を代入
            p.values[deletePos] = p.nodes[deletePos+1].values[0]
            //右兄弟から値を削除
            p.nodes[deletePos+1].removeHead()
        } else if deletePos > 0 && p.nodes[deletePos-1].dataCount >=p.t {
            //左兄弟ノードがあり要素が十分ある
            //対象子ノード先頭に要素を挿入
            p.nodes[deletePos].addHead(
                p.values[deletePos-1],
                p.nodes[deletePos-1].nodes[p.nodes[deletePos-1].dataCount])
            //現ノードに左兄弟から値を代入
            p.values[deletePos-1] = p.nodes[deletePos-1].values[p.nodes[deletePos-1].dataCount-1]
            //左兄弟から値を削除
            p.nodes[deletePos-1].removeTail()
        } else if deletePos < p.dataCount {
            //右兄弟ノードとマージ
            //対象子ノード、現在の値、右兄弟をマージして右兄弟に入れる
            p.nodes[deletePos+1] = mergeNodes(p.nodes[deletePos],p.values[deletePos],p.nodes[deletePos+1])
            //値（と左子）を削除
            //根ノードの時、valuesが無くなりnodes[0]だけ残る場合がある。
            p.deleteValue(deletePos)
        } else {
            //右兄弟が無い場合,左兄弟ノードとマージ
            //左兄弟、現在の値、対象子ノードをマージして対象子ノードに入れる
            p.nodes[deletePos] = mergeNodes(p.nodes[deletePos-1],p.values[deletePos-1],p.nodes[deletePos])
            //値（と左子）を削除
            //根ノードの時、valuesが無くなりnodes[0]だけ残る場合がある。
            p.deleteValue(deletePos-1)
            
            //削除した分ずれる
            deletePos = deletePos-1
        }
    }
    //再帰的に削除            
    return p.nodes[deletePos].Delete(deleteValue)
}

//左右ノードのマージ
//TODO test
func mergeNodes(leftNode *node,med nodeValueInteger,rightNode *node) *node{
    mergeNode := createNode(leftNode.t)
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
    j := mergeNode.dataCount
    for i:= 0;i<rightNode.dataCount;i++ {
        mergeNode.values[j + i] = rightNode.values[i]
        mergeNode.nodes[j + i] = rightNode.nodes[i]
        mergeNode.dataCount +=1
    }
    //最後のノードポインタ
    mergeNode.nodes[mergeNode.dataCount] = rightNode.nodes[rightNode.dataCount]
    return mergeNode
}
/*
 *ノードを分割する
 */
func(p *node) devideNode(devidePosition int) (nodeValueInteger,*node){
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
        p.values[i] = nodeValueInteger{0,nil}
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
    newNode := createNode(srcNode.t)
    newNode.nodes[0] = srcNode.nodes[devidePosition+1]
    for i,j:= devidePosition+1, 0 ; i<srcNode.dataCount;i,j = i+1,j+1{
        //データを移す
        newNode.values[j] = srcNode.values[i]
        newNode.nodes[j+1] = srcNode.nodes[i+1]
        newNode.dataCount +=1
    }    
    return newNode
}

//ノード内の操作対象箇所を検索する
func(p *node) getPosition(searchValue Integer) (bool,int) {
    return p.binarySearch(searchValue,0,p.dataCount-1)
}

//ノード内の操作対象箇所を線形検索する
func(p *node) linearSearch(searchValue Integer) (bool,int){
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
//ノード内の操作対象箇所を二分検索する
func(p *node) binarySearch(searchValue Integer,head int,tail int) (bool,int){
	return binarySearch(p.values,searchValue,head,tail)
}

//ノード内に値を挿入する
func(p *node) insertValue(insertPos int,insertNodeValue nodeValueInteger,newNode *node) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
        p.nodes[i+1] = p.nodes[i]   
    }
    p.values[insertPos] = insertNodeValue
    p.nodes[insertPos+1] = newNode
    p.dataCount += 1
}
//末尾に値を挿入する
func(p *node) addTail(insertNodeValue nodeValueInteger,newNode *node) {
    p.insertValue(p.dataCount,insertNodeValue,newNode)
}

//ノードの先頭に値を挿入する
//nodesも0番目に挿入されるので、insertValueと異なる。
func(p *node) addHead(insertNodeValue nodeValueInteger,newNode *node) {

    p.nodes[p.dataCount+1] = p.nodes[p.dataCount]
    for i:= p.dataCount;i > 0;i-- {
        p.values[i] = p.values[i-1]
        p.nodes[i] = p.nodes[i-1]   
    }
    p.values[0] = insertNodeValue
    p.nodes[0] = newNode
    p.dataCount += 1
}

//ノード内の値を削除する
func(p *node) deleteValue(deletePos int) ROWNUM {
    rows := len(p.values[deletePos].rows)
    for i:= deletePos ; i < p.dataCount-1;i++ {
        p.values[i] = p.values[i+1]
        p.nodes[i] = p.nodes[i+1]   
    }
    p.nodes[p.dataCount-1] = p.nodes[p.dataCount]
    
    //初期化
    p.values[p.dataCount-1] = nodeValueInteger{}
    p.nodes[p.dataCount] = nil
    
    p.dataCount -= 1
    return ROWNUM(rows)
}
//先頭の値を削除する
func(p *node) removeHead() {
    p.deleteValue(0)
}
//ノードの末尾の値を削除する
//nodesはdataCount番目が削除されるので、deleteVal。
func(p *node) removeTail() {
    p.values[p.dataCount-1] = nodeValueInteger{}
    p.nodes[p.dataCount] = nil
    p.dataCount -= 1
}
//ノード配下の最大値をgetする
//TODO test
func(p *node) getMaxValue() nodeValueInteger{
    if p.nodes[p.dataCount] != nil {
        return p.nodes[p.dataCount].getMaxValue()
    }
    return p.values[p.dataCount-1]
}
//ノード配下の最小値をgetする
//TODO test
func(p *node) getMinValue() nodeValueInteger{
    if p.nodes[0] != nil {
        return p.nodes[0].getMinValue()
    }
    return p.values[0]
}

//ノード内の状態を出力する
func(p *node) Show() string {
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
