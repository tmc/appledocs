// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

package securityfoundation

import "github.com/ebitengine/purego"

// lib holds the framework library handle
var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/SecurityFoundation.framework/SecurityFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
