// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCompositionTrackFormatDescriptionReplacement */


/* debug [class_header]: Header for AVCompositionTrackFormatDescriptionReplacement */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompositionTrackFormatDescriptionReplacement */
// An interface definition for the [CompositionTrackFormatDescriptionReplacement] class.
type ICompositionTrackFormatDescriptionReplacement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CompositionTrackFormatDescriptionReplacement */
	// properties:
	OriginalFormatDescription() FormatDescriptionRef /* not a class type */
	ReplacementFormatDescription() FormatDescriptionRef /* not a class type */
	FormatDescriptionReplacements() IAVCompositionTrackFormatDescriptionReplacement
	SetFormatDescriptionReplacements(value IAVCompositionTrackFormatDescriptionReplacement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompositionTrackFormatDescriptionReplacement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompositionTrackFormatDescriptionReplacement */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompositionTrackFormatDescriptionReplacement */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompositionTrackFormatDescriptionReplacement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompositionTrackFormatDescriptionReplacement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompositionTrackFormatDescriptionReplacement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompositionTrackFormatDescriptionReplacement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompositionTrackFormatDescriptionReplacement */

// The format description to replace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/originalFormatDescription
func (c_ CompositionTrackFormatDescriptionReplacement) OriginalFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("originalFormatDescription"))
	return rv
}/* debug [instance_properties/getter]: originalFormatDescription */


// The replacement format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/replacementFormatDescription
func (c_ CompositionTrackFormatDescriptionReplacement) ReplacementFormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("replacementFormatDescription"))
	return rv
}/* debug [instance_properties/getter]: replacementFormatDescription */


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrackFormatDescriptionReplacement) FormatDescriptionReplacements() IAVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[CompositionTrackFormatDescriptionReplacement](c_.ID, objc.Sel("formatDescriptionReplacements"))
	return rv
}/* debug [instance_properties/getter]: formatDescriptionReplacements */


// The replacement format descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontrack/formatdescriptionreplacements
func (c_ CompositionTrackFormatDescriptionReplacement) SetFormatDescriptionReplacements(value IAVCompositionTrackFormatDescriptionReplacement) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescriptionReplacements:"), value)
}/* debug [instance_properties/setter]: formatDescriptionReplacements */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCompositionTrackFormatDescriptionReplacement */



