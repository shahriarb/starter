package packs

import "github.com/shahriarb/starter/common"

type DockerComposeYAMLContextBase struct {
	Services []*common.Service
	Dbs      []common.Database
}
