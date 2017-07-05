package queue

import (
	"testing"
)

var q = NewQueue()

func TestPush(t *testing.T) {
	q.Push("test")
	if q.Len() == 0 {
		t.Error("expected 1")
	}
}

func testPop(t *testing.T) {
	q.Push("test")
	pop := q.Pop()

	if pop != "test" {
		t.Error("expected 'test'")
	}

	if q.Len() != 0 {
		t.Error("expected 0")
	}
}

func testLen(t *testing.T) {
	if q.Len() == 0 {
		q.Push("test")
	}

	if q.Len() != 1 {
		t.Error("expected 1")
	}

}
