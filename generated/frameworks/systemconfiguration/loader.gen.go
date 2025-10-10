// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import "github.com/ebitengine/purego"

// lib holds the framework library handle
var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/SystemConfiguration.framework/SystemConfiguration", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
