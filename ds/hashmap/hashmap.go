package hashmap

import (
	"encoding/json"
	"golang-katas/ds/linkedlist/doublylinkedlist"
	"hash/fnv"
	"iter"
	"maps"
	"math"
)

const (
	defaultSize                = 17 // prime
	defaultLoadFactorThreshold = 0.75
)

func isPrime(n int) bool {
	if n < 0 {
		n = -n
	}
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int) int {
	for ; ; n++ {
		if isPrime(n) {
			return n
		}
	}
}

type entry[K string, V any] struct {
	key K
	val V
}

// HashMap is an implementation of a hash map using the FNV-1a algorithm.
type HashMap[K string, V any] struct {
	buckets             []*doublylinkedlist.LinkedList[*entry[K, V]]
	NumElements         int
	loadFactorThreshold float64
}

func (hm HashMap[K, V]) hash(key K, size uint32) uint {
	hashVal := fnv.New32a()
	hashVal.Write([]byte(key))
	return uint(hashVal.Sum32() % size)
}

// Size returns the number of elements (buckets) in the underlying slice of
// linked lists.
func (hm HashMap[K, V]) Size() int {
	return len(hm.buckets)
}

func (hm HashMap[K, V]) loadFactor() float64 {
	return float64(hm.NumElements) / float64(hm.Size())
}

// SetLoadFactorThreshold allows one to set the load factor
// that is used to decide when the hash map needs rehashing to
// a larger bucket size.
func (hm *HashMap[K, V]) SetLoadFactorThreshold(lf float64) {
	hm.loadFactorThreshold = lf
}

func (hm *HashMap[K, V]) resize() {
	curSize := len(hm.buckets)
	newSize := nextPrime(curSize*2 + 1)
	newBuckets := make([]*doublylinkedlist.LinkedList[*entry[K, V]], newSize)
	for i := range newSize {
		newBuckets[i] = doublylinkedlist.New[*entry[K, V]]()
	}
	for i := range curSize {
		for entry := range hm.buckets[i].IterVals() {
			hash := hm.hash(entry.key, uint32(newSize))
			newBuckets[hash].Append(entry)
		}
	}
	hm.buckets = newBuckets
}

// Put inserts the value val so that it can be accessed via
// the key key.
func (hm *HashMap[K, V]) Put(key K, val V) {
	entry := entry[K, V]{key, val}
	hash := hm.hash(key, uint32(len(hm.buckets)))
	bucketList := hm.buckets[hash]
	for entry := range bucketList.IterVals() {
		if entry.key == key {
			entry.val = val
			return
		}
	}
	bucketList.Append(&entry)
	hm.NumElements++
	if hm.loadFactor() > hm.loadFactorThreshold {
		hm.resize()
	}
}

// Get returns the value in the HashMap associated with
// the key. If the key does not exist in the HashMap,
// the zero value is returned, along with a false boolean
// indicating it could not be found.
func (hm *HashMap[K, V]) Get(key K) (V, bool) {
	var val V
	found := false
	hash := hm.hash(key, uint32(len(hm.buckets)))
	for entry := range hm.buckets[hash].IterVals() {
		if entry.key == key {
			val = entry.val
			found = true
			break
		}
	}
	return val, found
}

// Get returns the value in the HashMap associated with
// the key, and removes the key from the HashMap.
// If the key does not exist in the HashMap,
// the zero value is returned, along with a false boolean
// indicating it could not be found.
func (hm *HashMap[K, V]) Pop(key K) (V, bool) {
	idx := -1
	hash := hm.hash(key, uint32(len(hm.buckets)))
	list := hm.buckets[hash]
	for entry := range list.IterVals() {
		idx++
		if entry.key == key {
			_, err := list.Remove(idx)
			hm.NumElements--
			return entry.val, err == nil
		}
	}
	var val V
	return val, false
}

// Items returns an iterator over the key value pairs
// comprising the HashMap.
func (hm HashMap[K, V]) Items() iter.Seq2[K, V] {
	return func(yield func(key K, val V) bool) {
		for _, bucket := range hm.buckets {
			for entry := range bucket.IterVals() {
				if !yield(entry.key, entry.val) {
					return
				}
			}
		}
	}
}

func (hm HashMap[K, V]) toMap() map[K]V {
	asMap := map[K]V{}
	maps.Insert(asMap, hm.Items())
	return asMap
}

// String returns a JSON/python dictionary style representation
// of the key value pairs in the HashMap.
func (hm HashMap[K, V]) String() string {
	repr, _ := json.MarshalIndent(hm.toMap(), "", "  ")
	return string(repr)
}

// NewHashMap initializes a HashMap with defaultSize buckets,
// and returns a pointer to it.
func NewHashMap[K string, V any]() *HashMap[K, V] {
	buckets := make([]*doublylinkedlist.LinkedList[*entry[K, V]], defaultSize)
	for i := range len(buckets) {
		buckets[i] = doublylinkedlist.New[*entry[K, V]]()
	}
	return &HashMap[K, V]{
		buckets:             buckets,
		loadFactorThreshold: defaultLoadFactorThreshold,
	}
}
