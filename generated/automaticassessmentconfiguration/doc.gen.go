
// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

// Package automaticassessmentconfiguration provides Go bindings for the AutomaticAssessmentConfiguration framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AutomaticAssessmentConfiguration without requiring cgo.
package automaticassessmentconfiguration

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AutomaticAssessmentConfiguration.framework/AutomaticAssessmentConfiguration"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

