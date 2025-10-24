// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class METrackInfo */


/* debug [class_header]: Header for METrackInfo */
// The class instance for the [METrackInfo] class.
var (
	METrackInfoClass     _METrackInfoClass
	METrackInfoClassOnce sync.Once
)

func getMETrackInfoClass() _METrackInfoClass {
	METrackInfoClassOnce.Do(func() {
		METrackInfoClass = _METrackInfoClass{objc.GetClass("METrackInfo")}
	})
	return METrackInfoClass
}

type _METrackInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for METrackInfo */
// An interface definition for the [METrackInfo] class.
type IMETrackInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for METrackInfo */
	// properties:
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	FormatDescriptions() objc.IObject /* cross-framework: NSArray */
	Enabled() bool
	SetEnabled(value bool)
	MediaType() MediaType /* not a class type */
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	NaturalTimescale() TimeScale /* not a class type */
	SetNaturalTimescale(value TimeScale /* not a class type */)
	NominalFrameRate() unsafe.Pointer
	SetNominalFrameRate(value unsafe.Pointer)
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)
	TrackEdits() []foundation.Value
	SetTrackEdits(value []foundation.Value)
	TrackID() PersistentTrackID /* not a class type */
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for METrackInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for METrackInfo */
// Alloc allocates a new instance without initialization.
func (mc _METrackInfoClass) Alloc() METrackInfo {
	rv := objc.Send[METrackInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _METrackInfoClass) New() METrackInfo {
	rv := objc.Send[METrackInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ METrackInfo) Init() METrackInfo {
	rv := objc.Send[METrackInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ METrackInfo) Autorelease() METrackInfo {
	rv := objc.Send[METrackInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMETrackInfo creates a new METrackInfo instance.
func NewMETrackInfo() METrackInfo {
	return getMETrackInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for METrackInfo */
// An object that includes track properties parsed from the media asset.


// An object that includes track properties parsed from the media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo
type METrackInfo struct {
	objectivec.Object
}

// METrackInfoFrom constructs a [METrackInfo] from an unsafe.Pointer.
//
// An object that includes track properties parsed from the media asset.
func METrackInfoFrom(ptr unsafe.Pointer) METrackInfo {
	return METrackInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for METrackInfo */

// Creates a new track info object with the media type, track ID, and format descriptions that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/initWithMediaType:trackID:formatDescriptions:
func NewMETrackInfoWithMediaTypeTrackIDFormatDescriptions(mediaType MediaType /* not a class type */, trackID PersistentTrackID /* not a class type */, formatDescriptions objc.IObject /* cross-framework: NSArray */) METrackInfo {
	instance := getMETrackInfoClass().Alloc()
	rv := objc.Send[METrackInfo](instance.ID, objc.Sel("initWithMediaType:trackID:formatDescriptions:"), mediaType, trackID, formatDescriptions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMETrackInfoWithMediaTypeTrackIDFormatDescriptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for METrackInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for METrackInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for METrackInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for METrackInfo */

// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/extendedLanguageTag
func (m_ METrackInfo) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}/* debug [instance_properties/getter]: extendedLanguageTag */


// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/extendedLanguageTag
func (m_ METrackInfo) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}/* debug [instance_properties/setter]: extendedLanguageTag */


// An array of format descriptions for the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/formatDescriptions
func (m_ METrackInfo) FormatDescriptions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("formatDescriptions"))
	return rv
}/* debug [instance_properties/getter]: formatDescriptions */


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/isEnabled
func (m_ METrackInfo) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/isEnabled
func (m_ METrackInfo) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The media type of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/mediaType
func (m_ METrackInfo) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// Indicates the natural dimensions of the media data referenced by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalSize
func (m_ METrackInfo) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}/* debug [instance_properties/getter]: naturalSize */


// Indicates the natural dimensions of the media data referenced by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalSize
func (m_ METrackInfo) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}/* debug [instance_properties/setter]: naturalSize */


// The natural timescale of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) NaturalTimescale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](m_.ID, objc.Sel("naturalTimescale"))
	return rv
}/* debug [instance_properties/getter]: naturalTimescale */


// The natural timescale of the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m_ METrackInfo) SetNaturalTimescale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalTimescale:"), value)
}/* debug [instance_properties/setter]: naturalTimescale */


// The frame rate of the track in frames per second, as a 32-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) NominalFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalFrameRate"))
	return rv
}/* debug [instance_properties/getter]: nominalFrameRate */


// The frame rate of the track in frames per second, as a 32-bit floating point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m_ METrackInfo) SetNominalFrameRate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalFrameRate:"), value)
}/* debug [instance_properties/setter]: nominalFrameRate */


// Indicates the preferred affine display transform of the track media for visual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/preferredTransform
func (m_ METrackInfo) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// Indicates the preferred affine display transform of the track media for visual display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/preferredTransform
func (m_ METrackInfo) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredTransform:"), value)
}/* debug [instance_properties/setter]: preferredTransform */


// A Boolean value that indicates whether frame reordering occurs in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/requiresFrameReordering
func (m_ METrackInfo) RequiresFrameReordering() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresFrameReordering"))
	return rv
}/* debug [instance_properties/getter]: requiresFrameReordering */


// A Boolean value that indicates whether frame reordering occurs in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/requiresFrameReordering
func (m_ METrackInfo) SetRequiresFrameReordering(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresFrameReordering:"), value)
}/* debug [instance_properties/setter]: requiresFrameReordering */


// An array of edit segments for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) TrackEdits() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("trackEdits"))
	return rv
}/* debug [instance_properties/getter]: trackEdits */


// An array of edit segments for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackEdits
func (m_ METrackInfo) SetTrackEdits(value []foundation.Value) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackEdits:"), nsArray)
}/* debug [instance_properties/setter]: trackEdits */


// An integer that identifies the track within the media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackID
func (m_ METrackInfo) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the track is enabled by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/metrackinfo/isenabled
func (m_ METrackInfo) SetIsEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class METrackInfo */


