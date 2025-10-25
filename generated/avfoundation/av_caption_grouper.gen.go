// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionGrouper */


/* debug [class_header]: Header for AVCaptionGrouper */
// The class instance for the [CaptionGrouper] class.
var (
	CaptionGrouperClass     _CaptionGrouperClass
	CaptionGrouperClassOnce sync.Once
)

func getCaptionGrouperClass() _CaptionGrouperClass {
	CaptionGrouperClassOnce.Do(func() {
		CaptionGrouperClass = _CaptionGrouperClass{objc.GetClass("AVCaptionGrouper")}
	})
	return CaptionGrouperClass
}

type _CaptionGrouperClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionGrouper */
// An interface definition for the [CaptionGrouper] class.
type ICaptionGrouper interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionGrouper */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionGrouper */
	// methods:
	AddCaption(input IAVCaption)
	FlushAddedCaptionsIntoGroupsUpToTime(upToTime objc.IObject /* cross-framework: Time */) []CaptionGroup
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionGrouper */
// Alloc allocates a new instance without initialization.
func (cc _CaptionGrouperClass) Alloc() CaptionGrouper {
	rv := objc.Send[CaptionGrouper](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionGrouperClass) New() CaptionGrouper {
	rv := objc.Send[CaptionGrouper](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionGrouper) Init() CaptionGrouper {
	rv := objc.Send[CaptionGrouper](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionGrouper) Autorelease() CaptionGrouper {
	rv := objc.Send[CaptionGrouper](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionGrouper creates a new CaptionGrouper instance.
func NewCaptionGrouper() CaptionGrouper {
	return getCaptionGrouperClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionGrouper */
// An object that analyzes the temporal overlaps of caption objects to create caption groups for each span of concurrent captions.


// An object that analyzes the temporal overlaps of caption objects to create caption groups for each span of concurrent captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper
type CaptionGrouper struct {
	objectivec.Object
}

// CaptionGrouperFrom constructs a [CaptionGrouper] from an unsafe.Pointer.
//
// An object that analyzes the temporal overlaps of caption objects to create caption groups for each span of concurrent captions.
func CaptionGrouperFrom(ptr unsafe.Pointer) CaptionGrouper {
	return CaptionGrouper{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionGrouper *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionGrouper */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionGrouper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionGrouper */

// Adds a caption to the pending group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper/add(_:)
func (c_ CaptionGrouper) AddCaption(input IAVCaption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addCaption:"), input)
}/* debug [instance_methods/method]: AddCaption */


// Creates caption groups for the captions you enqueue up to the time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper/flushAddedCaptions(upTo:)
func (c_ CaptionGrouper) FlushAddedCaptionsIntoGroupsUpToTime(upToTime objc.IObject /* cross-framework: Time */) []CaptionGroup {
	rv := objc.Send[[]CaptionGroup](c_.ID, objc.Sel("flushAddedCaptionsIntoGroupsUpToTime:"), upToTime)
	return rv
}/* debug [instance_methods/method]: FlushAddedCaptionsIntoGroupsUpToTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionGrouper */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionGrouper */



