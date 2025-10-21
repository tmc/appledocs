// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

// Package safariservices provides Go bindings for the SafariServices framework.
//
// Enable web views and services in your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SafariServices without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices
package safariservices

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SafariServices.framework/SafariServices"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


