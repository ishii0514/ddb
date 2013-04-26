package db

import (
    "testing"
)

func TestTypeComp(t *testing.T) {
  int5 := Integer(5)
  int5_2 := Integer(5)
  if int5.comp(int5_2) != 0 {
        t.Error("illegal integer equqls")
    }

    int6 := Integer(6)
  if int5.comp(int6) != -1 {
        t.Error("illegal integer equqls")
    }

    varchar1 := Varchar("abc")
    varchar2 := Varchar("abc")
    varchar3 := Varchar("aac")
    varchar4 := Varchar("bbc")
    if varchar1.comp(varchar2) != 0 {
        t.Error("illegal varchar comp 0")
    }
    if varchar1.comp(varchar3) != 1 {
        t.Error("illegal varchar comp 1")
    }
    if varchar1.comp(varchar4) != -1 {
        t.Error("illegal varchar comp -1")
    }
}
/*
func TestTypeClear(t *testing.T) {
  int5 := Integer(5)
  int5.clear()
  if int5 != 0 {
        t.Error("illegal Integer clear")
    }

    varchar := Varchar("あいうえお")
    if varchar != "あいうえお" {
        t.Error("illegal Varchar before clear")
    }
  varchar.clear()
  if varchar != "" {
        t.Error("illegal Varchar clear")
    }
}
*/