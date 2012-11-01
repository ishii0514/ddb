// AVLTree のテスト
package db

import (
    "testing"
)

// CreateNewAVLTree のテスト
func TestCreateNewAVLTree(t *testing.T) {
  actual := CreateNewAVLTree()
  if actual.rootNode == nil || actual.dataCount != 0 {
    t.Error("failed to create avl tree.")
  }
}

