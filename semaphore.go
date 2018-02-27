// See http://www.golangpatterns.info/concurrency/semaphores
package main

// Semaphore is a semaphore for acquiring and releasing resources.
type Semaphore chan struct{}

// Acquire n resources.
func (s Semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

// Release n resources.
func (s Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
