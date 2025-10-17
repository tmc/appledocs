// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LinguisticTagger] class.
var LinguisticTaggerClass objc.Class

func init() {
	LinguisticTaggerClass = objc.GetClass("NSLinguisticTagger")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagger/enumerateTags(in:scheme:options:using:)
func (l_ LinguisticTagger) EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ Range, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	sel := objc.RegisterName("enumerateTagsInRange:scheme:options:usingBlock:")
	l_.ID.Send(sel, range_, tagScheme, opts, block)
}

