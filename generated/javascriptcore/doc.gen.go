// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

// Package javascriptcore provides Go bindings for the JavaScriptCore framework.
//
// Evaluate JavaScript programs from within an app, and support JavaScript scripting of your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to JavaScriptCore without requiring cgo.

// Evaluate JavaScript programs from within an app, and support JavaScript scripting of your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore
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

