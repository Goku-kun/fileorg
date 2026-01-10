package organizer

import (
	"testing"
	"time"
)

func TestExtensionStrategy_Categorize(t *testing.T) {
	tests := []struct {
		name     string
		file     FileInfo
		expected string
	}{
		{
			name:     "pdf file",
			file:     FileInfo{Ext: "pdf"},
			expected: "pdf",
		},
		{
			name:     "no extension",
			file:     FileInfo{Ext: ""},
			expected: "misc",
		},
		{
			name:     "jpeg file",
			file:     FileInfo{Ext: "jpeg"},
			expected: "jpeg",
		},
	}
	strategy := &ExtensionStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strategy.Categorize(tt.file)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSizeStrategy_Categorize(t *testing.T) {
	tests := []struct {
		name     string
		file     FileInfo
		expected string
	}{
		// Test: tiny file → small
		{
			name:     "tiny file",
			file:     FileInfo{Size: KB},
			expected: "small",
		},
		// Test: exactly 1MB - 1 byte → small
		{
			name:     "1MB file",
			file:     FileInfo{Size: MB - KB},
			expected: "small",
		},
		// Test: exactly 1MB → medium (boundary!)
		{
			name:     "1MB file",
			file:     FileInfo{Size: MB},
			expected: "medium",
		},
		// Test: 50MB → medium
		{
			name:     "50MB file",
			file:     FileInfo{Size: 50 * MB},
			expected: "medium",
		},
		// Test: 100MB → medium (boundary!)
		{
			name:     "100MB file",
			file:     FileInfo{Size: 100 * MB},
			expected: "medium",
		},
		// Test: 101MB → large
		{
			name:     "101MB file",
			file:     FileInfo{Size: 101 * MB},
			expected: "large",
		},
	}

	strategy := &SizeStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strategy.Categorize(tt.file)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestModifiedDateStrategy_Categorize(t *testing.T) {
	tests := []struct {
		name     string
		file     FileInfo
		expected string
	}{
		{
			name:     "January 2023",
			file:     FileInfo{ModTime: time.Date(2023, time.January, 15, 10, 30, 0, 0, time.UTC)},
			expected: "2023-01",
		},
		{
			name:     "December 2025",
			file:     FileInfo{ModTime: time.Date(2025, time.December, 31, 23, 59, 0, 0, time.UTC)},
			expected: "2025-12",
		},
		{
			name:     "February 2024 leap year",
			file:     FileInfo{ModTime: time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)},
			expected: "2024-02",
		},
		{
			name:     "June 2020",
			file:     FileInfo{ModTime: time.Date(2020, time.June, 1, 12, 0, 0, 0, time.UTC)},
			expected: "2020-06",
		},
	}

	strategy := &ModifiedDateStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strategy.Categorize(tt.file)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
