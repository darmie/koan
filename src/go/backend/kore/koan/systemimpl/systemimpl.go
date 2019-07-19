package systemimpl

// #include <kinc/system.h>
// #include <kinc/window.h>
// #include <kinc/display.h>
// #include <kinc/input/mouse.h>
// #include <kinc/input/pen.h>

import "C"
import (
	"unsafe"
)

// Todo: add platform specific CGO CFLAGS: -I./Kinc/Sources

var (
	// Needs3d needs 3D
	Needs3d            bool
	framebuffers       []interface{}
	keyboard           interface{}
	mouse              interface{}
	pen                interface{}
	gamepad1           interface{}
	gamepad2           interface{}
	gamepad3           interface{}
	gamepad4           interface{}
	surface            interface{}
	mouseLockListeners []interface{}
	// MouseX mouse X position
	MouseX int
	// MouseY mouse Y position
	MouseY int
)

// GetMouse get mouse
func GetMouse(num int) (interface{}, error) {
	if num != 0 {
		return nil, nil
	}

	return mouse, nil
}

// GetPen get pen
func GetPen(num int) (interface{}, error) {
	if num != 0 {
		return nil, nil
	}

	return pen, nil
}

// GetKeyboard get keyboard
func GetKeyboard(num int) (interface{}, error) {
	if num != 0 {
		return nil, nil
	}

	var keyboard interface{}

	return keyboard, nil
}

// GetTime get time
func GetTime() (float32, error) {
	time := C.double(C.kinc_time())

	if time != nil {
		return float32(time), nil
	}

	return nil, nil
}

// WindowWidth get window width
func WindowWidth(windowID int) int {
	width := C.int(C.kinc_window_width(C.int(windowID)))

	return int(width)
}

// WindowHeight get window height
func WindowHeight(windowID int) int {
	height := C.int(C.kinc_window_height(C.int(windowID)))
	return int(height)
}

// ScreenDPI get screen dpi
func ScreenDPI() int {
	display := C.int(C.kinc_primary_display())
	dpi := C.int(C.kinc_display_current_mode(display).pixels_per_inch)
	return int(dpi)
}

// GetVsync get vsync
func GetVsync() bool {
	return true
}

// GetRefreshRate refresh rate
func GetRefreshRate() int {
	return 60
}

//GetScreenRotation screen rotation
func GetScreenRotation() interface{} {
	return nil
}

// GetSystemID get system id
func GetSystemID() string {
	id := C.GoString(unsafe.Pointer(C.kinc_system_id()))
	return id
}

// Vibrate vibrate
func Vibrate(ms int) {
	C.kinc_vibrate(C.int(ms))
}

// GetLanguage get language
func GetLanguage() string {
	return C.GoString(C.kinc_language())
}

// RequestShutdown request shutdown
func RequestShutdown() bool {
	C.kinc_stop()
	return true
}

// Init initialize the system
func Init(options interface{}, callback func(interface{})) error {
	initKinc(options.title, options.width, options.height, options.window, options.framebuffer)

	return postInit(callback)
}

func initKinc(name string, width int, height int, winoptions interface{}, frame interface{}) {
	windops := (*C.kinc_window_options)(unsafe.Pointer(winoptions))
	framwops := (*C.kinc_framebuffer_options)(unsafe.Pointer(frame))
	C.kinc_init(C.char(name), C.int(width), C.int(height), windops, frame)
}

func postInit(callback func(interface{})) error {
	callback(nil)

	C.kinc_start()
}

// IsMouseLocked is mouse locked
func IsMouseLocked(windowID *int) bool {
	return C.kinc_mouse_is_locked() != 0
}

// LockMouse lock mouse
func LockMouse(windowID *int) {
	if windowID == nil {
		windowID = 0
	}
	if IsMouseLocked(windowID) {
		C.kinc_mouse_lock(C.int(windowID))
		for _, listener := range mouseLockListeners {
			listener()
		}
	}
}

// UnlockMouse unlock mouse
func UnlockMouse(windowID *int) {
	if windowID == nil {
		windowID = 0
	}
	if IsMouseLocked(windowID) {
		C.kinc_mouse_unlock(C.int(windowID))
		for _, listener := range mouseLockListeners {
			listener()
		}
	}
}

// CanLockMouse can lock mouse
func CanLockMouse(windowID int) bool {
	if windowID == nil {
		windowID = 0
	}
	return C.kinc_mouse_can_lock(C.int(windowID)) != 0
}

// NotifyOfMouseLockChange notify when mouse lock change
func NotifyOfMouseLockChange(fun func(), err func(), windowID *int) {
	if CanLockMouse(windowID) && fun != nil {
		mouseLockListeners = append(mouseLockListeners, fun)
	}
}

// RemoveFromMouseLockChange notify when mouse lock change
func RemoveFromMouseLockChange(fun func(), err func(), windowID *int) {
	if CanLockMouse(windowID) && fun != nil {
		for i, f := range mouseLockListeners {
			if mouseLockListeners[i] == fun {
				mouseLockListeners = append(mouseLockListeners[i:], mouseLockListeners[i+1:]...)
				break
			}
		}
	}
}

// HideSystemCursor hide cursor
func HideSystemCursor() {
	C.kinc_mouse_hide()
}

// ShowSystemCursor show cursor
func ShowSystemCursor() {
	C.kinc_mouse_show()
}

// Frame  framebuffer
func Frame() {}

// KeyDown key down event
func KeyDown(code int) {
	keyboard.sendDownEvent(code)
}

// KeyUp key up event
func KeyUp(code int) {
	keyboard.sendUpEvent(code)
}

// MouseDown mouse down event
func MouseDown(windowID int, button int, x int, y int) {
	MouseX = x
	MouseY = y

	mouse.SendDownEvent(windowID, button, x, y)
}

// MouseUp mouse up event
func MouseUp(windowID int, button int, x int, y int) {
	MouseX = x
	MouseY = y

	mouse.SendUpEvent(windowID, button, x, y)
}

// MouseMove mouse move event
func MouseMove(windowID int, x int, y int, movementX int, movementY int) {
	MouseX = x
	MouseY = y

	mouse.SendMoveEvent(windowID, x, y, movementX, movementX)
}

// MouseWheel mouse wheel event
func MouseWheel(windowID int, x int, y int, delta int) {
	MouseX = x
	MouseY = y

	mouse.SendWheelEvent(windowID, x, y, delta)
}

// MouseLeave mouse leave event
func MouseLeave(windowID int) {
	mouse.SendLeaveEvent(windowID)
}

// PenDown pen down event
func PenDown(windowID int, x int, y int, pressure int) {
	pen.sendDownEvent(windowID, x, y, pressure)
}

// PenUp pen up event
func PenUp(windowID int, x int, y int, pressure int) {
	pen.sendUpEvent(windowID, x, y, pressure)
}

// PenMove pen move event
func PenMove(windowID int, x int, y int, pressure int) {
	pen.sendMoveEvent(windowID, x, y, pressure)
}

// Gamepad1Axis gamepad1 axis event
func Gamepad1Axis(axis int, value float32) {
	gamepad1.sendAxisEvent(axis, value)
}

// Gamepad1Button gamepad1 button event
func Gamepad1Button(button int, value float32) {
	gamepad1.sendButtonEvent(button, value)
}

// Gamepad2Axis gamepad2 axis event
func Gamepad2Axis(axis int, value float32) {
	gamepad2.sendAxisEvent(axis, value)
}

// Gamepad2Button gamepad2 button event
func Gamepad2Button(button int, value float32) {
	gamepad2.sendButtonEvent(button, value)
}

// Gamepad3Axis gamepad3 axis event
func Gamepad3Axis(axis int, value float32) {
	gamepad3.sendAxisEvent(axis, value)
}

// Gamepad3Button gamepad3 button event
func Gamepad3Button(button int, value float32) {
	gamepad3.sendButtonEvent(button, value)
}

// Gamepad4Axis gamepad4 axis event
func Gamepad4Axis(axis int, value float32) {
	gamepad4.sendAxisEvent(axis, value)
}

// Gamepad4Button gamepad4 button event
func Gamepad4Button(button int, value float32) {
	gamepad4.sendButtonEvent(button, value)
}

// TouchStart touch start event
func TouchStart(index int, x int, y int) {
	surface.sendTouchStartEvent(index, x, y)
}

// TouchEnd touch end event
func TouchEnd(index int, x int, y int) {
	surface.sendTouchEndEvent(index, x, y)
}

// TouchMove touch move event
func TouchMove(index int, x int, y int) {
	surface.sendMoveEvent(index, x, y)
}
