package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "nigger"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "85fe8de475bc9884da850bb5ac9dedaa50a5f850"
	expectedPathName := "85fe8/de475/bc988/4da85/0bb5a/c9ded/aa50a/5f850"
	if pathKey.pathname != expectedPathName {
		t.Errorf("have : %s, want : %s\n", pathKey.pathname, expectedPathName)
	}
	if pathKey.original != expectedOriginalKey {
		t.Errorf("have : %s, want : %s\n", pathKey.original, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("black niggers bytes"))
	if err := s.writeStream("niggers", data); err != nil {
		t.Error(err)
	}
}
