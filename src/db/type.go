package db

import (
  "strconv"
)
//データ型インターフェース
type Type interface {
  comp(Type) int
  print() string
}

//行を表す型
type ROWNUM int32
//データ型
//integer型
type Integer int
//文字列型
type Varchar string

/*
 * p.comp(v) == 0 : equal
 * p.comp(v) > 0  : p > v
 * p.comp(v) < 0  : p < v
 */
func(p Integer) comp(value Type) int{
  return (int)(p - value.(Integer))
}
func(p Integer) print() string{
  return strconv.Itoa(int(p))
}

func(p Varchar) comp(value Type) int{
  if p == value.(Varchar) {
    return 0
  } else if p > value.(Varchar) {
    return 1
  }
  return -1
}
func(p Varchar) print() string{
  return string(p)
}

//文字列をInteger型に変換する
func StringtoInteger(value string) (Integer,error){
  v,err := strconv.Atoi(value)
  return Integer(v),err
}

//Integer型の無効値　とりあえず0
const INVALID_VALUE_INTEGER Integer = 0
//全型の無効値
const INVALID_VALUE string = "0"
