package db

import (
)

//エラー構造体
type DbError struct {
    Message string
}
func (e *DbError) Error() string {
    return e.Message
}


type Column interface {
	GetName() string
	SetName(name string)
}

type ColumnNumber struct{
    name string
    data []int
}

func (p *ColumnNumber) GetName() string { return p.name}
func (p *ColumnNumber) SetName(name string) { p.name = name}
func (p *ColumnNumber) DataCount() int { return len(p.data)}
func (p *ColumnNumber) Get(row int) (int,error) {
    if row >= p.DataCount(){
        return 0,&DbError{"out of range."}
    }
	return p.data[row],nil
}
func (p *ColumnNumber) Insert(data int) {p.data = append(p.data,data)}
//指定した行のデータを削除する
func (p *ColumnNumber) DeleteAt(row int) {
    if row >= p.DataCount(){
        return
    }
    p.data = append(p.data[:row],p.data[row+1:]...)
}
 