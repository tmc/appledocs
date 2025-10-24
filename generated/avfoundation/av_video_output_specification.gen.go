// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVVideoOutputSpecification */


/* debug [class_header]: Header for AVVideoOutputSpecification */
// The class instance for the [VideoOutputSpecification] class.
var (
	VideoOutputSpecificationClass     _VideoOutputSpecificationClass
	VideoOutputSpecificationClassOnce sync.Once
)

func getVideoOutputSpecificationClass() _VideoOutputSpecificationClass {
	VideoOutputSpecificationClassOnce.Do(func() {
		VideoOutputSpecificationClass = _VideoOutputSpecificationClass{objc.GetClass("AVVideoOutputSpecification")}
	})
	return VideoOutputSpecificationClass
}

type _VideoOutputSpecificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoOutputSpecification */
// An interface definition for the [VideoOutputSpecification] class.
type IVideoOutputSpecification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoOutputSpecification */
	// properties:
	DefaultOutputSettings() foundation.IDictionary
	SetDefaultOutputSettings(value foundation.IDictionary)
	DefaultPixelBufferAttributes() foundation.IDictionary
	SetDefaultPixelBufferAttributes(value foundation.IDictionary)
	PreferredTagCollections() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoOutputSpecification */
	// methods:
	SetOutputSettingsForTagCollection(outputSettings foundation.IDictionary, tagCollection TagCollectionRef /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoOutputSpecification */
// Alloc allocates a new instance without initialization.
func (vc _VideoOutputSpecificationClass) Alloc() VideoOutputSpecification {
	rv := objc.Send[VideoOutputSpecification](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoOutputSpecificationClass) New() VideoOutputSpecification {
	rv := objc.Send[VideoOutputSpecification](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoOutputSpecification) Init() VideoOutputSpecification {
	rv := objc.Send[VideoOutputSpecification](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoOutputSpecification) Autorelease() VideoOutputSpecification {
	rv := objc.Send[VideoOutputSpecification](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoOutputSpecification creates a new VideoOutputSpecification instance.
func NewVideoOutputSpecification() VideoOutputSpecification {
	return getVideoOutputSpecificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoOutputSpecification */
// An object that specifies the pixel buffer attributes and tag collections handled by a player video output.


// An object that specifies the pixel buffer attributes and tag collections handled by a player video output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification
type VideoOutputSpecification struct {
	objectivec.Object
}

// VideoOutputSpecificationFrom constructs a [VideoOutputSpecification] from an unsafe.Pointer.
//
// An object that specifies the pixel buffer attributes and tag collections handled by a player video output.
func VideoOutputSpecificationFrom(ptr unsafe.Pointer) VideoOutputSpecification {
	return VideoOutputSpecification{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoOutputSpecification */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/initWithTagCollections:
func NewVideoOutputSpecificationWithTagCollections(tagCollections objc.IObject /* cross-framework: NSArray */) VideoOutputSpecification {
	instance := getVideoOutputSpecificationClass().Alloc()
	rv := objc.Send[VideoOutputSpecification](instance.ID, objc.Sel("initWithTagCollections:"), tagCollections)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVideoOutputSpecificationWithTagCollections */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoOutputSpecification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoOutputSpecification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoOutputSpecification */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/setOutputSettings:forTagCollection:
func (v_ VideoOutputSpecification) SetOutputSettingsForTagCollection(outputSettings foundation.IDictionary, tagCollection TagCollectionRef /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOutputSettings:forTagCollection:"), outputSettings, tagCollection)
}/* debug [instance_methods/method]: SetOutputSettingsForTagCollection */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoOutputSpecification */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultOutputSettings
func (v_ VideoOutputSpecification) DefaultOutputSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](v_.ID, objc.Sel("defaultOutputSettings"))
	return rv
}/* debug [instance_properties/getter]: defaultOutputSettings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultOutputSettings
func (v_ VideoOutputSpecification) SetDefaultOutputSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDefaultOutputSettings:"), value)
}/* debug [instance_properties/setter]: defaultOutputSettings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultPixelBufferAttributes
func (v_ VideoOutputSpecification) DefaultPixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](v_.ID, objc.Sel("defaultPixelBufferAttributes"))
	return rv
}/* debug [instance_properties/getter]: defaultPixelBufferAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultPixelBufferAttributes
func (v_ VideoOutputSpecification) SetDefaultPixelBufferAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDefaultPixelBufferAttributes:"), value)
}/* debug [instance_properties/setter]: defaultPixelBufferAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/preferredTagCollections-2ikbd
func (v_ VideoOutputSpecification) PreferredTagCollections() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](v_.ID, objc.Sel("preferredTagCollections"))
	return rv
}/* debug [instance_properties/getter]: preferredTagCollections */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVVideoOutputSpecification */


