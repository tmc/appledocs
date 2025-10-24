// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCShareableContentInfo */


/* debug [class_header]: Header for SCShareableContentInfo */
// The class instance for the [ShareableContentInfo] class.
var (
	ShareableContentInfoClass     _ShareableContentInfoClass
	ShareableContentInfoClassOnce sync.Once
)

func getShareableContentInfoClass() _ShareableContentInfoClass {
	ShareableContentInfoClassOnce.Do(func() {
		ShareableContentInfoClass = _ShareableContentInfoClass{objc.GetClass("SCShareableContentInfo")}
	})
	return ShareableContentInfoClass
}

type _ShareableContentInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ShareableContentInfo */
// An interface definition for the [ShareableContentInfo] class.
type IShareableContentInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ShareableContentInfo */
	// properties:
	ContentRect() corefoundation.CGRect
	PointPixelScale() float32
	Style() ShareableContentStyle
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ShareableContentInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ShareableContentInfo */
// Alloc allocates a new instance without initialization.
func (sc _ShareableContentInfoClass) Alloc() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ShareableContentInfoClass) New() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ShareableContentInfo) Init() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ShareableContentInfo) Autorelease() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShareableContentInfo creates a new ShareableContentInfo instance.
func NewShareableContentInfo() ShareableContentInfo {
	return getShareableContentInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ShareableContentInfo */
// An instance that provides information for the content in a given stream.


// An instance that provides information for the content in a given stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo
type ShareableContentInfo struct {
	objectivec.Object
}

// ShareableContentInfoFrom constructs a [ShareableContentInfo] from an unsafe.Pointer.
//
// An instance that provides information for the content in a given stream.
func ShareableContentInfoFrom(ptr unsafe.Pointer) ShareableContentInfo {
	return ShareableContentInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ShareableContentInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ShareableContentInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ShareableContentInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ShareableContentInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ShareableContentInfo */

// The size and location of content for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/contentRect
func (s_ ShareableContentInfo) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("contentRect"))
	return rv
}/* debug [instance_properties/getter]: contentRect */


// The scaling from points to output pixel resolution for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/pointPixelScale
func (s_ ShareableContentInfo) PointPixelScale() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("pointPixelScale"))
	return rv
}/* debug [instance_properties/getter]: pointPixelScale */


// The current presentation style of the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/style
func (s_ ShareableContentInfo) Style() ShareableContentStyle {
	rv := objc.Send[ShareableContentStyle](s_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCShareableContentInfo */



