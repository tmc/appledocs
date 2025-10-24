// Code generated from Apple documentation for CoreTransferable. DO NOT EDIT.

// Package coretransferable provides Go bindings for the CoreTransferable framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreTransferable without requiring cgo.
package coretransferable

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreTransferable.framework/CoreTransferable"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
