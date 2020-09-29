package macd

import (
	"fmt"
	"testing"

	"github.com/cpurta/tatanka/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMACDCalculateError(t *testing.T) {
	macd := NewMACD(12, 26)

	periods := []*model.Period{
		{
			Close: 11.84,
		},
		{
			Close: 11.75,
		},
	}

	_, err := macd.Calculate(periods)

	assert.Equal(t, err, fmt.Errorf("must provided at least %d periods to calculate MACD indicator", 26))
}
