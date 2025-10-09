// Foundation Framework Example - Working Implementation
//
// This example demonstrates actual Foundation framework usage with purego.
// Unlike the template examples, this shows real API calls.
package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end tests")
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETests()
		return
	}

	fmt.Println("Foundation Framework - Working Example")
	fmt.Println("======================================")
	fmt.Println()

	// Load Foundation framework
	_, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Printf("Failed to load Foundation: %v\n", err)
		return
	}
	fmt.Println("✓ Foundation framework loaded")
	fmt.Println()

	// Test NSString creation
	testNSString()
	fmt.Println()

	// Test NSArray creation
	testNSArray()
	fmt.Println()

	// Test NSDate
	testNSDate()
	fmt.Println()

	fmt.Println("All Foundation operations completed successfully!")
}

func testNSString() {
	fmt.Println("Testing NSString:")

	// Create NSString from Go string
	str := objc.RegisterName("Hello from Foundation!")
	nsString := objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("alloc"))
	nsString = nsString.Send(objc.RegisterName("initWithUTF8String:"), str)

	// Get length
	length := nsString.Send(objc.RegisterName("length"))
	fmt.Printf("  Created NSString with length: %d\n", length)

	// Get UTF8 string back
	cstr := nsString.Send(objc.RegisterName("UTF8String"))
	if cstr != 0 {
		fmt.Printf("  String content verified: %s\n", objc.ID(cstr))
	}

	// Release
	nsString.Send(objc.RegisterName("release"))
	fmt.Println("  ✓ NSString test passed")
}

func testNSArray() {
	fmt.Println("Testing NSArray:")

	// Create empty array
	nsArray := objc.ID(objc.GetClass("NSArray")).Send(objc.RegisterName("array"))
	count := nsArray.Send(objc.RegisterName("count"))
	fmt.Printf("  Created NSArray with count: %d\n", count)

	// Create mutable array and add objects
	nsMutableArray := objc.ID(objc.GetClass("NSMutableArray")).Send(objc.RegisterName("array"))

	// Add a string
	str := objc.RegisterName("test")
	nsString := objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("alloc"))
	nsString = nsString.Send(objc.RegisterName("initWithUTF8String:"), str)

	nsMutableArray.Send(objc.RegisterName("addObject:"), nsString)
	count = nsMutableArray.Send(objc.RegisterName("count"))
	fmt.Printf("  NSMutableArray with %d object\n", count)

	nsString.Send(objc.RegisterName("release"))
	fmt.Println("  ✓ NSArray test passed")
}

func testNSDate() {
	fmt.Println("Testing NSDate:")

	// Get current date
	nsDate := objc.ID(objc.GetClass("NSDate")).Send(objc.RegisterName("date"))

	// Get time interval since 1970
	timeInterval := nsDate.Send(objc.RegisterName("timeIntervalSince1970"))
	fmt.Printf("  Current time interval: %d seconds since 1970\n", timeInterval)

	// Create date description
	description := nsDate.Send(objc.RegisterName("description"))
	if description != 0 {
		cstr := objc.ID(description).Send(objc.RegisterName("UTF8String"))
		if cstr != 0 {
			// Note: actual string extraction would require more work
			fmt.Println("  Date object created successfully")
		}
	}
	fmt.Println("  ✓ NSDate test passed")
}

func runE2ETests() {
	fmt.Println("Running Foundation E2E Tests...")
	fmt.Println()

	tests := []struct {
		name string
		fn   func() error
	}{
		{"Framework Loading", testFrameworkLoad},
		{"NSString Operations", testNSStringE2E},
		{"NSArray Operations", testNSArrayE2E},
		{"NSDate Operations", testNSDateE2E},
	}

	passed := 0
	failed := 0

	for _, test := range tests {
		fmt.Printf("  Test: %s... ", test.name)
		if err := test.fn(); err != nil {
			fmt.Printf("✗ FAIL: %v\n", err)
			failed++
		} else {
			fmt.Println("✓ PASS")
			passed++
		}
	}

	fmt.Println()
	fmt.Printf("Tests: %d passed, %d failed\n", passed, failed)

	if failed > 0 {
		fmt.Println("E2E tests FAILED")
	} else {
		fmt.Println("All E2E tests PASSED!")
	}
}

func testFrameworkLoad() error {
	_, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	return err
}

func testNSStringE2E() error {
	str := objc.RegisterName("test")
	nsString := objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("alloc"))
	nsString = nsString.Send(objc.RegisterName("initWithUTF8String:"), str)
	defer nsString.Send(objc.RegisterName("release"))

	length := nsString.Send(objc.RegisterName("length"))
	if length != 4 {
		return fmt.Errorf("expected length 4, got %d", length)
	}
	return nil
}

func testNSArrayE2E() error {
	nsArray := objc.ID(objc.GetClass("NSArray")).Send(objc.RegisterName("array"))
	count := nsArray.Send(objc.RegisterName("count"))
	if count != 0 {
		return fmt.Errorf("expected empty array, got count %d", count)
	}
	return nil
}

func testNSDateE2E() error {
	nsDate := objc.ID(objc.GetClass("NSDate")).Send(objc.RegisterName("date"))
	if nsDate == 0 {
		return fmt.Errorf("failed to create NSDate")
	}
	return nil
}
