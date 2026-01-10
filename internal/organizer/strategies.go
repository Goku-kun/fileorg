package organizer

const (
	KB = 1024
	MB = 1024 * KB

	smallLimit = 1 * MB
	largeLimit = 100 * MB
)

type Strategy interface {
	Categorize(file FileInfo) string
}

type ExtensionStrategy struct {
	// fields
}

func (s *ExtensionStrategy) Categorize(file FileInfo) string {
	if file.Ext == "" {
		return "misc"
	}
	return file.Ext
}

type ModifiedDateStrategy struct {
	// fields
}

func (s *ModifiedDateStrategy) Categorize(file FileInfo) string {
	modifiedTime := file.ModTime.Format("2006-01")
	return modifiedTime
}

type SizeStrategy struct {
	// fields
}

func (s *SizeStrategy) Categorize(file FileInfo) string {
	size := file.Size

	if size < smallLimit {
		return "small"
	}

	if size <= largeLimit {
		return "medium"
	}
	return "large"
}
