package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "someoneishere"
	pathname := CASPathTransformFunc(key)
	expectedPathName := "31edd/d6655/7457d/df20d/7c896/45a68/63bfd/b4678"

	if pathname != expectedPathName {
		t.Errorf("pathname is %q, expected %q", pathname, expectedPathName)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("hello world"))
	if err := s.writeStream("someone", data); err != nil {
		t.Error(err)
	}
}
