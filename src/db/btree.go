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

type node struct{
    values [MAX_NODE_NUM]nodeValue
    nodes  [MAX_NODE_NUM+1]*node
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
func(p *node) Insert(insertValue Integer,row ROWNUM){
	isMatch,insertPos := p.getPositionLinear(insertValue)
	
	//一致
    if isMatch {
        p.values[insertPos].rows = append(p.values[insertPos].rows,row)
    	return
    }
    //子ノード
    if  p.nodes[insertPos] != nil {
        p.nodes[insertPos].Insert(insertValue,row)        
        return
    }
   
    //新規データの挿入
    p.insertValue(insertPos,insertValue,row)
    if p.dataCount >= MAX_NODE_NUM {
        //木の分割
    }
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
func(p *node) insertValue(insertPos int,insertValue Integer,row ROWNUM) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i].key = p.values[i-1].key
        p.values[i].rows = p.values[i-1].rows
        p.nodes[i+1] = p.nodes[i]   
    }
    p.values[insertPos].key = insertValue
    p.values[insertPos].rows = []ROWNUM{row}
    p.nodes[insertPos+1] = nil
    p.dataCount += 1
}