// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass     _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass() _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass {
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass = _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent interface {
	IMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent
	// properties:
	BytesDownloaded() objc.IObject /* cross-framework: NSNumber */
	SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */)
	PlatformCode() objc.IObject /* cross-framework: NSNumber */
	SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */)
	ProgressPercent() objc.IObject /* cross-framework: NSNumber */
	SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t
type MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent struct {
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent
}

// MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	return MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent{
		MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent: MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass) New() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) Init() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent creates a new MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	return getMTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/bytesdownloaded
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) BytesDownloaded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("bytesDownloaded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/bytesdownloaded
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBytesDownloaded:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/platformcode
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) PlatformCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("platformCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/platformcode
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlatformCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/progresspercent
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) ProgressPercent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("progressPercent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/progresspercent
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressPercent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/softwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-73h3t/softwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}



