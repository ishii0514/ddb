package db

import (

)

// implements CompareTo(NodeValue)
type NodeValue nodeValue

// compares two keys. 
// returns negative if p.key < value.key, 
// positive if p.key > value.key,
// or 0 if p.key == value.key
func (p *NodeValue) CompareTo(value NodeValue) int {
	return int(p.key - value.key)
}

// node of AVLTree
type AVLNode struct {
	data NodeValue
	height int
	parent *AVLNode
	left  *AVLNode
	right *AVLNode
}

// makes a new node with given value and parent, and with nil child links.
// param: value
// param: parent
func CreateAVLNode(data NodeValue, parent *AVLNode) *AVLNode {
	return &AVLNode{data, 0, parent, nil, nil}
}

// returns the key of data
func (p *AVLNode) Key() Integer {
	return p.data.key
}

// returns the value of data
func (p *AVLNode) Value() []ROWNUM {
	return p.data.rows
}

// appends row value to rows
func (p *AVLNode) AppendValue(row ROWNUM) *AVLNode {
	p.data.rows = append(p.data.rows, row)
	return p
}

// AVL Tree
type AVLTreeInteger struct {
	// root node
	root *AVLNode
	// number of entries in the tree
	dataCount ROWNUM
}

// creates a new empty AVL Tree. 
func CreateAVLTreeInteger() *AVLTreeInteger {
	return &AVLTreeInteger{nil, 0}
}

// returns the number of entries
func (p *AVLTreeInteger) DataCount() ROWNUM {
	return p.dataCount
}

// returns key associated with given value(represents rowid)
func (p *AVLTreeInteger) Get(row ROWNUM) (Integer, error) {
	// not implemented
	return 0, nil
}

// compares two keys using the correct comparison method for this BalancedTree.

func (p *AVLTreeInteger) Compare(key1, key2 Integer) int {
	return int(key1 - key2)
}

// returns value associated with given key, or empty set if the map does not contain
// an entry for the key.
func (p *AVLTreeInteger) Search(key Integer) []ROWNUM {
	if node := p.GetEntry(key); node != nil {
		return node.Value()
	}
	// not found
	return []ROWNUM{}
}

// returns entry for given key, or nil if the map does not contain
// an entry for the key.
func(p *AVLTreeInteger) GetEntry(key Integer) *AVLNode {
	node := p.root
	for ; node != nil; {
		cmp := p.Compare(key, node.Key())
		if cmp == 0 {
			return node
		} else if cmp < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	// not found
	return nil
}

// associates the specified value with the specified key in this map.
// If the map previously contained a mapping for this key, 
// append the newest value and returned.
func (p *AVLTreeInteger) Insert(key Integer) ROWNUM {
	if p.InsertKey(key) != nil {
		return ROWNUM(1)
	}
	return ROWNUM(0)
}

func (p *AVLTreeInteger) InsertKey(key Integer) *AVLNode {
	return p.insert(key, p.dataCount + 1)
}

func (p *AVLTreeInteger) insert(key Integer, row ROWNUM) *AVLNode {
	t := p.root
	if t == nil {
		p.dataCount++
		p.root = p.Construct(NodeValue{key, []ROWNUM{row}}, nil)
		return p.root
	}
	return nil
}

// Construct a node in the tree.
func (p *AVLTreeInteger) Construct(data NodeValue, parent *AVLNode) *AVLNode {
	return CreateAVLNode(data, parent)
}

