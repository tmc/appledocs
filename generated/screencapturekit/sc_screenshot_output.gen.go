// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCScreenshotOutput */


/* debug [class_header]: Header for SCScreenshotOutput */
// The class instance for the [ScreenshotOutput] class.
var (
	ScreenshotOutputClass     _ScreenshotOutputClass
	ScreenshotOutputClassOnce sync.Once
)

func getScreenshotOutputClass() _ScreenshotOutputClass {
	ScreenshotOutputClassOnce.Do(func() {
		ScreenshotOutputClass = _ScreenshotOutputClass{objc.GetClass("SCScreenshotOutput")}
	})
	return ScreenshotOutputClass
}

type _ScreenshotOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScreenshotOutput */
// An interface definition for the [ScreenshotOutput] class.
type IScreenshotOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScreenshotOutput */
	// properties:
	FileURL() objc.IObject /* cross-framework: NSURL */
	SetFileURL(value objc.IObject /* cross-framework: NSURL */)
	HdrImage() ImageRef /* not a class type */
	SetHdrImage(value ImageRef /* not a class type */)
	SdrImage() ImageRef /* not a class type */
	SetSdrImage(value ImageRef /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScreenshotOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScreenshotOutput */
// Alloc allocates a new instance without initialization.
func (sc _ScreenshotOutputClass) Alloc() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScreenshotOutputClass) New() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenshotOutput) Init() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenshotOutput) Autorelease() ScreenshotOutput {
	rv := objc.Send[ScreenshotOutput](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenshotOutput creates a new ScreenshotOutput instance.
func NewScreenshotOutput() ScreenshotOutput {
	return getScreenshotOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScreenshotOutput */
// An object that contains all images requested by the client.


// An object that contains all images requested by the client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type ScreenshotOutput struct {
	objectivec.Object
}

// ScreenshotOutputFrom constructs a [ScreenshotOutput] from an unsafe.Pointer.
//
// An object that contains all images requested by the client.
func ScreenshotOutputFrom(ptr unsafe.Pointer) ScreenshotOutput {
	return ScreenshotOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScreenshotOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScreenshotOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScreenshotOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScreenshotOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScreenshotOutput */

// A URL property that specifies the location of the saved image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/fileURL
func (s_ ScreenshotOutput) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// A URL property that specifies the location of the saved image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/fileURL
func (s_ ScreenshotOutput) SetFileURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFileURL:"), value)
}/* debug [instance_properties/setter]: fileURL */


// An output property that specifies the high dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) HdrImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("hdrImage"))
	return rv
}/* debug [instance_properties/getter]: hdrImage */


// An output property that specifies the high dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s_ ScreenshotOutput) SetHdrImage(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHdrImage:"), value)
}/* debug [instance_properties/setter]: hdrImage */


// An output property that specifies the standard dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/sdrImage
func (s_ ScreenshotOutput) SdrImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](s_.ID, objc.Sel("sdrImage"))
	return rv
}/* debug [instance_properties/getter]: sdrImage */


// An output property that specifies the standard dynamic range version of the screenshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/sdrImage
func (s_ ScreenshotOutput) SetSdrImage(value ImageRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSdrImage:"), value)
}/* debug [instance_properties/setter]: sdrImage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCScreenshotOutput */



