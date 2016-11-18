package packs

import "github.com/shahriarb/starter/common"

type ServiceYAMLContextBase struct {
	Services []*common.Service
	Dbs      []common.Database
}
