package gmap

import "sync"

// StrAnyMap is a concurrent-safe map implementation with string keys and any-type values.
// It provides thread-safe operations through RWMutex synchronization.
// The zero value is not usable - use NewStrAnyMap or NewStrAnyMapFrom to create instances.
type StrAnyMap struct {
	mu   sync.RWMutex   // Mutex to protect concurrent access
	data map[string]any // Underlying data storage
}

// NewStrAnyMap creates and returns a new empty StrAnyMap.
// The optional `safe` parameter is currently unused but reserved for future concurrency control.
func NewStrAnyMap(safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		data: make(map[string]any), // Initialize empty map
	}
}

// NewStrAnyMapFrom creates a StrAnyMap from an existing map.
// Note: This uses the provided map directly (no deep copy), so external modifications
// to the original map may cause concurrency issues if the StrAnyMap is used concurrently.
func NewStrAnyMapFrom(data map[string]any, safe ...bool) *StrAnyMap {
	return &StrAnyMap{
		data: data, // Use provided map directly
	}
}

// Iterator provides read-only iteration over the map elements.
// The callback function `f` receives each key-value pair.
// Iteration stops if `f` returns false.
func (m *StrAnyMap) Iterator(f func(k string, v any) bool) {
	for k, v := range m.Map() { // Use Map() to get a safe copy
		if !f(k, v) {
			break
		}
	}
}

// Map returns a copy of the underlying map data.
// This ensures thread-safety by returning a copy rather than the original map.
func (m *StrAnyMap) Map() map[string]any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data := make(map[string]any, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny is an alias for Map() that returns a copy of the underlying data.
func (m *StrAnyMap) MapStrAny() map[string]any {
	return m.Map()
}

// MapCopy is another alias for Map() that returns a copy of the underlying data.
func (m *StrAnyMap) MapCopy() map[string]any {
	return m.Map()
}

// Set adds or updates a key-value pair in the map.
func (m *StrAnyMap) Set(key string, val any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data == nil {
		m.data = make(map[string]any)
	}
	m.data[key] = val
}

// Sets adds multiple key-value pairs to the map in a single operation.
func (m *StrAnyMap) Sets(data map[string]any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data == nil {
		m.data = data
	} else {
		for k, v := range data {
			m.data[k] = v
		}
	}
}

// Search looks for a key and returns its value along with an existence flag.
func (m *StrAnyMap) Search(key string) (value any, found bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.data != nil {
		value, found = m.data[key]
	}
	return
}

// Get returns the value associated with the given key.
// Returns nil if key doesn't exist.
func (m *StrAnyMap) Get(key string) (value any) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.data != nil {
		value = m.data[key]
	}
	return
}

// GetString returns the string value associated with the key.
// Returns empty string if key doesn't exist or value isn't a string.
func (m *StrAnyMap) GetString(key string) string {
	value := m.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

// Pop removes and returns an arbitrary key-value pair from the map.
func (m *StrAnyMap) Pop() (key string, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, value = range m.data {
		delete(m.data, key)
		return
	}
	return
}

// Pops removes and returns up to `size` key-value pairs from the map.
// If size is -1, removes and returns all items.
func (m *StrAnyMap) Pops(size int) map[string]any {
	m.mu.Lock()
	defer m.mu.Unlock()

	if size > len(m.data) || size == -1 {
		size = len(m.data)
	}
	if size == 0 {
		return nil
	}

	result := make(map[string]any, size)
	count := 0
	for k, v := range m.data {
		delete(m.data, k)
		result[k] = v
		count++
		if count == size {
			break
		}
	}
	return result
}

// SetIfNotExistFunc sets the value for key if it doesn't exist, using the result of function f.
// Returns true if the value was set, false if key already existed.
func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() any) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock is similar to SetIfNotExistFunc but executes the function f
// while holding the mutex lock for thread-safety.
func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() any) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Removes deletes multiple keys from the map in one operation.
func (m *StrAnyMap) Removes(keys []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
}

// Remove deletes a single key from the map and returns its value if it existed.
func (m *StrAnyMap) Remove(key string) (value any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data != nil {
		var ok bool
		if value, ok = m.data[key]; ok {
			delete(m.data, key)
		}
	}
	return
}

// Keys returns all keys in the map as a slice.
func (m *StrAnyMap) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values in the map as a slice.
func (m *StrAnyMap) Values() []any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	values := make([]any, 0, len(m.data))
	for _, value := range m.data {
		values = append(values, value)
	}
	return values
}

// Contains checks if a key exists in the map.
func (m *StrAnyMap) Contains(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.data != nil {
		_, ok := m.data[key]
		return ok
	}
	return false
}

// doSetWithLockCheck is an internal helper function that safely checks and sets values.
// It handles both direct values and value-generating functions.
func (m *StrAnyMap) doSetWithLockCheck(key string, value any) any {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data == nil {
		m.data = make(map[string]any)
	}

	// Return existing value if key exists
	if v, ok := m.data[key]; ok {
		return v
	}

	// Handle function values
	if f, ok := value.(func() any); ok {
		value = f()
	}

	// Set the new value
	if value != nil {
		m.data[key] = value
	}
	return value
}

// GetOrSetFuncLock gets the value for a key, or sets it using the provided function if it doesn't exist.
// The function is executed while holding the lock for thread-safety.
func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() any) any {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}
