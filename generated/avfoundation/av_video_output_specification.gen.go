// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [VideoOutputSpecification] class.
type IVideoOutputSpecification interface {
	objectivec.IObject
	

	// properties:
	DefaultOutputSettings() foundation.IDictionary
	SetDefaultOutputSettings(value foundation.IDictionary)
	DefaultPixelBufferAttributes() foundation.IDictionary
	SetDefaultPixelBufferAttributes(value foundation.IDictionary)
	PreferredTagCollections() objc.IObject /* cross-framework: NSArray */


	

	// methods:
	SetOutputSettingsForTagCollection(outputSettings foundation.IDictionary, tagCollection TagCollectionRef /* not a class type */)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/initWithTagCollections:
func NewVideoOutputSpecificationWithTagCollections(tagCollections objc.IObject /* cross-framework: NSArray */) VideoOutputSpecification {
	instance := getVideoOutputSpecificationClass().Alloc()
	rv := objc.Send[VideoOutputSpecification](instance.ID, objc.Sel("initWithTagCollections:"), tagCollections)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/setOutputSettings:forTagCollection:
func (v_ VideoOutputSpecification) SetOutputSettingsForTagCollection(outputSettings foundation.IDictionary, tagCollection TagCollectionRef /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOutputSettings:forTagCollection:"), outputSettings, tagCollection)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultOutputSettings
func (v_ VideoOutputSpecification) DefaultOutputSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](v_.ID, objc.Sel("defaultOutputSettings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultOutputSettings
func (v_ VideoOutputSpecification) SetDefaultOutputSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDefaultOutputSettings:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultPixelBufferAttributes
func (v_ VideoOutputSpecification) DefaultPixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](v_.ID, objc.Sel("defaultPixelBufferAttributes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultPixelBufferAttributes
func (v_ VideoOutputSpecification) SetDefaultPixelBufferAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDefaultPixelBufferAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/preferredTagCollections-2ikbd
func (v_ VideoOutputSpecification) PreferredTagCollections() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](v_.ID, objc.Sel("preferredTagCollections"))
	return rv
}







