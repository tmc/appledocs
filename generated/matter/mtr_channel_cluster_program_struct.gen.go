// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterProgramStruct */


/* debug [class_header]: Header for MTRChannelClusterProgramStruct */
// The class instance for the [MTRChannelClusterProgramStruct] class.
var (
	MTRChannelClusterProgramStructClass     _MTRChannelClusterProgramStructClass
	MTRChannelClusterProgramStructClassOnce sync.Once
)

func getMTRChannelClusterProgramStructClass() _MTRChannelClusterProgramStructClass {
	MTRChannelClusterProgramStructClassOnce.Do(func() {
		MTRChannelClusterProgramStructClass = _MTRChannelClusterProgramStructClass{objc.GetClass("MTRChannelClusterProgramStruct")}
	})
	return MTRChannelClusterProgramStructClass
}

type _MTRChannelClusterProgramStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterProgramStruct */
// An interface definition for the [MTRChannelClusterProgramStruct] class.
type IMTRChannelClusterProgramStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterProgramStruct */
	// properties:
	AudioLanguages() objc.IObject /* cross-framework: NSArray */
	SetAudioLanguages(value objc.IObject /* cross-framework: NSArray */)
	Channel() objc.IObject /* cross-framework: MTRChannelClusterChannelInfoStruct */
	SetChannel(value objc.IObject /* cross-framework: MTRChannelClusterChannelInfoStruct */)
	DescriptionString() objc.IObject /* cross-framework: NSString */
	SetDescriptionString(value objc.IObject /* cross-framework: NSString */)
	EndTime() objc.IObject /* cross-framework: NSNumber */
	SetEndTime(value objc.IObject /* cross-framework: NSNumber */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	ParentalGuidanceText() objc.IObject /* cross-framework: NSString */
	SetParentalGuidanceText(value objc.IObject /* cross-framework: NSString */)
	RecordingFlag() objc.IObject /* cross-framework: NSNumber */
	SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */)
	ReleaseDate() objc.IObject /* cross-framework: NSString */
	SetReleaseDate(value objc.IObject /* cross-framework: NSString */)
	SeriesInfo() IMTRChannelClusterSeriesInfoStruct
	SetSeriesInfo(value IMTRChannelClusterSeriesInfoStruct)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterProgramStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterProgramStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramStructClass) Alloc() MTRChannelClusterProgramStruct {
	rv := objc.Send[MTRChannelClusterProgramStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterProgramStructClass) New() MTRChannelClusterProgramStruct {
	rv := objc.Send[MTRChannelClusterProgramStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramStruct) Init() MTRChannelClusterProgramStruct {
	rv := objc.Send[MTRChannelClusterProgramStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramStruct) Autorelease() MTRChannelClusterProgramStruct {
	rv := objc.Send[MTRChannelClusterProgramStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramStruct creates a new MTRChannelClusterProgramStruct instance.
func NewMTRChannelClusterProgramStruct() MTRChannelClusterProgramStruct {
	return getMTRChannelClusterProgramStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterProgramStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct
type MTRChannelClusterProgramStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramStructFrom constructs a [MTRChannelClusterProgramStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramStruct {
	return MTRChannelClusterProgramStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterProgramStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterProgramStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterProgramStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterProgramStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterProgramStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) AudioLanguages() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("audioLanguages"))
	return rv
}/* debug [instance_properties/getter]: audioLanguages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) SetAudioLanguages(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioLanguages:"), value)
}/* debug [instance_properties/setter]: audioLanguages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/channel
func (m_ MTRChannelClusterProgramStruct) Channel() objc.IObject /* cross-framework: MTRChannelClusterChannelInfoStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/channel
func (m_ MTRChannelClusterProgramStruct) SetChannel(value objc.IObject /* cross-framework: MTRChannelClusterChannelInfoStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/descriptionstring
func (m_ MTRChannelClusterProgramStruct) DescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionString"))
	return rv
}/* debug [instance_properties/getter]: descriptionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/descriptionstring
func (m_ MTRChannelClusterProgramStruct) SetDescriptionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionString:"), value)
}/* debug [instance_properties/setter]: descriptionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/endtime
func (m_ MTRChannelClusterProgramStruct) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}/* debug [instance_properties/getter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/endtime
func (m_ MTRChannelClusterProgramStruct) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}/* debug [instance_properties/setter]: endTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/identifier
func (m_ MTRChannelClusterProgramStruct) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/identifier
func (m_ MTRChannelClusterProgramStruct) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/parentalguidancetext
func (m_ MTRChannelClusterProgramStruct) ParentalGuidanceText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("parentalGuidanceText"))
	return rv
}/* debug [instance_properties/getter]: parentalGuidanceText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/parentalguidancetext
func (m_ MTRChannelClusterProgramStruct) SetParentalGuidanceText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParentalGuidanceText:"), value)
}/* debug [instance_properties/setter]: parentalGuidanceText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/recordingflag
func (m_ MTRChannelClusterProgramStruct) RecordingFlag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("recordingFlag"))
	return rv
}/* debug [instance_properties/getter]: recordingFlag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/recordingflag
func (m_ MTRChannelClusterProgramStruct) SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}/* debug [instance_properties/setter]: recordingFlag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/releasedate
func (m_ MTRChannelClusterProgramStruct) ReleaseDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("releaseDate"))
	return rv
}/* debug [instance_properties/getter]: releaseDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/releasedate
func (m_ MTRChannelClusterProgramStruct) SetReleaseDate(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseDate:"), value)
}/* debug [instance_properties/setter]: releaseDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/seriesinfo
func (m_ MTRChannelClusterProgramStruct) SeriesInfo() IMTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("seriesInfo"))
	return rv
}/* debug [instance_properties/getter]: seriesInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/seriesinfo
func (m_ MTRChannelClusterProgramStruct) SetSeriesInfo(value IMTRChannelClusterSeriesInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeriesInfo:"), value)
}/* debug [instance_properties/setter]: seriesInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/starttime
func (m_ MTRChannelClusterProgramStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/starttime
func (m_ MTRChannelClusterProgramStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/subtitle
func (m_ MTRChannelClusterProgramStruct) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/subtitle
func (m_ MTRChannelClusterProgramStruct) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/title
func (m_ MTRChannelClusterProgramStruct) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramstruct/title
func (m_ MTRChannelClusterProgramStruct) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterProgramStruct */



