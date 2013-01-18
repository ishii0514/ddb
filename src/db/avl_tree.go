package db

import (
	"fmt"
)

// implements CompareTo(Elem)
type Elem nodeValue

// compares two keys. 
// returns negative if p.key < value.key, 
// positive if p.key > value.key,
// or 0 if p.key == value.key
func (p *Elem) CompareTo(value Elem) int {
	return int(p.key - value.key)
}

// return true or false
func AreEqual(s1, s2 []ROWNUM) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// node of AVLTree
type AVLNode struct {
	data   Elem
	height int
	parent *AVLNode
	left   *AVLNode
	right  *AVLNode
}

// makes a new node with given value and parent, and with nil child links.
// NOTE: the height of leaf node is 1.
// param: value
// param: parent
func CreateAVLNode(data Elem, parent *AVLNode) *AVLNode {
	// re-calculate ancestor's height
	height := 2
	ancestor := parent
	for ancestor != nil && height > ancestor.height {
		ancestor.height = height
		height = height + 1
		ancestor = ancestor.parent
	}
	return &AVLNode{data, 1, parent, nil, nil}
}

// returns true of false
func (this *AVLNode) Equals(obj *AVLNode) bool {
	if this == obj {
		return true
	}
	if obj == nil {
		return false
	}
	other := obj
	if this.Key() != other.Key() {
		return false
	}
	if this.height != other.height {
		return false
	}
	if this.parent == nil {
		if other.parent != nil {
			return false
		}
	} else if other.parent == nil || this.parent.Key() != other.parent.Key() {
		return false
	}
	if this.left == nil {
		if other.left != nil {
			return false
		}
	} else if !this.left.Equals(other.left) {
		return false
	}
	if this.right == nil {
		if other.right != nil {
			return false
		}
	} else if !this.right.Equals(other.right) {
		return false
	}
	if this.Value() == nil {
		if other.Value() != nil {
			return false
		}
	} else if !AreEqual(this.Value(), other.Value()) {
		return false
	}
	return true
}

func (p *AVLNode) String() string {
	return fmt.Sprintf("Node{key:%v value:%v height:%v left:%v right:%v}",
		p.Key(), p.Value(), p.height, p.left, p.right)
}

// Getters

// returns the key of data
func (p *AVLNode) Key() Integer {
	return p.data.key
}

// returns the value of data
func (p *AVLNode) Value() []ROWNUM {
	return p.data.rows
}

// returns height
func (p *AVLNode) Height() int {
	return p.height
}

// returns parent
func (p *AVLNode) Parent() *AVLNode {
	return p.parent
}

// returns left
func (p *AVLNode) Left() *AVLNode {
	return p.left
}

// returns right
func (p *AVLNode) Right() *AVLNode {
	return p.right
}

// AVL Tree
type AVLTreeInteger struct {
	// root node
	root *AVLNode
	// number of entries in the tree
	data_count ROWNUM
}

// creates a new empty AVL Tree. 
func CreateAVLTreeInteger() *AVLTreeInteger {
	return &AVLTreeInteger{nil, 0}
}

// returns the number of entries
func (p *AVLTreeInteger) DataCount() ROWNUM {
	return p.data_count
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
func (p *AVLTreeInteger) GetEntry(key Integer) *AVLNode {
	node := p.root
	for node != nil {
		cmp := p.Compare(key, node.Key())
		if cmp == 0 {
			return node
		}
		if cmp < 0 {
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
	added := false
	p.root, added = p.InsertKey(key)
	if added {
		p.root = p.resetHeight(p.root)
		p.data_count = p.data_count + 1
		return ROWNUM(1)
	}
	return ROWNUM(0)
}

func (p *AVLTreeInteger) InsertKey(key Integer) (*AVLNode, bool) {
	return p.insert(p.root, nil, key, p.data_count+1)
}

// return new element associated with given key, or nil
func (p *AVLTreeInteger) insert(node, parent *AVLNode, key Integer, row ROWNUM) (*AVLNode, bool) {
	if node == nil {
		return p.Construct(Elem{key, []ROWNUM{row}}, parent), true
	}

	cmp := p.Compare(key, node.Key())
	if cmp == 0 {
		node = p.AppendValue(node, row)
		return node, true
	}
	added := true
	if cmp < 0 {
		node.left, added = p.insert(node.left, node, key, row)
		if node.left != nil && node.left.parent == nil {
			node.left.parent = node
		}
		if added {
			return p.Rebalance(node), added
		}
	} else {
		node.right, added = p.insert(node.right, node, key, row)
		if node.right != nil && node.right.parent == nil {
			node.right.parent = node
		}
		if added {
			return p.Rebalance(node), added
		}
	}
	// unfortunately break loop
	return nil, false
}

// Construct a node in the tree.
func (p *AVLTreeInteger) Construct(data Elem, parent *AVLNode) *AVLNode {
	return CreateAVLNode(data, parent)
}

// Rebalance subtree
func (p *AVLTreeInteger) Rebalance(subroot *AVLNode) *AVLNode {
	newroot := subroot
	if subroot != nil {
		// calculate balance factor
		factor := p.Height(subroot.left) - p.Height(subroot.right)
		if factor < -1 {
			factor = p.Height(subroot.right.left) - p.Height(subroot.right.right)
			if factor == -1 {
				// RR case
			} else if factor == 1 {
				// RL case
				p.rotateRight(subroot.right)
			}
			newroot = p.rotateLeft(subroot)
		} else if factor > 1 {
			factor = p.Height(subroot.left.left) - p.Height(subroot.left.right)
			if factor == 1 {
				// LL case
			} else if factor == -1 {
				// LR case
				p.rotateLeft(subroot.left)
			}
			newroot = p.rotateRight(subroot)
		}
	}
	return newroot
}

// appends row value to rows
func (p *AVLTreeInteger) AppendValue(node *AVLNode, row ROWNUM) *AVLNode {
	if node != nil {
		node.data.rows = append(node.data.rows, row)
	}
	return node
}

func (p *AVLTreeInteger) Height(node *AVLNode) int {
	height := 0
	if node != nil {
		height = node.height
	}
	return height
}

// rotates left around node
func (p *AVLTreeInteger) rotateLeft(node *AVLNode) *AVLNode {
	right := node.right
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.parent = node.parent
	if node.parent == nil {
		p.root = right
	} else if node.parent.left == node {
		node.parent.left = right
	} else {
		node.parent.right = right
	}
	right.left = node
	node.parent = right
	return right
}

// rotates right around node
func (p *AVLTreeInteger) rotateRight(node *AVLNode) *AVLNode {
	left := node.left
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.parent = node.parent
	if node.parent == nil {
		p.root = left
	} else if node.parent.right == node {
		node.parent.right = left
	} else {
		node.parent.left = left
	}
	left.right = node
	node.parent = left
	return left
}

func (p *AVLTreeInteger) resetHeight(node *AVLNode) *AVLNode {
	if node != nil {
		node.left = p.resetHeight(node.left)
		node.right = p.resetHeight(node.right)
		h_left, h_right := p.Height(node.left), p.Height(node.right)
		if h_left > h_right {
			node.height = h_left + 1
		} else {
			node.height = h_right + 1
		}
	}
	return node
}

// Delete
// return deleted row count
func (p *AVLTreeInteger) Delete(key Integer) ROWNUM {
	node := p.GetEntry(key)
	if node == nil {
		return 0
	}
	deleted_count := ROWNUM(len(node.Value()))
	p.remove(node)
	p.data_count = p.data_count - deleted_count
	return deleted_count
}

// remove given node from tree
func (p *AVLTreeInteger) remove(node *AVLNode) {
}
