package prque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrque(t *testing.T) {
	queue := New(nil)
	assert.True(t, queue.Empty())
	queue.Push("item1", 1)
	queue.Push("item5", 5)
	queue.Push("item3", 3)
	queue.Push("item2", 2)
	queue.Push("item4", 4)
	assert.False(t, queue.Empty())
	assert.Equal(t, queue.Size(), 5)
	value, priority := queue.Pop()
	assert.Equal(t, value, "item5")
	assert.Equal(t, priority, int64(5))
	assert.Equal(t, queue.Size(), 4)

	value = queue.PopItem()
	assert.Equal(t, value, "item4")
	assert.Equal(t, queue.Size(), 3)

	queue.Remove(0) // remove item3
	value, priority = queue.Pop()
	assert.Equal(t, value, "item2")
	assert.Equal(t, priority, int64(2))
	assert.Equal(t, queue.Size(), 1)

	queue.Reset()
	assert.True(t, queue.Empty())
	assert.Equal(t, queue.Size(), 0)
}
