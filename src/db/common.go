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



