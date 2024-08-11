
package starter_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  starter "go_test_starter"
)

func TestSayHello(t *testing.T) {
  greeting := starter.SayHello("William")
  assert.Equal(t, "Hello William. Welcome!", greeting)
  
  another_greeting := starter.SayHello("asdf ghjkl")
  assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}


func TestOddOrEven(t *testing.T) {
  t.Run("Check Non Negative Numbers", func(t *testing.T) {
    assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
    assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
    assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
  t.Run("Check Negative Numbers", func(t *testing.T) {
    assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
    assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
  })
}

func TestFunc(t *testing.T) {
  // place common values to all tests
  t.Run("subtest 1", func(t *testing.T) {
    // place common values to all subtest 1
    t.Run("subsubtest a", func(t *testing.T) {
      // place common values to all subsubtest a
    })
    t.Run("subsubtest b", func(t *testing.T) {
      //
    })
  })
  t.Run("subtest 2", func(t *testing.T) {
    //
  })
}

