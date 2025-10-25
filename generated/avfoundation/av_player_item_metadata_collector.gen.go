// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemMetadataCollector */


/* debug [class_header]: Header for AVPlayerItemMetadataCollector */
// The class instance for the [PlayerItemMetadataCollector] class.
var (
	PlayerItemMetadataCollectorClass     _PlayerItemMetadataCollectorClass
	PlayerItemMetadataCollectorClassOnce sync.Once
)

func getPlayerItemMetadataCollectorClass() _PlayerItemMetadataCollectorClass {
	PlayerItemMetadataCollectorClassOnce.Do(func() {
		PlayerItemMetadataCollectorClass = _PlayerItemMetadataCollectorClass{objc.GetClass("AVPlayerItemMetadataCollector")}
	})
	return PlayerItemMetadataCollectorClass
}

type _PlayerItemMetadataCollectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemMetadataCollector */
// An interface definition for the [PlayerItemMetadataCollector] class.
type IPlayerItemMetadataCollector interface {
	IPlayerItemMediaDataCollector
	
/* debug [class_interface_properties]: Properties for PlayerItemMetadataCollector */
	// properties:
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemMetadataCollector */
	// methods:
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemMetadataCollector */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemMetadataCollectorClass) Alloc() PlayerItemMetadataCollector {
	rv := objc.Send[PlayerItemMetadataCollector](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemMetadataCollectorClass) New() PlayerItemMetadataCollector {
	rv := objc.Send[PlayerItemMetadataCollector](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemMetadataCollector) Init() PlayerItemMetadataCollector {
	rv := objc.Send[PlayerItemMetadataCollector](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemMetadataCollector) Autorelease() PlayerItemMetadataCollector {
	rv := objc.Send[PlayerItemMetadataCollector](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemMetadataCollector creates a new PlayerItemMetadataCollector instance.
func NewPlayerItemMetadataCollector() PlayerItemMetadataCollector {
	return getPlayerItemMetadataCollectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemMetadataCollector */
// An object used to capture the date range metadata defined for an HTTP Live Streaming asset.
//
// You can use the HLS tag to define date range metadata in a media playlist. This tag is useful for defining timed metadata for interstitial regions such as advertisements, but can be used to define any timed metadata needed by your stream. To access this metadata when the stream is played using an , you create an instance of , configure its delegate object (see ), and add it as a media data collector to the (see example). Creating an as shown in the example, will capture all metadata defined in your stream. If you would like to filter the output to only the metadata of interest, you can create an instance to filter by identifier and/or classifying labels using the initializer.


// An object used to capture the date range metadata defined for an HTTP Live Streaming asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector
type PlayerItemMetadataCollector struct {
	PlayerItemMediaDataCollector
}

// PlayerItemMetadataCollectorFrom constructs a [PlayerItemMetadataCollector] from an unsafe.Pointer.
//
// An object used to capture the date range metadata defined for an HTTP Live Streaming asset.
func PlayerItemMetadataCollectorFrom(ptr unsafe.Pointer) PlayerItemMetadataCollector {
	return PlayerItemMetadataCollector{
		PlayerItemMediaDataCollector: PlayerItemMediaDataCollectorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemMetadataCollector */

// Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/init(identifiers:classifyingLabels:)
func NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) PlayerItemMetadataCollector {
	instance := getPlayerItemMetadataCollectorClass().Alloc()
	rv := objc.Send[PlayerItemMetadataCollector](instance.ID, objc.Sel("initWithIdentifiers:classifyingLabels:"), identifiers, classifyingLabels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemMetadataCollector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemMetadataCollector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemMetadataCollector */

// Sets the delegate and a dispatch queue on which the delegate will be called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/setDelegate(_:queue:)
func (p_ PlayerItemMetadataCollector) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemMetadataCollector */

// Accesses the metadata collector’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/delegate
func (p_ PlayerItemMetadataCollector) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue on which the delegate’s methods are called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/delegateQueue
func (p_ PlayerItemMetadataCollector) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemMetadataCollector */


