package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPositiveNum(t *testing.T) {
	a := 3
	b := 3

	expected := 6

	result := Sum(a, b)

	assert.Equal(t, expected, result)

}
