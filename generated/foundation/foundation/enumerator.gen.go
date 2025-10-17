// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Enumerator] class.
var EnumeratorClass objc.Class

func init() {
	EnumeratorClass = objc.GetClass("NSEnumerator")
}

type Enumerator struct {
	objc.ID
}

func EnumeratorFrom(ptr unsafe.Pointer) Enumerator {
	return Enumerator{
		ID: objc.ID(ptr),
	}
}




