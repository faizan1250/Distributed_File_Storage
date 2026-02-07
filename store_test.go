package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "nigger"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "85fe8de475bc9884da850bb5ac9dedaa50a5f850"
	expectedPathName := "85fe8/de475/bc988/4da85/0bb5a/c9ded/aa50a/5f850"
	if pathKey.Pathname != expectedPathName {
		t.Errorf("have : %s, want : %s\n", pathKey.Pathname, expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have : %s, want : %s\n", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "niggers"
	data := []byte("black niggers bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "niggers"
	data := []byte("black niggers bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	if string(b) != string(data) {
		t.Errorf("want: %s, have: %s", data, b)
	}

	s.Delete(key)
}
