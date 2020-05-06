package constant

import "path"

const (
	DataDir = "data"
	TmpDir  = "tmp"
)

var (
	ProjectDir = path.Join(DataDir, "project")
)
