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
    var searchPos int = p.getPositionLinear(searchValue)

    if p.values[searchPos].key == searchValue {
    	return p.values[searchPos].rows
    }
    if  p.nodes[searchPos] != nil {
        return p.nodes[searchPos].Search(searchValue)
    }
    return []ROWNUM{}
}

//挿入
func(p *node) Insert(insertValue Integer,row ROWNUM){
	var searchPos int = p.getPositionLinear(insertValue)
	
    if p.values[searchPos].key == insertValue {
        p.values[searchPos].rows = append(p.values[searchPos].rows,row)
    	return
    }
    if  p.nodes[searchPos] != nil {
        p.nodes[searchPos].Insert(insertValue,row)        
        return
    }
   
    //新規データの挿入
    if p.dataCount < MAX_NODE_NUM {
        //本ノードに挿入
	    for i:= p.dataCount;i > searchPos;i-- {
	    	p.values[i].key = p.values[i-1].key
	    	p.values[i].rows = p.values[i-1].rows
	    	p.nodes[i+1] = p.nodes[i] 	
	    }
	    p.values[searchPos].key = insertValue
	    p.values[searchPos].rows = []ROWNUM{row}
	    p.nodes[searchPos] = nil
        return
    }
    //木の分岐
    
}

//ノード内の操作対象箇所を検索する
func(p *node) getPositionLinear(searchValue Integer) int {
	//線形探索
	for i:=0 ; i< p.dataCount;i++ {
        if p.values[i].key >= searchValue {
            return i
        }
    }
    return p.dataCount;
}