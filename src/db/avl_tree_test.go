package db

import (
	"testing"
)

type TestData struct {
	actual []ROWNUM
	expect []ROWNUM
}

func TestAVLTreeSearch(t *testing.T) {
	// create test tree
	tree := CreateAVLTreeInteger()

	node1 := CreateAVLNode(Elem{12, []ROWNUM{10}}, nil)
	node2 := CreateAVLNode(Elem{1, []ROWNUM{20, 30}}, node1)
	node3 := CreateAVLNode(Elem{20, []ROWNUM{40, 50}}, node1)
	node4 := CreateAVLNode(Elem{15, []ROWNUM{60, 70, 80}}, node3)

	node3.left = node4
	node1.left = node2
	node1.right = node3

	tree.root = node1

	// if the key exists, return the value associated with given key
	// else, return []ROWNUM{}	
	data := []TestData{
		TestData{tree.Search(12), []ROWNUM{10}},
		TestData{tree.Search(1), []ROWNUM{20, 30}},
		TestData{tree.Search(20), []ROWNUM{40, 50}},
		TestData{tree.Search(15), []ROWNUM{60, 70, 80}},
		TestData{tree.Search(35), []ROWNUM{}},
	}

	for _, v := range data {
		if !AreEqual(v.actual, v.expect) {
			t.Errorf("expect: %v, actual: %v", v.expect, v.actual)
		}
	}
}

func TestAVLTreeInsert(t *testing.T) {
	var tree *AVLTreeInteger = nil
	var act_row, exp_row ROWNUM
	var act_node, exp_node *AVLNode
	var act_cnt, exp_cnt ROWNUM
	exp_row = 1

	// if empty tree inserts the key, the root is created
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_node = tree.root
	exp_node = &AVLNode{Elem{20, []ROWNUM{1}}, 1, nil, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("expect: %v, actual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 1, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}

	// add to left
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(12)
	act_node = tree.root
	exp_node = &AVLNode{Elem{20, []ROWNUM{1}}, 2, nil, nil, nil}
	exp_node.left = &AVLNode{Elem{12, []ROWNUM{2}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 2, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}

	// add to right
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(30)
	act_node = tree.root
	exp_node = &AVLNode{Elem{20, []ROWNUM{1}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{30, []ROWNUM{2}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 2, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}
	
	// duplicate
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(30)
	act_row = tree.Insert(30)
	act_node = tree.root
	exp_node = &AVLNode{Elem{20, []ROWNUM{1}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{30, []ROWNUM{2, 3}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(10)
	act_row = tree.Insert(20)
	act_node = tree.root
	exp_node = &AVLNode{Elem{20, []ROWNUM{1, 3}}, 2, nil, nil, nil}
	exp_node.left = &AVLNode{Elem{10, []ROWNUM{2}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}
	
	// RR case
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(30)
	act_row = tree.Insert(40)
	act_node = tree.root
	exp_node = &AVLNode{Elem{30, []ROWNUM{2}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{40, []ROWNUM{3}}, 1, exp_node, nil, nil}
	exp_node.left = &AVLNode{Elem{20, []ROWNUM{1}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}

	// RL rotation
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(40)
	act_row = tree.Insert(30)
	act_node = tree.root
	exp_node = &AVLNode{Elem{30, []ROWNUM{3}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{40, []ROWNUM{2}}, 1, exp_node, nil, nil}
	exp_node.left = &AVLNode{Elem{20, []ROWNUM{1}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}

	// LL rotation
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(10)
	act_row = tree.Insert(5)
	act_node = tree.root
	exp_node = &AVLNode{Elem{10, []ROWNUM{2}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{20, []ROWNUM{1}}, 1, exp_node, nil, nil}
	exp_node.left = &AVLNode{Elem{5, []ROWNUM{3}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}
	
	// LR rotation
	tree = CreateAVLTreeInteger()
	act_row = tree.Insert(20)
	act_row = tree.Insert(5)
	act_row = tree.Insert(10)
	act_node = tree.root
	exp_node = &AVLNode{Elem{10, []ROWNUM{3}}, 2, nil, nil, nil}
	exp_node.right = &AVLNode{Elem{20, []ROWNUM{1}}, 1, exp_node, nil, nil}
	exp_node.left = &AVLNode{Elem{5, []ROWNUM{2}}, 1, exp_node, nil, nil}
	if exp_row != act_row {
		t.Errorf("expect: %v, actual: %v", exp_row, act_row)
	}
	if !act_node.Equals(exp_node) {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_node, act_node)
	}
	act_cnt, exp_cnt = 3, tree.DataCount()
	if act_cnt != exp_cnt {
		t.Errorf("\nexpect: %v, \nactual: %v", exp_cnt, act_cnt)
	}
}

