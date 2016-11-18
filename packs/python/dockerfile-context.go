package python

import "github.com/shahriarb/starter/packs"

type DockerfileContext struct {
	packs.DockerfileContextBase
	RequirementsTxt string
}
