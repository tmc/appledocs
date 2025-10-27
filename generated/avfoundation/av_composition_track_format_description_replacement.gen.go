// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CompositionTrackFormatDescriptionReplacement] class.
var (
	CompositionTrackFormatDescriptionReplacementClass     _CompositionTrackFormatDescriptionReplacementClass
	CompositionTrackFormatDescriptionReplacementClassOnce sync.Once
)

func getCompositionTrackFormatDescriptionReplacementClass() _CompositionTrackFormatDescriptionReplacementClass {
	CompositionTrackFormatDescriptionReplacementClassOnce.Do(func() {
		CompositionTrackFormatDescriptionReplacementClass = _CompositionTrackFormatDescriptionReplacementClass{objc.GetClass("AVCompositionTrackFormatDescriptionReplacement")}
	})
	return CompositionTrackFormatDescriptionReplacementClass
}

type _CompositionTrackFormatDescriptionReplacementClass struct {
	class objc.Class
}





// An interface definition for the [CompositionTrackFormatDescriptionReplacement] class.
type ICompositionTrackFormatDescriptionReplacement interface {
	objectivec.IObject
	

	// properties:
	OriginalFormatDescription() FormatDescriptionRef /* not a class type */
	ReplacementFormatDescription() FormatDescriptionRef /* not a class type */
	FormatDescriptionReplacements() IAVCompositionTrackFormatDescriptionReplacement
	SetFormatDescriptionReplacements(value IAVCompositionTrackFormatDescriptionReplacement)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackFormatDescriptionReplacementClass) Alloc() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CompositionTrackFormatDescriptionReplacementClass) New() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositionTrackFormatDescriptionReplacement) Init() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositionTrackFormatDescriptionReplacement) Autorelease() CompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositionTrackFormatDescriptionReplacement creates a new CompositionTrackFormatDescriptionReplacement instance.
func NewCompositionTrackFormatDescriptionReplacement() CompositionTrackFormatDescriptionReplacement {
	return getCompositionTrackFormatDescriptionReplacementClass().New()
}





// An object that represents a format description and its replacement.


// An object that represents a format description and its replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement
type CompositionTrackFormatDescriptionReplacement struct {
	objectivec.Object
}

// CompositionTrackFormatDescriptionReplacementFrom constructs a [CompositionTrackFormatDescriptionReplacement] from an unsafe.Pointer.
//
// An object that represents a format description and its replacement.
func CompositionTrackFormatDescriptionReplacementFrom(ptr unsafe.Pointer) CompositionTrackFormatDescriptionReplacement {
	return CompositionTrackFormatDescriptionReplacement{objectivec.Object{objc.ID(ptr)}}
}

























// The format description to replace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/originalFormatDescription
func (c_ CompositionTrackFormatDescriptionReplacement) OriginalFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("originalFormatDescription"))
	return rv
}


// The replacement format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/replacementFormatDescription
func (c_ CompositionTrackFormatDescriptionReplacement) ReplacementFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("replacementFormatDescription"))
	return rv
}


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrackFormatDescriptionReplacement) FormatDescriptionReplacements() IAVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrackFormatDescriptionReplacement) SetFormatDescriptionReplacements(value IAVCompositionTrackFormatDescriptionReplacement) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptionReplacements:"), value)
}








