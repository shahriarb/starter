package packs

import "github.com/shahriarb/starter/common"

type DockerfileContextBase struct {
	Version  string
	Framework string
	Packages *common.Lister
}
