// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Episode() objc.IObject /* cross-framework: NSString */
	SetEpisode(value objc.IObject /* cross-framework: NSString */)
	Season() objc.IObject /* cross-framework: NSString */
	SetSeason(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) Episode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("episode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) SetEpisode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpisode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) Season() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("season"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) SetSeason(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeason:"), value)
}



