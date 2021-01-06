package week06

import (
	"sync"
	"time"
)

type Config int64

type Bucket struct {
	Count int64
}

type Window struct {
	Mutex  *sync.RWMutex
	window []Bucket
	size   Config
}

func NewWindows(config Config) *Window {
	buckets := make([]Bucket, config)
	for offset := range buckets {
		buckets[offset] = Bucket{}
	}
	return &Window{window: buckets, size: config}
}

func (w *Window) getCurrentBucket() Bucket {
	size := int64(w.size)
	now := time.Now().Unix() % size
	return w.window[now]
}

func (w *Window) Increment(cur int64) {
	if cur <= 0 {
		return
	}
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	b := w.getCurrentBucket()
	b.Count += cur
}

func main() {
	config := Config(3)
	window := NewWindows(config)
	for _, request := range []int64{1, 2, 3, 4, 5} {
		window.Increment(request)
		time.Sleep(1 * time.Second)
	}
}
