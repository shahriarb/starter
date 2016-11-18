package python

import (
	"path/filepath"

	"github.com/shahriarb/starter/common"
	"github.com/shahriarb/starter/packs"
)

type Detector struct {
	packs.PackElement
}

func (d *Detector) Detect(rootDir string) bool {
	return common.FileExists(filepath.Join(rootDir, "requirements.txt"))
}
