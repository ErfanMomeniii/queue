package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFront(t *testing.T) {
	q := New()

	q.PushFront(4)
	assert.Equal(t, q.GetFront(), 4)

	q.PushFront("hi")
	assert.Equal(t, q.GetFront(), "hi")

	q.PopFront()
	assert.Equal(t, q.GetFront(), 4)

	q.PopFront()
	assert.Equal(t, q.GetFront(), nil)
}

func TestBack(t *testing.T) {
	q := New()

	q.PushBack(4)
	assert.Equal(t, q.GetBack(), 4)

	q.PushBack("hi")
	assert.Equal(t, q.GetBack(), "hi")

	q.PopBack()
	assert.Equal(t, q.GetBack(), 4)

	q.PopBack()
	assert.Equal(t, q.GetBack(), nil)
}

func TestCombine(t *testing.T) {
	q := New()

	q.PushBack(4)
	q.PopBack()
	assert.Equal(t, q.GetFront(), nil)
	assert.Equal(t, q.GetBack(), nil)

	q.PushBack("hi")
	q.PushFront(4)
	q.PushBack(3.14)
	assert.Equal(t, q.GetBack(), 3.14)
	assert.Equal(t, q.GetFront(), 4)
}