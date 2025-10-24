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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/bytesdownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) BytesDownloaded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("bytesDownloaded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/bytesdownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBytesDownloaded:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/platformcode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) PlatformCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("platformCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/platformcode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlatformCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/progresspercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) ProgressPercent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("progressPercent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/progresspercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressPercent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/softwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterdownloaderrorevent-2w5rw/softwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}



