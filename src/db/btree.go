package db

import (

)

//Bteeデータ構造
type BtreeInteger struct{
    rootNode *node
    dataCount ROWNUM
}
//探索
func(p *BtreeInteger) Search(searchValue Integer) []ROWNUM{
    return p.rootNode.Search(searchValue)
}

//挿入
func(p *BtreeInteger) Insert(insertValue Integer) ROWNUM{
    p.rootNode.Insert(insertValue,p.dataCount)
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

//挿入
func(p *node) Insert(insertValue Integer,row ROWNUM) (nodeValue,*node){
	isMatch,insertPos := p.getPositionLinear(insertValue)
	
	var returnNode nodeValue
	var newNode *node =nil
	//一致
    if isMatch {
        p.values[insertPos].rows = append(p.values[insertPos].rows,row)
    	return returnNode,nil
    }
    //子ノード
    if  p.nodes[insertPos] != nil {
        returnNode,newNode = p.nodes[insertPos].Insert(insertValue,row)
        if newNode == nil {
            //子ノードで分割が無ければリターン
            return returnNode,nil
        }
    }
   
    //TODO 子ノードからの値が返った場合考慮
    //insertPosも変わる
    
    //新規データの挿入
    p.insertValue(insertPos,insertValue,row,newNode)
    if p.dataCount >= MAX_NODE_NUM {
        //木の分割
        newNode := new(node)
        devPos := p.dataCount /2

        //親ノードに返す値        
        returnNode = p.values[devPos]
        //初期化
        p.values[devPos].insert(0,[]ROWNUM{})
        //データを移す
        newNode.nodes[0] = p.nodes[devPos+1]
        //初期化
        p.nodes[devPos+1] =nil
        

        for i,j:= devPos+1, 0 ; i<p.dataCount;i,j = i+1,j+1{
            newNode.values[j].insert(p.values[i].key,p.values[i].rows)
            newNode.nodes[j+1] = p.nodes[i+1]
            
            //初期化
            p.values[i].insert(0,[]ROWNUM{})
            p.nodes[i+1] = nil
        }
        
        return returnNode,newNode        
    }
    //分割なし
    return returnNode,nil
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
func(p *node) insertValue(insertPos int,insertValue Integer,row ROWNUM,newNode *node) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i].insert(p.values[i-1].key,p.values[i-1].rows)
        p.nodes[i+1] = p.nodes[i]   
    }
    p.values[insertPos].insert(insertValue,[]ROWNUM{row})
    p.nodes[insertPos+1] = newNode
    p.dataCount += 1
}

/*
 *データの挿入
 *
 */
func(p *nodeValue) insert(key Integer,rows []ROWNUM){
        p.key = key
        p.rows = rows
}