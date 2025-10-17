// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LinguisticTagger] class.
var linguisticTaggerClass = _LinguisticTaggerClass{objc.GetClass("NSLinguisticTagger")}

type _LinguisticTaggerClass struct {
	class objc.Class
}

// An interface definition for the [LinguisticTagger] class.
type ILinguisticTagger interface {
	objectivec.IObject
	EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ unsafe.Pointer, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
}

// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger

type LinguisticTagger struct {
	objectivec.Object
}

// LinguisticTaggerFrom constructs a [LinguisticTagger] from an unsafe.Pointer.
//
// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script.
func LinguisticTaggerFrom(ptr unsafe.Pointer) LinguisticTagger {
	return LinguisticTagger{objectivec.Object{objc.ID(ptr)}}
}

// Enumerates over a given range of the string and calls the specified block for each tag. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/enumerateTags(in:scheme:options:using:)
func (l_ LinguisticTagger) EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ unsafe.Pointer, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("enumerateTagsInRange:scheme:options:usingBlock:"), range_, tagScheme, opts, block)
}


