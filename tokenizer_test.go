package toks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	tok := NewTokenizer("2 + 2")
	assert.Equal(t, Number.token("2"), tok.NextToken())
	assert.Equal(t, Plus.token("+"), tok.NextToken())
	assert.Equal(t, Number.token("2"), tok.NextToken())
	assert.Equal(t, EOF.token(""), tok.NextToken())
}

func TestTokens(t *testing.T) {
	tok := NewTokenizer("2 + 2")
	assert.Equal(t, Number.token("2"), <-tok.Tokens())
	assert.Equal(t, Plus.token("+"), <-tok.Tokens())
	assert.Equal(t, Number.token("2"), <-tok.Tokens())
	assert.Equal(t, EOF.token(""), <-tok.Tokens())
}
