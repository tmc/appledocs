// Code generated from Apple documentation for AppleArchive. DO NOT EDIT.

package applearchive

/* debug [functions.gen.go]: Generating 144 functions for AppleArchive */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AppleArchive Functions (144 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AAArchiveStreamCancel func(unsafe.Pointer)
	_AAArchiveStreamClose func(unsafe.Pointer) int
	_AAArchiveStreamProcess func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAArchiveStreamReadBlob func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAArchiveStreamReadHeader func(unsafe.Pointer, unsafe.Pointer) int
	_AAArchiveStreamWriteBlob func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAArchiveStreamWriteHeader func(unsafe.Pointer, unsafe.Pointer) int
	_AAArchiveStreamWritePathList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) int
	_AAByteStreamCancel func(unsafe.Pointer)
	_AAByteStreamClose func(unsafe.Pointer) int
	_AAByteStreamPRead func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_AAByteStreamProcess func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAByteStreamPWrite func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_AAByteStreamRead func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_AAByteStreamSeek func(unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAByteStreamWrite func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_AACompressionOutputStreamOpen func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, int) unsafe.Pointer
	_AACompressionOutputStreamOpenExisting func(unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAConvertArchiveOutputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AACustomArchiveStreamOpen func() unsafe.Pointer
	_AACustomArchiveStreamSetCancelProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetCloseProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetData func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetReadBlobProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetReadHeaderProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetWriteBlobProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomArchiveStreamSetWriteHeaderProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamOpen func() unsafe.Pointer
	_AACustomByteStreamSetCancelProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetCloseProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetData func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetPReadProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetPWriteProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetReadProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetSeekProc func(unsafe.Pointer, unsafe.Pointer)
	_AACustomByteStreamSetWriteProc func(unsafe.Pointer, unsafe.Pointer)
	_AADecodeArchiveInputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AADecompressionInputStreamOpen func(unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AADecompressionRandomAccessInputStreamOpen func(unsafe.Pointer, uintptr, unsafe.Pointer, int) unsafe.Pointer
	_AAEncodeArchiveOutputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAEntryACLBlobAppendEntry func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAEntryACLBlobApplyToPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_AAEntryACLBlobClear func(unsafe.Pointer) int
	_AAEntryACLBlobCreate func() unsafe.Pointer
	_AAEntryACLBlobCreateWithEncodedData func(unsafe.Pointer, uintptr) unsafe.Pointer
	_AAEntryACLBlobCreateWithPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAEntryACLBlobDestroy func(unsafe.Pointer)
	_AAEntryACLBlobGetEncodedData func(unsafe.Pointer) unsafe.Pointer
	_AAEntryACLBlobGetEncodedSize func(unsafe.Pointer) uintptr
	_AAEntryACLBlobGetEntry func(unsafe.Pointer, uint32, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AAEntryACLBlobGetEntryCount func(unsafe.Pointer) uint32
	_AAEntryACLBlobRemoveEntry func(unsafe.Pointer, uint32) int
	_AAEntryACLBlobSetEntry func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAEntryXATBlobAppendEntry func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAEntryXATBlobApplyToPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_AAEntryXATBlobClear func(unsafe.Pointer) int
	_AAEntryXATBlobCreate func() unsafe.Pointer
	_AAEntryXATBlobCreateWithEncodedData func(unsafe.Pointer, uintptr) unsafe.Pointer
	_AAEntryXATBlobCreateWithPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAEntryXATBlobDestroy func(unsafe.Pointer)
	_AAEntryXATBlobGetEncodedData func(unsafe.Pointer) unsafe.Pointer
	_AAEntryXATBlobGetEncodedSize func(unsafe.Pointer) uintptr
	_AAEntryXATBlobGetEntry func(unsafe.Pointer, uint32, uintptr, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AAEntryXATBlobGetEntryCount func(unsafe.Pointer) uint32
	_AAEntryXATBlobRemoveEntry func(unsafe.Pointer, uint32) int
	_AAEntryXATBlobSetEntry func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAExtractArchiveOutputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAFieldKeySetClear func(unsafe.Pointer) int
	_AAFieldKeySetClone func(unsafe.Pointer) unsafe.Pointer
	_AAFieldKeySetContainsKey func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetCreate func() unsafe.Pointer
	_AAFieldKeySetCreateWithString func(unsafe.Pointer) unsafe.Pointer
	_AAFieldKeySetDestroy func(unsafe.Pointer)
	_AAFieldKeySetGetKey func(unsafe.Pointer, uint32) unsafe.Pointer
	_AAFieldKeySetGetKeyCount func(unsafe.Pointer) uint32
	_AAFieldKeySetInsertKey func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetInsertKeySet func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetRemoveKey func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetRemoveKeySet func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetSelectKeySet func(unsafe.Pointer, unsafe.Pointer) int
	_AAFieldKeySetSerialize func(unsafe.Pointer, uintptr, unsafe.Pointer) int
	_AAFileStreamOpenWithFD func(int, int) unsafe.Pointer
	_AAFileStreamOpenWithPath func(unsafe.Pointer, int, unsafe.Pointer) unsafe.Pointer
	_AAHeaderAssign func(unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderClear func(unsafe.Pointer) int
	_AAHeaderClone func(unsafe.Pointer) unsafe.Pointer
	_AAHeaderCreate func() unsafe.Pointer
	_AAHeaderCreateWithEncodedData func(uintptr, unsafe.Pointer) unsafe.Pointer
	_AAHeaderCreateWithPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAHeaderDestroy func(unsafe.Pointer)
	_AAHeaderGetEncodedData func(unsafe.Pointer) unsafe.Pointer
	_AAHeaderGetEncodedSize func(unsafe.Pointer) uintptr
	_AAHeaderGetFieldBlob func(unsafe.Pointer, uint32, []uint64, []uint64) int
	_AAHeaderGetFieldCount func(unsafe.Pointer) uint32
	_AAHeaderGetFieldHash func(unsafe.Pointer, uint32, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderGetFieldKey func(unsafe.Pointer, uint32) unsafe.Pointer
	_AAHeaderGetFieldString func(unsafe.Pointer, uint32, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderGetFieldTimespec func(unsafe.Pointer, uint32, unsafe.Pointer) int
	_AAHeaderGetFieldType func(unsafe.Pointer, uint32) int
	_AAHeaderGetFieldUInt func(unsafe.Pointer, uint32, []uint64) int
	_AAHeaderGetKeyIndex func(unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderGetPayloadSize func(unsafe.Pointer) uint64
	_AAHeaderRemoveField func(unsafe.Pointer, uint32) int
	_AAHeaderSetFieldBlob func(unsafe.Pointer, uint32, unsafe.Pointer, uint64) int
	_AAHeaderSetFieldFlag func(unsafe.Pointer, uint32, unsafe.Pointer) int
	_AAHeaderSetFieldHash func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderSetFieldString func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AAHeaderSetFieldTimespec func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) int
	_AAHeaderSetFieldUInt func(unsafe.Pointer, uint32, unsafe.Pointer, uint64) int
	_AAPathListCreateWithDirectoryContents func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AAPathListCreateWithPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAPathListDestroy func(unsafe.Pointer)
	_AAPathListNodeFirst func(unsafe.Pointer) uint64
	_AAPathListNodeGetPath func(unsafe.Pointer, uint64, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AAPathListNodeNext func(unsafe.Pointer, uint64) uint64
	_AARandomAccessByteStreamProcess func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, int) unsafe.Pointer
	_AASharedBufferPipeOpen func(unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AATempFileStreamOpen func() unsafe.Pointer
	_AEAAuthDataAppendEntry func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AEAAuthDataClear func(unsafe.Pointer) int
	_AEAAuthDataCreate func() unsafe.Pointer
	_AEAAuthDataCreateWithContext func(unsafe.Pointer) unsafe.Pointer
	_AEAAuthDataDestroy func(unsafe.Pointer)
	_AEAAuthDataGetEncodedData func(unsafe.Pointer) unsafe.Pointer
	_AEAAuthDataGetEncodedSize func(unsafe.Pointer) uintptr
	_AEAAuthDataGetEntry func(unsafe.Pointer, uint32, uintptr, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AEAAuthDataGetEntryCount func(unsafe.Pointer) uint32
	_AEAAuthDataRemoveEntry func(unsafe.Pointer, uint32) int
	_AEAAuthDataSetEntry func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AEAContextCreateWithEncryptedStream func(unsafe.Pointer) unsafe.Pointer
	_AEAContextCreateWithProfile func(unsafe.Pointer) unsafe.Pointer
	_AEAContextDecryptAttributes func(unsafe.Pointer) int
	_AEAContextDestroy func(unsafe.Pointer)
	_AEAContextGenerateFieldBlob func(unsafe.Pointer, unsafe.Pointer) int
	_AEAContextGetFieldBlob func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) int
	_AEAContextGetFieldUInt func(unsafe.Pointer, unsafe.Pointer) uint64
	_AEAContextSetFieldBlob func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) int
	_AEAContextSetFieldUInt func(unsafe.Pointer, unsafe.Pointer, uint64) int
	_AEADecryptionInputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AEADecryptionRandomAccessInputStreamOpen func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, int) unsafe.Pointer
	_AEAEncryptionOutputStreamCloseAndUpdateContext func(unsafe.Pointer, unsafe.Pointer) int
	_AEAEncryptionOutputStreamOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AEAEncryptionOutputStreamOpenExisting func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int) unsafe.Pointer
	_AEAStreamSign func(unsafe.Pointer, unsafe.Pointer) int
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AAArchiveStreamCancel, lib, "AAArchiveStreamCancel")
	tryRegister(&_AAArchiveStreamClose, lib, "AAArchiveStreamClose")
	tryRegister(&_AAArchiveStreamProcess, lib, "AAArchiveStreamProcess")
	tryRegister(&_AAArchiveStreamReadBlob, lib, "AAArchiveStreamReadBlob")
	tryRegister(&_AAArchiveStreamReadHeader, lib, "AAArchiveStreamReadHeader")
	tryRegister(&_AAArchiveStreamWriteBlob, lib, "AAArchiveStreamWriteBlob")
	tryRegister(&_AAArchiveStreamWriteHeader, lib, "AAArchiveStreamWriteHeader")
	tryRegister(&_AAArchiveStreamWritePathList, lib, "AAArchiveStreamWritePathList")
	tryRegister(&_AAByteStreamCancel, lib, "AAByteStreamCancel")
	tryRegister(&_AAByteStreamClose, lib, "AAByteStreamClose")
	tryRegister(&_AAByteStreamPRead, lib, "AAByteStreamPRead")
	tryRegister(&_AAByteStreamProcess, lib, "AAByteStreamProcess")
	tryRegister(&_AAByteStreamPWrite, lib, "AAByteStreamPWrite")
	tryRegister(&_AAByteStreamRead, lib, "AAByteStreamRead")
	tryRegister(&_AAByteStreamSeek, lib, "AAByteStreamSeek")
	tryRegister(&_AAByteStreamWrite, lib, "AAByteStreamWrite")
	tryRegister(&_AACompressionOutputStreamOpen, lib, "AACompressionOutputStreamOpen")
	tryRegister(&_AACompressionOutputStreamOpenExisting, lib, "AACompressionOutputStreamOpenExisting")
	tryRegister(&_AAConvertArchiveOutputStreamOpen, lib, "AAConvertArchiveOutputStreamOpen")
	tryRegister(&_AACustomArchiveStreamOpen, lib, "AACustomArchiveStreamOpen")
	tryRegister(&_AACustomArchiveStreamSetCancelProc, lib, "AACustomArchiveStreamSetCancelProc")
	tryRegister(&_AACustomArchiveStreamSetCloseProc, lib, "AACustomArchiveStreamSetCloseProc")
	tryRegister(&_AACustomArchiveStreamSetData, lib, "AACustomArchiveStreamSetData")
	tryRegister(&_AACustomArchiveStreamSetReadBlobProc, lib, "AACustomArchiveStreamSetReadBlobProc")
	tryRegister(&_AACustomArchiveStreamSetReadHeaderProc, lib, "AACustomArchiveStreamSetReadHeaderProc")
	tryRegister(&_AACustomArchiveStreamSetWriteBlobProc, lib, "AACustomArchiveStreamSetWriteBlobProc")
	tryRegister(&_AACustomArchiveStreamSetWriteHeaderProc, lib, "AACustomArchiveStreamSetWriteHeaderProc")
	tryRegister(&_AACustomByteStreamOpen, lib, "AACustomByteStreamOpen")
	tryRegister(&_AACustomByteStreamSetCancelProc, lib, "AACustomByteStreamSetCancelProc")
	tryRegister(&_AACustomByteStreamSetCloseProc, lib, "AACustomByteStreamSetCloseProc")
	tryRegister(&_AACustomByteStreamSetData, lib, "AACustomByteStreamSetData")
	tryRegister(&_AACustomByteStreamSetPReadProc, lib, "AACustomByteStreamSetPReadProc")
	tryRegister(&_AACustomByteStreamSetPWriteProc, lib, "AACustomByteStreamSetPWriteProc")
	tryRegister(&_AACustomByteStreamSetReadProc, lib, "AACustomByteStreamSetReadProc")
	tryRegister(&_AACustomByteStreamSetSeekProc, lib, "AACustomByteStreamSetSeekProc")
	tryRegister(&_AACustomByteStreamSetWriteProc, lib, "AACustomByteStreamSetWriteProc")
	tryRegister(&_AADecodeArchiveInputStreamOpen, lib, "AADecodeArchiveInputStreamOpen")
	tryRegister(&_AADecompressionInputStreamOpen, lib, "AADecompressionInputStreamOpen")
	tryRegister(&_AADecompressionRandomAccessInputStreamOpen, lib, "AADecompressionRandomAccessInputStreamOpen")
	tryRegister(&_AAEncodeArchiveOutputStreamOpen, lib, "AAEncodeArchiveOutputStreamOpen")
	tryRegister(&_AAEntryACLBlobAppendEntry, lib, "AAEntryACLBlobAppendEntry")
	tryRegister(&_AAEntryACLBlobApplyToPath, lib, "AAEntryACLBlobApplyToPath")
	tryRegister(&_AAEntryACLBlobClear, lib, "AAEntryACLBlobClear")
	tryRegister(&_AAEntryACLBlobCreate, lib, "AAEntryACLBlobCreate")
	tryRegister(&_AAEntryACLBlobCreateWithEncodedData, lib, "AAEntryACLBlobCreateWithEncodedData")
	tryRegister(&_AAEntryACLBlobCreateWithPath, lib, "AAEntryACLBlobCreateWithPath")
	tryRegister(&_AAEntryACLBlobDestroy, lib, "AAEntryACLBlobDestroy")
	tryRegister(&_AAEntryACLBlobGetEncodedData, lib, "AAEntryACLBlobGetEncodedData")
	tryRegister(&_AAEntryACLBlobGetEncodedSize, lib, "AAEntryACLBlobGetEncodedSize")
	tryRegister(&_AAEntryACLBlobGetEntry, lib, "AAEntryACLBlobGetEntry")
	tryRegister(&_AAEntryACLBlobGetEntryCount, lib, "AAEntryACLBlobGetEntryCount")
	tryRegister(&_AAEntryACLBlobRemoveEntry, lib, "AAEntryACLBlobRemoveEntry")
	tryRegister(&_AAEntryACLBlobSetEntry, lib, "AAEntryACLBlobSetEntry")
	tryRegister(&_AAEntryXATBlobAppendEntry, lib, "AAEntryXATBlobAppendEntry")
	tryRegister(&_AAEntryXATBlobApplyToPath, lib, "AAEntryXATBlobApplyToPath")
	tryRegister(&_AAEntryXATBlobClear, lib, "AAEntryXATBlobClear")
	tryRegister(&_AAEntryXATBlobCreate, lib, "AAEntryXATBlobCreate")
	tryRegister(&_AAEntryXATBlobCreateWithEncodedData, lib, "AAEntryXATBlobCreateWithEncodedData")
	tryRegister(&_AAEntryXATBlobCreateWithPath, lib, "AAEntryXATBlobCreateWithPath")
	tryRegister(&_AAEntryXATBlobDestroy, lib, "AAEntryXATBlobDestroy")
	tryRegister(&_AAEntryXATBlobGetEncodedData, lib, "AAEntryXATBlobGetEncodedData")
	tryRegister(&_AAEntryXATBlobGetEncodedSize, lib, "AAEntryXATBlobGetEncodedSize")
	tryRegister(&_AAEntryXATBlobGetEntry, lib, "AAEntryXATBlobGetEntry")
	tryRegister(&_AAEntryXATBlobGetEntryCount, lib, "AAEntryXATBlobGetEntryCount")
	tryRegister(&_AAEntryXATBlobRemoveEntry, lib, "AAEntryXATBlobRemoveEntry")
	tryRegister(&_AAEntryXATBlobSetEntry, lib, "AAEntryXATBlobSetEntry")
	tryRegister(&_AAExtractArchiveOutputStreamOpen, lib, "AAExtractArchiveOutputStreamOpen")
	tryRegister(&_AAFieldKeySetClear, lib, "AAFieldKeySetClear")
	tryRegister(&_AAFieldKeySetClone, lib, "AAFieldKeySetClone")
	tryRegister(&_AAFieldKeySetContainsKey, lib, "AAFieldKeySetContainsKey")
	tryRegister(&_AAFieldKeySetCreate, lib, "AAFieldKeySetCreate")
	tryRegister(&_AAFieldKeySetCreateWithString, lib, "AAFieldKeySetCreateWithString")
	tryRegister(&_AAFieldKeySetDestroy, lib, "AAFieldKeySetDestroy")
	tryRegister(&_AAFieldKeySetGetKey, lib, "AAFieldKeySetGetKey")
	tryRegister(&_AAFieldKeySetGetKeyCount, lib, "AAFieldKeySetGetKeyCount")
	tryRegister(&_AAFieldKeySetInsertKey, lib, "AAFieldKeySetInsertKey")
	tryRegister(&_AAFieldKeySetInsertKeySet, lib, "AAFieldKeySetInsertKeySet")
	tryRegister(&_AAFieldKeySetRemoveKey, lib, "AAFieldKeySetRemoveKey")
	tryRegister(&_AAFieldKeySetRemoveKeySet, lib, "AAFieldKeySetRemoveKeySet")
	tryRegister(&_AAFieldKeySetSelectKeySet, lib, "AAFieldKeySetSelectKeySet")
	tryRegister(&_AAFieldKeySetSerialize, lib, "AAFieldKeySetSerialize")
	tryRegister(&_AAFileStreamOpenWithFD, lib, "AAFileStreamOpenWithFD")
	tryRegister(&_AAFileStreamOpenWithPath, lib, "AAFileStreamOpenWithPath")
	tryRegister(&_AAHeaderAssign, lib, "AAHeaderAssign")
	tryRegister(&_AAHeaderClear, lib, "AAHeaderClear")
	tryRegister(&_AAHeaderClone, lib, "AAHeaderClone")
	tryRegister(&_AAHeaderCreate, lib, "AAHeaderCreate")
	tryRegister(&_AAHeaderCreateWithEncodedData, lib, "AAHeaderCreateWithEncodedData")
	tryRegister(&_AAHeaderCreateWithPath, lib, "AAHeaderCreateWithPath")
	tryRegister(&_AAHeaderDestroy, lib, "AAHeaderDestroy")
	tryRegister(&_AAHeaderGetEncodedData, lib, "AAHeaderGetEncodedData")
	tryRegister(&_AAHeaderGetEncodedSize, lib, "AAHeaderGetEncodedSize")
	tryRegister(&_AAHeaderGetFieldBlob, lib, "AAHeaderGetFieldBlob")
	tryRegister(&_AAHeaderGetFieldCount, lib, "AAHeaderGetFieldCount")
	tryRegister(&_AAHeaderGetFieldHash, lib, "AAHeaderGetFieldHash")
	tryRegister(&_AAHeaderGetFieldKey, lib, "AAHeaderGetFieldKey")
	tryRegister(&_AAHeaderGetFieldString, lib, "AAHeaderGetFieldString")
	tryRegister(&_AAHeaderGetFieldTimespec, lib, "AAHeaderGetFieldTimespec")
	tryRegister(&_AAHeaderGetFieldType, lib, "AAHeaderGetFieldType")
	tryRegister(&_AAHeaderGetFieldUInt, lib, "AAHeaderGetFieldUInt")
	tryRegister(&_AAHeaderGetKeyIndex, lib, "AAHeaderGetKeyIndex")
	tryRegister(&_AAHeaderGetPayloadSize, lib, "AAHeaderGetPayloadSize")
	tryRegister(&_AAHeaderRemoveField, lib, "AAHeaderRemoveField")
	tryRegister(&_AAHeaderSetFieldBlob, lib, "AAHeaderSetFieldBlob")
	tryRegister(&_AAHeaderSetFieldFlag, lib, "AAHeaderSetFieldFlag")
	tryRegister(&_AAHeaderSetFieldHash, lib, "AAHeaderSetFieldHash")
	tryRegister(&_AAHeaderSetFieldString, lib, "AAHeaderSetFieldString")
	tryRegister(&_AAHeaderSetFieldTimespec, lib, "AAHeaderSetFieldTimespec")
	tryRegister(&_AAHeaderSetFieldUInt, lib, "AAHeaderSetFieldUInt")
	tryRegister(&_AAPathListCreateWithDirectoryContents, lib, "AAPathListCreateWithDirectoryContents")
	tryRegister(&_AAPathListCreateWithPath, lib, "AAPathListCreateWithPath")
	tryRegister(&_AAPathListDestroy, lib, "AAPathListDestroy")
	tryRegister(&_AAPathListNodeFirst, lib, "AAPathListNodeFirst")
	tryRegister(&_AAPathListNodeGetPath, lib, "AAPathListNodeGetPath")
	tryRegister(&_AAPathListNodeNext, lib, "AAPathListNodeNext")
	tryRegister(&_AARandomAccessByteStreamProcess, lib, "AARandomAccessByteStreamProcess")
	tryRegister(&_AASharedBufferPipeOpen, lib, "AASharedBufferPipeOpen")
	tryRegister(&_AATempFileStreamOpen, lib, "AATempFileStreamOpen")
	tryRegister(&_AEAAuthDataAppendEntry, lib, "AEAAuthDataAppendEntry")
	tryRegister(&_AEAAuthDataClear, lib, "AEAAuthDataClear")
	tryRegister(&_AEAAuthDataCreate, lib, "AEAAuthDataCreate")
	tryRegister(&_AEAAuthDataCreateWithContext, lib, "AEAAuthDataCreateWithContext")
	tryRegister(&_AEAAuthDataDestroy, lib, "AEAAuthDataDestroy")
	tryRegister(&_AEAAuthDataGetEncodedData, lib, "AEAAuthDataGetEncodedData")
	tryRegister(&_AEAAuthDataGetEncodedSize, lib, "AEAAuthDataGetEncodedSize")
	tryRegister(&_AEAAuthDataGetEntry, lib, "AEAAuthDataGetEntry")
	tryRegister(&_AEAAuthDataGetEntryCount, lib, "AEAAuthDataGetEntryCount")
	tryRegister(&_AEAAuthDataRemoveEntry, lib, "AEAAuthDataRemoveEntry")
	tryRegister(&_AEAAuthDataSetEntry, lib, "AEAAuthDataSetEntry")
	tryRegister(&_AEAContextCreateWithEncryptedStream, lib, "AEAContextCreateWithEncryptedStream")
	tryRegister(&_AEAContextCreateWithProfile, lib, "AEAContextCreateWithProfile")
	tryRegister(&_AEAContextDecryptAttributes, lib, "AEAContextDecryptAttributes")
	tryRegister(&_AEAContextDestroy, lib, "AEAContextDestroy")
	tryRegister(&_AEAContextGenerateFieldBlob, lib, "AEAContextGenerateFieldBlob")
	tryRegister(&_AEAContextGetFieldBlob, lib, "AEAContextGetFieldBlob")
	tryRegister(&_AEAContextGetFieldUInt, lib, "AEAContextGetFieldUInt")
	tryRegister(&_AEAContextSetFieldBlob, lib, "AEAContextSetFieldBlob")
	tryRegister(&_AEAContextSetFieldUInt, lib, "AEAContextSetFieldUInt")
	tryRegister(&_AEADecryptionInputStreamOpen, lib, "AEADecryptionInputStreamOpen")
	tryRegister(&_AEADecryptionRandomAccessInputStreamOpen, lib, "AEADecryptionRandomAccessInputStreamOpen")
	tryRegister(&_AEAEncryptionOutputStreamCloseAndUpdateContext, lib, "AEAEncryptionOutputStreamCloseAndUpdateContext")
	tryRegister(&_AEAEncryptionOutputStreamOpen, lib, "AEAEncryptionOutputStreamOpen")
	tryRegister(&_AEAEncryptionOutputStreamOpenExisting, lib, "AEAEncryptionOutputStreamOpenExisting")
	tryRegister(&_AEAStreamSign, lib, "AEAStreamSign")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// AAArchiveStreamCancel is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamCancel
func AAArchiveStreamCancel(s unsafe.Pointer) {
	_AAArchiveStreamCancel(s)
}/* debug [functions.gen.go/function]: AAArchiveStreamCancel */

// AAArchiveStreamClose is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamClose
func AAArchiveStreamClose(s unsafe.Pointer) int {
	return _AAArchiveStreamClose(s)
}/* debug [functions.gen.go/function]: AAArchiveStreamClose */

// AAArchiveStreamProcess is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamProcess
func AAArchiveStreamProcess(istream unsafe.Pointer, ostream unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AAArchiveStreamProcess(istream, ostream, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAArchiveStreamProcess */

// AAArchiveStreamReadBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamReadBlob
func AAArchiveStreamReadBlob(s unsafe.Pointer, key unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr) int {
	return _AAArchiveStreamReadBlob(s, key, buf, nbyte)
}/* debug [functions.gen.go/function]: AAArchiveStreamReadBlob */

// AAArchiveStreamReadHeader is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamReadHeader
func AAArchiveStreamReadHeader(s unsafe.Pointer, header unsafe.Pointer) int {
	return _AAArchiveStreamReadHeader(s, header)
}/* debug [functions.gen.go/function]: AAArchiveStreamReadHeader */

// AAArchiveStreamWriteBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamWriteBlob
func AAArchiveStreamWriteBlob(s unsafe.Pointer, key unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr) int {
	return _AAArchiveStreamWriteBlob(s, key, buf, nbyte)
}/* debug [functions.gen.go/function]: AAArchiveStreamWriteBlob */

// AAArchiveStreamWriteHeader is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamWriteHeader
func AAArchiveStreamWriteHeader(s unsafe.Pointer, header unsafe.Pointer) int {
	return _AAArchiveStreamWriteHeader(s, header)
}/* debug [functions.gen.go/function]: AAArchiveStreamWriteHeader */

// AAArchiveStreamWritePathList is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAArchiveStreamWritePathList
func AAArchiveStreamWritePathList(s unsafe.Pointer, path_list unsafe.Pointer, key_set unsafe.Pointer, dir unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) int {
	return _AAArchiveStreamWritePathList(s, path_list, key_set, dir, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAArchiveStreamWritePathList */

// AAByteStreamCancel is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamCancel
func AAByteStreamCancel(s unsafe.Pointer) {
	_AAByteStreamCancel(s)
}/* debug [functions.gen.go/function]: AAByteStreamCancel */

// AAByteStreamClose is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamClose
func AAByteStreamClose(s unsafe.Pointer) int {
	return _AAByteStreamClose(s)
}/* debug [functions.gen.go/function]: AAByteStreamClose */

// AAByteStreamPRead is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamPRead
func AAByteStreamPRead(s unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr, offset unsafe.Pointer) unsafe.Pointer {
	return _AAByteStreamPRead(s, buf, nbyte, offset)
}/* debug [functions.gen.go/function]: AAByteStreamPRead */

// AAByteStreamProcess is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamProcess
func AAByteStreamProcess(istream unsafe.Pointer, ostream unsafe.Pointer) unsafe.Pointer {
	return _AAByteStreamProcess(istream, ostream)
}/* debug [functions.gen.go/function]: AAByteStreamProcess */

// AAByteStreamPWrite is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamPWrite
func AAByteStreamPWrite(s unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr, offset unsafe.Pointer) unsafe.Pointer {
	return _AAByteStreamPWrite(s, buf, nbyte, offset)
}/* debug [functions.gen.go/function]: AAByteStreamPWrite */

// AAByteStreamRead is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamRead
func AAByteStreamRead(s unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr) unsafe.Pointer {
	return _AAByteStreamRead(s, buf, nbyte)
}/* debug [functions.gen.go/function]: AAByteStreamRead */

// AAByteStreamSeek is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamSeek
func AAByteStreamSeek(s unsafe.Pointer, offset unsafe.Pointer, whence int) unsafe.Pointer {
	return _AAByteStreamSeek(s, offset, whence)
}/* debug [functions.gen.go/function]: AAByteStreamSeek */

// AAByteStreamWrite is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAByteStreamWrite
func AAByteStreamWrite(s unsafe.Pointer, buf unsafe.Pointer, nbyte uintptr) unsafe.Pointer {
	return _AAByteStreamWrite(s, buf, nbyte)
}/* debug [functions.gen.go/function]: AAByteStreamWrite */

// AACompressionOutputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionOutputStreamOpen
func AACompressionOutputStreamOpen(compressed_stream unsafe.Pointer, compression_algorithm unsafe.Pointer, block_size uintptr, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AACompressionOutputStreamOpen(compressed_stream, compression_algorithm, block_size, flags, n_threads)
}/* debug [functions.gen.go/function]: AACompressionOutputStreamOpen */

// AACompressionOutputStreamOpenExisting is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionOutputStreamOpenExisting
func AACompressionOutputStreamOpenExisting(compressed_stream unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AACompressionOutputStreamOpenExisting(compressed_stream, flags, n_threads)
}/* debug [functions.gen.go/function]: AACompressionOutputStreamOpenExisting */

// AAConvertArchiveOutputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAConvertArchiveOutputStreamOpen
func AAConvertArchiveOutputStreamOpen(stream unsafe.Pointer, insert_key_set unsafe.Pointer, remove_key_set unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AAConvertArchiveOutputStreamOpen(stream, insert_key_set, remove_key_set, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAConvertArchiveOutputStreamOpen */

// AACustomArchiveStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamOpen
func AACustomArchiveStreamOpen() unsafe.Pointer {
	return _AACustomArchiveStreamOpen()
}/* debug [functions.gen.go/function]: AACustomArchiveStreamOpen */

// AACustomArchiveStreamSetCancelProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetCancelProc
func AACustomArchiveStreamSetCancelProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetCancelProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetCancelProc */

// AACustomArchiveStreamSetCloseProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetCloseProc
func AACustomArchiveStreamSetCloseProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetCloseProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetCloseProc */

// AACustomArchiveStreamSetData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetData
func AACustomArchiveStreamSetData(s unsafe.Pointer, data unsafe.Pointer) {
	_AACustomArchiveStreamSetData(s, data)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetData */

// AACustomArchiveStreamSetReadBlobProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetReadBlobProc
func AACustomArchiveStreamSetReadBlobProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetReadBlobProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetReadBlobProc */

// AACustomArchiveStreamSetReadHeaderProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetReadHeaderProc
func AACustomArchiveStreamSetReadHeaderProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetReadHeaderProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetReadHeaderProc */

// AACustomArchiveStreamSetWriteBlobProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetWriteBlobProc
func AACustomArchiveStreamSetWriteBlobProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetWriteBlobProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetWriteBlobProc */

// AACustomArchiveStreamSetWriteHeaderProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetWriteHeaderProc
func AACustomArchiveStreamSetWriteHeaderProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetWriteHeaderProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomArchiveStreamSetWriteHeaderProc */

// AACustomByteStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamOpen
func AACustomByteStreamOpen() unsafe.Pointer {
	return _AACustomByteStreamOpen()
}/* debug [functions.gen.go/function]: AACustomByteStreamOpen */

// AACustomByteStreamSetCancelProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetCancelProc
func AACustomByteStreamSetCancelProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetCancelProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetCancelProc */

// AACustomByteStreamSetCloseProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetCloseProc
func AACustomByteStreamSetCloseProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetCloseProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetCloseProc */

// AACustomByteStreamSetData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetData
func AACustomByteStreamSetData(s unsafe.Pointer, data unsafe.Pointer) {
	_AACustomByteStreamSetData(s, data)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetData */

// AACustomByteStreamSetPReadProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetPReadProc
func AACustomByteStreamSetPReadProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetPReadProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetPReadProc */

// AACustomByteStreamSetPWriteProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetPWriteProc
func AACustomByteStreamSetPWriteProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetPWriteProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetPWriteProc */

// AACustomByteStreamSetReadProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetReadProc
func AACustomByteStreamSetReadProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetReadProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetReadProc */

// AACustomByteStreamSetSeekProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetSeekProc
func AACustomByteStreamSetSeekProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetSeekProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetSeekProc */

// AACustomByteStreamSetWriteProc is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomByteStreamSetWriteProc
func AACustomByteStreamSetWriteProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomByteStreamSetWriteProc(s, proc)
}/* debug [functions.gen.go/function]: AACustomByteStreamSetWriteProc */

// AADecodeArchiveInputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AADecodeArchiveInputStreamOpen
func AADecodeArchiveInputStreamOpen(stream unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AADecodeArchiveInputStreamOpen(stream, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AADecodeArchiveInputStreamOpen */

// AADecompressionInputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AADecompressionInputStreamOpen
func AADecompressionInputStreamOpen(compressed_stream unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AADecompressionInputStreamOpen(compressed_stream, flags, n_threads)
}/* debug [functions.gen.go/function]: AADecompressionInputStreamOpen */

// AADecompressionRandomAccessInputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AADecompressionRandomAccessInputStreamOpen
func AADecompressionRandomAccessInputStreamOpen(compressed_stream unsafe.Pointer, alloc_limit uintptr, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AADecompressionRandomAccessInputStreamOpen(compressed_stream, alloc_limit, flags, n_threads)
}/* debug [functions.gen.go/function]: AADecompressionRandomAccessInputStreamOpen */

// AAEncodeArchiveOutputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEncodeArchiveOutputStreamOpen
func AAEncodeArchiveOutputStreamOpen(stream unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AAEncodeArchiveOutputStreamOpen(stream, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAEncodeArchiveOutputStreamOpen */

// AAEntryACLBlobAppendEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobAppendEntry
func AAEntryACLBlobAppendEntry(acl unsafe.Pointer, ace unsafe.Pointer, qualifier_value unsafe.Pointer, qualifier_size uintptr) int {
	return _AAEntryACLBlobAppendEntry(acl, ace, qualifier_value, qualifier_size)
}/* debug [functions.gen.go/function]: AAEntryACLBlobAppendEntry */

// AAEntryACLBlobApplyToPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobApplyToPath
func AAEntryACLBlobApplyToPath(acl unsafe.Pointer, dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) int {
	return _AAEntryACLBlobApplyToPath(acl, dir, path, flags)
}/* debug [functions.gen.go/function]: AAEntryACLBlobApplyToPath */

// AAEntryACLBlobClear is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobClear
func AAEntryACLBlobClear(acl unsafe.Pointer) int {
	return _AAEntryACLBlobClear(acl)
}/* debug [functions.gen.go/function]: AAEntryACLBlobClear */

// AAEntryACLBlobCreate is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobCreate
func AAEntryACLBlobCreate() unsafe.Pointer {
	return _AAEntryACLBlobCreate()
}/* debug [functions.gen.go/function]: AAEntryACLBlobCreate */

// AAEntryACLBlobCreateWithEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobCreateWithEncodedData
func AAEntryACLBlobCreateWithEncodedData(data unsafe.Pointer, data_size uintptr) unsafe.Pointer {
	return _AAEntryACLBlobCreateWithEncodedData(data, data_size)
}/* debug [functions.gen.go/function]: AAEntryACLBlobCreateWithEncodedData */

// AAEntryACLBlobCreateWithPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobCreateWithPath
func AAEntryACLBlobCreateWithPath(dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AAEntryACLBlobCreateWithPath(dir, path, flags)
}/* debug [functions.gen.go/function]: AAEntryACLBlobCreateWithPath */

// AAEntryACLBlobDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobDestroy
func AAEntryACLBlobDestroy(acl unsafe.Pointer) {
	_AAEntryACLBlobDestroy(acl)
}/* debug [functions.gen.go/function]: AAEntryACLBlobDestroy */

// AAEntryACLBlobGetEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobGetEncodedData
func AAEntryACLBlobGetEncodedData(acl unsafe.Pointer) unsafe.Pointer {
	return _AAEntryACLBlobGetEncodedData(acl)
}/* debug [functions.gen.go/function]: AAEntryACLBlobGetEncodedData */

// AAEntryACLBlobGetEncodedSize is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobGetEncodedSize
func AAEntryACLBlobGetEncodedSize(acl unsafe.Pointer) uintptr {
	return _AAEntryACLBlobGetEncodedSize(acl)
}/* debug [functions.gen.go/function]: AAEntryACLBlobGetEncodedSize */

// AAEntryACLBlobGetEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobGetEntry
func AAEntryACLBlobGetEntry(acl unsafe.Pointer, i uint32, ace unsafe.Pointer, qualifier_capacity uintptr, qualifier_value unsafe.Pointer, qualifier_size unsafe.Pointer) int {
	return _AAEntryACLBlobGetEntry(acl, i, ace, qualifier_capacity, qualifier_value, qualifier_size)
}/* debug [functions.gen.go/function]: AAEntryACLBlobGetEntry */

// AAEntryACLBlobGetEntryCount is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobGetEntryCount
func AAEntryACLBlobGetEntryCount(acl unsafe.Pointer) uint32 {
	return _AAEntryACLBlobGetEntryCount(acl)
}/* debug [functions.gen.go/function]: AAEntryACLBlobGetEntryCount */

// AAEntryACLBlobRemoveEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobRemoveEntry
func AAEntryACLBlobRemoveEntry(acl unsafe.Pointer, i uint32) int {
	return _AAEntryACLBlobRemoveEntry(acl, i)
}/* debug [functions.gen.go/function]: AAEntryACLBlobRemoveEntry */

// AAEntryACLBlobSetEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobSetEntry
func AAEntryACLBlobSetEntry(acl unsafe.Pointer, i uint32, ace unsafe.Pointer, qualifier_value unsafe.Pointer, qualifier_size uintptr) int {
	return _AAEntryACLBlobSetEntry(acl, i, ace, qualifier_value, qualifier_size)
}/* debug [functions.gen.go/function]: AAEntryACLBlobSetEntry */

// AAEntryXATBlobAppendEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobAppendEntry
func AAEntryXATBlobAppendEntry(xat unsafe.Pointer, key unsafe.Pointer, data unsafe.Pointer, data_size uintptr) int {
	return _AAEntryXATBlobAppendEntry(xat, key, data, data_size)
}/* debug [functions.gen.go/function]: AAEntryXATBlobAppendEntry */

// AAEntryXATBlobApplyToPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobApplyToPath
func AAEntryXATBlobApplyToPath(xat unsafe.Pointer, dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) int {
	return _AAEntryXATBlobApplyToPath(xat, dir, path, flags)
}/* debug [functions.gen.go/function]: AAEntryXATBlobApplyToPath */

// AAEntryXATBlobClear is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobClear
func AAEntryXATBlobClear(xat unsafe.Pointer) int {
	return _AAEntryXATBlobClear(xat)
}/* debug [functions.gen.go/function]: AAEntryXATBlobClear */

// AAEntryXATBlobCreate is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobCreate
func AAEntryXATBlobCreate() unsafe.Pointer {
	return _AAEntryXATBlobCreate()
}/* debug [functions.gen.go/function]: AAEntryXATBlobCreate */

// AAEntryXATBlobCreateWithEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobCreateWithEncodedData
func AAEntryXATBlobCreateWithEncodedData(data unsafe.Pointer, data_size uintptr) unsafe.Pointer {
	return _AAEntryXATBlobCreateWithEncodedData(data, data_size)
}/* debug [functions.gen.go/function]: AAEntryXATBlobCreateWithEncodedData */

// AAEntryXATBlobCreateWithPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobCreateWithPath
func AAEntryXATBlobCreateWithPath(dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AAEntryXATBlobCreateWithPath(dir, path, flags)
}/* debug [functions.gen.go/function]: AAEntryXATBlobCreateWithPath */

// AAEntryXATBlobDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobDestroy
func AAEntryXATBlobDestroy(xat unsafe.Pointer) {
	_AAEntryXATBlobDestroy(xat)
}/* debug [functions.gen.go/function]: AAEntryXATBlobDestroy */

// AAEntryXATBlobGetEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobGetEncodedData
func AAEntryXATBlobGetEncodedData(xat unsafe.Pointer) unsafe.Pointer {
	return _AAEntryXATBlobGetEncodedData(xat)
}/* debug [functions.gen.go/function]: AAEntryXATBlobGetEncodedData */

// AAEntryXATBlobGetEncodedSize is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobGetEncodedSize
func AAEntryXATBlobGetEncodedSize(xat unsafe.Pointer) uintptr {
	return _AAEntryXATBlobGetEncodedSize(xat)
}/* debug [functions.gen.go/function]: AAEntryXATBlobGetEncodedSize */

// AAEntryXATBlobGetEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobGetEntry
func AAEntryXATBlobGetEntry(xat unsafe.Pointer, i uint32, key_capacity uintptr, key unsafe.Pointer, key_length unsafe.Pointer, data_capacity uintptr, data unsafe.Pointer, data_size unsafe.Pointer) int {
	return _AAEntryXATBlobGetEntry(xat, i, key_capacity, key, key_length, data_capacity, data, data_size)
}/* debug [functions.gen.go/function]: AAEntryXATBlobGetEntry */

// AAEntryXATBlobGetEntryCount is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobGetEntryCount
func AAEntryXATBlobGetEntryCount(xat unsafe.Pointer) uint32 {
	return _AAEntryXATBlobGetEntryCount(xat)
}/* debug [functions.gen.go/function]: AAEntryXATBlobGetEntryCount */

// AAEntryXATBlobRemoveEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobRemoveEntry
func AAEntryXATBlobRemoveEntry(xat unsafe.Pointer, i uint32) int {
	return _AAEntryXATBlobRemoveEntry(xat, i)
}/* debug [functions.gen.go/function]: AAEntryXATBlobRemoveEntry */

// AAEntryXATBlobSetEntry is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryXATBlobSetEntry
func AAEntryXATBlobSetEntry(xat unsafe.Pointer, i uint32, key unsafe.Pointer, data unsafe.Pointer, data_size uintptr) int {
	return _AAEntryXATBlobSetEntry(xat, i, key, data, data_size)
}/* debug [functions.gen.go/function]: AAEntryXATBlobSetEntry */

// AAExtractArchiveOutputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAExtractArchiveOutputStreamOpen
func AAExtractArchiveOutputStreamOpen(dir unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AAExtractArchiveOutputStreamOpen(dir, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAExtractArchiveOutputStreamOpen */

// AAFieldKeySetClear is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetClear
func AAFieldKeySetClear(key_set unsafe.Pointer) int {
	return _AAFieldKeySetClear(key_set)
}/* debug [functions.gen.go/function]: AAFieldKeySetClear */

// AAFieldKeySetClone is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetClone
func AAFieldKeySetClone(key_set unsafe.Pointer) unsafe.Pointer {
	return _AAFieldKeySetClone(key_set)
}/* debug [functions.gen.go/function]: AAFieldKeySetClone */

// AAFieldKeySetContainsKey is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetContainsKey
func AAFieldKeySetContainsKey(key_set unsafe.Pointer, key unsafe.Pointer) int {
	return _AAFieldKeySetContainsKey(key_set, key)
}/* debug [functions.gen.go/function]: AAFieldKeySetContainsKey */

// AAFieldKeySetCreate is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetCreate
func AAFieldKeySetCreate() unsafe.Pointer {
	return _AAFieldKeySetCreate()
}/* debug [functions.gen.go/function]: AAFieldKeySetCreate */

// AAFieldKeySetCreateWithString is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetCreateWithString
func AAFieldKeySetCreateWithString(s unsafe.Pointer) unsafe.Pointer {
	return _AAFieldKeySetCreateWithString(s)
}/* debug [functions.gen.go/function]: AAFieldKeySetCreateWithString */

// AAFieldKeySetDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetDestroy
func AAFieldKeySetDestroy(key_set unsafe.Pointer) {
	_AAFieldKeySetDestroy(key_set)
}/* debug [functions.gen.go/function]: AAFieldKeySetDestroy */

// AAFieldKeySetGetKey is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetGetKey
func AAFieldKeySetGetKey(key_set unsafe.Pointer, i uint32) unsafe.Pointer {
	return _AAFieldKeySetGetKey(key_set, i)
}/* debug [functions.gen.go/function]: AAFieldKeySetGetKey */

// AAFieldKeySetGetKeyCount is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetGetKeyCount
func AAFieldKeySetGetKeyCount(key_set unsafe.Pointer) uint32 {
	return _AAFieldKeySetGetKeyCount(key_set)
}/* debug [functions.gen.go/function]: AAFieldKeySetGetKeyCount */

// AAFieldKeySetInsertKey is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetInsertKey
func AAFieldKeySetInsertKey(key_set unsafe.Pointer, key unsafe.Pointer) int {
	return _AAFieldKeySetInsertKey(key_set, key)
}/* debug [functions.gen.go/function]: AAFieldKeySetInsertKey */

// AAFieldKeySetInsertKeySet is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetInsertKeySet
func AAFieldKeySetInsertKeySet(key_set unsafe.Pointer, s unsafe.Pointer) int {
	return _AAFieldKeySetInsertKeySet(key_set, s)
}/* debug [functions.gen.go/function]: AAFieldKeySetInsertKeySet */

// AAFieldKeySetRemoveKey is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetRemoveKey
func AAFieldKeySetRemoveKey(key_set unsafe.Pointer, key unsafe.Pointer) int {
	return _AAFieldKeySetRemoveKey(key_set, key)
}/* debug [functions.gen.go/function]: AAFieldKeySetRemoveKey */

// AAFieldKeySetRemoveKeySet is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetRemoveKeySet
func AAFieldKeySetRemoveKeySet(key_set unsafe.Pointer, s unsafe.Pointer) int {
	return _AAFieldKeySetRemoveKeySet(key_set, s)
}/* debug [functions.gen.go/function]: AAFieldKeySetRemoveKeySet */

// AAFieldKeySetSelectKeySet is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetSelectKeySet
func AAFieldKeySetSelectKeySet(key_set unsafe.Pointer, s unsafe.Pointer) int {
	return _AAFieldKeySetSelectKeySet(key_set, s)
}/* debug [functions.gen.go/function]: AAFieldKeySetSelectKeySet */

// AAFieldKeySetSerialize is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldKeySetSerialize
func AAFieldKeySetSerialize(key_set unsafe.Pointer, capacity uintptr, s unsafe.Pointer) int {
	return _AAFieldKeySetSerialize(key_set, capacity, s)
}/* debug [functions.gen.go/function]: AAFieldKeySetSerialize */

// AAFileStreamOpenWithFD is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFileStreamOpenWithFD
func AAFileStreamOpenWithFD(fd int, automatic_close int) unsafe.Pointer {
	return _AAFileStreamOpenWithFD(fd, automatic_close)
}/* debug [functions.gen.go/function]: AAFileStreamOpenWithFD */

// AAFileStreamOpenWithPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFileStreamOpenWithPath
func AAFileStreamOpenWithPath(path unsafe.Pointer, open_flags int, open_mode unsafe.Pointer) unsafe.Pointer {
	return _AAFileStreamOpenWithPath(path, open_flags, open_mode)
}/* debug [functions.gen.go/function]: AAFileStreamOpenWithPath */

// AAHeaderAssign is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderAssign
func AAHeaderAssign(header unsafe.Pointer, from_header unsafe.Pointer) int {
	return _AAHeaderAssign(header, from_header)
}/* debug [functions.gen.go/function]: AAHeaderAssign */

// AAHeaderClear is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderClear
func AAHeaderClear(header unsafe.Pointer) int {
	return _AAHeaderClear(header)
}/* debug [functions.gen.go/function]: AAHeaderClear */

// AAHeaderClone is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderClone
func AAHeaderClone(header unsafe.Pointer) unsafe.Pointer {
	return _AAHeaderClone(header)
}/* debug [functions.gen.go/function]: AAHeaderClone */

// AAHeaderCreate is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderCreate
func AAHeaderCreate() unsafe.Pointer {
	return _AAHeaderCreate()
}/* debug [functions.gen.go/function]: AAHeaderCreate */

// AAHeaderCreateWithEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderCreateWithEncodedData
func AAHeaderCreateWithEncodedData(data_size uintptr, data unsafe.Pointer) unsafe.Pointer {
	return _AAHeaderCreateWithEncodedData(data_size, data)
}/* debug [functions.gen.go/function]: AAHeaderCreateWithEncodedData */

// AAHeaderCreateWithPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderCreateWithPath
func AAHeaderCreateWithPath(key_set unsafe.Pointer, dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AAHeaderCreateWithPath(key_set, dir, path, flags)
}/* debug [functions.gen.go/function]: AAHeaderCreateWithPath */

// AAHeaderDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderDestroy
func AAHeaderDestroy(header unsafe.Pointer) {
	_AAHeaderDestroy(header)
}/* debug [functions.gen.go/function]: AAHeaderDestroy */

// AAHeaderGetEncodedData is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetEncodedData
func AAHeaderGetEncodedData(header unsafe.Pointer) unsafe.Pointer {
	return _AAHeaderGetEncodedData(header)
}/* debug [functions.gen.go/function]: AAHeaderGetEncodedData */

// AAHeaderGetEncodedSize is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetEncodedSize
func AAHeaderGetEncodedSize(header unsafe.Pointer) uintptr {
	return _AAHeaderGetEncodedSize(header)
}/* debug [functions.gen.go/function]: AAHeaderGetEncodedSize */

// AAHeaderGetFieldBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldBlob
func AAHeaderGetFieldBlob(header unsafe.Pointer, i uint32, size []uint64, offset []uint64) int {
	return _AAHeaderGetFieldBlob(header, i, size, offset)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldBlob */

// AAHeaderGetFieldCount is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldCount
func AAHeaderGetFieldCount(header unsafe.Pointer) uint32 {
	return _AAHeaderGetFieldCount(header)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldCount */

// AAHeaderGetFieldHash is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldHash
func AAHeaderGetFieldHash(header unsafe.Pointer, i uint32, capacity uintptr, hash_function unsafe.Pointer, value unsafe.Pointer) int {
	return _AAHeaderGetFieldHash(header, i, capacity, hash_function, value)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldHash */

// AAHeaderGetFieldKey is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldKey
func AAHeaderGetFieldKey(header unsafe.Pointer, i uint32) unsafe.Pointer {
	return _AAHeaderGetFieldKey(header, i)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldKey */

// AAHeaderGetFieldString is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldString
func AAHeaderGetFieldString(header unsafe.Pointer, i uint32, capacity uintptr, value unsafe.Pointer, length unsafe.Pointer) int {
	return _AAHeaderGetFieldString(header, i, capacity, value, length)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldString */

// AAHeaderGetFieldTimespec is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldTimespec
func AAHeaderGetFieldTimespec(header unsafe.Pointer, i uint32, value unsafe.Pointer) int {
	return _AAHeaderGetFieldTimespec(header, i, value)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldTimespec */

// AAHeaderGetFieldType is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldType
func AAHeaderGetFieldType(header unsafe.Pointer, i uint32) int {
	return _AAHeaderGetFieldType(header, i)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldType */

// AAHeaderGetFieldUInt is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetFieldUInt
func AAHeaderGetFieldUInt(header unsafe.Pointer, i uint32, value []uint64) int {
	return _AAHeaderGetFieldUInt(header, i, value)
}/* debug [functions.gen.go/function]: AAHeaderGetFieldUInt */

// AAHeaderGetKeyIndex is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetKeyIndex
func AAHeaderGetKeyIndex(header unsafe.Pointer, key unsafe.Pointer) int {
	return _AAHeaderGetKeyIndex(header, key)
}/* debug [functions.gen.go/function]: AAHeaderGetKeyIndex */

// AAHeaderGetPayloadSize is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderGetPayloadSize
func AAHeaderGetPayloadSize(header unsafe.Pointer) uint64 {
	return _AAHeaderGetPayloadSize(header)
}/* debug [functions.gen.go/function]: AAHeaderGetPayloadSize */

// AAHeaderRemoveField is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderRemoveField
func AAHeaderRemoveField(header unsafe.Pointer, i uint32) int {
	return _AAHeaderRemoveField(header, i)
}/* debug [functions.gen.go/function]: AAHeaderRemoveField */

// AAHeaderSetFieldBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldBlob
func AAHeaderSetFieldBlob(header unsafe.Pointer, i uint32, key unsafe.Pointer, size uint64) int {
	return _AAHeaderSetFieldBlob(header, i, key, size)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldBlob */

// AAHeaderSetFieldFlag is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldFlag
func AAHeaderSetFieldFlag(header unsafe.Pointer, i uint32, key unsafe.Pointer) int {
	return _AAHeaderSetFieldFlag(header, i, key)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldFlag */

// AAHeaderSetFieldHash is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldHash
func AAHeaderSetFieldHash(header unsafe.Pointer, i uint32, key unsafe.Pointer, hash_function unsafe.Pointer, value unsafe.Pointer) int {
	return _AAHeaderSetFieldHash(header, i, key, hash_function, value)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldHash */

// AAHeaderSetFieldString is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldString
func AAHeaderSetFieldString(header unsafe.Pointer, i uint32, key unsafe.Pointer, value unsafe.Pointer, length uintptr) int {
	return _AAHeaderSetFieldString(header, i, key, value, length)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldString */

// AAHeaderSetFieldTimespec is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldTimespec
func AAHeaderSetFieldTimespec(header unsafe.Pointer, i uint32, key unsafe.Pointer, value unsafe.Pointer) int {
	return _AAHeaderSetFieldTimespec(header, i, key, value)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldTimespec */

// AAHeaderSetFieldUInt is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldUInt
func AAHeaderSetFieldUInt(header unsafe.Pointer, i uint32, key unsafe.Pointer, value uint64) int {
	return _AAHeaderSetFieldUInt(header, i, key, value)
}/* debug [functions.gen.go/function]: AAHeaderSetFieldUInt */

// AAPathListCreateWithDirectoryContents is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListCreateWithDirectoryContents
func AAPathListCreateWithDirectoryContents(dir unsafe.Pointer, path unsafe.Pointer, msg_data unsafe.Pointer, msg_proc unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AAPathListCreateWithDirectoryContents(dir, path, msg_data, msg_proc, flags, n_threads)
}/* debug [functions.gen.go/function]: AAPathListCreateWithDirectoryContents */

// AAPathListCreateWithPath is a AppleArchive function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListCreateWithPath
func AAPathListCreateWithPath(dir unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _AAPathListCreateWithPath(dir, path)
}/* debug [functions.gen.go/function]: AAPathListCreateWithPath */

// AAPathListDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListDestroy
func AAPathListDestroy(path_list unsafe.Pointer) {
	_AAPathListDestroy(path_list)
}/* debug [functions.gen.go/function]: AAPathListDestroy */

// AAPathListNodeFirst is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListNodeFirst
func AAPathListNodeFirst(path_list unsafe.Pointer) uint64 {
	return _AAPathListNodeFirst(path_list)
}/* debug [functions.gen.go/function]: AAPathListNodeFirst */

// AAPathListNodeGetPath is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListNodeGetPath
func AAPathListNodeGetPath(path_list unsafe.Pointer, node uint64, path_capacity uintptr, path unsafe.Pointer, path_length unsafe.Pointer) int {
	return _AAPathListNodeGetPath(path_list, node, path_capacity, path, path_length)
}/* debug [functions.gen.go/function]: AAPathListNodeGetPath */

// AAPathListNodeNext is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAPathListNodeNext
func AAPathListNodeNext(path_list unsafe.Pointer, node uint64) uint64 {
	return _AAPathListNodeNext(path_list, node)
}/* debug [functions.gen.go/function]: AAPathListNodeNext */

// AARandomAccessByteStreamProcess is a AppleArchive function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AARandomAccessByteStreamProcess
func AARandomAccessByteStreamProcess(istream unsafe.Pointer, ostream unsafe.Pointer, max_offset unsafe.Pointer, block_size uintptr, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AARandomAccessByteStreamProcess(istream, ostream, max_offset, block_size, flags, n_threads)
}/* debug [functions.gen.go/function]: AARandomAccessByteStreamProcess */

// AASharedBufferPipeOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AASharedBufferPipeOpen
func AASharedBufferPipeOpen(ostream unsafe.Pointer, istream unsafe.Pointer, buffer_capacity uintptr) int {
	return _AASharedBufferPipeOpen(ostream, istream, buffer_capacity)
}/* debug [functions.gen.go/function]: AASharedBufferPipeOpen */

// AATempFileStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AATempFileStreamOpen
func AATempFileStreamOpen() unsafe.Pointer {
	return _AATempFileStreamOpen()
}/* debug [functions.gen.go/function]: AATempFileStreamOpen */

// AEAAuthDataAppendEntry is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataAppendEntry
func AEAAuthDataAppendEntry(auth_data unsafe.Pointer, key unsafe.Pointer, data unsafe.Pointer, data_size uintptr) int {
	return _AEAAuthDataAppendEntry(auth_data, key, data, data_size)
}/* debug [functions.gen.go/function]: AEAAuthDataAppendEntry */

// AEAAuthDataClear is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataClear
func AEAAuthDataClear(auth_data unsafe.Pointer) int {
	return _AEAAuthDataClear(auth_data)
}/* debug [functions.gen.go/function]: AEAAuthDataClear */

// AEAAuthDataCreate is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataCreate
func AEAAuthDataCreate() unsafe.Pointer {
	return _AEAAuthDataCreate()
}/* debug [functions.gen.go/function]: AEAAuthDataCreate */

// AEAAuthDataCreateWithContext is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataCreateWithContext
func AEAAuthDataCreateWithContext(context unsafe.Pointer) unsafe.Pointer {
	return _AEAAuthDataCreateWithContext(context)
}/* debug [functions.gen.go/function]: AEAAuthDataCreateWithContext */

// AEAAuthDataDestroy is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataDestroy
func AEAAuthDataDestroy(auth_data unsafe.Pointer) {
	_AEAAuthDataDestroy(auth_data)
}/* debug [functions.gen.go/function]: AEAAuthDataDestroy */

// AEAAuthDataGetEncodedData is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataGetEncodedData
func AEAAuthDataGetEncodedData(auth_data unsafe.Pointer) unsafe.Pointer {
	return _AEAAuthDataGetEncodedData(auth_data)
}/* debug [functions.gen.go/function]: AEAAuthDataGetEncodedData */

// AEAAuthDataGetEncodedSize is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataGetEncodedSize
func AEAAuthDataGetEncodedSize(auth_data unsafe.Pointer) uintptr {
	return _AEAAuthDataGetEncodedSize(auth_data)
}/* debug [functions.gen.go/function]: AEAAuthDataGetEncodedSize */

// AEAAuthDataGetEntry is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataGetEntry
func AEAAuthDataGetEntry(auth_data unsafe.Pointer, i uint32, key_capacity uintptr, key unsafe.Pointer, key_length unsafe.Pointer, data_capacity uintptr, data unsafe.Pointer, data_size unsafe.Pointer) int {
	return _AEAAuthDataGetEntry(auth_data, i, key_capacity, key, key_length, data_capacity, data, data_size)
}/* debug [functions.gen.go/function]: AEAAuthDataGetEntry */

// AEAAuthDataGetEntryCount is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataGetEntryCount
func AEAAuthDataGetEntryCount(auth_data unsafe.Pointer) uint32 {
	return _AEAAuthDataGetEntryCount(auth_data)
}/* debug [functions.gen.go/function]: AEAAuthDataGetEntryCount */

// AEAAuthDataRemoveEntry is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataRemoveEntry
func AEAAuthDataRemoveEntry(auth_data unsafe.Pointer, i uint32) int {
	return _AEAAuthDataRemoveEntry(auth_data, i)
}/* debug [functions.gen.go/function]: AEAAuthDataRemoveEntry */

// AEAAuthDataSetEntry is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAAuthDataSetEntry
func AEAAuthDataSetEntry(auth_data unsafe.Pointer, i uint32, key unsafe.Pointer, data unsafe.Pointer, data_size uintptr) int {
	return _AEAAuthDataSetEntry(auth_data, i, key, data, data_size)
}/* debug [functions.gen.go/function]: AEAAuthDataSetEntry */

// AEAContextCreateWithEncryptedStream is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextCreateWithEncryptedStream
func AEAContextCreateWithEncryptedStream(encrypted_stream unsafe.Pointer) unsafe.Pointer {
	return _AEAContextCreateWithEncryptedStream(encrypted_stream)
}/* debug [functions.gen.go/function]: AEAContextCreateWithEncryptedStream */

// AEAContextCreateWithProfile is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextCreateWithProfile
func AEAContextCreateWithProfile(profile unsafe.Pointer) unsafe.Pointer {
	return _AEAContextCreateWithProfile(profile)
}/* debug [functions.gen.go/function]: AEAContextCreateWithProfile */

// AEAContextDecryptAttributes is a AppleArchive function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextDecryptAttributes
func AEAContextDecryptAttributes(context unsafe.Pointer) int {
	return _AEAContextDecryptAttributes(context)
}/* debug [functions.gen.go/function]: AEAContextDecryptAttributes */

// AEAContextDestroy is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextDestroy
func AEAContextDestroy(context unsafe.Pointer) {
	_AEAContextDestroy(context)
}/* debug [functions.gen.go/function]: AEAContextDestroy */

// AEAContextGenerateFieldBlob is a AppleArchive function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextGenerateFieldBlob
func AEAContextGenerateFieldBlob(context unsafe.Pointer, field unsafe.Pointer) int {
	return _AEAContextGenerateFieldBlob(context, field)
}/* debug [functions.gen.go/function]: AEAContextGenerateFieldBlob */

// AEAContextGetFieldBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextGetFieldBlob
func AEAContextGetFieldBlob(context unsafe.Pointer, field unsafe.Pointer, representation unsafe.Pointer, buf_capacity uintptr, buf unsafe.Pointer, buf_size unsafe.Pointer) int {
	return _AEAContextGetFieldBlob(context, field, representation, buf_capacity, buf, buf_size)
}/* debug [functions.gen.go/function]: AEAContextGetFieldBlob */

// AEAContextGetFieldUInt is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextGetFieldUInt
func AEAContextGetFieldUInt(context unsafe.Pointer, field unsafe.Pointer) uint64 {
	return _AEAContextGetFieldUInt(context, field)
}/* debug [functions.gen.go/function]: AEAContextGetFieldUInt */

// AEAContextSetFieldBlob is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextSetFieldBlob
func AEAContextSetFieldBlob(context unsafe.Pointer, field unsafe.Pointer, representation unsafe.Pointer, buf unsafe.Pointer, buf_size uintptr) int {
	return _AEAContextSetFieldBlob(context, field, representation, buf, buf_size)
}/* debug [functions.gen.go/function]: AEAContextSetFieldBlob */

// AEAContextSetFieldUInt is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextSetFieldUInt
func AEAContextSetFieldUInt(context unsafe.Pointer, field unsafe.Pointer, value uint64) int {
	return _AEAContextSetFieldUInt(context, field, value)
}/* debug [functions.gen.go/function]: AEAContextSetFieldUInt */

// AEADecryptionInputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEADecryptionInputStreamOpen
func AEADecryptionInputStreamOpen(encrypted_stream unsafe.Pointer, context unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AEADecryptionInputStreamOpen(encrypted_stream, context, flags, n_threads)
}/* debug [functions.gen.go/function]: AEADecryptionInputStreamOpen */

// AEADecryptionRandomAccessInputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEADecryptionRandomAccessInputStreamOpen
func AEADecryptionRandomAccessInputStreamOpen(encrypted_stream unsafe.Pointer, context unsafe.Pointer, alloc_limit uintptr, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AEADecryptionRandomAccessInputStreamOpen(encrypted_stream, context, alloc_limit, flags, n_threads)
}/* debug [functions.gen.go/function]: AEADecryptionRandomAccessInputStreamOpen */

// AEAEncryptionOutputStreamCloseAndUpdateContext is a AppleArchive function.
//
// Added in macOS 11.3.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAEncryptionOutputStreamCloseAndUpdateContext
func AEAEncryptionOutputStreamCloseAndUpdateContext(stream unsafe.Pointer, context unsafe.Pointer) int {
	return _AEAEncryptionOutputStreamCloseAndUpdateContext(stream, context)
}/* debug [functions.gen.go/function]: AEAEncryptionOutputStreamCloseAndUpdateContext */

// AEAEncryptionOutputStreamOpen is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAEncryptionOutputStreamOpen
func AEAEncryptionOutputStreamOpen(encrypted_stream unsafe.Pointer, context unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AEAEncryptionOutputStreamOpen(encrypted_stream, context, flags, n_threads)
}/* debug [functions.gen.go/function]: AEAEncryptionOutputStreamOpen */

// AEAEncryptionOutputStreamOpenExisting is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAEncryptionOutputStreamOpenExisting
func AEAEncryptionOutputStreamOpenExisting(encrypted_stream unsafe.Pointer, context unsafe.Pointer, flags unsafe.Pointer, n_threads int) unsafe.Pointer {
	return _AEAEncryptionOutputStreamOpenExisting(encrypted_stream, context, flags, n_threads)
}/* debug [functions.gen.go/function]: AEAEncryptionOutputStreamOpenExisting */

// AEAStreamSign is a AppleArchive function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAStreamSign
func AEAStreamSign(encrypted_stream unsafe.Pointer, context unsafe.Pointer) int {
	return _AEAStreamSign(encrypted_stream, context)
}/* debug [functions.gen.go/function]: AEAStreamSign */




