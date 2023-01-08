/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sync

// Note that sync.Mutex works exactly like a Semaphore of size 1:
// A sync.Mutex still uses channels to guard access to shared resources;
// i.e., there is no magic behind-the-scenes.

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	// Writes will only succeed if there is room in s.
	s <- struct{}{}
}

// TryLock is like Lock, but it immediately returns whether
// it was able to lock or not without waiting.
func (s Semaphore) TryLock() bool {
	// Select with default case:
	// if no cases are ready just fall in the default block.
	select {
	case s <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s Semaphore) Unlock() {
	// Make room for other users of the semaphore.
	<-s
}

type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)
	// Allow for exactly one read.
	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	// Read from a closed chan always succeeds.
	// This only blocks during initialization.
	_, ok := <-o
	if !ok {
		// Channel is closed, early return.
		// f must have already returned.
		return
	}

	// Call f.
	// Only one goroutine will get here
	// as there is only one value in the channel.
	f()

	// This unleashes all waiting goroutines and future ones.
	close(o)
}
