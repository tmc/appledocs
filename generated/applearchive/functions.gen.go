// Code generated from Apple documentation for AppleArchive. DO NOT EDIT.

package applearchive

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// AppleArchive Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AACustomArchiveStreamSetWriteBlobProc func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAEntryACLBlobCreateWithPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AAHeaderSetFieldUInt func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AEAContextCreateWithEncryptedStream func(unsafe.Pointer) unsafe.Pointer
	_AEAContextGenerateFieldBlob func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AEAContextGetFieldBlob func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AACustomArchiveStreamSetWriteBlobProc, lib, "AACustomArchiveStreamSetWriteBlobProc")
	tryRegister(&_AAEntryACLBlobCreateWithPath, lib, "AAEntryACLBlobCreateWithPath")
	tryRegister(&_AAHeaderSetFieldUInt, lib, "AAHeaderSetFieldUInt")
	tryRegister(&_AEAContextCreateWithEncryptedStream, lib, "AEAContextCreateWithEncryptedStream")
	tryRegister(&_AEAContextGenerateFieldBlob, lib, "AEAContextGenerateFieldBlob")
	tryRegister(&_AEAContextGetFieldBlob, lib, "AEAContextGetFieldBlob")
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



// AACustomArchiveStreamSetWriteBlobProc is a AppleArchive function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACustomArchiveStreamSetWriteBlobProc
func AACustomArchiveStreamSetWriteBlobProc(s unsafe.Pointer, proc unsafe.Pointer) {
	_AACustomArchiveStreamSetWriteBlobProc(s, proc)
	}


// AAEntryACLBlobCreateWithPath is a AppleArchive function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryACLBlobCreateWithPath
func AAEntryACLBlobCreateWithPath(dir unsafe.Pointer, path unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AAEntryACLBlobCreateWithPath(dir, path, flags)
	}


// AAHeaderSetFieldUInt is a AppleArchive function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHeaderSetFieldUInt
func AAHeaderSetFieldUInt(header unsafe.Pointer, i unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _AAHeaderSetFieldUInt(header, i, key, value)
	}


// AEAContextCreateWithEncryptedStream is a AppleArchive function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextCreateWithEncryptedStream
func AEAContextCreateWithEncryptedStream(encrypted_stream unsafe.Pointer) unsafe.Pointer {
	return _AEAContextCreateWithEncryptedStream(encrypted_stream)
	}


// AEAContextGenerateFieldBlob is a AppleArchive function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextGenerateFieldBlob
func AEAContextGenerateFieldBlob(context unsafe.Pointer, field unsafe.Pointer) unsafe.Pointer {
	return _AEAContextGenerateFieldBlob(context, field)
	}


// AEAContextGetFieldBlob is a AppleArchive function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextGetFieldBlob
func AEAContextGetFieldBlob(context unsafe.Pointer, field unsafe.Pointer, representation unsafe.Pointer, buf_capacity unsafe.Pointer, buf unsafe.Pointer, buf_size unsafe.Pointer) unsafe.Pointer {
	return _AEAContextGetFieldBlob(context, field, representation, buf_capacity, buf, buf_size)
	}




