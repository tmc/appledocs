package main

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// nsArrayCount returns the count of an NSArray (works for any objc.ID that responds to count)
func nsArrayCount(arr unsafe.Pointer) int {
	if arr == nil {
		return 0
	}
	return int(objc.ID(arr).Send(objc.RegisterName("count")))
}

// nsArrayObjectAt returns the object at index i in an NSArray
func nsArrayObjectAt(arr unsafe.Pointer, i int) objc.ID {
	// Use generated foundation.Array method
	array := foundation.ArrayFrom(arr)
	return objc.ID(array.ObjectAtIndex(uint(i)))
}

// nsStringToGo converts an NSString (as objc.ID) to a Go string
func nsStringToGo(str objc.ID) string {
	if str == 0 {
		return ""
	}
	return objc.Send[string](str, objc.RegisterName("UTF8String"))
}

// nsStringPtrToGo converts an NSString pointer (unsafe.Pointer) to a Go string
func nsStringPtrToGo(str unsafe.Pointer) string {
	return nsStringToGo(objc.ID(str))
}

// nsErrorToGo converts an NSError (objc.ID) to a Go error
func nsErrorToGo(err objc.ID) error {
	if err == 0 {
		return nil
	}
	desc := err.Send(objc.RegisterName("localizedDescription"))
	if desc == 0 {
		return errors.New("unknown error")
	}
	errMsg := objc.Send[string](desc, objc.RegisterName("UTF8String"))
	return errors.New(errMsg)
}

// newCompletionHandler creates a completion handler block and returns the block and error channel
func newCompletionHandler() (objc.Block, <-chan error) {
	errChan := make(chan error, 1)
	block := objc.NewBlock(func(b objc.Block, err objc.ID) {
		errChan <- nsErrorToGo(err)
	})
	return block, errChan
}

// emptyNSArray creates an empty NSArray
func emptyNSArray() objc.ID {
	arrayClass := objc.GetClass("NSArray")
	return objc.ID(arrayClass).Send(objc.RegisterName("array"))
}

// awaitCompletion waits for a completion handler and formats errors
func awaitCompletion(done <-chan error, operation string) error {
	if err := <-done; err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	return nil
}
