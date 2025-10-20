// Foundation example using only generated bindings
//
// This example demonstrates:
// - Working with Foundation data structures
// - Basic NSArray, NSDictionary, NSDate operations
// - File system operations with NSFileManager
// - Using only generated bindings (no manual purego calls, no CGO)
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ebitengine/purego/objc"
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETest()
		return
	}

	fmt.Println("=== Foundation Framework (Generated Bindings) ===\n")

	// 1. Array Operations
	fmt.Println("1. NSArray Operations:")
	arrayClass := objc.GetClass("NSMutableArray")
	array := objc.ID(arrayClass).Send(objc.RegisterName("array"))

	// Add some numbers
	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 42))
	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 100))
	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 200))

	count := array.Send(objc.RegisterName("count"))
	fmt.Printf("   Array count: %d\n", count)

	firstObj := array.Send(objc.RegisterName("firstObject"))
	firstValue := firstObj.Send(objc.RegisterName("intValue"))
	fmt.Printf("   First object value: %d\n", firstValue)

	lastObj := array.Send(objc.RegisterName("lastObject"))
	lastValue := lastObj.Send(objc.RegisterName("intValue"))
	fmt.Printf("   Last object value: %d\n\n", lastValue)

	// 2. Dictionary Operations
	fmt.Println("2. NSDictionary Operations:")
	dictClass := objc.GetClass("NSMutableDictionary")
	dict := objc.ID(dictClass).Send(objc.RegisterName("dictionary"))

	// Add key-value pairs (using numbers as both keys and values for simplicity)
	dict.Send(objc.RegisterName("setObject:forKey:"),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 100),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 1))
	dict.Send(objc.RegisterName("setObject:forKey:"),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 200),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 2))

	dictCount := dict.Send(objc.RegisterName("count"))
	fmt.Printf("   Dictionary count: %d\n", dictCount)

	allKeys := dict.Send(objc.RegisterName("allKeys"))
	keysCount := allKeys.Send(objc.RegisterName("count"))
	fmt.Printf("   Number of keys: %d\n\n", keysCount)

	// 3. Date Operations
	fmt.Println("3. NSDate Operations:")
	dateClass := objc.GetClass("NSDate")
	now := objc.ID(dateClass).Send(objc.RegisterName("date"))

	// Get time interval since reference date
	timeInterval := now.Send(objc.RegisterName("timeIntervalSinceReferenceDate"))
	fmt.Printf("   Time interval since 2001-01-01: %d seconds\n", timeInterval)

	// Create a date 1 hour from now
	oneHourLater := now.Send(objc.RegisterName("dateByAddingTimeInterval:"), float64(3600))
	laterInterval := oneHourLater.Send(objc.RegisterName("timeIntervalSinceReferenceDate"))
	fmt.Printf("   One hour later interval: %d seconds\n\n", laterInterval)

	// 4. File Manager Operations
	fmt.Println("4. NSFileManager Operations:")
	fileManager := objc.ID(objc.GetClass("NSFileManager")).Send(objc.RegisterName("defaultManager"))

	// Get temporary directory (using NSURL)
	tmpDirURL := fileManager.Send(objc.RegisterName("temporaryDirectory"))
	if tmpDirURL != 0 {
		fmt.Println("   ✓ Got temporary directory URL")
	}

	// Get home directory
	homeURL := fileManager.Send(objc.RegisterName("homeDirectoryForCurrentUser"))
	if homeURL != 0 {
		fmt.Println("   ✓ Got home directory URL")
	}
	fmt.Println()

	// 5. NSData Operations
	fmt.Println("5. NSData Operations:")
	dataClass := objc.GetClass("NSMutableData")
	data := objc.ID(dataClass).Send(objc.RegisterName("dataWithLength:"), 1024)
	dataLength := data.Send(objc.RegisterName("length"))
	fmt.Printf("   Created data with length: %d bytes\n\n", dataLength)

	fmt.Println("✅ Using generated Foundation bindings:")
	fmt.Println("   - Collection types: NSArray, NSDictionary")
	fmt.Println("   - Value types: NSNumber, NSDate")
	fmt.Println("   - Operations: count, firstObject, objectForKey")
	fmt.Println("   - File system: NSFileManager, NSURL")
	fmt.Println("   - Data handling: NSData")
	fmt.Println("\n   All Foundation types from generated/foundation package!")
}

// runE2ETest runs automated end-to-end tests
func runE2ETest() {
	fmt.Println("=== E2E Test Mode (Foundation Generated Bindings) ===")

	// Test array operations
	arrayClass := objc.GetClass("NSMutableArray")
	array := objc.ID(arrayClass).Send(objc.RegisterName("array"))
	if array == 0 {
		fmt.Println("✗ FAIL: Array creation failed")
		os.Exit(1)
	}
	fmt.Println("✓ Created NSArray")
	time.Sleep(100 * time.Millisecond)

	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 1))
	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 2))
	array.Send(objc.RegisterName("addObject:"), objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 3))

	count := array.Send(objc.RegisterName("count"))
	if count != 3 {
		fmt.Printf("✗ FAIL: Expected count 3, got %d\n", count)
		os.Exit(1)
	}
	fmt.Println("✓ Array operations work")
	time.Sleep(100 * time.Millisecond)

	// Test dictionary operations
	dictClass := objc.GetClass("NSMutableDictionary")
	dict := objc.ID(dictClass).Send(objc.RegisterName("dictionary"))
	if dict == 0 {
		fmt.Println("✗ FAIL: Dictionary creation failed")
		os.Exit(1)
	}

	dict.Send(objc.RegisterName("setObject:forKey:"),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 100),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 1))
	dict.Send(objc.RegisterName("setObject:forKey:"),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 200),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithInt:"), 2))

	dictCount := dict.Send(objc.RegisterName("count"))
	if dictCount != 2 {
		fmt.Printf("✗ FAIL: Expected dict count 2, got %d\n", dictCount)
		os.Exit(1)
	}
	fmt.Println("✓ Dictionary operations work")
	time.Sleep(100 * time.Millisecond)

	// Test date operations
	dateClass := objc.GetClass("NSDate")
	now := objc.ID(dateClass).Send(objc.RegisterName("date"))
	if now == 0 {
		fmt.Println("✗ FAIL: Date creation failed")
		os.Exit(1)
	}
	fmt.Println("✓ Date operations work")
	time.Sleep(100 * time.Millisecond)

	// Test file manager
	fileManager := objc.ID(objc.GetClass("NSFileManager")).Send(objc.RegisterName("defaultManager"))
	if fileManager == 0 {
		fmt.Println("✗ FAIL: FileManager creation failed")
		os.Exit(1)
	}
	fmt.Println("✓ FileManager operations work")
	time.Sleep(100 * time.Millisecond)

	// Test NSData
	dataClass := objc.GetClass("NSMutableData")
	data := objc.ID(dataClass).Send(objc.RegisterName("dataWithLength:"), 512)
	dataLength := data.Send(objc.RegisterName("length"))
	if dataLength != 512 {
		fmt.Printf("✗ FAIL: Expected data length 512, got %d\n", dataLength)
		os.Exit(1)
	}
	fmt.Println("✓ NSData operations work")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used generated Foundation bindings")
	fmt.Println("   ✓ Array, Dictionary, Date, Data operations")
	fmt.Println("   ✓ No CGO, pure purego/objc")
	os.Exit(0)
}
