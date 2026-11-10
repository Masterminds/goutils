package goutils

import (
	"sync"
	"testing"
)

func TestRandomConcurrent(t *testing.T) {
	const workers = 8
	start := make(chan struct{})
	var workersDone sync.WaitGroup
	workersDone.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer workersDone.Done()
			<-start
			for j := 0; j < 100; j++ {
				value, err := Random(32, 'a', 'z'+1, false, false)
				if err != nil {
					t.Errorf("Random: %v", err)
					return
				}
				if len(value) != 32 {
					t.Errorf("Random returned %d characters, want 32", len(value))
				}
				for _, ch := range value {
					if ch < 'a' || ch > 'z' {
						t.Errorf("Random returned character %q outside requested range", ch)
					}
				}
			}
		}()
	}
	close(start)
	workersDone.Wait()
}
