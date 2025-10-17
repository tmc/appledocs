// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MorphologyCustomPronoun] class.
var MorphologyCustomPronounClass objc.Class

func init() {
	MorphologyCustomPronounClass = objc.GetClass("NSMorphologyCustomPronoun")
}

type MorphologyCustomPronoun struct {
	objc.ID
}

func MorphologyCustomPronounFrom(ptr unsafe.Pointer) MorphologyCustomPronoun {
	return MorphologyCustomPronoun{
		ID: objc.ID(ptr),
	}
}




