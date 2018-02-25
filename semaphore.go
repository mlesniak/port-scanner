// See http://www.golangpatterns.info/concurrency/semaphores
package main


// Empty is a dummy content.
type Empty int

// Semaphore is a semaphore for acquiring and releasing resources.
type Semaphore chan Empty


// Acquire n resources.
func (s Semaphore) Acquire(n int) {
    e := Empty(0)
    for i := 0; i < n; i++ {
        s <- e
    }
}

// Release n resources.
func (s Semaphore) Release(n int) {
    for i := 0; i < n; i++ {
        <-s
    }
}




