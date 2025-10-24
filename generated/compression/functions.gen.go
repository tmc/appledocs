// Code generated from Apple documentation for Compression. DO NOT EDIT.

package compression

/* debug [functions.gen.go]: Generating 7 functions for Compression */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Compression Functions (7 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_compression_decode_buffer func(unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) uintptr
	_compression_decode_scratch_buffer_size func(unsafe.Pointer) uintptr
	_compression_encode_buffer func(unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) uintptr
	_compression_encode_scratch_buffer_size func(unsafe.Pointer) uintptr
	_compression_stream_destroy func(unsafe.Pointer) unsafe.Pointer
	_compression_stream_init func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_compression_stream_process func(unsafe.Pointer, int) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_compression_decode_buffer, lib, "compression_decode_buffer")
	tryRegister(&_compression_decode_scratch_buffer_size, lib, "compression_decode_scratch_buffer_size")
	tryRegister(&_compression_encode_buffer, lib, "compression_encode_buffer")
	tryRegister(&_compression_encode_scratch_buffer_size, lib, "compression_encode_scratch_buffer_size")
	tryRegister(&_compression_stream_destroy, lib, "compression_stream_destroy")
	tryRegister(&_compression_stream_init, lib, "compression_stream_init")
	tryRegister(&_compression_stream_process, lib, "compression_stream_process")
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



// Decompresses the contents of a source buffer into a destination buffer.
//
// Added in macOS 10.11.
// Decompresses the contents of a source buffer into a destination buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_decode_buffer(_:_:_:_:_:_:)
func compression_decode_buffer(dst_buffer unsafe.Pointer, dst_size uintptr, src_buffer unsafe.Pointer, src_size uintptr, scratch_buffer unsafe.Pointer, algorithm unsafe.Pointer) uintptr {
	return _compression_decode_buffer(dst_buffer, dst_size, src_buffer, src_size, scratch_buffer, algorithm)
}/* debug [functions.gen.go/function]: compression_decode_buffer */

// Returns the required decompression scratch buffer size for the selected algorithm.
//
// Added in macOS 10.11.
// Returns the required decompression scratch buffer size for the selected algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_decode_scratch_buffer_size(_:)
func compression_decode_scratch_buffer_size(algorithm unsafe.Pointer) uintptr {
	return _compression_decode_scratch_buffer_size(algorithm)
}/* debug [functions.gen.go/function]: compression_decode_scratch_buffer_size */

// Compresses the contents of a source buffer into a destination buffer.
//
// Added in macOS 10.11.
// Compresses the contents of a source buffer into a destination buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_encode_buffer(_:_:_:_:_:_:)
func compression_encode_buffer(dst_buffer unsafe.Pointer, dst_size uintptr, src_buffer unsafe.Pointer, src_size uintptr, scratch_buffer unsafe.Pointer, algorithm unsafe.Pointer) uintptr {
	return _compression_encode_buffer(dst_buffer, dst_size, src_buffer, src_size, scratch_buffer, algorithm)
}/* debug [functions.gen.go/function]: compression_encode_buffer */

// Returns the required compression scratch buffer size for the selected algorithm.
//
// Added in macOS 10.11.
// Returns the required compression scratch buffer size for the selected algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_encode_scratch_buffer_size(_:)
func compression_encode_scratch_buffer_size(algorithm unsafe.Pointer) uintptr {
	return _compression_encode_scratch_buffer_size(algorithm)
}/* debug [functions.gen.go/function]: compression_encode_scratch_buffer_size */

// Frees any memory allocated by stream initialization function.
//
// Added in macOS 10.11.
// Frees any memory allocated by stream initialization function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_stream_destroy(_:)
func compression_stream_destroy(stream unsafe.Pointer) unsafe.Pointer {
	return _compression_stream_destroy(stream)
}/* debug [functions.gen.go/function]: compression_stream_destroy */

// Initializes a compression stream for either compression or decompression.
//
// Added in macOS 10.11.
// Initializes a compression stream for either compression or decompression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_stream_init(_:_:_:)
func compression_stream_init(stream unsafe.Pointer, operation unsafe.Pointer, algorithm unsafe.Pointer) unsafe.Pointer {
	return _compression_stream_init(stream, operation, algorithm)
}/* debug [functions.gen.go/function]: compression_stream_init */

// Performs compression or decompression using an initialized compression stream structure.
//
// Added in macOS 10.11.
// Performs compression or decompression using an initialized compression stream structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Compression/compression_stream_process(_:_:)
func compression_stream_process(stream unsafe.Pointer, flags int) unsafe.Pointer {
	return _compression_stream_process(stream, flags)
}/* debug [functions.gen.go/function]: compression_stream_process */




