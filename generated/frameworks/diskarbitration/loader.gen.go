// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import "github.com/ebitengine/purego"

// lib holds the framework library handle
var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/DiskArbitration.framework/DiskArbitration", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
