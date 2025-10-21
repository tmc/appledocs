// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHAdjustmentData] class.
var (
	PHAdjustmentDataClass     _PHAdjustmentDataClass
	PHAdjustmentDataClassOnce sync.Once
)

func getPHAdjustmentDataClass() _PHAdjustmentDataClass {
	PHAdjustmentDataClassOnce.Do(func() {
		PHAdjustmentDataClass = _PHAdjustmentDataClass{objc.GetClass("PHAdjustmentData")}
	})
	return PHAdjustmentDataClass
}

type _PHAdjustmentDataClass struct {
	class objc.Class
}

// An interface definition for the [PHAdjustmentData] class.
type IPHAdjustmentData interface {
	objectivec.IObject
}

// A description of the edits made to an asset’s photo, video, or Live Photo content, which allows your app to reconstruct or revert the effects of prior editing sessions.
//
// When a user edits an asset, Photos saves a object along with the modified image or video data. This object provides an application-defined “recipe” you can use to reconstruct the edit. For example, if your app applies filters to a photo, you might create adjustment data that identifies which filters the user picked, the parameters for each, and the order to apply the filters in. Later, the user can resume working with those filters and parameters by using your app or another app that understands your adjustment data format. When iCloud Photos is enabled, a user can revert or resume edits made on a different device. You work with adjustment data when editing an asset, using either the method or a photo extension view controller that implements the protocol. When you begin an edit (through a object), examine the editing input’s property to decide whether the last edit made to the asset is compatible with your app. If so, you can allow the user to resume working with that edit. If not, you can make further edits to the last saved version of the photo. When you commit an edit (through a object), provide a new adjustment whose data represents the changes your app made. For each asset, Photos stores only one object, representing the edit made to the asset’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAdjustmentData
type PHAdjustmentData struct {
	objectivec.Object
}

// PHAdjustmentDataFrom constructs a [PHAdjustmentData] from an unsafe.Pointer.
//
// A description of the edits made to an asset’s photo, video, or Live Photo content, which allows your app to reconstruct or revert the effects of prior editing sessions.
func PHAdjustmentDataFrom(ptr unsafe.Pointer) PHAdjustmentData {
	return PHAdjustmentData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAdjustmentDataClass) Alloc() PHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAdjustmentDataClass) New() PHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAdjustmentData) Init() PHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAdjustmentData) Autorelease() PHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAdjustmentData creates a new PHAdjustmentData instance.
func NewPHAdjustmentData() PHAdjustmentData {
	return getPHAdjustmentDataClass().New()
}


// Initializes an adjustment object with the specified format and data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAdjustmentData/init(formatIdentifier:formatVersion:data:)
func NewPHAdjustmentDataWithFormatIdentifierFormatVersionData(formatIdentifier string, formatVersion string, data unsafe.Pointer) PHAdjustmentData {
	instance := getPHAdjustmentDataClass().Alloc()
	rv := objc.Send[PHAdjustmentData](instance.ID, objc.Sel("initWithFormatIdentifier:formatVersion:data:"), objc.String(formatIdentifier), objc.String(formatVersion), data)
	rv.Autorelease()
	return rv
}


// Data that contains the information necessary to reconstruct the adjustment.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAdjustmentData/data
func (p_ PHAdjustmentData) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("data"))
	return rv
}

// A string uniquely identifying the format of the adjustment data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAdjustmentData/formatIdentifier
func (p_ PHAdjustmentData) FormatIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("formatIdentifier"))
	return rv
}

// A version number for the adjustment data format.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAdjustmentData/formatVersion
func (p_ PHAdjustmentData) FormatVersion() string {
	rv := objc.Send[string](p_.ID, objc.Sel("formatVersion"))
	return rv
}


