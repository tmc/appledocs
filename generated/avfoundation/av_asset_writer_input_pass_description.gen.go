// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetWriterInputPassDescription */


/* debug [class_header]: Header for AVAssetWriterInputPassDescription */
// The class instance for the [AssetWriterInputPassDescription] class.
var (
	AssetWriterInputPassDescriptionClass     _AssetWriterInputPassDescriptionClass
	AssetWriterInputPassDescriptionClassOnce sync.Once
)

func getAssetWriterInputPassDescriptionClass() _AssetWriterInputPassDescriptionClass {
	AssetWriterInputPassDescriptionClassOnce.Do(func() {
		AssetWriterInputPassDescriptionClass = _AssetWriterInputPassDescriptionClass{objc.GetClass("AVAssetWriterInputPassDescription")}
	})
	return AssetWriterInputPassDescriptionClass
}

type _AssetWriterInputPassDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetWriterInputPassDescription */
// An interface definition for the [AssetWriterInputPassDescription] class.
type IAssetWriterInputPassDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetWriterInputPassDescription */
	// properties:
	SourceTimeRanges() []foundation.Value
	CanPerformMultiplePasses() bool
	SetCanPerformMultiplePasses(value bool)
	CurrentPassDescription() IAVAssetWriterInputPassDescription
	SetCurrentPassDescription(value IAVAssetWriterInputPassDescription)
	PerformsMultiPassEncodingIfSupported() bool
	SetPerformsMultiPassEncodingIfSupported(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetWriterInputPassDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetWriterInputPassDescription */
// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputPassDescriptionClass) Alloc() AssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetWriterInputPassDescriptionClass) New() AssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputPassDescription) Init() AssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputPassDescription) Autorelease() AssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputPassDescription creates a new AssetWriterInputPassDescription instance.
func NewAssetWriterInputPassDescription() AssetWriterInputPassDescription {
	return getAssetWriterInputPassDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetWriterInputPassDescription */
// An object that defines the interface to query for the requirements of the current pass.


// An object that defines the interface to query for the requirements of the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPassDescription
type AssetWriterInputPassDescription struct {
	objectivec.Object
}

// AssetWriterInputPassDescriptionFrom constructs a [AssetWriterInputPassDescription] from an unsafe.Pointer.
//
// An object that defines the interface to query for the requirements of the current pass.
func AssetWriterInputPassDescriptionFrom(ptr unsafe.Pointer) AssetWriterInputPassDescription {
	return AssetWriterInputPassDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetWriterInputPassDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetWriterInputPassDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetWriterInputPassDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetWriterInputPassDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetWriterInputPassDescription */

// An array of time ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPassDescription/sourceTimeRanges
func (a_ AssetWriterInputPassDescription) SourceTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](a_.ID, objc.Sel("sourceTimeRanges"))
	return rv
}/* debug [instance_properties/getter]: sourceTimeRanges */


// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInputPassDescription) CanPerformMultiplePasses() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePasses"))
	return rv
}/* debug [instance_properties/getter]: canPerformMultiplePasses */


// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInputPassDescription) SetCanPerformMultiplePasses(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePasses:"), value)
}/* debug [instance_properties/setter]: canPerformMultiplePasses */


// An object that describes the requirements for the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInputPassDescription) CurrentPassDescription() IAVAssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](a_.ID, objc.Sel("currentPassDescription"))
	return rv
}/* debug [instance_properties/getter]: currentPassDescription */


// An object that describes the requirements for the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInputPassDescription) SetCurrentPassDescription(value IAVAssetWriterInputPassDescription) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPassDescription:"), value)
}/* debug [instance_properties/setter]: currentPassDescription */


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInputPassDescription) PerformsMultiPassEncodingIfSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}/* debug [instance_properties/getter]: performsMultiPassEncodingIfSupported */


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInputPassDescription) SetPerformsMultiPassEncodingIfSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}/* debug [instance_properties/setter]: performsMultiPassEncodingIfSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetWriterInputPassDescription */



