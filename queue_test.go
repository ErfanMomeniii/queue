package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFront(t *testing.T) {
	q := New()

	q.PushFront(4)
	assert.Equal(t, q.Front(), 4)

	q.PushFront("hi")
	assert.Equal(t, q.Front(), "hi")

	q.PopFront()
	assert.Equal(t, q.Front(), 4)

	q.PopFront()
	assert.Equal(t, q.Front(), nil)
}

func TestBack(t *testing.T) {
	q := New()

	q.PushBack(4)
	assert.Equal(t, q.Back(), 4)

	q.PushBack("hi")
	assert.Equal(t, q.Back(), "hi")

	q.PopBack()
	assert.Equal(t, q.Back(), 4)

	q.PopBack()
	assert.Equal(t, q.Back(), nil)
}

func TestCombine(t *testing.T) {
	q := New()

	q.PushBack(4)
	q.PopBack()
	assert.Equal(t, q.Front(), nil)
	assert.Equal(t, q.Back(), nil)

	q.PushBack("hi")
	q.PushFront(4)
	q.PushBack(3.14)
	assert.Equal(t, q.Back(), 3.14)
	assert.Equal(t, q.Front(), 4)
}

func TestSize(t *testing.T) {
	q := New()
	assert.Equal(t, q.Size(), 0)

	q.PushBack(2)
	q.PushBack("hi")
	assert.Equal(t, q.Size(), 2)

	q.PopFront()
	assert.Equal(t, q.Size(), 1)
}

func TestList(t *testing.T) {
	q := New()
	assert.Equal(t, q.List(), []any(nil))

	q.PushBack(2)
	q.PushBack("hi")
	assert.Equal(t, q.List(), []any{2, "hi"})

	q.PopFront()
	assert.Equal(t, q.List(), []any{"hi"})
}
