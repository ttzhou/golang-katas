package assert_test

import (
	"golang-katas/internal/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert := assert.New(t)

	t.Run("Equals())", func(t *testing.T) {
		assert.That(5).Equals(5)
		assert.That("a").Equals(5)
		assert.That("a").IsNil()
		assert.That("a").IsTrue()
		assert.That(true).IsTrue()
		assert.That(1).IsTrue()
	})
}
