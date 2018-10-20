package vkb

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa
 #import <Foundation/Foundation.h>
 CGEventRef Create(int k){
	CGEventRef event = CGEventCreateKeyboardEvent (NULL, (CGKeyCode)k, true);
	return event;
 }
 void KeyTap(CGEventRef event){
	CGEventPost(kCGAnnotatedSessionEventTap, event);
	CFRelease(event);
 }
 void AddActionKey(CGEventFlags type,CGEventRef event){
 	CGEventSetFlags(event, type);
 }
*/
import "C"


const (
    AShift    = C.kCGEventFlagMaskAlphaShift
    VK_SHIFT  = C.kCGEventFlagMaskShift
    VK_CTRL   = C.kCGEventFlagMaskControl
    VK_ALT    = C.kCGEventFlagMaskAlternate
    VK_CMD    = C.kCGEventFlagMaskCommand
    Help      = C.kCGEventFlagMaskHelp
    VK_FN     = C.kCGEventFlagMaskSecondaryFn
    NumPad    = C.kCGEventFlagMaskNumericPad
    Coalesced = C.kCGEventFlagMaskNonCoalesced
)


var K2VK = map[int]int{
    K_Control:      VK_CTRL,
    K_RightShift:   VK_SHIFT,
    K_RightControl: VK_CTRL,
    K_Command:      VK_CMD,
    K_Shift:        VK_SHIFT,
    VK_ALT:         VK_ALT,
}

// apply the keystrokes
// make the virtual keyboard start typing
func (e *VirtKB) Type() (err error) {
    for _, key := range e.KeyStrokes {
        if e.isModifier(key) {
            e.toggleMod(key)
        } else {
            e.tap(key)
        }
    }

    return e.ReleaseMods()
}


// release any modifier keys that are still being pressed
func (e *VirtKB) ReleaseMods() error {
    initVKB(e)
    return nil
}


// close any open connection
func (e *VirtKB) Close() {
   // this does nothing on mac it is only here to keep the api consistent
}


// initialize the keyboard for the platform
func initVKB(e *VirtKB) (err error) {
    e.Modifiers = map[int]bool {
        K_Control:      false,
        K_RightShift:   false,
        K_RightControl: false,
        K_Command:      false,
        K_Shift:        false,
        VK_ALT:         false,
    }
    return
}


// set the modifier state by key code
func (e *VirtKB) toggleMod(key) {
    e.Modifiers[key] = !e.Modifiers[key]
}


// apply any active modifiers to the keystroke event
func (e *VirtKB) applyMods(event C.CGEventRef) {
    for mod, on := range e.Modifiers {
        if on {
            C.AddActionKey(K2VK[mod], event)
        }
    }
}


// tap a key
func (e *VirtKB) tap(key int) {
    event := C.Create(C.int(key))
    e.applyMods(event)

    C.KeyTap(event)
}