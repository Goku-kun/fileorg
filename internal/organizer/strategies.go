package organizer

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
