package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestShouldMockObj(t *testing.T) {
	assert := assert.New(t)

	// mock
	random := new(randomMock)
	random.On("Random", 100).Return([]byte("hello"))

	assert.NotNil(random)
	assert.Equal([]byte("hello"), random.Random(100))
}

type Random interface {
	Random(limit int) int
}

type randomMock struct {
	mock.Mock
}

func (o randomMock) Random(limit int) []byte {
	args := o.Called(limit)
	return args.Get(0).([]byte)
}
