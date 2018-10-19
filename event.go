package vkb

import (
    "os"
)

type Event struct {
    KeyStrokes []int
    Modifiers  map[int]bool

    // this file descriptor is only used for linux but i want to keep it per event anyway
    Fd *os.File
}

// create a new keyboard event
func NewEvent() (*Event, error) {
    e := &Event{}
    err := e.Reset()

    return e, err
}

// clear the keystrokes to set the event back to new
func (e *Event) Reset() error {
    e.KeyStrokes = []int{}
    err := initVKB(e)

    return err
}

// add  a single keystroke to the slice
func (e *Event) AddStroke(stroke int) *Event {
    e.KeyStrokes = append(e.KeyStrokes, stroke)

    return e
}

// Add keystrokes
func (e *Event) AddStrokes(strokes []int) *Event {
    e.KeyStrokes = append(e.KeyStrokes, strokes...)

    return e
}

// Replace Keystrokes
func (e *Event) SetKeystrokes(strokes []int) *Event {
    e.KeyStrokes = strokes

    return e
}

// check if the keycode belongs to a modifier key
func (e *Event) isModifier(key int) bool {
    _, ok := e.Modifiers[key]
    return ok
}
