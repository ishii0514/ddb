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
    values [MAX_NODE_NUM-1]nodeValue
    nodes  [MAX_NODE_NUM]*node
    dataCount int
}

type nodeValue struct{
    key Integer
    rows []ROWNUM
}

//探索
func(p *node) Search(searchValue Integer) []ROWNUM{
    return p.linearSearch(searchValue)
}

//線形探索によるサーチ
func(p *node) linearSearch(searchValue Integer) []ROWNUM{
    var i int =0
    for ; i< p.dataCount;i++ {
        if p.values[i].key == searchValue {
            return p.values[i].rows
        }
        if p.values[i].key > searchValue {
            break
        }
    }
    if  p.nodes[i] != nil {
        return p.nodes[i].linearSearch(searchValue)
    }
    return []ROWNUM{}
}

//挿入
func(p *node) Insert(insertValue Integer,row ROWNUM) {
    p.linearInsert(insertValue,row)
}
//線形探索によるInsert
func(p *node) linearInsert(insertValue Integer,row ROWNUM){
    var i int =0
    for ; i< p.dataCount;i++ {
        if p.values[i].key == insertValue {
            p.values[i].rows = append(p.values[i].rows,row)
            return
        }
        if p.values[i].key > insertValue {
            break
        }
    }
    if  p.nodes[i] != nil {
        p.nodes[i].linearInsert(insertValue,row)        
        return
    }
    
    //新規データの挿入
    /*
    if(p.dataCount < MAX_NODE_NUM {
        //本ノードに挿入
        return
    }
    */
    //木の分岐
    
}