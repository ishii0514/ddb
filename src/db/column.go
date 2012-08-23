package db

import (

)


type ColumnNumber struct{
    name string
}
func (p *ColumnNumber) Get() string { return p.name}
func (p *ColumnNumber) Set(name string) { p.name = name}

func Get() string{return "hello db"}