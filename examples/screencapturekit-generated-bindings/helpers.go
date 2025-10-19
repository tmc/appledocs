package main

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// nsArrayCount returns the count of an NSArray (works for any objc.ID that responds to count)
func nsArrayCount(arr unsafe.Pointer) int {
	if arr == nil {
		return 0
	}
	return int(objc.ID(arr).Send(objc.RegisterName("count")))
}

// nsArrayObjectAt returns the object at index i in an NSArray
func nsArrayObjectAt(arr unsafe.Pointer, i int) objc.ID {
	// Use generated foundation.Array method
	array := foundation.ArrayFrom(arr)
	return objc.ID(array.ObjectAtIndex(uint(i)))
}

// nsStringToGo converts an NSString (as objc.ID) to a Go string
func nsStringToGo(str objc.ID) string {
	if str == 0 {
		return ""
	}
	return objc.Send[string](str, objc.RegisterName("UTF8String"))
}

// nsStringPtrToGo converts an NSString pointer (unsafe.Pointer) to a Go string
func nsStringPtrToGo(str unsafe.Pointer) string {
	return nsStringToGo(objc.ID(str))
}
