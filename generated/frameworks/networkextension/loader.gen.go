// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import "github.com/ebitengine/purego"

// lib holds the framework library handle
var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/NetworkExtension.framework/NetworkExtension", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
