package benchmarks

import (
	"container/list"
	"github.com/erfanmomeniii/queue"
	"testing"
)

func BenchmarkCompare(b *testing.B) {
	b.Logf("Adding an element to the queue")
	b.Run("container/list", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l := list.New()
				l.PushFront("hi")
			}
		})
	})
	b.Run("erfanmomeniii/queue", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q := queue.New()
				q.PushFront("hi")
			}
		})
	})

	b.Logf("Removing an element from the queue")
	b.Run("container/list", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l := list.New()
				l.PushFront(4)
				l.Remove(l.Front())
			}
		})
	})
	b.Run("erfanmomeniii/queue", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q := queue.New()
				q.PushFront(4)
				q.PopFront()
			}
		})
	})
}
