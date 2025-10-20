// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var morphologyPronounClass _MorphologyPronounClass

func init() {
	morphologyPronounClass = _MorphologyPronounClass{objc.GetClass("NSMorphologyPronoun")}
}

type _MorphologyPronounClass struct {
	class objc.Class
}

type MorphologyPronoun struct {
	objc.ID
}

func MorphologyPronounFrom(ptr unsafe.Pointer) MorphologyPronoun {
	return MorphologyPronoun{
		ID: objc.ID(ptr),
	}
}




