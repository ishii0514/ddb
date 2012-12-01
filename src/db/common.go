package db

import (
	"strconv"
)

//データ型
//integer型
type Integer int
//文字列型
type Varchar string
//文字列をInteger型に変換する
func StringtoInteger(value string) (Integer,error){
	v,err := strconv.Atoi(value)
	return Integer(v),err
}

//行を表す型
type ROWNUM int32

//Integer型の無効値　とりあえず0
const INVALID_VALUE_INTEGER Integer = 0
//全型の無効値
const INVALID_VALUE string = "0"


//エラー構造体
type DbError struct {
    Message string
}
func (e *DbError) Error() string {
    return e.Message
}

//カラム型を表す定数
type ColumnType int
const (
	COLUMN_TYPE_INTEGER ColumnType = iota
	COLUMN_TYPE_STRING  ColumnType = iota
)

//ttreeのMergeを表す定数
type MergeType int
const (
	MERGE_TYPE_NONE MergeType = iota	//0
	MERGE_TYPE_LEFT MergeType = iota	//1
	MERGE_TYPE_RIGHT  MergeType = iota	//2
	MERGE_TYPE_BOTH  MergeType = iota	//3
)
