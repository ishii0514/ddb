package db

import (

)

//Bteeデータ構造
type BtreeInteger struct{
    data node
}

const NODE_NUM int = 255

type node struct{
    values [NODE_NUM-1]data
    nodes  [NODE_NUM]*node
    dataCount int
}

type data struct{
    key Integer
    rows []ROWNUM
}

func(p *node) Search(searchValue Integer) []ROWNUM{
    for i:=0; i< p.dataCount;i++ {
        if p.values[i].key == searchValue {
            return p.values[i].rows
        }
        if p.values[i].key >= searchValue {
            return p.nodes[i].Search(searchValue)
        }
    }
    return []ROWNUM{}
}
