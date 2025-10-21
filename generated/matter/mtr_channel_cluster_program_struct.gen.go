// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) AudioLanguages() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("audioLanguages"))
	return rv
}


// SetAudioLanguages sets the value of the audioLanguages property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/audioLanguages
func (m_ MTRChannelClusterProgramStruct) SetAudioLanguages(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioLanguages:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/castList
func (m_ MTRChannelClusterProgramStruct) CastList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("castList"))
	return rv
}


// SetCastList sets the value of the castList property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/castList
func (m_ MTRChannelClusterProgramStruct) SetCastList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCastList:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/categoryList
func (m_ MTRChannelClusterProgramStruct) CategoryList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("categoryList"))
	return rv
}


// SetCategoryList sets the value of the categoryList property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/categoryList
func (m_ MTRChannelClusterProgramStruct) SetCategoryList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCategoryList:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/channel
func (m_ MTRChannelClusterProgramStruct) Channel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/channel
func (m_ MTRChannelClusterProgramStruct) SetChannel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/descriptionString
func (m_ MTRChannelClusterProgramStruct) DescriptionString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("descriptionString"))
	return rv
}


// SetDescriptionString sets the value of the descriptionString property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/descriptionString
func (m_ MTRChannelClusterProgramStruct) SetDescriptionString(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionString:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/endTime
func (m_ MTRChannelClusterProgramStruct) EndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endTime"))
	return rv
}


// SetEndTime sets the value of the endTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/endTime
func (m_ MTRChannelClusterProgramStruct) SetEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/identifier
func (m_ MTRChannelClusterProgramStruct) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/identifier
func (m_ MTRChannelClusterProgramStruct) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/parentalGuidanceText
func (m_ MTRChannelClusterProgramStruct) ParentalGuidanceText() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parentalGuidanceText"))
	return rv
}


// SetParentalGuidanceText sets the value of the parentalGuidanceText property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/parentalGuidanceText
func (m_ MTRChannelClusterProgramStruct) SetParentalGuidanceText(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParentalGuidanceText:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/ratings
func (m_ MTRChannelClusterProgramStruct) Ratings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("ratings"))
	return rv
}


// SetRatings sets the value of the ratings property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/ratings
func (m_ MTRChannelClusterProgramStruct) SetRatings(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRatings:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/recordingFlag
func (m_ MTRChannelClusterProgramStruct) RecordingFlag() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("recordingFlag"))
	return rv
}


// SetRecordingFlag sets the value of the recordingFlag property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/recordingFlag
func (m_ MTRChannelClusterProgramStruct) SetRecordingFlag(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecordingFlag:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/releaseDate
func (m_ MTRChannelClusterProgramStruct) ReleaseDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("releaseDate"))
	return rv
}


// SetReleaseDate sets the value of the releaseDate property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/releaseDate
func (m_ MTRChannelClusterProgramStruct) SetReleaseDate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReleaseDate:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/seriesInfo
func (m_ MTRChannelClusterProgramStruct) SeriesInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("seriesInfo"))
	return rv
}


// SetSeriesInfo sets the value of the seriesInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/seriesInfo
func (m_ MTRChannelClusterProgramStruct) SetSeriesInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeriesInfo:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/startTime
func (m_ MTRChannelClusterProgramStruct) StartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/startTime
func (m_ MTRChannelClusterProgramStruct) SetStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/subtitle
func (m_ MTRChannelClusterProgramStruct) Subtitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/subtitle
func (m_ MTRChannelClusterProgramStruct) SetSubtitle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/title
func (m_ MTRChannelClusterProgramStruct) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramStruct/title
func (m_ MTRChannelClusterProgramStruct) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}


