package vkb

import (
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