package layout

import (
    "errors"
)

type Layout struct {
    Keys      map[byte]int
    ShiftKeys map[byte]int
}


// get the char code for the given character
func (l *Layout) GetCode(char byte) (code int, isShift bool, err error) {
    var ok bool
    if code, ok = l.Keys[char]; ok {
        return
    }
    if code, ok = l.ShiftKeys[char]; ok {
        isShift = true
        return
    }

    err = errors.New("char not found in layout")
    return
}


// check if a character is on the shift layer
func (l *Layout) IsShifCode(char byte) (exists bool) {
    _, exists = l.ShiftKeys[char]
    return
}


// check if a character exists in the chosen layout
func (l *Layout) CharExists(char byte) (exists bool) {
    if _, exists = l.Keys[char]; exists {
        return
    }

    _, exists = l.ShiftKeys[char]
    return
}

