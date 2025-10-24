// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders


import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MetalPerformanceShaders Functions (19 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MPSSupportsMTLDevice func(unsafe.Pointer,   MTLDevice) unsafe.Pointer
	_MPSImageBatchIncrementReadCount func(unsafe.Pointer,   [ MPSImage ] _,   Int) unsafe.Pointer
	_MPSStateBatchIncrementReadCount func(unsafe.Pointer,   [ MPSState ]? _,   Int) unsafe.Pointer
	_MPSStateBatchSynchronize func(unsafe.Pointer,   [ MPSState ] _,   any  MTLCommandBuffer) unsafe.Pointer
	_MPSImageBatchSynchronize func(unsafe.Pointer,   [ MPSImage ] _,   any  MTLCommandBuffer) unsafe.Pointer
	_MPSFindIntegerDivisionParams func(unsafe.Pointer,   UInt16) unsafe.Pointer
	_MPSImageBatchResourceSize func(unsafe.Pointer,   [ MPSImage ]) unsafe.Pointer
	_MPSStateBatchResourceSize func(unsafe.Pointer,   [ MPSState ]?) unsafe.Pointer
	_MPSGetCustomKernelBatchedDestinationIndex func(unsafe.Pointer,   MPSCustomKernelArgumentCount) unsafe.Pointer
	_MPSGetCustomKernelBatchedSourceIndex func(unsafe.Pointer,   MPSCustomKernelArgumentCount _,   UInt _,   UInt) unsafe.Pointer
	_MPSGetCustomKernelBroadcastSourceIndex func(unsafe.Pointer,   MPSCustomKernelArgumentCount _,   UInt _,   UInt) unsafe.Pointer
	_MPSGetCustomKernelMaxBatchSize func(unsafe.Pointer,   MPSCustomKernelArgumentCount _,   UInt) unsafe.Pointer
	_MPSHintTemporaryMemoryHighWaterMark func(unsafe.Pointer,   any  MTLCommandBuffer _,   Int) unsafe.Pointer
	_MPSSetHeapCacheDuration func(unsafe.Pointer,   any  MTLCommandBuffer _,   Double) unsafe.Pointer
	_MPSImageBatchIterate func(unsafe.Pointer,   [ MPSImage ] _,   @escaping MPSImage Int) unsafe.Pointer
	_MPSGetPreferredDevice func(unsafe.Pointer,   MPSDeviceOptions) unsafe.Pointer
	_MPSGetImageType func(unsafe.Pointer,   MPSImage) unsafe.Pointer
	_MPSSizeofMPSDataType func(unsafe.Pointer,   MPSDataType) unsafe.Pointer
	_MPSDataTypeBitsCount func(unsafe.Pointer,   MPSDataType) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MPSSupportsMTLDevice, lib, "MPSSupportsMTLDevice")
	tryRegister(&_MPSImageBatchIncrementReadCount, lib, "MPSImageBatchIncrementReadCount")
	tryRegister(&_MPSStateBatchIncrementReadCount, lib, "MPSStateBatchIncrementReadCount")
	tryRegister(&_MPSStateBatchSynchronize, lib, "MPSStateBatchSynchronize")
	tryRegister(&_MPSImageBatchSynchronize, lib, "MPSImageBatchSynchronize")
	tryRegister(&_MPSFindIntegerDivisionParams, lib, "MPSFindIntegerDivisionParams")
	tryRegister(&_MPSImageBatchResourceSize, lib, "MPSImageBatchResourceSize")
	tryRegister(&_MPSStateBatchResourceSize, lib, "MPSStateBatchResourceSize")
	tryRegister(&_MPSGetCustomKernelBatchedDestinationIndex, lib, "MPSGetCustomKernelBatchedDestinationIndex")
	tryRegister(&_MPSGetCustomKernelBatchedSourceIndex, lib, "MPSGetCustomKernelBatchedSourceIndex")
	tryRegister(&_MPSGetCustomKernelBroadcastSourceIndex, lib, "MPSGetCustomKernelBroadcastSourceIndex")
	tryRegister(&_MPSGetCustomKernelMaxBatchSize, lib, "MPSGetCustomKernelMaxBatchSize")
	tryRegister(&_MPSHintTemporaryMemoryHighWaterMark, lib, "MPSHintTemporaryMemoryHighWaterMark")
	tryRegister(&_MPSSetHeapCacheDuration, lib, "MPSSetHeapCacheDuration")
	tryRegister(&_MPSImageBatchIterate, lib, "MPSImageBatchIterate")
	tryRegister(&_MPSGetPreferredDevice, lib, "MPSGetPreferredDevice")
	tryRegister(&_MPSGetImageType, lib, "MPSGetImageType")
	tryRegister(&_MPSSizeofMPSDataType, lib, "MPSSizeofMPSDataType")
	tryRegister(&_MPSDataTypeBitsCount, lib, "MPSDataTypeBitsCount")
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



// Determines whether the Metal Performance Shaders framework supports a Metal device.
//
// Added in macOS 10.13.
// Determines whether the Metal Performance Shaders framework supports a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/1618849-mpssupportsmtldevice
func MPSSupportsMTLDevice(device unsafe.Pointer, p1   MTLDevice) unsafe.Pointer {
	return _MPSSupportsMTLDevice(device, p1)
}

// Increments or decrements the read count of an image batch by a specified amount.
//
// Added in macOS 10.13.4.
// Increments or decrements the read count of an image batch by a specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2951916-mpsimagebatchincrementreadcount
func MPSImageBatchIncrementReadCount(batch unsafe.Pointer, amount   [ MPSImage ] _, p2   Int) unsafe.Pointer {
	return _MPSImageBatchIncrementReadCount(batch, amount, p2)
}

// Increments or decrements the read count of a state batch by a specified amount.
//
// Added in macOS 10.13.4.
// Increments or decrements the read count of a state batch by a specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2951920-mpsstatebatchincrementreadcount
func MPSStateBatchIncrementReadCount(batch unsafe.Pointer, amount   [ MPSState ]? _, p2   Int) unsafe.Pointer {
	return _MPSStateBatchIncrementReadCount(batch, amount, p2)
}

// Removes any copy of the specified state batch from the device's caches, and, if needed, invalidates any CPU caches.
//
// Added in macOS 10.13.4.
// Removes any copy of the specified state batch from the device's caches, and, if needed, invalidates any CPU caches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2953920-mpsstatebatchsynchronize
func MPSStateBatchSynchronize(batch unsafe.Pointer, cmdBuf   [ MPSState ] _, p2   any  MTLCommandBuffer) unsafe.Pointer {
	return _MPSStateBatchSynchronize(batch, cmdBuf, p2)
}

// Removes any copy of the specified image batch from the device's caches, and, if needed, invalidates any CPU caches.
//
// Added in macOS 10.13.4.
// Removes any copy of the specified image batch from the device's caches, and, if needed, invalidates any CPU caches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2953951-mpsimagebatchsynchronize
func MPSImageBatchSynchronize(batch unsafe.Pointer, cmdBuf   [ MPSImage ] _, p2   any  MTLCommandBuffer) unsafe.Pointer {
	return _MPSImageBatchSynchronize(batch, cmdBuf, p2)
}

// Returns the integer division parameters for a specified divisor.
//
// Added in macOS 10.13.4.
// Returns the integer division parameters for a specified divisor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2954868-mpsfindintegerdivisionparams
func MPSFindIntegerDivisionParams(divisor unsafe.Pointer, p1   UInt16) unsafe.Pointer {
	return _MPSFindIntegerDivisionParams(divisor, p1)
}

// Returns the number of bytes used to allocate the specified image batch.
//
// Added in macOS 10.14.
// Returns the number of bytes used to allocate the specified image batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2980727-mpsimagebatchresourcesize
func MPSImageBatchResourceSize(batch unsafe.Pointer, p1   [ MPSImage ]) unsafe.Pointer {
	return _MPSImageBatchResourceSize(batch, p1)
}

// Returns the number of bytes used to allocate the specified state batch.
//
// Added in macOS 10.14.
// Returns the number of bytes used to allocate the specified state batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2980728-mpsstatebatchresourcesize
func MPSStateBatchResourceSize(batch unsafe.Pointer, p1   [ MPSState ]?) unsafe.Pointer {
	return _MPSStateBatchResourceSize(batch, p1)
}

// Returns the index of the first destination texture argument.
//
// Added in macOS 10.14.
// Returns the index of the first destination texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990481-mpsgetcustomkernelbatcheddestina
func MPSGetCustomKernelBatchedDestinationIndex(c unsafe.Pointer, p1   MPSCustomKernelArgumentCount) unsafe.Pointer {
	return _MPSGetCustomKernelBatchedDestinationIndex(c, p1)
}

// Returns the index of the specified batched source texture argument.
//
// Added in macOS 10.14.
// Returns the index of the specified batched source texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990482-mpsgetcustomkernelbatchedsourcei
func MPSGetCustomKernelBatchedSourceIndex(c unsafe.Pointer, sourceIndex   MPSCustomKernelArgumentCount _, MPSMaxTextures   UInt _, p3   UInt) unsafe.Pointer {
	return _MPSGetCustomKernelBatchedSourceIndex(c, sourceIndex, MPSMaxTextures, p3)
}

// Returns the index of the specified nonbatched source texture argument.
//
// Added in macOS 10.14.
// Returns the index of the specified nonbatched source texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990483-mpsgetcustomkernelbroadcastsourc
func MPSGetCustomKernelBroadcastSourceIndex(c unsafe.Pointer, sourceIndex   MPSCustomKernelArgumentCount _, MPSMaxTextures   UInt _, p3   UInt) unsafe.Pointer {
	return _MPSGetCustomKernelBroadcastSourceIndex(c, sourceIndex, MPSMaxTextures, p3)
}

// Returns the maximum allowed batch size.
//
// Added in macOS 10.14.
// Returns the maximum allowed batch size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990484-mpsgetcustomkernelmaxbatchsize
func MPSGetCustomKernelMaxBatchSize(c unsafe.Pointer, MPSMaxTextures   MPSCustomKernelArgumentCount _, p2   UInt) unsafe.Pointer {
	return _MPSGetCustomKernelMaxBatchSize(c, MPSMaxTextures, p2)
}

// Triggers Metal Performance Shaders to prefetch a Metal heap of the indicated size into its internal cache.
//
// Added in macOS 10.14.
// Triggers Metal Performance Shaders to prefetch a Metal heap of the indicated size into its internal cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990485-mpshinttemporarymemoryhighwaterm
func MPSHintTemporaryMemoryHighWaterMark(cmdBuf unsafe.Pointer, bytes   any  MTLCommandBuffer _, p2   Int) unsafe.Pointer {
	return _MPSHintTemporaryMemoryHighWaterMark(cmdBuf, bytes, p2)
}

// Sets the timeout after which unused cached Metal heaps are released.
//
// Added in macOS 10.14.
// Sets the timeout after which unused cached Metal heaps are released.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/2990486-mpssetheapcacheduration
func MPSSetHeapCacheDuration(cmdBuf unsafe.Pointer, seconds   any  MTLCommandBuffer _, p2   Double) unsafe.Pointer {
	return _MPSSetHeapCacheDuration(cmdBuf, seconds, p2)
}

// Executes a callback block once for each unique image in a batch.
//
// Added in macOS 10.15.
// Executes a callback block once for each unique image in a batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/3019319-mpsimagebatchiterate
func MPSImageBatchIterate(batch unsafe.Pointer, iteratorBlock   [ MPSImage ] _, p2   @escaping MPSImage Int) unsafe.Pointer {
	return _MPSImageBatchIterate(batch, iteratorBlock, p2)
}

// MPSGetPreferredDevice is a MetalPerformanceShaders function.
//
// Added in macOS 10.14.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/3088918-mpsgetpreferreddevice
func MPSGetPreferredDevice(options unsafe.Pointer, p1   MPSDeviceOptions) unsafe.Pointer {
	return _MPSGetPreferredDevice(options, p1)
}

// MPSGetImageType is a MetalPerformanceShaders function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/3131717-mpsgetimagetype
func MPSGetImageType(image unsafe.Pointer, p1   MPSImage) unsafe.Pointer {
	return _MPSGetImageType(image, p1)
}

// MPSSizeofMPSDataType is a MetalPerformanceShaders function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/4092019-mpssizeofmpsdatatype
func MPSSizeofMPSDataType(t unsafe.Pointer, p1   MPSDataType) unsafe.Pointer {
	return _MPSSizeofMPSDataType(t, p1)
}

// MPSDataTypeBitsCount is a MetalPerformanceShaders function.
//
// Added in macOS 14.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/4316495-mpsdatatypebitscount
func MPSDataTypeBitsCount(t unsafe.Pointer, p1   MPSDataType) unsafe.Pointer {
	return _MPSDataTypeBitsCount(t, p1)
}




