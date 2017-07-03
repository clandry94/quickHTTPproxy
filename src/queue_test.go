package queue

import (
	"testing"
)

func createSampleQueue() Queue {
	q := make(Queue, 0)
	return q
}

func TestPush(t *testing.T) {
	queue := createSampleQueue()
	queue.Push("test")
	if queue.Len() == 0 {
		t.Error("expected 1")
	}
}

func testPop(t *testing.T) {
	queue := createSampleQueue()
	queue.Push("test")
	pop := queue.Pop()

	if pop != "test" {
		t.Error("expected 'test'")
	}

	if queue.Len() != 0 {
		t.Error("expected 0")
	}
}

func testLen(t *testing.T) {

}
