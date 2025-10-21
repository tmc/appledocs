// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterSeriesInfoStruct] class.
var (
	MTRChannelClusterSeriesInfoStructClass     _MTRChannelClusterSeriesInfoStructClass
	MTRChannelClusterSeriesInfoStructClassOnce sync.Once
)

func getMTRChannelClusterSeriesInfoStructClass() _MTRChannelClusterSeriesInfoStructClass {
	MTRChannelClusterSeriesInfoStructClassOnce.Do(func() {
		MTRChannelClusterSeriesInfoStructClass = _MTRChannelClusterSeriesInfoStructClass{objc.GetClass("MTRChannelClusterSeriesInfoStruct")}
	})
	return MTRChannelClusterSeriesInfoStructClass
}

type _MTRChannelClusterSeriesInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterSeriesInfoStruct] class.
type IMTRChannelClusterSeriesInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct
type MTRChannelClusterSeriesInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterSeriesInfoStructFrom constructs a [MTRChannelClusterSeriesInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterSeriesInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterSeriesInfoStruct {
	return MTRChannelClusterSeriesInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterSeriesInfoStructClass) Alloc() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterSeriesInfoStructClass) New() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterSeriesInfoStruct) Init() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterSeriesInfoStruct) Autorelease() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterSeriesInfoStruct creates a new MTRChannelClusterSeriesInfoStruct instance.
func NewMTRChannelClusterSeriesInfoStruct() MTRChannelClusterSeriesInfoStruct {
	return getMTRChannelClusterSeriesInfoStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) Episode() string {
	rv := objc.Send[string](m_.ID, objc.Sel("episode"))
	return rv
}


// SetEpisode sets the value of the episode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) SetEpisode(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpisode:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) Season() string {
	rv := objc.Send[string](m_.ID, objc.Sel("season"))
	return rv
}


// SetSeason sets the value of the season property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) SetSeason(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeason:"), objc.String(value))
}



