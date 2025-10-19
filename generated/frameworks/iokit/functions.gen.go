// Code generated from Apple documentation for IOKit. DO NOT EDIT.

package iokit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// IOKit Functions (116 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOCFSerialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODestroyPlugInInterface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCreatePlugInInterfaceForService func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCatalogueGetData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryCreateIterator func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryCopyFromPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSGetNotificationFromMessage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCFUnserialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOURLWriteDataAndPropertiesToResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallStructMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueDequeue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryCreateCFProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueSetNotificationPort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryCreateCFProperties func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryCreateIterator func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectGetKernelRetainCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryIteratorExitEntry func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetLocationInPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap5 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap2 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceAddMatchingNotification func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceRequestProbe func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetParentIterator func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectCopyBundleIdentifierForClass func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectMapMemory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOIteratorReset func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceAddNotification func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueDataAvailable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallAsyncStructMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCatalogueSendData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap4 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntrySetCFProperties func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceNameMatching func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallAsyncMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectGetService func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOKitWaitQuiet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetParentEntry func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOKitGetBusyState func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectGetUserRetainCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetNameInPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueEnqueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOBSDNameMatching func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap6 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceGetMatchingServices func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueAllocateNotificationPort func() unsafe.Pointer
	_IORegistryEntryGetChildEntry func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOURLCreatePropertyFromResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectConformsTo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectRelease func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectUnmapMemory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceAuthorize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceGetMatchingService func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntrySearchCFProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectSetNotificationPort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOIteratorIsValid func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectIsEqualTo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceWaitQuiet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortGetRunLoopSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceGetBusyState func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectAddClient func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectRelease func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectCopySuperclassForClass func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceClose func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueuePeek func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOMasterPort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceOFPathToBSDName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCatalogueTerminate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryInPlane func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap0 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceMatchPropertyTable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceMatching func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODataQueueWaitForAvailableData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCreateReceivePort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCatalogueReset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetChildIterator func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectSetCFProperties func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOOpenFirmwarePathMatching func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryGetRegistryEntryID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectAddRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOIteratorNext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCFUnserializeWithSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortDestroy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectGetClass func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectUnmapMemory64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectRetain func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODispatchCalloutFromMessage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectCopyClass func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallScalarMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectSetCFProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryFromPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap1 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryIteratorEnterEntry func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOObjectGetRetainCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectTrap3 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOURLCreateDataAndPropertiesFromResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryCopyPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectMapMemory64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceAddInterestNotification func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortGetMachPort func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCFUnserializeBinary func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryGetRootEntry func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOServiceOpenAsFileDescriptor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntryIDMatching func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORegistryEntrySetCFProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOConnectCallAsyncScalarMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOCatalogueModuleLoaded func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IONotificationPortSetImportanceReceiver func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IORPCMessageFromMach func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOMainPort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOCFSerialize, lib, "IOCFSerialize")
	tryRegister(&_IODestroyPlugInInterface, lib, "IODestroyPlugInInterface")
	tryRegister(&_IOCreatePlugInInterfaceForService, lib, "IOCreatePlugInInterfaceForService")
	tryRegister(&_IORegistryEntryGetPath, lib, "IORegistryEntryGetPath")
	tryRegister(&_IOCatalogueGetData, lib, "IOCatalogueGetData")
	tryRegister(&_IORegistryCreateIterator, lib, "IORegistryCreateIterator")
	tryRegister(&_IOConnectCallMethod, lib, "IOConnectCallMethod")
	tryRegister(&_IORegistryEntryCopyFromPath, lib, "IORegistryEntryCopyFromPath")
	tryRegister(&_IORegistryEntryGetProperty, lib, "IORegistryEntryGetProperty")
	tryRegister(&_OSGetNotificationFromMessage, lib, "OSGetNotificationFromMessage")
	tryRegister(&_IOCFUnserialize, lib, "IOCFUnserialize")
	tryRegister(&_IOURLWriteDataAndPropertiesToResource, lib, "IOURLWriteDataAndPropertiesToResource")
	tryRegister(&_IOConnectCallStructMethod, lib, "IOConnectCallStructMethod")
	tryRegister(&_IODataQueueDequeue, lib, "IODataQueueDequeue")
	tryRegister(&_IORegistryEntryCreateCFProperty, lib, "IORegistryEntryCreateCFProperty")
	tryRegister(&_IODataQueueSetNotificationPort, lib, "IODataQueueSetNotificationPort")
	tryRegister(&_IORegistryEntryCreateCFProperties, lib, "IORegistryEntryCreateCFProperties")
	tryRegister(&_IORegistryEntryCreateIterator, lib, "IORegistryEntryCreateIterator")
	tryRegister(&_IORegistryEntryGetName, lib, "IORegistryEntryGetName")
	tryRegister(&_IOObjectGetKernelRetainCount, lib, "IOObjectGetKernelRetainCount")
	tryRegister(&_IORegistryIteratorExitEntry, lib, "IORegistryIteratorExitEntry")
	tryRegister(&_IORegistryEntryGetLocationInPlane, lib, "IORegistryEntryGetLocationInPlane")
	tryRegister(&_IOConnectTrap5, lib, "IOConnectTrap5")
	tryRegister(&_IOConnectTrap2, lib, "IOConnectTrap2")
	tryRegister(&_IOServiceAddMatchingNotification, lib, "IOServiceAddMatchingNotification")
	tryRegister(&_IOServiceRequestProbe, lib, "IOServiceRequestProbe")
	tryRegister(&_IORegistryEntryGetParentIterator, lib, "IORegistryEntryGetParentIterator")
	tryRegister(&_IOObjectCopyBundleIdentifierForClass, lib, "IOObjectCopyBundleIdentifierForClass")
	tryRegister(&_IOConnectMapMemory, lib, "IOConnectMapMemory")
	tryRegister(&_IOIteratorReset, lib, "IOIteratorReset")
	tryRegister(&_IOServiceAddNotification, lib, "IOServiceAddNotification")
	tryRegister(&_IODataQueueDataAvailable, lib, "IODataQueueDataAvailable")
	tryRegister(&_IOConnectCallAsyncStructMethod, lib, "IOConnectCallAsyncStructMethod")
	tryRegister(&_IOCatalogueSendData, lib, "IOCatalogueSendData")
	tryRegister(&_IOConnectTrap4, lib, "IOConnectTrap4")
	tryRegister(&_IORegistryEntrySetCFProperties, lib, "IORegistryEntrySetCFProperties")
	tryRegister(&_IOServiceNameMatching, lib, "IOServiceNameMatching")
	tryRegister(&_IOConnectCallAsyncMethod, lib, "IOConnectCallAsyncMethod")
	tryRegister(&_IOConnectGetService, lib, "IOConnectGetService")
	tryRegister(&_IOKitWaitQuiet, lib, "IOKitWaitQuiet")
	tryRegister(&_IORegistryEntryGetParentEntry, lib, "IORegistryEntryGetParentEntry")
	tryRegister(&_IOKitGetBusyState, lib, "IOKitGetBusyState")
	tryRegister(&_IOObjectGetUserRetainCount, lib, "IOObjectGetUserRetainCount")
	tryRegister(&_IORegistryEntryGetNameInPlane, lib, "IORegistryEntryGetNameInPlane")
	tryRegister(&_IONotificationPortCreate, lib, "IONotificationPortCreate")
	tryRegister(&_IODataQueueEnqueue, lib, "IODataQueueEnqueue")
	tryRegister(&_IOBSDNameMatching, lib, "IOBSDNameMatching")
	tryRegister(&_IOConnectTrap6, lib, "IOConnectTrap6")
	tryRegister(&_IOServiceGetMatchingServices, lib, "IOServiceGetMatchingServices")
	tryRegister(&_IODataQueueAllocateNotificationPort, lib, "IODataQueueAllocateNotificationPort")
	tryRegister(&_IORegistryEntryGetChildEntry, lib, "IORegistryEntryGetChildEntry")
	tryRegister(&_IOURLCreatePropertyFromResource, lib, "IOURLCreatePropertyFromResource")
	tryRegister(&_IOObjectConformsTo, lib, "IOObjectConformsTo")
	tryRegister(&_IOConnectRelease, lib, "IOConnectRelease")
	tryRegister(&_IOServiceOpen, lib, "IOServiceOpen")
	tryRegister(&_IOConnectUnmapMemory, lib, "IOConnectUnmapMemory")
	tryRegister(&_IOServiceAuthorize, lib, "IOServiceAuthorize")
	tryRegister(&_IOServiceGetMatchingService, lib, "IOServiceGetMatchingService")
	tryRegister(&_IORegistryEntrySearchCFProperty, lib, "IORegistryEntrySearchCFProperty")
	tryRegister(&_IOConnectSetNotificationPort, lib, "IOConnectSetNotificationPort")
	tryRegister(&_IOIteratorIsValid, lib, "IOIteratorIsValid")
	tryRegister(&_IOObjectIsEqualTo, lib, "IOObjectIsEqualTo")
	tryRegister(&_IOServiceWaitQuiet, lib, "IOServiceWaitQuiet")
	tryRegister(&_IONotificationPortSetDispatchQueue, lib, "IONotificationPortSetDispatchQueue")
	tryRegister(&_IONotificationPortGetRunLoopSource, lib, "IONotificationPortGetRunLoopSource")
	tryRegister(&_IOServiceGetBusyState, lib, "IOServiceGetBusyState")
	tryRegister(&_IOConnectAddClient, lib, "IOConnectAddClient")
	tryRegister(&_IOObjectRelease, lib, "IOObjectRelease")
	tryRegister(&_IOObjectCopySuperclassForClass, lib, "IOObjectCopySuperclassForClass")
	tryRegister(&_IOServiceClose, lib, "IOServiceClose")
	tryRegister(&_IODataQueuePeek, lib, "IODataQueuePeek")
	tryRegister(&_IOMasterPort, lib, "IOMasterPort")
	tryRegister(&_IOServiceOFPathToBSDName, lib, "IOServiceOFPathToBSDName")
	tryRegister(&_IOCatalogueTerminate, lib, "IOCatalogueTerminate")
	tryRegister(&_IORegistryEntryInPlane, lib, "IORegistryEntryInPlane")
	tryRegister(&_IOConnectTrap0, lib, "IOConnectTrap0")
	tryRegister(&_IOServiceMatchPropertyTable, lib, "IOServiceMatchPropertyTable")
	tryRegister(&_IOServiceMatching, lib, "IOServiceMatching")
	tryRegister(&_IODataQueueWaitForAvailableData, lib, "IODataQueueWaitForAvailableData")
	tryRegister(&_IOCreateReceivePort, lib, "IOCreateReceivePort")
	tryRegister(&_IOCatalogueReset, lib, "IOCatalogueReset")
	tryRegister(&_IORegistryEntryGetChildIterator, lib, "IORegistryEntryGetChildIterator")
	tryRegister(&_IOConnectSetCFProperties, lib, "IOConnectSetCFProperties")
	tryRegister(&_IOOpenFirmwarePathMatching, lib, "IOOpenFirmwarePathMatching")
	tryRegister(&_IORegistryEntryGetRegistryEntryID, lib, "IORegistryEntryGetRegistryEntryID")
	tryRegister(&_IOConnectAddRef, lib, "IOConnectAddRef")
	tryRegister(&_IOIteratorNext, lib, "IOIteratorNext")
	tryRegister(&_IOCFUnserializeWithSize, lib, "IOCFUnserializeWithSize")
	tryRegister(&_IONotificationPortDestroy, lib, "IONotificationPortDestroy")
	tryRegister(&_IOObjectGetClass, lib, "IOObjectGetClass")
	tryRegister(&_IOConnectUnmapMemory64, lib, "IOConnectUnmapMemory64")
	tryRegister(&_IOObjectRetain, lib, "IOObjectRetain")
	tryRegister(&_IODispatchCalloutFromMessage, lib, "IODispatchCalloutFromMessage")
	tryRegister(&_IOObjectCopyClass, lib, "IOObjectCopyClass")
	tryRegister(&_IOConnectCallScalarMethod, lib, "IOConnectCallScalarMethod")
	tryRegister(&_IOConnectSetCFProperty, lib, "IOConnectSetCFProperty")
	tryRegister(&_IORegistryEntryFromPath, lib, "IORegistryEntryFromPath")
	tryRegister(&_IOConnectTrap1, lib, "IOConnectTrap1")
	tryRegister(&_IORegistryIteratorEnterEntry, lib, "IORegistryIteratorEnterEntry")
	tryRegister(&_IOObjectGetRetainCount, lib, "IOObjectGetRetainCount")
	tryRegister(&_IOConnectTrap3, lib, "IOConnectTrap3")
	tryRegister(&_IOURLCreateDataAndPropertiesFromResource, lib, "IOURLCreateDataAndPropertiesFromResource")
	tryRegister(&_IORegistryEntryCopyPath, lib, "IORegistryEntryCopyPath")
	tryRegister(&_IOConnectMapMemory64, lib, "IOConnectMapMemory64")
	tryRegister(&_IOServiceAddInterestNotification, lib, "IOServiceAddInterestNotification")
	tryRegister(&_IONotificationPortGetMachPort, lib, "IONotificationPortGetMachPort")
	tryRegister(&_IOCFUnserializeBinary, lib, "IOCFUnserializeBinary")
	tryRegister(&_IORegistryGetRootEntry, lib, "IORegistryGetRootEntry")
	tryRegister(&_IOServiceOpenAsFileDescriptor, lib, "IOServiceOpenAsFileDescriptor")
	tryRegister(&_IORegistryEntryIDMatching, lib, "IORegistryEntryIDMatching")
	tryRegister(&_IORegistryEntrySetCFProperty, lib, "IORegistryEntrySetCFProperty")
	tryRegister(&_IOConnectCallAsyncScalarMethod, lib, "IOConnectCallAsyncScalarMethod")
	tryRegister(&_IOCatalogueModuleLoaded, lib, "IOCatalogueModuleLoaded")
	tryRegister(&_IONotificationPortSetImportanceReceiver, lib, "IONotificationPortSetImportanceReceiver")
	tryRegister(&_IORPCMessageFromMach, lib, "IORPCMessageFromMach")
	tryRegister(&_IOMainPort, lib, "IOMainPort")
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


// IOCFSerialize is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1403329-iocfserialize
func IOCFSerialize(object unsafe.Pointer, options unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOCFSerialize(object, options, p2)
	}


// IODestroyPlugInInterface is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1412425-iodestroyplugininterface
func IODestroyPlugInInterface(interface_ unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IODestroyPlugInInterface(interface_, p1)
	}


// IOCreatePlugInInterfaceForService is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1412429-iocreateplugininterfaceforservic
func IOCreatePlugInInterfaceForService(service unsafe.Pointer, pluginType unsafe.Pointer, interfaceType unsafe.Pointer, theInterface unsafe.Pointer, theScore unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _IOCreatePlugInInterfaceForService(service, pluginType, interfaceType, theInterface, theScore, p5)
	}


// Create a path for a registry entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514229-ioregistryentrygetpath
func IORegistryEntryGetPath(entry unsafe.Pointer, plane unsafe.Pointer, path unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetPath(entry, plane, path, p3)
	}


// IOCatalogueGetData is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514233-iocataloguegetdata
func IOCatalogueGetData(mainPort unsafe.Pointer, flag unsafe.Pointer, buffer unsafe.Pointer, size unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOCatalogueGetData(mainPort, flag, buffer, size, p4)
	}


// Create an iterator rooted at the registry root. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514238-ioregistrycreateiterator
func IORegistryCreateIterator(mainPort unsafe.Pointer, plane unsafe.Pointer, options unsafe.Pointer, iterator unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryCreateIterator(mainPort, plane, options, iterator, p4)
	}


// IOConnectCallMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514240-ioconnectcallmethod
func IOConnectCallMethod(connection unsafe.Pointer, selector unsafe.Pointer, input unsafe.Pointer, inputCnt unsafe.Pointer, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output unsafe.Pointer, outputCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer, p10 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallMethod(connection, selector, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt, p10)
	}


// IORegistryEntryCopyFromPath is a IOKit function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514248-ioregistryentrycopyfrompath
func IORegistryEntryCopyFromPath(mainPort unsafe.Pointer, path unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryCopyFromPath(mainPort, path, p2)
	}


// IORegistryEntryGetProperty is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514254-ioregistryentrygetproperty
func IORegistryEntryGetProperty(entry unsafe.Pointer, propertyName unsafe.Pointer, buffer unsafe.Pointer, size unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetProperty(entry, propertyName, buffer, size, p4)
	}


// OSGetNotificationFromMessage is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514263-osgetnotificationfrommessage
func OSGetNotificationFromMessage(msg unsafe.Pointer, index unsafe.Pointer, type_ unsafe.Pointer, reference unsafe.Pointer, content unsafe.Pointer, size unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _OSGetNotificationFromMessage(msg, index, type_, reference, content, size, p6)
	}


// IOCFUnserialize is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514265-iocfunserialize
func IOCFUnserialize(buffer unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, errorString unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOCFUnserialize(buffer, allocator, options, errorString, p4)
	}


// IOURLWriteDataAndPropertiesToResource is a IOKit function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514272-iourlwritedataandpropertiestores
func IOURLWriteDataAndPropertiesToResource(url unsafe.Pointer, dataToWrite unsafe.Pointer, propertiesToWrite unsafe.Pointer, errorCode unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOURLWriteDataAndPropertiesToResource(url, dataToWrite, propertiesToWrite, errorCode, p4)
	}


// IOConnectCallStructMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514274-ioconnectcallstructmethod
func IOConnectCallStructMethod(connection unsafe.Pointer, selector unsafe.Pointer, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallStructMethod(connection, selector, inputStruct, inputStructCnt, outputStruct, outputStructCnt, p6)
	}


// Dequeues the next available entry on the queue and copies it into the given data pointer. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514287-iodataqueuedequeue
func IODataQueueDequeue(dataQueue unsafe.Pointer, data unsafe.Pointer, dataSize unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueueDequeue(dataQueue, data, dataSize, p3)
	}


// Create a CF representation of a registry entry's property. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514293-ioregistryentrycreatecfproperty
func IORegistryEntryCreateCFProperty(entry unsafe.Pointer, key unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryCreateCFProperty(entry, key, allocator, options, p4)
	}


// Creates a simple mach message targeting the mach port specified in port. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514301-iodataqueuesetnotificationport
func IODataQueueSetNotificationPort(dataQueue unsafe.Pointer, notifyPort unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueueSetNotificationPort(dataQueue, notifyPort, p2)
	}


// Create a CF dictionary representation of a registry entry's property table. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514310-ioregistryentrycreatecfpropertie
func IORegistryEntryCreateCFProperties(entry unsafe.Pointer, properties unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryCreateCFProperties(entry, properties, allocator, options, p4)
	}


// Create an iterator rooted at a given registry entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514318-ioregistryentrycreateiterator
func IORegistryEntryCreateIterator(entry unsafe.Pointer, plane unsafe.Pointer, options unsafe.Pointer, iterator unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryCreateIterator(entry, plane, options, iterator, p4)
	}


// Returns a C-string name assigned to a registry entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514323-ioregistryentrygetname
func IORegistryEntryGetName(entry unsafe.Pointer, name unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetName(entry, name, p2)
	}


// Returns kernel retain count of an IOKit object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514325-ioobjectgetkernelretaincount
func IOObjectGetKernelRetainCount(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectGetKernelRetainCount(object, p1)
	}


// Exits a level of recursion, restoring the current entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514334-ioregistryiteratorexitentry
func IORegistryIteratorExitEntry(iterator unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryIteratorExitEntry(iterator, p1)
	}


// Returns a C-string location assigned to a registry entry, in a specified plane. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514340-ioregistryentrygetlocationinplan
func IORegistryEntryGetLocationInPlane(entry unsafe.Pointer, plane unsafe.Pointer, location unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetLocationInPlane(entry, plane, location, p3)
	}


// IOConnectTrap5 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514346-ioconnecttrap5
func IOConnectTrap5(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p7 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap5(connect, index, p1, p2, p3, p4, p5, p7)
	}


// IOConnectTrap2 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514354-ioconnecttrap2
func IOConnectTrap2(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap2(connect, index, p1, p2, p4)
	}


// Look up registered IOService objects that match a matching dictionary, and install a notification request of new IOServices that match. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514362-ioserviceaddmatchingnotification
func IOServiceAddMatchingNotification(notifyPort unsafe.Pointer, notificationType unsafe.Pointer, matching unsafe.Pointer, callback unsafe.Pointer, refCon unsafe.Pointer, notification unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceAddMatchingNotification(notifyPort, notificationType, matching, callback, refCon, notification, p6)
	}


// A request to rescan a bus for device changes. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514364-ioservicerequestprobe
func IOServiceRequestProbe(service unsafe.Pointer, options unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceRequestProbe(service, options, p2)
	}


// Returns an iterator over a registry entry’s parent entries in a plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514366-ioregistryentrygetparentiterator
func IORegistryEntryGetParentIterator(entry unsafe.Pointer, plane unsafe.Pointer, iterator unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetParentIterator(entry, plane, iterator, p3)
	}


// Return the bundle identifier of the given class. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514375-ioobjectcopybundleidentifierforc
func IOObjectCopyBundleIdentifierForClass(classname unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectCopyBundleIdentifierForClass(classname, p1)
	}


// Map hardware or shared memory into the caller's task. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514377-ioconnectmapmemory
func IOConnectMapMemory(connect unsafe.Pointer, memoryType unsafe.Pointer, intoTask unsafe.Pointer, atAddress unsafe.Pointer, ofSize unsafe.Pointer, options unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectMapMemory(connect, memoryType, intoTask, atAddress, ofSize, options, p6)
	}


// Resets an iteration back to the beginning. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514379-ioiteratorreset
func IOIteratorReset(iterator unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOIteratorReset(iterator, p1)
	}


// IOServiceAddNotification is a IOKit function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514382-ioserviceaddnotification
func IOServiceAddNotification(mainPort unsafe.Pointer, notificationType unsafe.Pointer, matching unsafe.Pointer, wakePort unsafe.Pointer, reference unsafe.Pointer, notification unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceAddNotification(mainPort, notificationType, matching, wakePort, reference, notification, p6)
	}


// Used to determine if more data is avilable on the queue. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514386-iodataqueuedataavailable
func IODataQueueDataAvailable(dataQueue unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueueDataAvailable(dataQueue, p1)
	}


// IOConnectCallAsyncStructMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514403-ioconnectcallasyncstructmethod
func IOConnectCallAsyncStructMethod(connection unsafe.Pointer, selector unsafe.Pointer, wake_port unsafe.Pointer, reference unsafe.Pointer, referenceCnt unsafe.Pointer, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer, p9 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallAsyncStructMethod(connection, selector, wake_port, reference, referenceCnt, inputStruct, inputStructCnt, outputStruct, outputStructCnt, p9)
	}


// IOCatalogueSendData is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514405-iocataloguesenddata
func IOCatalogueSendData(mainPort unsafe.Pointer, flag unsafe.Pointer, buffer unsafe.Pointer, size unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOCatalogueSendData(mainPort, flag, buffer, size, p4)
	}


// IOConnectTrap4 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514410-ioconnecttrap4
func IOConnectTrap4(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap4(connect, index, p1, p2, p3, p4, p6)
	}


// Set CF container based properties in a registry entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514414-ioregistryentrysetcfproperties
func IORegistryEntrySetCFProperties(entry unsafe.Pointer, properties unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntrySetCFProperties(entry, properties, p2)
	}


// Create a matching dictionary that specifies an IOService name match. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514416-ioservicenamematching
func IOServiceNameMatching(name unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceNameMatching(name, p1)
	}


// IOConnectCallAsyncMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514418-ioconnectcallasyncmethod
func IOConnectCallAsyncMethod(connection unsafe.Pointer, selector unsafe.Pointer, wake_port unsafe.Pointer, reference unsafe.Pointer, referenceCnt unsafe.Pointer, input unsafe.Pointer, inputCnt unsafe.Pointer, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output unsafe.Pointer, outputCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer, p13 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallAsyncMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt, p13)
	}


// Returns the IOService a connect handle was opened on. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514438-ioconnectgetservice
func IOConnectGetService(connect unsafe.Pointer, service unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectGetService(connect, service, p2)
	}


// Wait for a all IOServices' busyState to be zero. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514440-iokitwaitquiet
func IOKitWaitQuiet(mainPort unsafe.Pointer, waitTime unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOKitWaitQuiet(mainPort, waitTime, p2)
	}


// Returns the first parent of a registry entry in a plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514454-ioregistryentrygetparententry
func IORegistryEntryGetParentEntry(entry unsafe.Pointer, plane unsafe.Pointer, parent unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetParentEntry(entry, plane, parent, p3)
	}


// Returns the busyState of all IOServices. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514460-iokitgetbusystate
func IOKitGetBusyState(mainPort unsafe.Pointer, busyState unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOKitGetBusyState(mainPort, busyState, p2)
	}


// Returns the retain count for the current process of an IOKit object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514464-ioobjectgetuserretaincount
func IOObjectGetUserRetainCount(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectGetUserRetainCount(object, p1)
	}


// Returns a C-string name assigned to a registry entry, in a specified plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514475-ioregistryentrygetnameinplane
func IORegistryEntryGetNameInPlane(entry unsafe.Pointer, plane unsafe.Pointer, name unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetNameInPlane(entry, plane, name, p3)
	}


// Creates and returns a notification object for receiving IOKit notifications of new devices or state changes. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514480-ionotificationportcreate
func IONotificationPortCreate(mainPort unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortCreate(mainPort, p1)
	}


// Enqueues a new entry on the queue. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514482-iodataqueueenqueue
func IODataQueueEnqueue(dataQueue unsafe.Pointer, data unsafe.Pointer, dataSize unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueueEnqueue(dataQueue, data, dataSize, p3)
	}


// Create a matching dictionary that specifies an IOService match based on BSD device name. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514486-iobsdnamematching
func IOBSDNameMatching(mainPort unsafe.Pointer, options unsafe.Pointer, bsdName unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOBSDNameMatching(mainPort, options, bsdName, p3)
	}


// IOConnectTrap6 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514493-ioconnecttrap6
func IOConnectTrap6(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer, p8 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap6(connect, index, p1, p2, p3, p4, p5, p6, p8)
	}


// Look up registered IOService objects that match a matching dictionary. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514494-ioservicegetmatchingservices
func IOServiceGetMatchingServices(mainPort unsafe.Pointer, matching unsafe.Pointer, existing unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceGetMatchingServices(mainPort, matching, existing, p3)
	}


// Allocates and returns a new mach port able to receive data available notifications from an IODataQueue. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514495-iodataqueueallocatenotificationp
func IODataQueueAllocateNotificationPort() unsafe.Pointer {
	return _IODataQueueAllocateNotificationPort()
	}


// Returns the first child of a registry entry in a plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514496-ioregistryentrygetchildentry
func IORegistryEntryGetChildEntry(entry unsafe.Pointer, plane unsafe.Pointer, child unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetChildEntry(entry, plane, child, p3)
	}


// IOURLCreatePropertyFromResource is a IOKit function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514499-iourlcreatepropertyfromresource
func IOURLCreatePropertyFromResource(alloc unsafe.Pointer, url unsafe.Pointer, property unsafe.Pointer, errorCode unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOURLCreatePropertyFromResource(alloc, url, property, errorCode, p4)
	}


// Performs an OSDynamicCast operation on an IOKit object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514505-ioobjectconformsto
func IOObjectConformsTo(object unsafe.Pointer, className unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectConformsTo(object, className, p2)
	}


// Remove a reference to the connect handle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514511-ioconnectrelease
func IOConnectRelease(connect unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectRelease(connect, p1)
	}


// A request to create a connection to an IOService. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514515-ioserviceopen
func IOServiceOpen(service unsafe.Pointer, owningTask unsafe.Pointer, type_ unsafe.Pointer, connect unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceOpen(service, owningTask, type_, connect, p4)
	}


// Remove a mapping made with IOConnectMapMemory. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514527-ioconnectunmapmemory
func IOConnectUnmapMemory(connect unsafe.Pointer, memoryType unsafe.Pointer, fromTask unsafe.Pointer, atAddress unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectUnmapMemory(connect, memoryType, fromTask, atAddress, p4)
	}


// IOServiceAuthorize is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514533-ioserviceauthorize
func IOServiceAuthorize(service unsafe.Pointer, options unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceAuthorize(service, options, p2)
	}


// Look up a registered IOService object that matches a matching dictionary. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514535-ioservicegetmatchingservice
func IOServiceGetMatchingService(mainPort unsafe.Pointer, matching unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceGetMatchingService(mainPort, matching, p2)
	}


// Create a CF representation of a registry entry's property. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514537-ioregistryentrysearchcfproperty
func IORegistryEntrySearchCFProperty(entry unsafe.Pointer, plane unsafe.Pointer, key unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntrySearchCFProperty(entry, plane, key, allocator, options, p5)
	}


// Set a port to receive family specific notifications. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514541-ioconnectsetnotificationport
func IOConnectSetNotificationPort(connect unsafe.Pointer, type_ unsafe.Pointer, port unsafe.Pointer, reference unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectSetNotificationPort(connect, type_, port, reference, p4)
	}


// Checks an iterator is still valid. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514556-ioiteratorisvalid
func IOIteratorIsValid(iterator unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOIteratorIsValid(iterator, p1)
	}


// Checks two object handles to see if they represent the same kernel object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514563-ioobjectisequalto
func IOObjectIsEqualTo(object unsafe.Pointer, anObject unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectIsEqualTo(object, anObject, p2)
	}


// Wait for an IOService's busyState to be zero. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514573-ioservicewaitquiet
func IOServiceWaitQuiet(service unsafe.Pointer, waitTime unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceWaitQuiet(service, waitTime, p2)
	}


// Sets a dispatch queue to be used to listen for notifications. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514596-ionotificationportsetdispatchque
func IONotificationPortSetDispatchQueue(notify unsafe.Pointer, queue unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortSetDispatchQueue(notify, queue, p2)
	}


// Returns a CFRunLoopSource to be used to listen for notifications. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514599-ionotificationportgetrunloopsour
func IONotificationPortGetRunLoopSource(notify unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortGetRunLoopSource(notify, p1)
	}


// Returns the busyState of an IOService. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514607-ioservicegetbusystate
func IOServiceGetBusyState(service unsafe.Pointer, busyState unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceGetBusyState(service, busyState, p2)
	}


// Inform a connection of a second connection. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514609-ioconnectaddclient
func IOConnectAddClient(connect unsafe.Pointer, client unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectAddClient(connect, client, p2)
	}


// Releases an object handle previously returned by IOKitLib. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514627-ioobjectrelease
func IOObjectRelease(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectRelease(object, p1)
	}


// Return the superclass name of the given class. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514635-ioobjectcopysuperclassforclass
func IOObjectCopySuperclassForClass(classname unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectCopySuperclassForClass(classname, p1)
	}


// Close a connection to an IOService and destroy the connect handle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514646-ioserviceclose
func IOServiceClose(connect unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceClose(connect, p1)
	}


// Used to peek at the next entry on the queue. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514649-iodataqueuepeek
func IODataQueuePeek(dataQueue unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueuePeek(dataQueue, p1)
	}


// Returns the mach port used to initiate communication with IOKit. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514652-iomasterport
func IOMasterPort(bootstrapPort unsafe.Pointer, mainPort unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOMasterPort(bootstrapPort, mainPort, p2)
	}


// IOServiceOFPathToBSDName is a IOKit function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514661-ioserviceofpathtobsdname
func IOServiceOFPathToBSDName(mainPort unsafe.Pointer, openFirmwarePath unsafe.Pointer, bsdName unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceOFPathToBSDName(mainPort, openFirmwarePath, bsdName, p3)
	}


// IOCatalogueTerminate is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514665-iocatalogueterminate
func IOCatalogueTerminate(mainPort unsafe.Pointer, flag unsafe.Pointer, description unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOCatalogueTerminate(mainPort, flag, description, p3)
	}


// Determines if the registry entry is attached in a plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514668-ioregistryentryinplane
func IORegistryEntryInPlane(entry unsafe.Pointer, plane unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryInPlane(entry, plane, p2)
	}


// IOConnectTrap0 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514674-ioconnecttrap0
func IOConnectTrap0(connect unsafe.Pointer, index unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap0(connect, index, p2)
	}


// Match an IOService objects with matching dictionary. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514685-ioservicematchpropertytable
func IOServiceMatchPropertyTable(service unsafe.Pointer, matching unsafe.Pointer, matches unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceMatchPropertyTable(service, matching, matches, p3)
	}


// Create a matching dictionary that specifies an IOService class match. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514687-ioservicematching
func IOServiceMatching(name unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceMatching(name, p1)
	}


// Wait for an incoming dataAvailable message on the given notifyPort. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514696-iodataqueuewaitforavailabledata
func IODataQueueWaitForAvailableData(dataQueue unsafe.Pointer, notificationPort unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IODataQueueWaitForAvailableData(dataQueue, notificationPort, p2)
	}


// Creates and returns a mach port suitable for receiving IOKit messages of the specified type. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514698-iocreatereceiveport
func IOCreateReceivePort(msgType unsafe.Pointer, recvPort unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOCreateReceivePort(msgType, recvPort, p2)
	}


// IOCatalogueReset is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514702-iocataloguereset
func IOCatalogueReset(mainPort unsafe.Pointer, flag unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOCatalogueReset(mainPort, flag, p2)
	}


// Returns an iterator over a registry entry’s child entries in a plane. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514703-ioregistryentrygetchilditerator
func IORegistryEntryGetChildIterator(entry unsafe.Pointer, plane unsafe.Pointer, iterator unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetChildIterator(entry, plane, iterator, p3)
	}


// Set CF container based properties on a connection. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514713-ioconnectsetcfproperties
func IOConnectSetCFProperties(connect unsafe.Pointer, properties unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectSetCFProperties(connect, properties, p2)
	}


// IOOpenFirmwarePathMatching is a IOKit function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514715-ioopenfirmwarepathmatching
func IOOpenFirmwarePathMatching(mainPort unsafe.Pointer, options unsafe.Pointer, path unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOOpenFirmwarePathMatching(mainPort, options, path, p3)
	}


// Returns an ID for the registry entry that is global to all tasks. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514719-ioregistryentrygetregistryentryi
func IORegistryEntryGetRegistryEntryID(entry unsafe.Pointer, entryID unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryGetRegistryEntryID(entry, entryID, p2)
	}


// Adds a reference to the connect handle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514739-ioconnectaddref
func IOConnectAddRef(connect unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectAddRef(connect, p1)
	}


// Returns the next object in an iteration. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514741-ioiteratornext
func IOIteratorNext(iterator unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOIteratorNext(iterator, p1)
	}


// IOCFUnserializeWithSize is a IOKit function. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514745-iocfunserializewithsize
func IOCFUnserializeWithSize(buffer unsafe.Pointer, bufferSize unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, errorString unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _IOCFUnserializeWithSize(buffer, bufferSize, allocator, options, errorString, p5)
	}


// Destroys a notification object created with IONotificationPortCreate. Also destroys any mach_port's or CFRunLoopSources obatined from   or  [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514751-ionotificationportdestroy
func IONotificationPortDestroy(notify unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortDestroy(notify, p1)
	}


// Return the class name of an IOKit object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514756-ioobjectgetclass
func IOObjectGetClass(object unsafe.Pointer, className unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectGetClass(object, className, p2)
	}


// Remove a mapping made with IOConnectMapMemory64. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514760-ioconnectunmapmemory64
func IOConnectUnmapMemory64(connect unsafe.Pointer, memoryType unsafe.Pointer, fromTask unsafe.Pointer, atAddress unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectUnmapMemory64(connect, memoryType, fromTask, atAddress, p4)
	}


// Retains an object handle previously returned by IOKitLib. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514769-ioobjectretain
func IOObjectRetain(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectRetain(object, p1)
	}


// Dispatches callback notifications from a mach message. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514775-iodispatchcalloutfrommessage
func IODispatchCalloutFromMessage(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IODispatchCalloutFromMessage(unused, msg, reference, p3)
	}


// Return the class name of an IOKit object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514781-ioobjectcopyclass
func IOObjectCopyClass(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectCopyClass(object, p1)
	}


// IOConnectCallScalarMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514793-ioconnectcallscalarmethod
func IOConnectCallScalarMethod(connection unsafe.Pointer, selector unsafe.Pointer, input unsafe.Pointer, inputCnt unsafe.Pointer, output unsafe.Pointer, outputCnt unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallScalarMethod(connection, selector, input, inputCnt, output, outputCnt, p6)
	}


// Set a CF container based property on a connection. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514796-ioconnectsetcfproperty
func IOConnectSetCFProperty(connect unsafe.Pointer, propertyName unsafe.Pointer, property unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectSetCFProperty(connect, propertyName, property, p3)
	}


// Looks up a registry entry by path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514802-ioregistryentryfrompath
func IORegistryEntryFromPath(mainPort unsafe.Pointer, path unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryFromPath(mainPort, path, p2)
	}


// IOConnectTrap1 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514816-ioconnecttrap1
func IOConnectTrap1(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap1(connect, index, p1, p3)
	}


// Recurse into the current entry in the registry iteration. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514822-ioregistryiteratorenterentry
func IORegistryIteratorEnterEntry(iterator unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryIteratorEnterEntry(iterator, p1)
	}


// Returns kernel retain count of an IOKit object. Identical to IOObjectGetKernelRetainCount() but available prior to Mac OS 10.6. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514824-ioobjectgetretaincount
func IOObjectGetRetainCount(object unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IOObjectGetRetainCount(object, p1)
	}


// IOConnectTrap3 is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514833-ioconnecttrap3
func IOConnectTrap3(connect unsafe.Pointer, index unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectTrap3(connect, index, p1, p2, p3, p5)
	}


// IOURLCreateDataAndPropertiesFromResource is a IOKit function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514836-iourlcreatedataandpropertiesfrom
func IOURLCreateDataAndPropertiesFromResource(alloc unsafe.Pointer, url unsafe.Pointer, resourceData unsafe.Pointer, properties unsafe.Pointer, desiredProperties unsafe.Pointer, errorCode unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode, p6)
	}


// IORegistryEntryCopyPath is a IOKit function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514853-ioregistryentrycopypath
func IORegistryEntryCopyPath(entry unsafe.Pointer, plane unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryCopyPath(entry, plane, p2)
	}


// Map hardware or shared memory into the caller's task. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514862-ioconnectmapmemory64
func IOConnectMapMemory64(connect unsafe.Pointer, memoryType unsafe.Pointer, intoTask unsafe.Pointer, atAddress unsafe.Pointer, ofSize unsafe.Pointer, options unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectMapMemory64(connect, memoryType, intoTask, atAddress, ofSize, options, p6)
	}


// Register for notification of state changes in an IOService. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514866-ioserviceaddinterestnotification
func IOServiceAddInterestNotification(notifyPort unsafe.Pointer, service unsafe.Pointer, interestType unsafe.Pointer, callback unsafe.Pointer, refCon unsafe.Pointer, notification unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceAddInterestNotification(notifyPort, service, interestType, callback, refCon, notification, p6)
	}


// Returns a mach_port to be used to listen for notifications. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514875-ionotificationportgetmachport
func IONotificationPortGetMachPort(notify unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortGetMachPort(notify, p1)
	}


// IOCFUnserializeBinary is a IOKit function. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514876-iocfunserializebinary
func IOCFUnserializeBinary(buffer unsafe.Pointer, bufferSize unsafe.Pointer, allocator unsafe.Pointer, options unsafe.Pointer, errorString unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _IOCFUnserializeBinary(buffer, bufferSize, allocator, options, errorString, p5)
	}


// Return a handle to the registry root. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514878-ioregistrygetrootentry
func IORegistryGetRootEntry(mainPort unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryGetRootEntry(mainPort, p1)
	}


// IOServiceOpenAsFileDescriptor is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514879-ioserviceopenasfiledescriptor
func IOServiceOpenAsFileDescriptor(service unsafe.Pointer, oflag unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOServiceOpenAsFileDescriptor(service, oflag, p2)
	}


// Create a matching dictionary that specifies an IOService match based on a registry entry ID. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514880-ioregistryentryidmatching
func IORegistryEntryIDMatching(entryID unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntryIDMatching(entryID, p1)
	}


// Set a CF container based property in a registry entry. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514882-ioregistryentrysetcfproperty
func IORegistryEntrySetCFProperty(entry unsafe.Pointer, propertyName unsafe.Pointer, property unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _IORegistryEntrySetCFProperty(entry, propertyName, property, p3)
	}


// IOConnectCallAsyncScalarMethod is a IOKit function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514884-ioconnectcallasyncscalarmethod
func IOConnectCallAsyncScalarMethod(connection unsafe.Pointer, selector unsafe.Pointer, wake_port unsafe.Pointer, reference unsafe.Pointer, referenceCnt unsafe.Pointer, input unsafe.Pointer, inputCnt unsafe.Pointer, output unsafe.Pointer, outputCnt unsafe.Pointer, p9 unsafe.Pointer) unsafe.Pointer {
	return _IOConnectCallAsyncScalarMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, output, outputCnt, p9)
	}


// IOCatalogueModuleLoaded is a IOKit function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/1514886-iocataloguemoduleloaded
func IOCatalogueModuleLoaded(mainPort unsafe.Pointer, name unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOCatalogueModuleLoaded(mainPort, name, p2)
	}


// IONotificationPortSetImportanceReceiver is a IOKit function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/2870065-ionotificationportsetimportancer
func IONotificationPortSetImportanceReceiver(notify unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _IONotificationPortSetImportanceReceiver(notify, p1)
	}


// IORPCMessageFromMach is a IOKit function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/3325691-iorpcmessagefrommach
func IORPCMessageFromMach(msg unsafe.Pointer, reply unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IORPCMessageFromMach(msg, reply, p2)
	}


// IOMainPort is a IOKit function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/iokit/3753260-iomainport
func IOMainPort(bootstrapPort unsafe.Pointer, mainPort unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _IOMainPort(bootstrapPort, mainPort, p2)
	}



