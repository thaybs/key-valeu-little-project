package main

import ( // impots libraries
	"encoding/json"
	"io/ioutil" // writes e reads files
	"sync"      // synchronization (mutex)
)

// KVStore is a structure that represents the key-value store
type KVStore struct {
	store map[string]string // Stores the key-value pairs
	mutex sync.RWMutex      // Mutex to ensure safe concurrent access
}

//NewKVStore creates a new instance of KVStore
func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]string), //Initializa the storage map
	}
}

// Set adds or updates a key-value pair in the store
func (kv *KVStore) Set(key, value string) {
	kv.mutex.Lock()         // Lock the mutex for writing
	defer kv.mutex.Unlock() // Unlock the mutex after the operation
	kv.store[key] = value   // Add or update the value in the map
}

// Get retrieves the value for a given key
func (kv *KVStore) Get(key string) (string, bool) {
	kv.mutex.RLock()               // Lock the mutex for reading
	defer kv.mutex.RUnlock()       // Unlock mutex
	value, exists := kv.store[key] // Retrieve the value and check if the key exists
	return value, exists           // Return the value and wheter it exists
}

//Delete removes a key-value pair from the store
func (kv *KVStore) Delete(key string) {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	delete(kv.store, key) // Remove the key from the map
}

// SaveToFile saves the current state of the store to a file
func (kv *KVStore) SaveToFile(filename string) error {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()
	data, err := json.Marshal(kv.store) // Enciode the map to JSON
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644) // Write the JSON to a file
}

// LoadFromFile loads the state of the store from a file
func (kv *KVStore) LoadFromFile(filename string) error {
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	data, err := ioutil.ReadFile(filename) // Read the data from the file
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &kv.store) // Decode the JSON into the map
}
