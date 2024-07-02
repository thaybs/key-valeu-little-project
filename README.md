# Go Key-Value Store

A simple in-memory key-value store implemented in Go. This project demonstrates basic operations such as setting, getting, and deleting key-value pairs with concurrency safety.

## Features

- Add or update key-value pairs.
- Retrieve values by key.
- Delete key-value pairs.
- Concurrency safe.

## Usage

### Installation

First, clone the repository:


```
git clone https://github.com/your-username/go-key-value-store.git
cd go-key-value-store
```

### Running the Store

You can run the example usage by executing:

```sh
go run main.go 
```

### Example Code

Here's how you can use the key-value store:

```go
package main

import (
    "fmt"
    "sync"
)

// KVStore represents a thread-safe key-value store
type KVStore struct {
    store map[string]string
    mutex sync.RWMutex
}

// NewKVStore creates a new KVStore
func NewKVStore() *KVStore {
    return &KVStore{
        store: make(map[string]string),
    }
}

// Set adds or updates a key-value pair in the store
func (kv *KVStore) Set(key, value string) {
    kv.mutex.Lock()
    defer kv.mutex.Unlock()
    kv.store[key] = value
}

// Get retrieves the value for a given key
func (kv *KVStore) Get(key string) (string, bool) {
    kv.mutex.RLock()
    defer kv.mutex.RUnlock()
    value, exists := kv.store[key]
    return value, exists
}

// Delete removes a key-value pair from the store
func (kv *KVStore) Delete(key string) {
    kv.mutex.Lock()
    defer kv.mutex.Unlock()
    delete(kv.store, key)
}

func main() {
    store := NewKVStore()
    
    store.Set("name", "Google")
    value, exists := store.Get("name")
    if exists {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Not found")
    }
    
    store.Delete("name")
    value, exists = store.Get("name")
    if exists {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Not found")
    }
}
```

## Contributing

Feel free to fork this project, submit pull requests, and contribute to its development.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.