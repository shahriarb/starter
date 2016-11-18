package node

import "github.com/shahriarb/starter/packs"

type Analysis struct {
	packs.AnalysisBase
	DockerComposeYAMLContext *DockerComposeYAMLContext
	ServiceYAMLContext *ServiceYAMLContext
	DockerfileContext  *DockerfileContext
}
