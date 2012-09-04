package db

import (
)

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
func (p *ColumnNumber) Insert(data int) {p.data = append(p.data,data)}
func (p *ColumnNumber) Get(row int) int {
	return p.data[row]
}
func (p *ColumnNumber) DataCount() int { return len(p.data)} 