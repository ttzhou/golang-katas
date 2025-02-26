package hashmap_test

import (
	"golang-katas/ds/hashmap"
	"golang-katas/internal/assert"
	"strings"
	"testing"
)

func TestHashmap(t *testing.T) {

	assert := assert.New(t)

	hm := hashmap.NewHashMap[string, int]()
	hm.Put("a", 5)
	val, ok := hm.Get("a")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(5)

	hm.Put("a", 5)

	val, ok = hm.Get("a")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(5)

	hm.Put("a", 6)
	val, ok = hm.Get("a")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(6)

	hm.Put("b", 8)
	val, ok = hm.Get("b")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(8)

	val, ok = hm.Pop("b")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(8)

	_, ok = hm.Pop("b")
	assert.That(ok).IsFalse()

	hm.SetLoadFactorThreshold(0.01)
	hm.Put("b", 7)
	val, ok = hm.Get("b")
	assert.That(ok).IsTrue()
	assert.That(val).Equals(7)

	for range hm.Items() {
		break
	}

	for key, val := range hm.Items() {
		retrievedVal, ok := hm.Get(key)
		assert.That(ok).IsTrue()
		assert.That(retrievedVal).Equals(val)
	}

	assert.That(strings.Contains(hm.String(), `"a": 6`))
	assert.That(strings.Contains(hm.String(), `"b": 7`))
}
