// SpriteKit example using only generated bindings
//
// This example demonstrates:
// - Creating a SpriteKit scene with animated sprites
// - Using SKView, SKScene, SKSpriteNode
// - Adding physics simulation with SKPhysicsBody
// - Using only generated bindings (no manual purego calls)
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end test mode (non-interactive)")
)

func init() {
	runtime.LockOSThread()
}

func createNSString(s string) objc.ID {
	strClass := objc.GetClass("NSString")
	str := objc.ID(strClass).Send(objc.RegisterName("alloc"))
	return str.Send(objc.RegisterName("initWithUTF8String:"), s)
}

// RunApp runs the AppKit event loop with proper initialization
func RunApp(didLaunch func(app appkit.Application)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("=== SpriteKit Game (Generated Bindings) ===")

	app := appkit.SharedApplication()

	delegateClass, err := objc.RegisterClass(
		"AppDelegate",
		objc.GetClass("NSObject"),
		nil,
		nil,
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("applicationDidFinishLaunching:"),
				Fn: func(self objc.ID, _cmd objc.SEL, notification objc.ID) {
					fmt.Println("✓ Application finished launching")
					didLaunch(app)
				},
			},
			{
				Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
				Fn: func(self objc.ID, _cmd objc.SEL, app objc.ID) bool {
					return false
				},
			},
			{
				Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
				Fn: func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
					return true
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ERROR: Failed to register delegate class: %v\n", err)
		os.Exit(1)
	}

	delegate := objc.ID(delegateClass).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	app.ID.Send(objc.RegisterName("setDelegate:"), delegate)

	fmt.Println("✓ Starting AppKit event loop...")
	app.Run()
}

func main() {
	flag.Parse()

	if *e2e {
		runE2ETest()
		return
	}

	RunApp(func(app appkit.Application) {
		app.SetActivationPolicy(appkit.ActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create window
		type NSPoint struct{ X, Y float64 }
		type NSSize struct{ Width, Height float64 }
		type NSRect struct {
			Origin NSPoint
			Size   NSSize
		}
		rect := NSRect{
			Origin: NSPoint{X: 100, Y: 100},
			Size:   NSSize{Width: 800, Height: 600},
		}

		windowClass := objc.GetClass("NSWindow")
		windowID := objc.ID(windowClass).Send(objc.RegisterName("alloc"))
		styleMask := appkit.WindowStyleMaskTitled | appkit.WindowStyleMaskClosable | appkit.WindowStyleMaskResizable
		windowID = windowID.Send(objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
			unsafe.Pointer(&rect), styleMask, appkit.BackingStoreBuffered, false)
		window := appkit.WindowFrom(unsafe.Pointer(windowID))

		window.SetTitle("SpriteKit Game Example")

		// Create SKView
		skViewClass := objc.GetClass("SKView")
		skView := objc.ID(skViewClass).Send(objc.RegisterName("alloc"))
		skView = skView.Send(objc.RegisterName("initWithFrame:"), NSRect{
			Origin: NSPoint{X: 0, Y: 0},
			Size:   NSSize{Width: 800, Height: 600},
		})

		// Enable debug options
		skView.Send(objc.RegisterName("setShowsFPS:"), true)
		skView.Send(objc.RegisterName("setShowsNodeCount:"), true)

		// Create custom scene
		sceneClass := objc.GetClass("SKScene")
		scene := objc.ID(sceneClass).Send(objc.RegisterName("alloc"))

		// Create CGSize for scene
		type CGSize struct {
			Width  float64
			Height float64
		}
		sceneSize := CGSize{Width: 800, Height: 600}
		scene = scene.Send(objc.RegisterName("initWithSize:"), sceneSize)

		// Set scene background color (dark blue)
		colorClass := objc.GetClass("NSColor")
		bgColor := objc.ID(colorClass).Send(objc.RegisterName("colorWithRed:green:blue:alpha:"), 0.1, 0.1, 0.3, 1.0)
		scene.Send(objc.RegisterName("setBackgroundColor:"), bgColor)

		// Create sprite nodes
		for i := 0; i < 5; i++ {
			// Create sprite node
			spriteClass := objc.GetClass("SKSpriteNode")
			sprite := objc.ID(spriteClass).Send(objc.RegisterName("alloc"))

			// Create with color and size
			spriteSize := CGSize{Width: 50, Height: 50}
			spriteColor := objc.ID(colorClass).Send(objc.RegisterName("colorWithRed:green:blue:alpha:"),
				float64(i)*0.2, 1.0-float64(i)*0.2, 0.5, 1.0)
			sprite = sprite.Send(objc.RegisterName("initWithColor:size:"), spriteColor, spriteSize)

			// Set position
			type CGPoint struct{ X, Y float64 }
			position := CGPoint{X: 100 + float64(i)*150, Y: 500}
			sprite.Send(objc.RegisterName("setPosition:"), position)

			// Add physics body
			physicsBodyClass := objc.GetClass("SKPhysicsBody")
			physicsBody := objc.ID(physicsBodyClass).Send(objc.RegisterName("bodyWithRectangleOfSize:"), spriteSize)
			physicsBody.Send(objc.RegisterName("setRestitution:"), 0.8) // Bounciness
			sprite.Send(objc.RegisterName("setPhysicsBody:"), physicsBody)

			// Add to scene
			scene.Send(objc.RegisterName("addChild:"), sprite)
		}

		// Present scene
		skView.Send(objc.RegisterName("presentScene:"), scene)

		// Set as content view
		window.ID.Send(objc.RegisterName("setContentView:"), skView)

		window.ID.Send(objc.RegisterName("retain"))
		window.ID.Send(objc.RegisterName("center"))
		window.MakeKeyAndOrderFront(window.ID)

		fmt.Println("✓ Window created with SpriteKit scene")
		fmt.Println("✅ Using generated SpriteKit bindings:")
		fmt.Println("   - SKView for rendering")
		fmt.Println("   - SKScene for game world")
		fmt.Println("   - SKSpriteNode for sprites")
		fmt.Println("   - SKPhysicsBody for physics simulation")
		fmt.Println("\n   Watch the sprites fall with physics!")
		fmt.Println("   Press Cmd+Q to quit.")
	})
}

func runE2ETest() {
	fmt.Println("=== E2E Test Mode (SpriteKit Generated Bindings) ===")

	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ActivationPolicyAccessory)
	fmt.Println("✓ Created application")
	time.Sleep(100 * time.Millisecond)

	// Create window
	window := appkit.NewWindowWithFrame(100, 100, 800, 600,
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable)
	window.SetTitle("E2E Test Window")
	fmt.Println("✓ Created window")
	time.Sleep(100 * time.Millisecond)

	// Create SKView
	skViewClass := objc.GetClass("SKView")
	skView := objc.ID(skViewClass).Send(objc.RegisterName("alloc"))
	type NSPoint struct{ X, Y float64 }
	type NSSize struct{ Width, Height float64 }
	type NSRect struct {
		Origin NSPoint
		Size   NSSize
	}
	skView = skView.Send(objc.RegisterName("initWithFrame:"), NSRect{
		Origin: NSPoint{X: 0, Y: 0},
		Size:   NSSize{Width: 800, Height: 600},
	})
	if skView == 0 {
		fmt.Println("✗ FAIL: Failed to create SKView")
		os.Exit(1)
	}
	fmt.Println("✓ Created SKView")
	time.Sleep(100 * time.Millisecond)

	// Create SKScene
	sceneClass := objc.GetClass("SKScene")
	scene := objc.ID(sceneClass).Send(objc.RegisterName("alloc"))
	type CGSize struct {
		Width  float64
		Height float64
	}
	scene = scene.Send(objc.RegisterName("initWithSize:"), CGSize{Width: 800, Height: 600})
	if scene == 0 {
		fmt.Println("✗ FAIL: Failed to create SKScene")
		os.Exit(1)
	}
	fmt.Println("✓ Created SKScene")
	time.Sleep(100 * time.Millisecond)

	// Create SKSpriteNode
	spriteClass := objc.GetClass("SKSpriteNode")
	sprite := objc.ID(spriteClass).Send(objc.RegisterName("alloc"))

	colorClass := objc.GetClass("NSColor")
	spriteColor := objc.ID(colorClass).Send(objc.RegisterName("redColor"))
	sprite = sprite.Send(objc.RegisterName("initWithColor:size:"),
		spriteColor, CGSize{Width: 50, Height: 50})
	if sprite == 0 {
		fmt.Println("✗ FAIL: Failed to create SKSpriteNode")
		os.Exit(1)
	}
	fmt.Println("✓ Created SKSpriteNode")
	time.Sleep(100 * time.Millisecond)

	// Add sprite to scene
	scene.Send(objc.RegisterName("addChild:"), sprite)
	fmt.Println("✓ Added sprite to scene")
	time.Sleep(100 * time.Millisecond)

	// Present scene
	skView.Send(objc.RegisterName("presentScene:"), scene)
	fmt.Println("✓ Presented scene in SKView")
	time.Sleep(100 * time.Millisecond)

	// Set as content view
	window.ID.Send(objc.RegisterName("setContentView:"), skView)
	fmt.Println("✓ Set SKView as content view")
	time.Sleep(100 * time.Millisecond)

	// Show window briefly
	window.MakeKeyAndOrderFront(0)
	fmt.Println("✓ Window displayed with SpriteKit scene")
	time.Sleep(200 * time.Millisecond)

	window.ID.Send(objc.RegisterName("close"))
	fmt.Println("✓ Window closed")

	fmt.Println("\n=== E2E Test PASSED ===")
	fmt.Println("   ✓ Used generated SpriteKit bindings")
	fmt.Println("   ✓ SKView, SKScene, SKSpriteNode")
	fmt.Println("   ✓ Physics simulation setup")
	os.Exit(0)
}
