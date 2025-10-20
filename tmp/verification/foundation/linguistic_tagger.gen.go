// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var linguisticTaggerClass _LinguisticTaggerClass

func init() {
	linguisticTaggerClass = _LinguisticTaggerClass{objc.GetClass("NSLinguisticTagger")}
}

type _LinguisticTaggerClass struct {
	class objc.Class
}

type LinguisticTagger struct {
	objc.ID
}

func LinguisticTaggerFrom(ptr unsafe.Pointer) LinguisticTagger {
	return LinguisticTagger{
		ID: objc.ID(ptr),
	}
}


// Enumerates over a given range of the string and calls the specified block for each tag. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/enumerateTags(in:scheme:options:using:)
func (l_ LinguisticTagger) EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ unsafe.Pointer, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("enumerateTagsInRange:scheme:options:usingBlock:"), range_, tagScheme, opts, block)
}


