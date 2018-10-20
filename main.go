package vkb

import (
    "os"
    "github.com/indeedhat/vkb/layout"
)

type VirtKB struct {
    KeyStrokes []int
    Modifiers  map[int]bool
    Layout    *layout.Layout

    // this file descriptor is only used for linux but i want to keep it per event anyway
    Fd *os.File
}

// create a new keyboard event
func NewVKB() (*VirtKB, error) {
    e := &VirtKB{}
    err := e.Reset()

    return e, err
}

// clear the keystrokes to set the event back to new
func (e *VirtKB) Reset() error {
    e.KeyStrokes = []int{}
    err := initVKB(e)

    return err
}

// add  a single keystroke to the slice
func (e *VirtKB) AddStroke(stroke int) *VirtKB {
    e.KeyStrokes = append(e.KeyStrokes, stroke)

    return e
}

// Add keystrokes
func (e *VirtKB) AddStrokes(strokes []int) *VirtKB {
    e.KeyStrokes = append(e.KeyStrokes, strokes...)

    return e
}

// Replace Keystrokes
func (e *VirtKB) SetKeystrokes(strokes []int) *VirtKB {
    e.KeyStrokes = strokes

    return e
}

// check if the keycode belongs to a modifier key
func (e *VirtKB) isModifier(key int) bool {
    _, ok := e.Modifiers[key]
    return ok
}
