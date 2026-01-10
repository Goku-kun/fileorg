package organizer

import (
	"os"
	"path/filepath"
	"strings"
)

// Config holds the settings for organizing files.
// SourceDir must be a valid, readable directory path.
type Config struct {
	SourceDir string
	Strategy  Strategy
	DryRun    bool
	Verbose   bool
}

// Results hold the display summary information, along with any errors collected during processing the files
type Result struct {
	FoldersCreated int
	FilesMoved     int
	FilesSkipped   int
	Errors         []error
}

type Organizer struct {
	result Result
	cfg    Config
}

func NewOrganizer(cfg Config) *Organizer {
	return &Organizer{
		cfg:    cfg,
		result: Result{},
	}
}

func (o *Organizer) scanFiles() ([]FileInfo, error) {
	dir := o.cfg.SourceDir
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := make([]FileInfo, 0)

	for _, entry := range entries {

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if info.IsDir() {
			continue
		}

		hiddenFile := strings.HasPrefix(strings.ToLower(info.Name()), ".")
		if hiddenFile {
			continue
		}

		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(info.Name()), "."))
		fullPath := filepath.Join(dir, info.Name())
		fileInfo := FileInfo{
			Name: info.Name(),
			Size: info.Size(),
			Path: fullPath,
			Ext:  ext,
		}

		files = append(files, fileInfo)
	}

	return files, err
}

func (o *Organizer) groupByStrategy(files []FileInfo) map[string][]FileInfo {

	folders := make(map[string][]FileInfo)

	for _, file := range files {
		folder := o.cfg.Strategy.Categorize(file)
		folders[folder] = append(folders[folder], file)
	}

	return folders
}
