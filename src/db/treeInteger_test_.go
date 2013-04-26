package db

import (
  "testing"
//  "math/rand"
//  "strconv"
)

func TestInsertValueTInteger(t *testing.T) {
  root := createTnodeInteger(5)
  root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
  if root.dataCount != 1 {
        t.Error("illegal dataCount 1")
    }
    if root.values[0].key != 1 {
        t.Error("illegal key 1")
    }

    root.insertValue(1,nodeValueInteger{5,[]ROWNUM{1}})
  if root.dataCount != 2 {
        t.Error("illegal dataCount 2")
    }
    if root.values[1].key != 5 {
        t.Error("illegal key 5")
    }

    root.insertValue(1,nodeValueInteger{3,[]ROWNUM{1}})
  if root.dataCount != 3 {
        t.Error("illegal dataCount 3")
    }
    if root.values[0].key != 1 {
        t.Error("illegal key 1")
    }
    if root.values[1].key != 3 {
        t.Error("illegal key 3")
    }
    if root.values[2].key != 5 {
        t.Error("illegal key 5")
    }

}
func TestInsertValueT2Integer(t *testing.T) {
  root := createTnodeInteger(5)
  root.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
  if root.dataCount != 3 {
        t.Error("illegal dataCount 3")
    }
    if root.values[0].key != 1 {
        t.Error("illegal key 1")
    }
    if root.values[1].key != 3 {
        t.Error("illegal key 3")
    }
    if root.values[2].key != 5 {
        t.Error("illegal key 5")
    }

}
func TestDeleteValueTInteger(t *testing.T) {
  root := createTnodeInteger(5)
  root.insertValue(0,nodeValueInteger{9,[]ROWNUM{1}})
  root.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})
  root.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
    if root.dataCount != 5 {
        t.Error("illegal dataCount 5")
    }

    if root.deleteValue(0) != 1 {
        t.Error("illegal rownum 1")
    }
    if root.dataCount != 4 {
        t.Error("illegal rownum 4")
    }

    if root.deleteValue(2) != 1 {
        t.Error("illegal rownum 1")
    }
    if root.dataCount != 3 {
        t.Error("illegal rownum 3")
    }

    if root.values[0].key != 3 {
        t.Error("illegal key 3")
    }
    if root.values[1].key != 5 {
        t.Error("illegal key 5")
    }
    if root.values[2].key != 9 {
        t.Error("illegal key 9")
    }
}
func TestCanMergeChildNodeInteger(t *testing.T) {
  root := createTnodeInteger(5)
  left := createTnodeInteger(5)
  right := createTnodeInteger(5)
  if root.canMergeChildNode() != MERGE_TYPE_NONE {
        t.Error("illegal merge none")
    }

  root.leftNode = left
  root.rightNode = right
  left.parentNode = root
  right.parentNode = root
  if root.canMergeChildNode() != MERGE_TYPE_BOTH {
        t.Error("illegal merge both")
    }

    root.insertValue(0,nodeValueInteger{9,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})

    left.insertValue(0,nodeValueInteger{9,[]ROWNUM{1}})
    left.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    if root.canMergeChildNode() != MERGE_TYPE_BOTH {
        t.Error("illegal merge both2")
    }
    left.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})
    if root.canMergeChildNode() != MERGE_TYPE_RIGHT {
        t.Error("illegal merge right")
    }

    right.insertValue(0,nodeValueInteger{9,[]ROWNUM{1}})
    right.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    if root.canMergeChildNode() != MERGE_TYPE_RIGHT {
        t.Error("illegal merge right2")
    }
    right.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})
    if root.canMergeChildNode() != MERGE_TYPE_NONE {
        t.Error("illegal merge none")
    }
}
func TestMaxMinValueInteger(t *testing.T) {
  root := createTnodeInteger(5)
  root.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
  if root.maxValue() != 5 {
        t.Error("illegal maxvalue 5")
    }
    if root.minValue() != 1 {
        t.Error("illegal minvalue 1")
    }
}
func TestPopNodeValueInteger(t *testing.T) {
  root := createTnodeInteger(5)
  root.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})
  root.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})

    popValue := root.popNodeValue(1)
  if  popValue.key != 3 {
        t.Error("illegal pop value 3")
    }

    if root.dataCount != 3 {
        t.Error("illegal dataCount 3")
    }
    if root.values[0].key != 1 {
        t.Error("illegal key 3")
    }
    if root.values[1].key != 5 {
        t.Error("illegal key 5")
    }
    if root.values[2].key != 7 {
        t.Error("illegal key 7")
    }

}

func TestMergeFromLeftNodeInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    left := createTnodeInteger(5)
    left.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    left.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
    root.leftNode = left
    left.parentNode = root

    root.mergeFromLeftNode()
  if  root.dataCount != 4 {
        t.Error("illegal dataCount 4")
    }
    if  root.leftNode != nil {
        t.Error("illegal left nil")
    }
    if  root.rightNode != nil {
        t.Error("illegal right nil")
    }

    if root.values[0].key != 1 {
        t.Error("illegal key 1")
    }
    if root.values[1].key != 3 {
        t.Error("illegal key 3")
    }
    if root.values[2].key != 4 {
        t.Error("illegal key 4")
    }
    if root.values[3].key != 8 {
        t.Error("illegal key 8")
    }

}
func TestMergeFromRightNodeInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    right := createTnodeInteger(5)
    right.insertValue(0,nodeValueInteger{20,[]ROWNUM{1}})
    right.insertValue(0,nodeValueInteger{10,[]ROWNUM{1}})
    root.rightNode = right
    right.parentNode = root

    root.mergeFromRightNode()
  if  root.dataCount != 4 {
        t.Error("illegal dataCount 4")
    }
    if  root.leftNode != nil {
        t.Error("illegal left nil")
    }
    if  root.rightNode != nil {
        t.Error("illegal right nil")
    }

    if root.values[0].key != 4 {
        t.Error("illegal key 4")
    }
    if root.values[1].key != 8 {
        t.Error("illegal key 8")
    }
    if root.values[2].key != 10 {
        t.Error("illegal key 20")
    }
    if root.values[3].key != 20 {
        t.Error("illegal key 20")
    }

}
func TestMergeTailInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    src := createTnodeInteger(5)
    src.insertValue(0,nodeValueInteger{20,[]ROWNUM{1}})
    src.insertValue(0,nodeValueInteger{10,[]ROWNUM{1}})


    root.mergeTail(src,1)
  if  root.dataCount != 3 {
        t.Error("illegal dataCount 3")
    }

    if root.values[0].key != 4 {
        t.Error("illegal key 4")
    }
    if root.values[1].key != 8 {
        t.Error("illegal key 8")
    }
    if root.values[2].key != 20 {
        t.Error("illegal key 20")
    }

}
func TestMergeHeadInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    src := createTnodeInteger(5)
    src.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    src.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    src.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})


    root.mergeHead(src,1)
  if  root.dataCount != 4 {
        t.Error("illegal dataCount 4")
    }

    if root.values[0].key != 2 {
        t.Error("illegal key 2")
    }
    if root.values[1].key != 3 {
        t.Error("illegal key 3")
    }
    if root.values[2].key != 4 {
        t.Error("illegal key 4")
    }
    if root.values[3].key != 8 {
        t.Error("illegal key 8")
    }

}

func TestClearTInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})


    root.clear(2)
  if  root.dataCount != 2 {
        t.Error("illegal dataCount 2")
    }

    if root.values[0].key != 1 {
        t.Error("illegal key 1")
    }
    if root.values[1].key != 2 {
        t.Error("illegal key 2")
    }
}
func TestRotationLLInteger(t *testing.T) {
  bl := createTnodeInteger(5)
    bl.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
  b := createTnodeInteger(5)
    b.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    br := createTnodeInteger(5)
    br.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    a := createTnodeInteger(5)
    a.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    ar := createTnodeInteger(5)
    ar.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})

    bl.parentNode = b
    b.leftNode = bl

    br.parentNode = b
    b.rightNode = br

    a.leftNode = b
    b.parentNode = a

    a.rightNode = ar
    ar.parentNode = a

    newRoot := rotationLLInteger(a)

    //root
    if newRoot.values[0].key != 2 {
      t.Error("illegal root 2")
    }
    if newRoot.parentNode != nil {
      t.Error("illegal root parent")
    }

    //左子
    if newRoot.leftNode.values[0].key != 1 {
      t.Error("illegal left")
    }
    if newRoot.leftNode.parentNode != newRoot {
      t.Error("illegal left parent")
    }
    //右子
    if newRoot.rightNode.values[0].key != 4 {
      t.Error("illegal right")
    }
    if newRoot.rightNode.parentNode != newRoot {
      t.Error("illegal right parent")
    }

    //右左子
    if newRoot.rightNode.leftNode.values[0].key != 3 {
      t.Error("illegal right left")
    }
    if newRoot.rightNode.leftNode.parentNode != newRoot.rightNode {
      t.Error("illegal right left parent")
    }

    //右右子
    if newRoot.rightNode.rightNode.values[0].key != 5 {
      t.Error("illegal right right")
    }
    if newRoot.rightNode.rightNode.parentNode != newRoot.rightNode {
      t.Error("illegal right right parent")
    }
}
func TestRotationRRInteger(t *testing.T) {
  br := createTnodeInteger(5)
    br.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
  b := createTnodeInteger(5)
    b.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    bl := createTnodeInteger(5)
    bl.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    a := createTnodeInteger(5)
    a.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    al := createTnodeInteger(5)
    al.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})

    bl.parentNode = b
    b.leftNode = bl

    br.parentNode = b
    b.rightNode = br

    a.rightNode = b
    b.parentNode = a

    a.leftNode = al
    al.parentNode = a

    newRoot := rotationRRInteger(a)

    //root
    if newRoot.values[0].key != 4 {
      t.Error("illegal root 4")
    }
    if newRoot.parentNode != nil {
      t.Error("illegal root parent")
    }

    //右子
    if newRoot.rightNode.values[0].key != 5 {
      t.Error("illegal right")
    }
    if newRoot.rightNode.parentNode != newRoot {
      t.Error("illegal right parent")
    }
    //左子
    if newRoot.leftNode.values[0].key != 2 {
      t.Error("illegal left")
    }
    if newRoot.leftNode.parentNode != newRoot {
      t.Error("illegal left parent")
    }

    //左右子
    if newRoot.leftNode.rightNode.values[0].key != 3 {
      t.Error("illegal left right")
    }
    if newRoot.leftNode.rightNode.parentNode != newRoot.leftNode {
      t.Error("illegal left right parent")
    }

    //左左子
    if newRoot.leftNode.leftNode.values[0].key != 1 {
      t.Error("illegal left left")
    }
    if newRoot.leftNode.leftNode.parentNode != newRoot.leftNode {
      t.Error("illegal left left parent")
    }
}
func TestRotationLRInteger(t *testing.T) {
  bl := createTnodeInteger(5)
    bl.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})
  b := createTnodeInteger(5)
    b.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    cl := createTnodeInteger(5)
    cl.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    c := createTnodeInteger(5)
    c.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    cr := createTnodeInteger(5)
    cr.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    a := createTnodeInteger(5)
    a.insertValue(0,nodeValueInteger{6,[]ROWNUM{1}})
    ar := createTnodeInteger(5)
    ar.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})

    cl.parentNode = c
    c.leftNode = cl

    cr.parentNode = c
    c.rightNode = cr

    bl.parentNode = b
    b.leftNode = bl

    c.parentNode = b
    b.rightNode = c

    a.leftNode = b
    b.parentNode = a

    a.rightNode = ar
    ar.parentNode = a

    newRoot := rotationLRInteger(a)

    //root
    if newRoot.values[0].key != 4 {
      t.Error("illegal root 4")
    }
    if newRoot.parentNode != nil {
      t.Error("illegal root parent")
    }

    //左子
    if newRoot.leftNode.values[0].key != 2 {
      t.Error("illegal left")
    }
    if newRoot.leftNode.parentNode != newRoot {
      t.Error("illegal left parent")
    }
    //左左子
    if newRoot.leftNode.leftNode.values[0].key != 1 {
      t.Error("illegal left left")
    }
    if newRoot.leftNode.leftNode.parentNode != newRoot.leftNode {
      t.Error("illegal left left parent")
    }
    //左右子
    if newRoot.leftNode.rightNode.values[0].key != 3 {
      t.Error("illegal left right")
    }
    if newRoot.leftNode.rightNode.parentNode != newRoot.leftNode {
      t.Error("illegal left right parent")
    }

    //右子
    if newRoot.rightNode.values[0].key != 6 {
      t.Error("illegal right")
    }
    if newRoot.rightNode.parentNode != newRoot {
      t.Error("illegal right parent")
    }

    //右左子
    if newRoot.rightNode.leftNode.values[0].key != 5 {
      t.Error("illegal right left")
    }
    if newRoot.rightNode.leftNode.parentNode != newRoot.rightNode {
      t.Error("illegal right left parent")
    }

    //右右子
    if newRoot.rightNode.rightNode.values[0].key != 7 {
      t.Error("illegal right right")
    }
    if newRoot.rightNode.rightNode.parentNode != newRoot.rightNode {
      t.Error("illegal right right parent")
    }
}
func TestRotationRLInteger(t *testing.T) {
  br := createTnodeInteger(5)
    br.insertValue(0,nodeValueInteger{7,[]ROWNUM{1}})
  b := createTnodeInteger(5)
    b.insertValue(0,nodeValueInteger{6,[]ROWNUM{1}})
    cr := createTnodeInteger(5)
    cr.insertValue(0,nodeValueInteger{5,[]ROWNUM{1}})
    c := createTnodeInteger(5)
    c.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    cl := createTnodeInteger(5)
    cl.insertValue(0,nodeValueInteger{3,[]ROWNUM{1}})
    a := createTnodeInteger(5)
    a.insertValue(0,nodeValueInteger{2,[]ROWNUM{1}})
    al := createTnodeInteger(5)
    al.insertValue(0,nodeValueInteger{1,[]ROWNUM{1}})

    cr.parentNode = c
    c.rightNode = cr

    cl.parentNode = c
    c.leftNode = cl


    br.parentNode = b
    b.rightNode = br

    c.parentNode = b
    b.leftNode = c

    a.rightNode = b
    b.parentNode = a

    a.leftNode = al
    al.parentNode = a

    newRoot := rotationRLInteger(a)

    //root
    if newRoot.values[0].key != 4 {
      t.Error("illegal root 4")
    }
    if newRoot.parentNode != nil {
      t.Error("illegal root parent")
    }

    //左子
    if newRoot.leftNode.values[0].key != 2 {
      t.Error("illegal left")
    }
    if newRoot.leftNode.parentNode != newRoot {
      t.Error("illegal left parent")
    }
    //左左子
    if newRoot.leftNode.leftNode.values[0].key != 1 {
      t.Error("illegal left left")
    }
    if newRoot.leftNode.leftNode.parentNode != newRoot.leftNode {
      t.Error("illegal left left parent")
    }
    //左右子
    if newRoot.leftNode.rightNode.values[0].key != 3 {
      t.Error("illegal left right")
    }
    if newRoot.leftNode.rightNode.parentNode != newRoot.leftNode {
      t.Error("illegal left right parent")
    }

    //右子
    if newRoot.rightNode.values[0].key != 6 {
      t.Error("illegal right")
    }
    if newRoot.rightNode.parentNode != newRoot {
      t.Error("illegal right parent")
    }

    //右左子
    if newRoot.rightNode.leftNode.values[0].key != 5 {
      t.Error("illegal right left")
    }
    if newRoot.rightNode.leftNode.parentNode != newRoot.rightNode {
      t.Error("illegal right left parent")
    }

    //右右子
    if newRoot.rightNode.rightNode.values[0].key != 7 {
      t.Error("illegal right right")
    }
    if newRoot.rightNode.rightNode.parentNode != newRoot.rightNode {
      t.Error("illegal right right parent")
    }
}
func TestGetPositionTInteger(t *testing.T) {
  root := createTnodeInteger(3)
  root.Insert(nodeValueInteger{8,[]ROWNUM{1}})
  root.Insert(nodeValueInteger{7,[]ROWNUM{1}})
  root.Insert(nodeValueInteger{6,[]ROWNUM{1}})

  match,pos := root.getPosition(5)
  if match != false {
    t.Error("illegal match")
  }
  if pos != 0 {
    t.Error("illegal pos")
  }
}
func TestDepthInteger(t *testing.T) {
  root := createTnodeInteger(5)
    root.insertValue(0,nodeValueInteger{8,[]ROWNUM{1}})
    root.insertValue(0,nodeValueInteger{4,[]ROWNUM{1}})
    right := createTnodeInteger(5)
    right.insertValue(0,nodeValueInteger{20,[]ROWNUM{1}})
    right.insertValue(0,nodeValueInteger{10,[]ROWNUM{1}})
    root.rightNode = right
    right.parentNode = root

  if root.depth() != 1 {
    t.Error("illegal root depth")
  }
  if right.depth() != 0 {
    t.Error("illegal right depth")
  }
  left := createTnodeInteger(5)
  root.leftNode = left
    left.parentNode = root
  leftleft := createTnodeInteger(5)
  left.leftNode = leftleft
    leftleft.parentNode = left
  if root.depth() != 2 {
    t.Error("illegal root depth")
  }
}
func TestShowTInteger(t *testing.T) {
  root := createTnodeInteger(3)
  _,root = root.Insert(nodeValueInteger{8,[]ROWNUM{1}})
  _,root = root.Insert(nodeValueInteger{7,[]ROWNUM{1}})
  _,root = root.Insert(nodeValueInteger{5,[]ROWNUM{1}})
  _,root = root.Insert(nodeValueInteger{6,[]ROWNUM{1}})

  res := root.Show()
  exp := "[6(1),7(1),8(1),]\n"
  exp += "l:-[5(1),]\n"
  if res != exp {
    t.Error("illegal insert")
  }

  _,root = root.Insert(nodeValueInteger{4,[]ROWNUM{1}})
  _,root = root.Insert(nodeValueInteger{3,[]ROWNUM{1}})
  _,root = root.Insert(nodeValueInteger{2,[]ROWNUM{1}})
  root.Insert(nodeValueInteger{1,[]ROWNUM{1}})
  root.Insert(nodeValueInteger{0,[]ROWNUM{1}})
  root.Insert(nodeValueInteger{9,[]ROWNUM{1}})
  res = root.Show()
  //print(res)
  exp = "[3(1),4(1),5(1),]\n"
  exp += "l:-[0(1),1(1),2(1),]\n"
  exp += "r:-[6(1),7(1),8(1),]\n"
  exp += "r:--[9(1),]\n"
  if res != exp {
    t.Error("illegal insert 2")
  }
}
func TestTreeInsertInteger(t *testing.T) {
  tree := CreateTtreeInteger(3)
  tree.Insert(8)
  tree.Insert(7)
  tree.Insert(5)
  tree.Insert(6)

  res := tree.Show()
  exp := "[6(1),7(1),8(1),]\n"
  exp += "l:-[5(1),]\n"
  if res != exp {
    t.Error("illegal insert")
  }
  tree.Insert(4)
  tree.Insert(3)
  tree.Insert(2)
  tree.Insert(1)
  tree.Insert(0)
  tree.Insert(9)

  res = tree.Show()
  //print(res)
  exp = "[3(1),4(1),5(1),]\n"
  exp += "l:-[0(1),1(1),2(1),]\n"
  exp += "r:-[6(1),7(1),8(1),]\n"
  exp += "r:--[9(1),]\n"
  if res != exp {
    t.Error("illegal insert 2")
  }
}

func TestTreeDeleteRRLotationInteger(t *testing.T) {
  tree := CreateTtreeInteger(3)
  tree.Insert(8)
  tree.Insert(7)
  tree.Insert(5)
  tree.Insert(6)
  tree.Insert(4)
  tree.Insert(3)
  tree.Insert(2)
  tree.Insert(1)
  tree.Insert(0)
  tree.Insert(9)

  res := tree.Show()
  //print(res)
  exp := "[3(1),4(1),5(1),]\n"
  exp += "l:-[0(1),1(1),2(1),]\n"
  exp += "r:-[6(1),7(1),8(1),]\n"
  exp += "r:--[9(1),]\n"
  if res != exp {
    t.Error("illegal insert")
  }

  tree.Delete(2)
  tree.Delete(1)
  tree.Delete(0)
  res = tree.Show()
  //print(res)
  exp = "[6(1),7(1),8(1),]\n"
  exp += "l:-[3(1),4(1),5(1),]\n"
  exp += "r:-[9(1),]\n"
  if res != exp {
    t.Error("illegal delete")
  }
}
func TestTreeDeleteLRLotationInteger(t *testing.T) {
  tree := CreateTtreeInteger(3)
  tree.Insert(9)
  tree.Insert(10)
  tree.Insert(11)
  tree.Insert(12)
  tree.Insert(3)
  tree.Insert(4)
  tree.Insert(5)
  tree.Insert(6)
  tree.Insert(7)
  tree.Insert(8)

  res := tree.Show()
  //print(res)
  exp := "[9(1),10(1),11(1),]\n"
  exp += "l:-[3(1),4(1),5(1),]\n"
  exp += "r:--[6(1),7(1),8(1),]\n"
  exp += "r:-[12(1),]\n"
  if res != exp {
    t.Error("illegal insert")
  }

  tree.Delete(12)

  res = tree.Show()
  //print(res)
  exp = "[6(1),7(1),8(1),]\n"
  exp += "l:-[3(1),4(1),5(1),]\n"
  exp += "r:-[9(1),10(1),11(1),]\n"
  if res != exp {
    t.Error("illegal insert")
  }

}
