package db

import (

)

type ColumnNumber struct{
    Name string
}
func (p *ColumnNumber) Get() string { return p.Name}
func (p *ColumnNumber) Set(name string) { p.Name = name}