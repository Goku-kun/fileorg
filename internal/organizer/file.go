package organizer

import "time"

type FileInfo struct {
	Name    string
	Size    int64
	Path    string
	ModTime time.Time
	Ext     string
}
