package db

import (

)

// AVLTree データ構造
type AVLTreeInteger struct {
  rootNode *AVLNode
  dataCount ROWNUM
}

// AVLNode 構造体
type AVLNode struct {
  // 値
  values nodeValue
  // 左子ノード
  left *AVLNode
  // 右子ノード
  right *AVLNode
  // データ数
  dataCount int
}

// AVLTreeInteger の生成
func CreateNewAVLTree() AVLTreeInteger {
  return AVLTreeInteger{rootNode: new(AVLNode), dataCount: 0}
}
