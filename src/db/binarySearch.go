package db

import (

)

func binarySearchInteger(values []nodeValueInteger,searchValue Integer,head int,tail int) (bool,int){
    //再帰なし
    for ;; {
      if head > tail {
        return false,head
      }
      pivot := (head+tail)/2
      if values[pivot].key == searchValue {
        return true,pivot
      } else if values[pivot].key > searchValue{
        tail = pivot-1
      } else {
         head = pivot+1
      }
    }
    return false,head
}
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
func binarySearchInt(values []nodeValueInt,searchValue int,head int,tail int) (bool,int){
    //再帰なし
    for ;; {
      if head > tail {
        return false,head
      }
      pivot := (head+tail)/2
      if values[pivot].key == searchValue {
        return true,pivot
      } else if values[pivot].key > searchValue {
        tail = pivot-1
      } else {
        head = pivot+1
      }
    }
    return false,head
}
func binarySearchIntArray(values []Integer,searchValue Integer,head int,tail int) (bool,int){
    //再帰なし
    for ;; {
      if head > tail {
        return false,head
      }
      pivot := (head+tail)/2
      if values[pivot] == searchValue {
        return true,pivot
      } else if values[pivot] > searchValue {
        tail = pivot-1
      } else {
        head = pivot+1
      }
    }
    return false,head
}