package db

import (

)

type nodeValue struct{
  key Type
  rows []ROWNUM
}

func(p* nodeValue)clear(){
  p.key= nil
  p.rows =nil
}

type nodeValueInteger struct{
    key Integer
    rows []ROWNUM
}


type nodeValueString struct{
    key Varchar
    rows []ROWNUM
}
