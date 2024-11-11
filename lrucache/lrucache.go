package lrucache

type LruCache struct {
	// Add your fields here
}

// Creates a new LruCache with the given capacity.
func New(capacity int) *LruCache {
	return nil
}

// Get returns value associated with the key.
//
// The second value is a bool that is true if the key exists in the cache,
// and false if not.
func (l *LruCache) Get(key int) (int, bool) {
	return 0, false
}

// Set updates value associated with the key.
//
// If there is no key in the cache new (key, value) pair is created.
func (l *LruCache) Set(key, value int) {
}

// Range calls function f on all elements of the cache
// in increasing access time order.
//
// Stops earlier if f returns false.
func (l *LruCache) Range(f func(key, value int) bool) {
}

// Clear removes all elements from the cache.
func (l *LruCache) Clear() {
}
