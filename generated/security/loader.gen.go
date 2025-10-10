// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import "github.com/ebitengine/purego"

// lib holds the framework library handle
var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/Security.framework/Security", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
