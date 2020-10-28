package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Helpers_FormatLatency(t *testing.T) {
	t.Parallel()

	cases := []struct {
		d        time.Duration
		expected time.Duration
	}{
		{time.Millisecond * 123456, time.Millisecond * 123450},
		{time.Millisecond * 12340, time.Millisecond * 12340},
		{time.Microsecond * 123456, time.Microsecond * 123450},
		{time.Microsecond * 123450, time.Microsecond * 123450},
		{time.Nanosecond * 123456, time.Nanosecond * 123450},
		{time.Nanosecond * 123450, time.Nanosecond * 123450},
		{time.Nanosecond * 123, time.Nanosecond * 123},
	}

	for _, tc := range cases {
		t.Run(tc.d.String(), func(t *testing.T) {
			assert.Equal(t, formatLatency(tc.d), tc.expected)
		})
	}
}
