package hello_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/snktagarwal/go-helloworld"
)

func TestSayPlease(t *testing.T) {
    inst := hello.SayHello{}
    assert.Equal(t, inst.SayPlease(), "SayPlease")
}