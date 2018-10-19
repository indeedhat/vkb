package vkb

import (
    "syscall"
)

const (
    KEYUP = 0x0002
)

// initialize the u32 lib
var dll    = syscall.NewLazyDLL("user32.dll")
var virtKb = dll.NewProc("keybd_event")

// execute the keystrokes
func (e *Event) Execute() (err error) {
    for _, key := range e.KeyStrokes {
        if e.isModifier(key) {
            if e.Modifiers[key] {
                release(key)
                e.Modifiers[key] = false
            } else {
                press(key)
                e.Modifiers[key] = true
            }
        } else {
            tap(key)
        }
    }

    e.ReleaseMods()
    return nil
}

// release any modifier keys that are still being pressed
func (e *Event) ReleaseMods() error {
    for key, state := range e.Modifiers {
        if state {
            release(key)
            e.Modifiers[key] = false
        }
    }

    return nil
}

// initialize the keyboard for the platform
// for windows nothing needs to be initialised so just set the mod map
func initVKB(e *Event) error {
    e.Modifiers = map[int]bool{
        K_SHIFT:    false,
        K_CTRL:     false,
        K_ALT:      false,
        K_LSHIFT:   false,
        K_RSHIFT:   false,
        K_LCONTROL: false,
        K_RCONTROL: false,
    }

    return nil
}

// press down a key
func press(key int) {
    vkey := key + 0x80
    virtKb.Call(uintptr(key), uintptr(vkey), 0, 0)
}

// release a key
func release(key int) {
    vkey := key + 0x80
    virtKb.Call(uintptr(key), uintptr(vkey), KEYUP, 0)
}

// tap a key
func tap(key int) {
    press(key)
    release(key)
}