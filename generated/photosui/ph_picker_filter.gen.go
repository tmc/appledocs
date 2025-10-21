// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHPickerFilter] class.
var (
	PHPickerFilterClass     _PHPickerFilterClass
	PHPickerFilterClassOnce sync.Once
)

func getPHPickerFilterClass() _PHPickerFilterClass {
	PHPickerFilterClassOnce.Do(func() {
		PHPickerFilterClass = _PHPickerFilterClass{objc.GetClass("PHPickerFilter")}
	})
	return PHPickerFilterClass
}

type _PHPickerFilterClass struct {
	class objc.Class
}

// An interface definition for the [PHPickerFilter] class.
type IPHPickerFilter interface {
	objectivec.IObject
}

// A type that defines the filter to apply to the photo library.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class
type PHPickerFilter struct {
	objectivec.Object
}

// PHPickerFilterFrom constructs a [PHPickerFilter] from an unsafe.Pointer.
//
// A type that defines the filter to apply to the photo library.
func PHPickerFilterFrom(ptr unsafe.Pointer) PHPickerFilter {
	return PHPickerFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHPickerFilterClass) Alloc() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHPickerFilterClass) New() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHPickerFilter) Init() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHPickerFilter) Autorelease() PHPickerFilter {
	rv := objc.Send[PHPickerFilter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHPickerFilter creates a new PHPickerFilter instance.
func NewPHPickerFilter() PHPickerFilter {
	return getPHPickerFilterClass().New()
}


// Creates a new filter that includes only the filters you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/allFilterMatchingSubfilters:
func (pc _PHPickerFilterClass) AllFilterMatchingSubfilters(subfilters unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("allFilterMatchingSubfilters:"), subfilters)
	return rv
}

// Creates a new filter by combining the filters in the array.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/anyFilterMatchingSubfilters:
func (pc _PHPickerFilterClass) AnyFilterMatchingSubfilters(subfilters unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("anyFilterMatchingSubfilters:"), subfilters)
	return rv
}

// Creates a new filter that excludes the filter you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/notFilterOfSubfilter:
func (pc _PHPickerFilterClass) NotFilterOfSubfilter(subfilter unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("notFilterOfSubfilter:"), subfilter)
	return rv
}

// Creates a new filter by using the playback style you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/playbackStyleFilter:
func (pc _PHPickerFilterClass) PlaybackStyleFilter(playbackStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("playbackStyleFilter:"), playbackStyle)
	return rv
}

// A filter that represents assets with multiple high-speed photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/burstsFilter
func (pc _PHPickerFilterClass) BurstsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("burstsFilter"))
	return rv
}
// A filter that represents videos with a shallow depth of field and focus transitions.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/cinematicVideosFilter
func (pc _PHPickerFilterClass) CinematicVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("cinematicVideosFilter"))
	return rv
}
// A filter that represents photos with depth information.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/depthEffectPhotosFilter
func (pc _PHPickerFilterClass) DepthEffectPhotosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("depthEffectPhotosFilter"))
	return rv
}
// A filter that represents images, and includes Live Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/imagesFilter
func (pc _PHPickerFilterClass) ImagesFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("imagesFilter"))
	return rv
}
// A filter that represents Live Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/livePhotosFilter
func (pc _PHPickerFilterClass) LivePhotosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("livePhotosFilter"))
	return rv
}
// A filter that represents panorama photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/panoramasFilter
func (pc _PHPickerFilterClass) PanoramasFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("panoramasFilter"))
	return rv
}
// A filter that represents screen recordings.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/screenRecordingsFilter
func (pc _PHPickerFilterClass) ScreenRecordingsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("screenRecordingsFilter"))
	return rv
}
// A filter that represents screenshots.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/screenshotsFilter
func (pc _PHPickerFilterClass) ScreenshotsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("screenshotsFilter"))
	return rv
}
// A filter that represents slow-motion videos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/slomoVideosFilter
func (pc _PHPickerFilterClass) SlomoVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("slomoVideosFilter"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/spatialMediaFilter
func (pc _PHPickerFilterClass) SpatialMediaFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("spatialMediaFilter"))
	return rv
}
// A filter that represents time-lapse videos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/timelapseVideosFilter
func (pc _PHPickerFilterClass) TimelapseVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("timelapseVideosFilter"))
	return rv
}
// A filter that represents video assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/videosFilter
func (pc _PHPickerFilterClass) VideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("videosFilter"))
	return rv
}
// A filter that represents assets with multiple high-speed photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/burstsFilter
func (p_ PHPickerFilter) BurstsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("burstsFilter"))
	return rv
}

// A filter that represents videos with a shallow depth of field and focus transitions.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/cinematicVideosFilter
func (p_ PHPickerFilter) CinematicVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cinematicVideosFilter"))
	return rv
}

// A filter that represents photos with depth information.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/depthEffectPhotosFilter
func (p_ PHPickerFilter) DepthEffectPhotosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("depthEffectPhotosFilter"))
	return rv
}

// A filter that represents images, and includes Live Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/imagesFilter
func (p_ PHPickerFilter) ImagesFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("imagesFilter"))
	return rv
}

// A filter that represents Live Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/livePhotosFilter
func (p_ PHPickerFilter) LivePhotosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("livePhotosFilter"))
	return rv
}

// A filter that represents panorama photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/panoramasFilter
func (p_ PHPickerFilter) PanoramasFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("panoramasFilter"))
	return rv
}

// A filter that represents screen recordings.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/screenRecordingsFilter
func (p_ PHPickerFilter) ScreenRecordingsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("screenRecordingsFilter"))
	return rv
}

// A filter that represents screenshots.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/screenshotsFilter
func (p_ PHPickerFilter) ScreenshotsFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("screenshotsFilter"))
	return rv
}

// A filter that represents slow-motion videos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/slomoVideosFilter
func (p_ PHPickerFilter) SlomoVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("slomoVideosFilter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/spatialMediaFilter
func (p_ PHPickerFilter) SpatialMediaFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("spatialMediaFilter"))
	return rv
}

// A filter that represents time-lapse videos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/timelapseVideosFilter
func (p_ PHPickerFilter) TimelapseVideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timelapseVideosFilter"))
	return rv
}

// A filter that represents video assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerFilter-c.class/videosFilter
func (p_ PHPickerFilter) VideosFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videosFilter"))
	return rv
}



