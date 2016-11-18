package packs

import "github.com/shahriarb/starter/common"

type AnalysisBase struct {
	PackName string
	LanguageVersion string
	GitURL    string
	GitBranch string
	Framework string
	FrameworkVersion string
	SupportedLanguageVersions []string
	Messages common.Lister
	Databases []string
	ListOfStartCommands []string
}
