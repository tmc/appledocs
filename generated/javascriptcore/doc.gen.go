
// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

// Package javascriptcore provides Go bindings for the JavaScriptCore framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to JavaScriptCore without requiring cgo.
package javascriptcore

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/JavaScriptCore.framework/JavaScriptCore"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

