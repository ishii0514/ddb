package db

import (

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

//探索
func(p *BtreeInteger) Search(searchValue Integer) []ROWNUM{
    return p.rootNode.Search(searchValue)
}

//挿入
func(p *BtreeInteger) Insert(insertValue Integer) ROWNUM{
    newNodeValue,newChildNode := p.rootNode.Insert(insertValue,p.dataCount)
    if newChildNode != nil {
        //ルートノードの分割
        newRootNode := new(node)
        newRootNode.values[0] = newNodeValue
        newRootNode.nodes[0] = p.rootNode
        newRootNode.nodes[1] = newChildNode
        newRootNode.dataCount = 1
        //ルートノードの付け替え
        p.rootNode = newRootNode
    }
    p.dataCount += 1
    return ROWNUM(1)
}



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

//TODO リファクタ
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
        //木の分割
        newNode := new(node)
        devPos := p.dataCount /2

        //親ノードに返す値        
        returnNode := p.values[devPos]
        //初期化
        p.values[devPos] = nodeValue{0,[]ROWNUM{}}
        //データを移す
        newNode.nodes[0] = p.nodes[devPos+1]
        //初期化
        p.nodes[devPos+1] =nil
        

        for i,j:= devPos+1, 0 ; i<p.dataCount;i,j = i+1,j+1{
            //データを移す
            newNode.values[j] = p.values[i]
            newNode.nodes[j+1] = p.nodes[i+1]
            
            //初期化
            p.values[i] = nodeValue{0,[]ROWNUM{}}
            p.nodes[i+1] = nil
        }
        //データ数
        newNode.dataCount = p.dataCount - (devPos+1)
        p.dataCount = devPos 
        
        return returnNode,newNode        
    }
    //分割なし
    return nodeValue{},nil
}

//ノード内の操作対象箇所を検索する
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
