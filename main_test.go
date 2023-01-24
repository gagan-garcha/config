package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetValue(t *testing.T) {

	const path = "test_files"

	info := Run(path)

	tests := []struct {
		name        string
		shouldError bool
		key         string
		value       any
	}{
		{
			name:        "successful",
			shouldError: false,
			key:         "environment",
			value:       "development",
		},
		{
			name:        "successful",
			shouldError: false,
			key:         "database",
			value: map[string]interface{}{
				"host":     "127.0.0.1",
				"port":     float64(3306),
				"username": "divido",
				"password": "divido",
			},
		},
		{
			name:        "error",
			shouldError: true,
			key:         "database.kk",
			value:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			val, err := info.get(test.key)

			assert.DeepEqual(t, test.value, val)
			assert.Equal(t, test.shouldError, err != nil)

		})
	}

}
