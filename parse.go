package vkb

import (
    "errors"
    "github.com/indeedhat/vkb/layout"
    "github.com/indeedhat/vkb/keys"
)


// Parse a string of text into keystrokes based on the given keyboard layout
// As long as the keyboard layout has been created using the vkb/keys constants then this should work
// cross platform without issue
// NOTE: all layouts in vkb/layout are created this way
func (e *VirtKB) ParseString(layout layout.Layout, text string, ignoreMissing bool) (err error) {
    code      := 0
    shift     := false
    prevShift := false
    for i, char := range text {
        if code, shift, err = layout.GetCode(char); nil == err {
            // toggle shift if necessary
            if shift != prevShift {
                e.AddStroke(keys.MOD_SHIFT)
            }

            e.AddStroke(code)
        } else if !ignoreMissing {
            // remove keypresses already added by this parse operation
            e.KeyStrokes = e.KeyStrokes[:len(e.KeyStrokes) - i + 1]
            return
        }

        prevShift = shift
    }

    // release shift if the last character of the parsed text was a shift key
    // this will prevent any extra text added to the keyboard before its run from being in the wrong case
    if prevShift {
        e.AddStroke(keys.MOD_SHIFT)
    }

    return
}


// run the parser with the layout already assigned to the keyboard
// returns error if a layout is not set
func (e *VirtKB) ParseWithAssignedLayout(text string, ignoreMissing bool) (err error) {
    if nil == e.Layout {
        return errors.New("no layout assigned to vkb")
    }

    return e.ParseString(*e.Layout, text, ignoreMissing)
}


// assign a layout to the keyboard
func (e *VirtKB) AssignLayout(layout *layout.Layout) (*VirtKB) {
    e.Layout = layout

    return e
}