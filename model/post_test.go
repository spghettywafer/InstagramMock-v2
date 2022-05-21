package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShortDescription(t *testing.T) {
	arg := []struct {
		request  Post
		expected string
	}{
		{
			request:  Post{Description: "This description length is more than 15 char"},
			expected: "This description...",
		},
		{
			request:  Post{Description: "Short desc"},
			expected: "Short desc",
		},
	}

	for i, post := range arg {
		actual := post.request.GetShortDescription()
		assert.Equal(t, actual, arg[i].expected)
	}
}
