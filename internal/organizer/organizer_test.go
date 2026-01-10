package organizer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSafePath(t *testing.T) {
	tests := []struct {
		name          string
		existingFiles []string
		inputName     string
		expectedName  string
	}{
		{
			name:          "no collision",
			existingFiles: []string{},
			inputName:     "test.txt",
			expectedName:  "test.txt",
		},
		{
			name:          "one collision",
			existingFiles: []string{"test.txt"},
			inputName:     "test.txt",
			expectedName:  "test_1.txt",
		},
		{
			name:          "two collisions",
			existingFiles: []string{"test.txt", "test_1.txt"},
			inputName:     "test.txt",
			expectedName:  "test_2.txt",
		},
		{
			name:          "many collisions",
			existingFiles: []string{"doc.pdf", "doc_1.pdf", "doc_2.pdf", "doc_3.pdf"},
			inputName:     "doc.pdf",
			expectedName:  "doc_4.pdf",
		},
		{
			name:          "no extension with collision",
			existingFiles: []string{"README"},
			inputName:     "README",
			expectedName:  "README_1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp directory (auto-cleaned after test)
			tempDir := t.TempDir()

			// Create existing files
			for _, filename := range tt.existingFiles {
				path := filepath.Join(tempDir, filename)
				err := os.WriteFile(path, []byte("test"), 0644)
				if err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
			}

			// Call safePath
			got := safePath(tempDir, tt.inputName)
			expected := filepath.Join(tempDir, tt.expectedName)

			if got != expected {
				t.Errorf("got %q, want %q", got, expected)
			}
		})
	}
}

