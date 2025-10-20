// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MorphologyCustomPronounClass _MorphologyCustomPronounClass

func init() {
	MorphologyCustomPronounClass = _MorphologyCustomPronounClass{objc.GetClass("NSMorphologyCustomPronoun")}
}

type _MorphologyCustomPronounClass struct {
	class objc.Class
}

type MorphologyCustomPronoun struct {
	objc.ID
}

func MorphologyCustomPronounFrom(ptr unsafe.Pointer) MorphologyCustomPronoun {
	return MorphologyCustomPronoun{
		ID: objc.ID(ptr),
	}
}




