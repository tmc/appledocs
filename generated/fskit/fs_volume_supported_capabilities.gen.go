// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FSVolumeSupportedCapabilities] class.
var (
	FSVolumeSupportedCapabilitiesClass     _FSVolumeSupportedCapabilitiesClass
	FSVolumeSupportedCapabilitiesClassOnce sync.Once
)

func getFSVolumeSupportedCapabilitiesClass() _FSVolumeSupportedCapabilitiesClass {
	FSVolumeSupportedCapabilitiesClassOnce.Do(func() {
		FSVolumeSupportedCapabilitiesClass = _FSVolumeSupportedCapabilitiesClass{objc.GetClass("FSVolumeSupportedCapabilities")}
	})
	return FSVolumeSupportedCapabilitiesClass
}

type _FSVolumeSupportedCapabilitiesClass struct {
	class objc.Class
}





// An interface definition for the [FSVolumeSupportedCapabilities] class.
type IFSVolumeSupportedCapabilities interface {
	objectivec.IObject
	

	// properties:
	CaseFormat() FSVolumeCaseFormat
	SetCaseFormat(value FSVolumeCaseFormat)
	DoesNotSupportImmutableFiles() bool
	SetDoesNotSupportImmutableFiles(value bool)
	DoesNotSupportRootTimes() bool
	SetDoesNotSupportRootTimes(value bool)
	DoesNotSupportSettingFilePermissions() bool
	SetDoesNotSupportSettingFilePermissions(value bool)
	DoesNotSupportVolumeSizes() bool
	SetDoesNotSupportVolumeSizes(value bool)
	Supports2TBFiles() bool
	SetSupports2TBFiles(value bool)
	Supports64BitObjectIDs() bool
	SetSupports64BitObjectIDs(value bool)
	SupportsActiveJournal() bool
	SetSupportsActiveJournal(value bool)
	SupportsDocumentID() bool
	SetSupportsDocumentID(value bool)
	SupportsFastStatFS() bool
	SetSupportsFastStatFS(value bool)
	SupportsHardLinks() bool
	SetSupportsHardLinks(value bool)
	SupportsHiddenFiles() bool
	SetSupportsHiddenFiles(value bool)
	SupportsJournal() bool
	SetSupportsJournal(value bool)
	SupportsOpenDenyModes() bool
	SetSupportsOpenDenyModes(value bool)
	SupportsPersistentObjectIDs() bool
	SetSupportsPersistentObjectIDs(value bool)
	SupportsSharedSpace() bool
	SetSupportsSharedSpace(value bool)
	SupportsSparseFiles() bool
	SetSupportsSparseFiles(value bool)
	SupportsSymbolicLinks() bool
	SetSupportsSymbolicLinks(value bool)
	SupportsVolumeGroups() bool
	SetSupportsVolumeGroups(value bool)
	SupportsZeroRuns() bool
	SetSupportsZeroRuns(value bool)
	SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities
	SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities)
	VolumeStatistics() IFSStatFSResult
	SetVolumeStatistics(value IFSStatFSResult)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FSVolumeSupportedCapabilitiesClass) Alloc() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSVolumeSupportedCapabilitiesClass) New() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSVolumeSupportedCapabilities) Init() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSVolumeSupportedCapabilities) Autorelease() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeSupportedCapabilities creates a new FSVolumeSupportedCapabilities instance.
func NewFSVolumeSupportedCapabilities() FSVolumeSupportedCapabilities {
	return getFSVolumeSupportedCapabilitiesClass().New()
}





// A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.


// A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities
type FSVolumeSupportedCapabilities struct {
	objectivec.Object
}

// FSVolumeSupportedCapabilitiesFrom constructs a [FSVolumeSupportedCapabilities] from an unsafe.Pointer.
//
// A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.
func FSVolumeSupportedCapabilitiesFrom(ptr unsafe.Pointer) FSVolumeSupportedCapabilities {
	return FSVolumeSupportedCapabilities{objectivec.Object{objc.ID(ptr)}}
}

























// A value that indicates the volume’s support for case sensitivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/caseFormat
func (f_ FSVolumeSupportedCapabilities) CaseFormat() FSVolumeCaseFormat {
	rv := objc.Send[FSVolumeCaseFormat](f_.ID, objc.Sel("caseFormat"))
	return rv
}


// A value that indicates the volume’s support for case sensitivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/caseFormat
func (f_ FSVolumeSupportedCapabilities) SetCaseFormat(value FSVolumeCaseFormat) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCaseFormat:"), value)
}


// A Boolean property that indicates the volume doesn’t support immutable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportImmutableFiles
func (f_ FSVolumeSupportedCapabilities) DoesNotSupportImmutableFiles() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("doesNotSupportImmutableFiles"))
	return rv
}


// A Boolean property that indicates the volume doesn’t support immutable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportImmutableFiles
func (f_ FSVolumeSupportedCapabilities) SetDoesNotSupportImmutableFiles(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDoesNotSupportImmutableFiles:"), value)
}


// A Boolan property that indicates the volume doesn’t store reliable times for the root directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportRootTimes
func (f_ FSVolumeSupportedCapabilities) DoesNotSupportRootTimes() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("doesNotSupportRootTimes"))
	return rv
}


// A Boolan property that indicates the volume doesn’t store reliable times for the root directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportRootTimes
func (f_ FSVolumeSupportedCapabilities) SetDoesNotSupportRootTimes(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDoesNotSupportRootTimes:"), value)
}


// A Boolean property that indicates the volume doesn’t set file permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportSettingFilePermissions
func (f_ FSVolumeSupportedCapabilities) DoesNotSupportSettingFilePermissions() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("doesNotSupportSettingFilePermissions"))
	return rv
}


// A Boolean property that indicates the volume doesn’t set file permissions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportSettingFilePermissions
func (f_ FSVolumeSupportedCapabilities) SetDoesNotSupportSettingFilePermissions(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDoesNotSupportSettingFilePermissions:"), value)
}


// A Boolean property that indicates the volume doesn’t support certain volume size reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportVolumeSizes
func (f_ FSVolumeSupportedCapabilities) DoesNotSupportVolumeSizes() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("doesNotSupportVolumeSizes"))
	return rv
}


// A Boolean property that indicates the volume doesn’t support certain volume size reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportVolumeSizes
func (f_ FSVolumeSupportedCapabilities) SetDoesNotSupportVolumeSizes(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDoesNotSupportVolumeSizes:"), value)
}


// A Boolean property that indicates whether the volume supports file sizes larger than 4GB, and potentially up to 2TB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports2TBFiles
func (f_ FSVolumeSupportedCapabilities) Supports2TBFiles() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supports2TBFiles"))
	return rv
}


// A Boolean property that indicates whether the volume supports file sizes larger than 4GB, and potentially up to 2TB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports2TBFiles
func (f_ FSVolumeSupportedCapabilities) SetSupports2TBFiles(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupports2TBFiles:"), value)
}


// A Boolean property that indicates whether the volume supports 64-bit object IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports64BitObjectIDs
func (f_ FSVolumeSupportedCapabilities) Supports64BitObjectIDs() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supports64BitObjectIDs"))
	return rv
}


// A Boolean property that indicates whether the volume supports 64-bit object IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports64BitObjectIDs
func (f_ FSVolumeSupportedCapabilities) SetSupports64BitObjectIDs(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupports64BitObjectIDs:"), value)
}


// A Boolean property that indicates whether the volume currently uses a journal for speeding recovery after an unplanned shutdown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsActiveJournal
func (f_ FSVolumeSupportedCapabilities) SupportsActiveJournal() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsActiveJournal"))
	return rv
}


// A Boolean property that indicates whether the volume currently uses a journal for speeding recovery after an unplanned shutdown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsActiveJournal
func (f_ FSVolumeSupportedCapabilities) SetSupportsActiveJournal(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsActiveJournal:"), value)
}


// A Boolean property that indicates whether the volume supports document IDs for document revisions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsDocumentID
func (f_ FSVolumeSupportedCapabilities) SupportsDocumentID() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsDocumentID"))
	return rv
}


// A Boolean property that indicates whether the volume supports document IDs for document revisions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsDocumentID
func (f_ FSVolumeSupportedCapabilities) SetSupportsDocumentID(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsDocumentID:"), value)
}


// A Boolean property that indicates whether the volume supports fast results when fetching file system statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsFastStatFS
func (f_ FSVolumeSupportedCapabilities) SupportsFastStatFS() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsFastStatFS"))
	return rv
}


// A Boolean property that indicates whether the volume supports fast results when fetching file system statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsFastStatFS
func (f_ FSVolumeSupportedCapabilities) SetSupportsFastStatFS(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsFastStatFS:"), value)
}


// A Boolean property that indicates whether the volume supports hard links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHardLinks
func (f_ FSVolumeSupportedCapabilities) SupportsHardLinks() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsHardLinks"))
	return rv
}


// A Boolean property that indicates whether the volume supports hard links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHardLinks
func (f_ FSVolumeSupportedCapabilities) SetSupportsHardLinks(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsHardLinks:"), value)
}


// A Boolean property that indicates whether the volume supports hidden files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHiddenFiles
func (f_ FSVolumeSupportedCapabilities) SupportsHiddenFiles() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsHiddenFiles"))
	return rv
}


// A Boolean property that indicates whether the volume supports hidden files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHiddenFiles
func (f_ FSVolumeSupportedCapabilities) SetSupportsHiddenFiles(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsHiddenFiles:"), value)
}


// A Boolean property that indicates whether the volume supports a journal used to speed recovery in case of unplanned restart, such as a power outage or crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsJournal
func (f_ FSVolumeSupportedCapabilities) SupportsJournal() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsJournal"))
	return rv
}


// A Boolean property that indicates whether the volume supports a journal used to speed recovery in case of unplanned restart, such as a power outage or crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsJournal
func (f_ FSVolumeSupportedCapabilities) SetSupportsJournal(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsJournal:"), value)
}


// A Boolean property that indicates whether the volume supports open deny modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsOpenDenyModes
func (f_ FSVolumeSupportedCapabilities) SupportsOpenDenyModes() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsOpenDenyModes"))
	return rv
}


// A Boolean property that indicates whether the volume supports open deny modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsOpenDenyModes
func (f_ FSVolumeSupportedCapabilities) SetSupportsOpenDenyModes(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsOpenDenyModes:"), value)
}


// A Boolean property that indicates whether the volume supports persistent object identifiers and can look up file system objects by their IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsPersistentObjectIDs
func (f_ FSVolumeSupportedCapabilities) SupportsPersistentObjectIDs() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsPersistentObjectIDs"))
	return rv
}


// A Boolean property that indicates whether the volume supports persistent object identifiers and can look up file system objects by their IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsPersistentObjectIDs
func (f_ FSVolumeSupportedCapabilities) SetSupportsPersistentObjectIDs(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsPersistentObjectIDs:"), value)
}


// A Boolean property that indicates whether the volume supports multiple logical file systems that share space in a single “partition.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSharedSpace
func (f_ FSVolumeSupportedCapabilities) SupportsSharedSpace() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsSharedSpace"))
	return rv
}


// A Boolean property that indicates whether the volume supports multiple logical file systems that share space in a single “partition.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSharedSpace
func (f_ FSVolumeSupportedCapabilities) SetSupportsSharedSpace(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsSharedSpace:"), value)
}


// A Boolean property that indicates whether the volume supports sparse files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSparseFiles
func (f_ FSVolumeSupportedCapabilities) SupportsSparseFiles() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsSparseFiles"))
	return rv
}


// A Boolean property that indicates whether the volume supports sparse files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSparseFiles
func (f_ FSVolumeSupportedCapabilities) SetSupportsSparseFiles(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsSparseFiles:"), value)
}


// A Boolean property that indicates whether the volume supports symbolic links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSymbolicLinks
func (f_ FSVolumeSupportedCapabilities) SupportsSymbolicLinks() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsSymbolicLinks"))
	return rv
}


// A Boolean property that indicates whether the volume supports symbolic links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSymbolicLinks
func (f_ FSVolumeSupportedCapabilities) SetSupportsSymbolicLinks(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsSymbolicLinks:"), value)
}


// A Boolean property that indicates whether the volume supports volume groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsVolumeGroups
func (f_ FSVolumeSupportedCapabilities) SupportsVolumeGroups() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsVolumeGroups"))
	return rv
}


// A Boolean property that indicates whether the volume supports volume groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsVolumeGroups
func (f_ FSVolumeSupportedCapabilities) SetSupportsVolumeGroups(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsVolumeGroups:"), value)
}


// A Boolean property that indicates whether the volume supports zero runs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsZeroRuns
func (f_ FSVolumeSupportedCapabilities) SupportsZeroRuns() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("supportsZeroRuns"))
	return rv
}


// A Boolean property that indicates whether the volume supports zero runs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsZeroRuns
func (f_ FSVolumeSupportedCapabilities) SetSupportsZeroRuns(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportsZeroRuns:"), value)
}


// A property that provides the supported capabilities of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/supportedvolumecapabilities
func (f_ FSVolumeSupportedCapabilities) SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("supportedVolumeCapabilities"))
	return rv
}


// A property that provides the supported capabilities of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/supportedvolumecapabilities
func (f_ FSVolumeSupportedCapabilities) SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportedVolumeCapabilities:"), value)
}


// A property that provides up-to-date statistics of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/volumestatistics
func (f_ FSVolumeSupportedCapabilities) VolumeStatistics() IFSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("volumeStatistics"))
	return rv
}


// A property that provides up-to-date statistics of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/volumestatistics
func (f_ FSVolumeSupportedCapabilities) SetVolumeStatistics(value IFSStatFSResult) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setVolumeStatistics:"), value)
}








