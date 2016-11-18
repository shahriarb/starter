package python

import "github.com/shahriarb/starter/packs"

type Analysis struct {
	packs.AnalysisBase

	ServiceYAMLContext *ServiceYAMLContext
	DockerfileContext  *DockerfileContext
}
