// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MorphologyPronoun] class.
var MorphologyPronounClass objc.Class

func init() {
	MorphologyPronounClass = objc.GetClass("NSMorphologyPronoun")
}

type MorphologyPronoun struct {
	objc.ID
}

func MorphologyPronounFrom(ptr unsafe.Pointer) MorphologyPronoun {
	return MorphologyPronoun{
		ID: objc.ID(ptr),
	}
}




