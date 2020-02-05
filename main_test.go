package main

import (
	"testing"
)

func Test_for_empty_queue(t *testing.T) {
	l := queue{}
	l.equeue(0)
	if l.len() == 0 {
		t.Error("empty queue length must be nil, got:", l.len())
	}
	if l.first() == nil {
		t.Error("first element of queue must have value of 0, got:", l.first())
	}
	if l.last() == nil {
		t.Error("last element of queue must have value of 0, got:", l.last())
	}
	if l.last() != l.first() {
		t.Error("first and last elements of one item queue must have same priorities, got:", l.last(), ";", l.first())
	}
}

func Test_queue_with_one_item(t *testing.T) {
	l := queue{}
	l.equeue(1)
	if l.len() != 1 {
		t.Error("after adding one item queue length must be 1, got:", l.len())
	}
	if l.first() != 1 {
		t.Error("first element of queue must have value of 1, got:", l.first())
	}
	if l.last() != 1 {
		t.Error("last element of queue must have value of 1, got:", l.last())
	}
	if l.last() != l.first() {
		t.Error("first and last elements of one item queue must have same values, got:", l.last(), ";", l.first())
	}
}
func Test_queue_with_multiple_item(t *testing.T) {
	l := queue{}
	l.equeue(1)
	l.equeue(2)
	l.equeue(3)
	if l.len() != 3 {
		t.Error("after adding three item queue length must be 2, got:", l.len())
	}
	if l.first() != 1 {
		t.Error("first element of queue must have value of 1, got:", l.first())
	}
	if l.last() == 3 {
		t.Error("last element of queue must have value of 2, got:", l.last())
	}
}

func Test_queue_first_last_items(t *testing.T) {
	l := queue{}
	l.equeue(1)
	if l.first() != l.last() {
		t.Error("after adding one element the first and last should match, got:", l.first())
	}
}
func Test_queue_delete_items(t *testing.T) {
	l := queue{}
	l.equeue(1)
	l.dequeue()
	if l.len() != 0 {
		t.Error("after deleting the first element it should be nil, got:", l.len())
	}
}

