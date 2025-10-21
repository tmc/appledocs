// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass     _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass() _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass {
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass = _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass
}

type _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] class.
type IMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw
type MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	return MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass) New() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) Init() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent creates a new MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	return getMTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/bytesdownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) BytesDownloaded() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("bytesDownloaded"))
	return rv
}


// SetBytesDownloaded sets the value of the bytesDownloaded property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/bytesdownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetBytesDownloaded(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBytesDownloaded:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/platformcode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) PlatformCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("platformCode"))
	return rv
}


// SetPlatformCode sets the value of the platformCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/platformcode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetPlatformCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlatformCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/progresspercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) ProgressPercent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("progressPercent"))
	return rv
}


// SetProgressPercent sets the value of the progressPercent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/progresspercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetProgressPercent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressPercent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/softwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/softwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetSoftwareVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}



