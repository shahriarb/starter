package packs

import "github.com/shahriarb/starter/common"

type Analyzer interface {
	FillServices(*[]*common.Service) error
	HasPackage(pack string) bool
	GuessFramework() string
}
