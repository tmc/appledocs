// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRChannelClusterProgramStruct] class.
type IMTRChannelClusterProgramStruct interface {
	objectivec.IObject
	// properties:
	AudioLanguages() objc.IObject /* cross-framework: NSArray */
	SetAudioLanguages(value objc.IObject /* cross-framework: NSArray */)
	CastList() objc.IObject /* cross-framework: NSArray */
	SetCastList(value objc.IObject /* cross-framework: NSArray */)
	CategoryList() objc.IObject /* cross-framework: NSArray */
	SetCategoryList(value objc.IObject /* cross-framework: NSArray */)
	Channel() IMTRChannelClusterChannelInfoStruct
	SetChannel(value IMTRChannelClusterChannelInfoStruct)
	DescriptionString() objc.IObject /* cross-framework: NSString */
	SetDescriptionString(value objc.IObject /* cross-framework: NSString */)
	EndTime() objc.IObject /* cross-framework: NSNumber */
	SetEndTime(value objc.IObject /* cross-framework: NSNumber */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	ParentalGuidanceText() objc.IObject /* cross-framework: NSString */
	SetParentalGuidanceText(value objc.IObject /* cross-framework: NSString */)
	Ratings() objc.IObject /* cross-framework: NSArray */
	SetRatings(value objc.IObject /* cross-framework: NSArray */)
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct
type MTRChannelClusterProgramStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramStructFrom constructs a [MTRChannelClusterProgramStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramStruct {
	return MTRChannelClusterProgramStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramStructClass) Alloc() MTRChannelClusterProgramStruct {
	rv := objc.Send[MTRChannelClusterProgramStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) AudioLanguages() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("audioLanguages"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) SetAudioLanguages(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioLanguages:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/castList
func (m_ MTRChannelClusterProgramStruct) CastList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("castList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/castList
func (m_ MTRChannelClusterProgramStruct) SetCastList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCastList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/categoryList
func (m_ MTRChannelClusterProgramStruct) CategoryList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("categoryList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/categoryList
func (m_ MTRChannelClusterProgramStruct) SetCategoryList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCategoryList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/channel
func (m_ MTRChannelClusterProgramStruct) Channel() IMTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](m_.ID, objc.Sel("channel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/channel
func (m_ MTRChannelClusterProgramStruct) SetChannel(value IMTRChannelClusterChannelInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/descriptionString
func (m_ MTRChannelClusterProgramStruct) DescriptionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/descriptionString
func (m_ MTRChannelClusterProgramStruct) SetDescriptionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/endTime
func (m_ MTRChannelClusterProgramStruct) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/endTime
func (m_ MTRChannelClusterProgramStruct) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/identifier
func (m_ MTRChannelClusterProgramStruct) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/identifier
func (m_ MTRChannelClusterProgramStruct) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/parentalGuidanceText
func (m_ MTRChannelClusterProgramStruct) ParentalGuidanceText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("parentalGuidanceText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/parentalGuidanceText
func (m_ MTRChannelClusterProgramStruct) SetParentalGuidanceText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParentalGuidanceText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/ratings
func (m_ MTRChannelClusterProgramStruct) Ratings() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("ratings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/ratings
func (m_ MTRChannelClusterProgramStruct) SetRatings(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRatings:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/recordingFlag
func (m_ MTRChannelClusterProgramStruct) RecordingFlag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("recordingFlag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/recordingFlag
func (m_ MTRChannelClusterProgramStruct) SetRecordingFlag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/releaseDate
func (m_ MTRChannelClusterProgramStruct) ReleaseDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("releaseDate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/releaseDate
func (m_ MTRChannelClusterProgramStruct) SetReleaseDate(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseDate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/seriesInfo
func (m_ MTRChannelClusterProgramStruct) SeriesInfo() IMTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("seriesInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/seriesInfo
func (m_ MTRChannelClusterProgramStruct) SetSeriesInfo(value IMTRChannelClusterSeriesInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeriesInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/startTime
func (m_ MTRChannelClusterProgramStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/startTime
func (m_ MTRChannelClusterProgramStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/subtitle
func (m_ MTRChannelClusterProgramStruct) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/subtitle
func (m_ MTRChannelClusterProgramStruct) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/title
func (m_ MTRChannelClusterProgramStruct) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/title
func (m_ MTRChannelClusterProgramStruct) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}



