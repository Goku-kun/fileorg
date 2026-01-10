package organizer

import (
	"fmt"
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

// Organizer orchestrates the scanning, grouping by strategy and moving the files
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
			o.result.FilesSkipped++
			continue
		}

		hiddenFile := strings.HasPrefix(strings.ToLower(info.Name()), ".")
		if hiddenFile {
			o.result.FilesSkipped++
			continue
		}

		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(info.Name()), "."))
		fullPath := filepath.Join(dir, info.Name())
		fileInfo := FileInfo{
			Name: info.Name(),
			Size: info.Size(),
			Path: fullPath,
			Ext:  ext,
			ModTime: info.ModTime(),
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

func (o *Organizer) moveFiles(folders map[string][]FileInfo) {

	for folder, files := range folders {
		folderPath := filepath.Join(o.cfg.SourceDir, folder)
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			if o.cfg.DryRun {
				fmt.Println("[Dry Run] Creating folder:", folderPath)
			} else {
				fmt.Println("Creating folder:", folderPath)

				err := os.MkdirAll(folderPath, 0755)
				if err != nil {
					o.result.Errors = append(o.result.Errors, err)
					continue
				}
				o.result.FoldersCreated++
			}
		}

		for _, file := range files {
			newPath := safePath(folderPath, file.Name)
			if o.cfg.DryRun {
				fmt.Printf("[Dry Run] Moving file: %s -> %s\n", file.Path, newPath)
			} else {
				fmt.Printf("Moving file: %s -> %s\n", file.Path, newPath)
				err := os.Rename(file.Path, newPath)
				if err != nil {
					o.result.Errors = append(o.result.Errors, err)
					continue
				}
				o.result.FilesMoved++
			}

		}
	}
}

func (o *Organizer) Organize() Result {
	file, err := o.scanFiles()
	if err != nil {
		o.result.Errors = append(o.result.Errors, err)
		return o.result
	}

	folders := o.groupByStrategy(file)
	o.moveFiles(folders)

	return o.result
}

func safePath(dir, filename string) string {
	originalPath := filepath.Join(dir, filename)
	if _, err := os.Stat(originalPath); os.IsNotExist(err) {
		return originalPath
	}

	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	for i := 1; ; i++ {
		newFilename := fmt.Sprintf("%s_%d%s", name, i, ext)
		newPath := filepath.Join(dir, newFilename)
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
	}
}

func (r Result) PrintSummary() {
	fmt.Println("Organization Summary:")
	fmt.Printf("Folders Created: %d\n", r.FoldersCreated)
	fmt.Printf("Files Moved: %d\n", r.FilesMoved)
	fmt.Printf("Files Skipped: %d\n", r.FilesSkipped)

	if len(r.Errors) > 0 {
		fmt.Println("Errors encountered during organization:")
		for _, err := range r.Errors {
			fmt.Printf("- %v\n", err)
		}
	} else {
		fmt.Println("No errors encountered.")
	}

}
