// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNRenderingSessionFrameAttributes] class.
var (
	CNRenderingSessionFrameAttributesClass     _CNRenderingSessionFrameAttributesClass
	CNRenderingSessionFrameAttributesClassOnce sync.Once
)

func getCNRenderingSessionFrameAttributesClass() _CNRenderingSessionFrameAttributesClass {
	CNRenderingSessionFrameAttributesClassOnce.Do(func() {
		CNRenderingSessionFrameAttributesClass = _CNRenderingSessionFrameAttributesClass{objc.GetClass("CNRenderingSessionFrameAttributes")}
	})
	return CNRenderingSessionFrameAttributesClass
}

type _CNRenderingSessionFrameAttributesClass struct {
	class objc.Class
}

// An interface definition for the [CNRenderingSessionFrameAttributes] class.
type ICNRenderingSessionFrameAttributes interface {
	objectivec.IObject
	// properties:
	// methods:
}

// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
//
// The object exposes properties such as focus disparity and f-stop. It initializes these to the values that the original recorded movie used for that frame. However, you can change them before rendering to adjust focus and aperture.


// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes
type CNRenderingSessionFrameAttributes struct {
	objectivec.Object
}

// CNRenderingSessionFrameAttributesFrom constructs a [CNRenderingSessionFrameAttributes] from an unsafe.Pointer.
//
// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
func CNRenderingSessionFrameAttributesFrom(ptr unsafe.Pointer) CNRenderingSessionFrameAttributes {
	return CNRenderingSessionFrameAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionFrameAttributesClass) Alloc() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNRenderingSessionFrameAttributesClass) New() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNRenderingSessionFrameAttributes) Init() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNRenderingSessionFrameAttributes) Autorelease() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNRenderingSessionFrameAttributes creates a new CNRenderingSessionFrameAttributes instance.
func NewCNRenderingSessionFrameAttributes() CNRenderingSessionFrameAttributes {
	return getCNRenderingSessionFrameAttributesClass().New()
}



// Initializes the rendering frame attributes from a sample buffer read from a Cinematic metadata track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/initWithSampleBuffer:sessionAttributes:
func NewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes(sampleBuffer SampleBufferRef /* not a class type */, sessionAttributes ICNRenderingSessionAttributes) CNRenderingSessionFrameAttributes {
	instance := getCNRenderingSessionFrameAttributesClass().Alloc()
	rv := objc.Send[CNRenderingSessionFrameAttributes](instance.ID, objc.Sel("initWithSampleBuffer:sessionAttributes:"), sampleBuffer, sessionAttributes)
	rv.Autorelease()
	return rv
}


// Initializes the rendering frame attributes from a timed metadata group read from a Cinematic metadata track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/initWithTimedMetadataGroup:sessionAttributes:
func NewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes(metadataGroup objc.IObject /* cross-framework TimedMetadataGroup */, sessionAttributes ICNRenderingSessionAttributes) CNRenderingSessionFrameAttributes {
	instance := getCNRenderingSessionFrameAttributesClass().Alloc()
	rv := objc.Send[CNRenderingSessionFrameAttributes](instance.ID, objc.Sel("initWithTimedMetadataGroup:sessionAttributes:"), metadataGroup, sessionAttributes)
	rv.Autorelease()
	return rv
}



