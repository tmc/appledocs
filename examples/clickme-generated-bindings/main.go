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
	"github.com/tmc/appledocs/generated/coregraphics"
	localobjc "github.com/tmc/appledocs/generated/objc"
)

var (
	e2e          = flag.Bool("e2e", false, "run end-to-end tests")
	clickCount   int
	counterLabel appkit.TextField
)

func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()

	app := appkit.ApplicationFrom(appkit.ApplicationClass.SharedApplication())
	app.SetActivationPolicy(unsafe.Pointer(uintptr(0)))

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 100, Y: 100}, Size: coregraphics.CGSize{Width: 400, Height: 300}},
		appkit.WindowStyleMask(1|2), appkit.BackingStoreType(2), false)
	window.SetTitle("Hello from Generated Bindings!")
	view := window.ContentView()

	counterLabel = appkit.NewTextField()
	counterLabel.SetFrame(coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 100, Y: 200}, Size: coregraphics.CGSize{Width: 200, Height: 40}})
	counterLabel.SetStringValue("Clicks: 0")
	counterLabel.SetEditable(false)
	counterLabel.SetBezeled(false)
	counterLabel.SetDrawsBackground(false)
	appkit.ControlFrom(unsafe.Pointer(counterLabel.ID)).SetAlignment(unsafe.Pointer(uintptr(2)))
	appkit.ViewFrom(view).AddSubview(unsafe.Pointer(counterLabel.ID))

	handler := createButtonHandler(*e2e)
	button := appkit.NewButtonWithTitleTargetAction("Click Me!", localobjc.ID(handler), localobjc.Sel("buttonClicked:"))
	button.SetFrame(coregraphics.CGRect{Origin: coregraphics.CGPoint{X: 150, Y: 120}, Size: coregraphics.CGSize{Width: 100, Height: 40}})
	appkit.ViewFrom(view).AddSubview(unsafe.Pointer(button.ID))

	window.MakeKeyAndOrderFront(localobjc.ID(0))
	app.ActivateIgnoringOtherApps(true)

	if *e2e {
		go func() {
			time.Sleep(100 * time.Millisecond)
			button.PerformClick(unsafe.Pointer(uintptr(0)))
			time.Sleep(100 * time.Millisecond)
			button.PerformClick(unsafe.Pointer(uintptr(0)))
			time.Sleep(100 * time.Millisecond)
			if clickCount == 2 && counterLabel.StringValue() == "Clicks: 2" {
				fmt.Println("=== E2E PASSED ===")
			} else {
				fmt.Printf("=== E2E FAILED: count=%d label=%s ===\n", clickCount, counterLabel.StringValue())
			}
			os.Exit(0)
		}()
	}

	app.Run()
}

func createButtonHandler(isE2E bool) localobjc.ID {
	if class := localobjc.GetClass("ButtonHandler"); class == 0 {
		clicked := func(self localobjc.ID, _cmd localobjc.SEL, sender localobjc.ID) {
			clickCount++
			counterLabel.SetStringValue(fmt.Sprintf("Clicks: %d", clickCount))
			if !isE2E {
				fmt.Printf("Clicked! Count: %d\n", clickCount)
			}
		}
		localobjc.RegisterClass("ButtonHandler", localobjc.GetClass("NSObject"), nil, nil,
			[]localobjc.MethodDef{{Cmd: localobjc.Sel("buttonClicked:"), Fn: clicked}})
	}
	return localobjc.Send[localobjc.ID](localobjc.Send[localobjc.ID](localobjc.ID(localobjc.GetClass("ButtonHandler")), localobjc.Sel("alloc")), localobjc.Sel("init"))
}
