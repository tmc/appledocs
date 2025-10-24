// Code generated from Apple documentation for AppleNewsAPI. DO NOT EDIT.

// Package applenewsapi provides Go bindings for the AppleNewsAPI framework.
//
// Publish and manage Apple News Format articles.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppleNewsAPI without requiring cgo.
//
// See: https://developer.apple.com/documentation/AppleNewsAPI
package applenewsapi

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppleNewsAPI.framework/AppleNewsAPI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

