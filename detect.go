package main

import (
	"fmt"
	"github.com/shahriarb/starter/common"
	"github.com/shahriarb/starter/packs"
	"github.com/shahriarb/starter/packs/node"
	"github.com/shahriarb/starter/packs/php"
	"github.com/shahriarb/starter/packs/ruby"
)

func Detect(rootDir string) (packs.Pack, error) {
	ruby := ruby.Pack{}
	node := node.Pack{}
	php := php.Pack{}
	detectors := []packs.Detector{ruby.Detector(), node.Detector(), php.Detector()}
	var packs []packs.Pack

	for _, d := range detectors {
		if d.Detect(rootDir) {
			packs = append(packs, d.GetPack())
			common.PrintlnL0("Found %s application", d.GetPack().Name())
		}
	}

	if len(packs) == 0 {
		return nil, fmt.Errorf("Could not detect any of the supported frameworks")
	} else if len(packs) > 1 {
		return nil, fmt.Errorf("More than one framework detected")
	} else {
		return packs[0], nil
	}
}
