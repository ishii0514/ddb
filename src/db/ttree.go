package db

import (
    "strings"
    "strconv"
)
//ttreeのMergeを表す定数
type MergeType int
const (
  MERGE_TYPE_NONE MergeType = iota  //0
  MERGE_TYPE_LEFT MergeType = iota  //1
  MERGE_TYPE_RIGHT  MergeType = iota  //2
  MERGE_TYPE_BOTH  MergeType = iota  //3
)
//ttreeのMergeを表す定数
type ChildType int
const (
  NONE ChildType = iota  //0
  LEFT ChildType = iota  //1
  RIGHT ChildType = iota  //2

)

func binarySearch(values []nodeValue,searchValue Type,head int,tail int) (bool,int){
    //再帰なし
    for ;; {
      if head > tail {
        return false,head
      }
      pivot := (head+tail)/2
      if values[pivot].key.comp(searchValue) == 0 {
        return true,pivot
      } else if values[pivot].key.comp(searchValue) > 0 {
        tail = pivot-1
      } else {
         head = pivot+1
      }
    }
    return false,head
}

//Tteeデータ構造
type Ttree struct{
    rootNode *tnode
    dataCount ROWNUM
    rowid ROWNUM
}
//Ttreeコンストラクタ
func CreateTtree(t int) Ttree {
    return Ttree{rootNode:createTnode(t),dataCount:0,rowid:0}
}

//データ数
func(p *Ttree) DataCount() ROWNUM{
    return p.dataCount
}

//行指定（ダミー）
func(p *Ttree) Get(row ROWNUM) (Type,error){
    return nil,nil
}
//探索
func(p *Ttree) Search(searchValue Type) []ROWNUM{
    return p.rootNode.Search(searchValue)
}
//挿入
func(p *Ttree) Insert(insertValue Type) ROWNUM{
  _,p.rootNode = p.rootNode.Insert(nodeValue{insertValue,[]ROWNUM{p.rowid}})
    p.dataCount += 1
    p.rowid += 1
    return ROWNUM(1)
}
func(p *Ttree) Delete(deleteValue Type) ROWNUM{
  var deleteRows ROWNUM = 0
    deleteRows,_,p.rootNode = p.rootNode.Delete(deleteValue)
    p.dataCount = p.dataCount - deleteRows
    return deleteRows
}
//探索
func(p *Ttree) Show() string{
    return p.rootNode.Show()
}
//Tツリーのノード
type tnode struct{
    //ノードサイズ
    t int
    //データ数
    dataCount int
    //値
    values []nodeValue

    //parentノード
    parentNode  *tnode
    //leftノード
    leftNode  *tnode
    //rightノード
    rightNode  *tnode
}

//Tノードコンストラクタ
func createTnode(t int) *tnode{
    newNode := new(tnode)
    newNode.values = make([]nodeValue,t)
    newNode.t = t
    newNode.dataCount = 0
    newNode.parentNode = nil
    newNode.leftNode = nil
    newNode.rightNode = nil
    return newNode
}
//探索
//TODO test
func(p *tnode) Search(searchValue Type) []ROWNUM{
  //再帰なし版
  node := p
  for ; ; {
    if node.leftNode != nil && searchValue.comp(node.minValue()) < 0  {
      node = node.leftNode
      continue
    }
    if node.rightNode != nil && searchValue.comp(node.maxValue()) >0 {
      node = node.rightNode
      continue
    }
    break
  }
  isMatch,pos := node.getPosition(searchValue)
  if isMatch == true {
    return node.values[pos].rows
  }
  return []ROWNUM{}
}
func(p *tnode) Search_(searchValue Type) []ROWNUM{
    if p.leftNode != nil && searchValue.comp(p.minValue())<0  {
    return p.leftNode.Search(searchValue)
  }
  if p.rightNode != nil && searchValue.comp(p.maxValue())>0 {
    return p.rightNode.Search(searchValue)
  }
  isMatch,pos := p.getPosition(searchValue)
  if isMatch == true {
    return p.values[pos].rows
  }
  return []ROWNUM{}
}
//Tnodeインサート
//戻り値　ノード追加発生,新たなルートノード
//TODO リファクタ
func(p *tnode) Insert(insertNodeValue nodeValue) (bool,*tnode) {
  if p.leftNode != nil && insertNodeValue.key.comp(p.minValue())<0  {
    add,_ := p.leftNode.Insert(insertNodeValue)
    if add {
      return rebalance(p)
    }
    return false,p
  }
  if p.rightNode != nil && insertNodeValue.key.comp(p.maxValue())>0 {
    add,_ := p.rightNode.Insert(insertNodeValue)
    if add {
      return rebalance(p)
    }
    return false,p
  }
  isMatch,pos := p.getPosition(insertNodeValue.key)
  if isMatch == true {
    p.values[pos].rows = append(p.values[pos].rows,insertNodeValue.rows...)
    return false,p
  }

  //新規データ
  if p.IsOverFlow() == false {
    //オーバーフローなし
    p.insertValue(pos,insertNodeValue)
    return false,p
  }

  //オーバフローする
  if pos == 0 {
    p.createLeftNode()
    p.leftNode.Insert(insertNodeValue)
    return true,p
  }else if pos == p.dataCount {
    p.createRightNode()
    p.rightNode.Insert(insertNodeValue)
    return true,p
  } else {
    //minimumを取得して左ノードに再帰的にインサート
    minNode := p.popNodeValue(0)
    p.insertValue(pos-1,insertNodeValue)
    if p.leftNode == nil {
      p.createLeftNode()
      p.leftNode.Insert(minNode)
      return true,p
    }
    add,_ := p.leftNode.Insert(minNode)
    if add {
      return rebalance(p)
    }
    return false,p
  }
  return false,p
}


//Tnode削除
//TODO test
func(p *tnode) Delete(deleteValue Type) (ROWNUM,bool,*tnode) {
  if p.leftNode != nil && deleteValue.comp(p.minValue())<0  {
    deleteNum,del,_ := p.leftNode.Delete(deleteValue)
    delflg,newRoot := p.doAfterChildDelete(LEFT,del)
    return deleteNum,delflg,newRoot
  }
  if p.rightNode != nil && deleteValue.comp(p.maxValue())>0 {
    deleteNum,del,_ := p.rightNode.Delete(deleteValue)
    delflg,newRoot := p.doAfterChildDelete(RIGHT,del)
    return deleteNum,delflg,newRoot
  }
  isMatch,pos := p.getPosition(deleteValue)
  if isMatch == false {
    //該当データなし
    return 0,false,p
  }
  //削除
  deleteNum := p.deleteValue(pos)
  if p.IsInternalNode() && p.IsUnderFlow()  {
    //underflow時処理
    glb,del,_ := p.leftNode.popMaxNode()
    p.insertValue(0,glb)
    delflg,newRoot := p.doAfterChildDelete(LEFT,del)
    return deleteNum,delflg,newRoot
  }
  del := p.halfLeafMerge()
  return deleteNum,del,p
}
//最大値をpop
//TODO test
func(p *tnode) popMaxNode() (nodeValue,bool,*tnode){
  if p.rightNode != nil {
    retNode,del,_ := p.rightNode.popMaxNode()
    delflg,newRoot := p.doAfterChildDelete(RIGHT,del)
    return retNode,delflg,newRoot
  }
  retNode := p.popNodeValue(p.dataCount-1)
  del := p.halfLeafMerge()
  return retNode,del,p
}

//harfnodeのマージ
//戻り値　ノードの削除発生
func(p *tnode) halfLeafMerge() bool{
  isDel := false
  if p.IsHalfLeaf() {
    //half-leaf
    switch p.canMergeChildNode() {
      case MERGE_TYPE_LEFT:
        p.mergeFromLeftNode()
        isDel = true
      case MERGE_TYPE_RIGHT:
        p.mergeFromRightNode()
        isDel = true
    }
  }
  return isDel
}
//子ノードのDelete実行後処理
//空になったリーフの削除、リバランス
func (p *tnode) doAfterChildDelete(child ChildType,del bool) (bool,*tnode){
  newRoot := p
  delflg := del
  if child == LEFT {
    if p.leftNode.dataCount == 0 {
      p.leftNode = nil
      delflg = true
    }
  } else if child == RIGHT {
    if p.rightNode.dataCount == 0 {
      p.rightNode = nil
      delflg = true
    }
  }
  //リバランス
  if delflg {
    _,newRoot = rebalance(p)
  }
  return delflg,newRoot
}
//次に挿入したらOverFlowするか
func(p *tnode) IsOverFlow() bool{
  return p.dataCount  == p.t
}
//under flowしているか
func(p *tnode) IsUnderFlow() bool{
  return p.dataCount  <= p.t-3
}
func(p *tnode) IsInternalNode() bool{
  return p.leftNode != nil && p.rightNode != nil
}
func(p *tnode) IsLeafNode() bool{
  return p.leftNode == nil && p.rightNode == nil
}
func(p *tnode) IsHalfLeaf() bool{
  return p.IsInternalNode() == false && p.IsLeafNode() == false
}
//マージできる子ノードの有無
//0なし 1左　2右　3両方
func(p *tnode) canMergeChildNode() MergeType{
  canMerge := MERGE_TYPE_NONE
  if p.leftNode != nil && p.dataCount + p.leftNode.dataCount  <= p.t {
    canMerge += MERGE_TYPE_LEFT
  }
  if p.rightNode != nil && p.dataCount + p.rightNode.dataCount  <= p.t {
    canMerge += MERGE_TYPE_RIGHT
  }
  return canMerge
}
func(p *tnode) maxValue() Type{
  return p.values[p.dataCount-1].key
}
func(p *tnode) minValue() Type{
  return p.values[0].key
}
//指定ポジションをpop
func(p *tnode) popNodeValue(pos int) nodeValue{
  minNodeValue := p.values[pos]
  p.deleteValue(pos)
  return minNodeValue
}
//ノード内に値を挿入する
func(p *tnode) insertValue(insertPos int,insertNodeValue nodeValue) {
    for i:= p.dataCount;i > insertPos;i-- {
        p.values[i] = p.values[i-1]
    }
    p.values[insertPos] = insertNodeValue
    p.dataCount += 1
}
//ノード内の値を削除する
func(p *tnode) deleteValue(deletePos int) ROWNUM {
    rows := len(p.values[deletePos].rows)
    for i:= deletePos ; i < p.dataCount-1;i++ {
        p.values[i] = p.values[i+1]
    }
    //初期化
    p.values[p.dataCount-1] = nodeValue{}
    p.dataCount -= 1
    return ROWNUM(rows)
}
//ノード内の操作対象箇所を検索する
func(p *tnode) getPosition(searchValue Type) (bool,int) {
    return binarySearch(p.values,searchValue,0,p.dataCount-1)
}
//左ノードを作る
func(p *tnode) createLeftNode(){
    p.leftNode = createTnode(p.t)
    p.leftNode.parentNode = p
}
//右ノードを作る
func(p *tnode) createRightNode(){
    p.rightNode = createTnode(p.t)
    p.rightNode.parentNode = p
}

//左子ノードからマージ
func (p *tnode) mergeFromLeftNode(){
  p.mergeHead(p.leftNode,0)
  p.rightNode = p.leftNode.rightNode
  p.leftNode = p.leftNode.leftNode

}
//右子ノードからマージ
func (p *tnode) mergeFromRightNode(){
  p.mergeTail(p.rightNode,0)
  p.leftNode = p.rightNode.leftNode
  p.rightNode = p.rightNode.rightNode
}
//対象ノードを後ろ側にマージする
//srcNodeのstart番目以降をマージ
func (p *tnode) mergeTail(srcNode *tnode,start int){
  cnt := p.dataCount
  for i:= start; i < srcNode.dataCount ; i++{
    p.values[cnt+i-start] = srcNode.values[i]
  }
  p.dataCount = p.dataCount + srcNode.dataCount - start

  srcNode.clear(start)
}
//対象ノードを前側にマージする
//srcNodeのstart番目以降をマージ
func (p *tnode) mergeHead(srcNode *tnode,start int){
  cnt := srcNode.dataCount - start
  for i:= 0; i < p.dataCount ; i++{
    p.values[cnt+i] = p.values[i]
  }
  for i:= start; i < srcNode.dataCount ; i++{
    p.values[i-start] = srcNode.values[i]
  }
  p.dataCount = p.dataCount + srcNode.dataCount - start

  srcNode.clear(start)
}
//start番目以降をクリアする
func (p *tnode) clear(start int){
  for i:= start ;i < p.dataCount;i++ {
    p.values[i].clear()
  }
  p.dataCount =start
}
func (p *tnode) leftDepth() int{
  if p.leftNode != nil {
    return p.leftNode.depth() + 1
  }
  return 0
}
func (p *tnode) rightDepth() int{
  if p.rightNode != nil {
    return p.rightNode.depth() + 1
  }
  return 0
}
//木の深さを取得する
func (p *tnode) depth() int{
  leftDepth := p.leftDepth()
  rightDepth := p.rightDepth()
  //深い方
  if leftDepth > rightDepth {
    return leftDepth
  }
  return rightDepth
}

//ノード内の状態を出力する
func(p *tnode) Show() string {
    return p.showPadding(0)
}
func(p *tnode) showPadding(pad int) string {
    res := ""
    padding := strings.Repeat("-", pad)

    res += padding + "["
    for i:= 0;i < p.dataCount;i++ {
        res += toString(p.values[i].key) + "("
        res += strconv.Itoa(len(p.values[i].rows))
        res += "),"
    }
    res += "]\n"
    if p.leftNode !=nil {
      res += "l:" + p.leftNode.showPadding(pad+1)
    }
    if p.rightNode !=nil {
      res += "r:" + p.rightNode.showPadding(pad+1)
    }
    return res
}
func toString(value interface{}) string {
    res := ""
    switch val := value.(type){
        case Integer: res = strconv.Itoa(int(val))
        case Varchar: res =  string(val)
    }
    return res
}

//LLローテーション
func rotationLL(root *tnode) *tnode{
  //新たにrootになる
  newRoot := root.leftNode

  //左子を付け替え
  root.leftNode = newRoot.rightNode
  if root.leftNode != nil {
    root.leftNode.parentNode = root
  }

  //親を付け替え
  newRoot.parentNode = root.parentNode
  root.parentNode = newRoot
  newRoot.rightNode = root

  return newRoot
}
//RRローテーション
func rotationRR(root *tnode) *tnode{
  //新たにrootになる
  newRoot := root.rightNode

  //右子を付け替え
  root.rightNode = newRoot.leftNode
  if root.rightNode != nil {
    root.rightNode.parentNode = root
  }

  //親を付け替え
  newRoot.parentNode = root.parentNode
  root.parentNode = newRoot
  newRoot.leftNode = root

  return newRoot
}
//LRローテーション
func rotationLR(root *tnode) *tnode{
  //新たにrootになる
  newRoot := root.leftNode.rightNode
  //新たなleftNode
  newLeft := root.leftNode

  //leftNode(B)の付け替え
  newLeft.rightNode = newRoot.leftNode
  if newRoot.leftNode != nil {
    newRoot.leftNode.parentNode = newLeft
  }

  //rootNode(A)の付け替え
  root.leftNode = newRoot.rightNode
  if newRoot.rightNode != nil {
    newRoot.rightNode.parentNode = root
  }

  //newRoot(C)の付け替え
  newRoot.parentNode = root.parentNode

  newRoot.leftNode = newLeft
  newLeft.parentNode = newRoot

  newRoot.rightNode = root
  root.parentNode = newRoot

  //special lotation
  if newRoot.dataCount == 1 {
    newRoot.mergeHead(newLeft,1)
  }

  return newRoot
}
//RLローテーション
func rotationRL(root *tnode) *tnode{
  //新たにrootになる
  newRoot := root.rightNode.leftNode
  //新たなrightNode
  newRight := root.rightNode

  //rightNode(B)の付け替え
  newRight.leftNode = newRoot.rightNode
  if newRoot.rightNode != nil {
    newRoot.rightNode.parentNode = newRight
  }
  //rootNode(A)の付け替え
  root.rightNode = newRoot.leftNode
  if newRoot.leftNode != nil {
    newRoot.leftNode.parentNode = root
  }
  //newRoot(C)の付け替え
  newRoot.parentNode = root.parentNode

  newRoot.rightNode = newRight
  newRight.parentNode = newRoot

  newRoot.leftNode = root
  root.parentNode = newRoot

  //special lotation
  if newRoot.dataCount == 1 {
    newRoot.mergeTail(newRight,1)
  }

  return newRoot
}

//リバランスチェックし必要ならリバランス
//戻り値　リバランスしたらfalse,新しいroot
func rebalance(root *tnode) (bool,*tnode){
  newRoot :=root
  def := root.leftDepth() - root.rightDepth()
  if def > -2 && def < 2 {
    //差が2以内
    return true,root
  }
  //親ノード退避
  parent := root.parentNode

  if def > 0 {
    //左が深い
    if root.leftNode.leftDepth() >= root.leftNode.rightDepth() {
      //LL回転
      newRoot = rotationLL(root)
    } else {
      //LR回転
      newRoot = rotationLR(root)
    }
  }else {
    //右が深い
    if root.rightNode.rightDepth() >= root.rightNode.leftDepth() {
      //RR回転
      newRoot = rotationRR(root)
    } else {
      //RL回転
      newRoot = rotationRL(root)
    }
  }

  //親ノードのポインタ付け替え
  if parent != nil {
    if parent.leftNode == root {
      parent.leftNode = newRoot
    } else if parent.rightNode == root {
      parent.rightNode = newRoot
    }
  }
  return false,newRoot
}
