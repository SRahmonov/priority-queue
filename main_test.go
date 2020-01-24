package main

import (
	"testing"
)

func Test_queue_len(t *testing.T) {
	l := queue{}
	l.equeue(0)
	if l.len() != 1 {
		t.Error("after adding one element size must be 1, got:",l.len())
	}
}

func Test_queue_first_last_items(t *testing.T) {
	l := queue{}
	l.equeue(1)
	if l.first() != l.last(){
		t.Error("after adding one element the first and last should match, got:", l.first())
	}
}
func Test_queue_delete_items(t *testing.T)  {
	l := queue{}
	l.equeue(1)
	l.dequeue()
	if l.first() != nil{
		t.Error("after deleting the first element it should be empty, got:",l.first())
	}
}