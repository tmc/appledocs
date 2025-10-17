// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Orthography] class.
var OrthographyClass objc.Class

func init() {
	OrthographyClass = objc.GetClass("NSOrthography")
}

type Orthography struct {
	objc.ID
}

func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{
		ID: objc.ID(ptr),
	}
}



