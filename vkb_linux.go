package vkb

/*
 #include <linux/uinput.h>
*/
import "C"
import (
    "encoding/binary"
    "os"
)

type uinput_user_dev C.struct_uinput_user_dev
type timeval C.struct_timeval
type input_event C.struct_input_event

var fd *os.File

const (
    EV_KEY               = C.EV_KEY
    EV_SYN               = C.EV_SYN
    UI_SET_EVBIT         = C.UI_SET_EVBIT
    UI_SET_KEYBIT        = C.UI_SET_KEYBIT
    UI_DEV_CREATE        = C.UI_DEV_CREATE
    UI_DEV_DESTROY       = C.UI_DEV_DESTROY
    UINPUT_MAX_NAME_SIZE = C.UINPUT_MAX_NAME_SIZE

)

// execute the keystrokes
func (e *Event) Execute() (err error){
    for _, key := range e.KeyStrokes {
       if e.isModifier(key) {
           if e.Modifiers[key] {
               if err = e.release(key); nil != err {
                   return
               }
               e.Modifiers[key] = false
           } else {
               if err = e.press(key); nil != err {
                   return
               }
               e.Modifiers[key] = true
           }
       } else {
           if err = e.tap(key); nil != err {
               return
           }
       }
       if err = e.sync(); nil != err {
           return
       }
    }

    err = e.ReleaseMods()
    return
}

// release any modifier keys that are still being pressed
func (e *Event) ReleaseMods() error {
    for key, state := range e.Modifiers {
       if state {
           if err := e.release(key); nil != err {
               return
           }
           e.Modifiers[key] = false
       }
    }

    return
}

// press down a key
func (e *Event) press(key int) error {
    ev := input_event{
        _type: EV_KEY,
        code:  c.__u16(key),
        value: 1,
    }

    return binary.Write(e.Fd, binary.LittleEndian, &ev)
}

// release a key
func (e *Event) release(key int) error {
    ev := input_event{
        _type: EV_KEY,
        code:  c.__u16(key),
        value: 0,
    }

    return binary.Write(e.Fd, binary.LittleEndian, &ev)
}

func (e *Event) sync() error {
    ev := input_event{
        _type: EV_SYN,
        code:  0,
        value: 0,
    }

    return binary.Write(e.Fd, binary.LittleEndian, &ev)
}

// tap a key
func (e *Event) tap(key int) error {
    if err := e.press(key); nil != err {
        return err
    }

    return e.release(key)
}

func initVKB(e *Event) (err error) {
    // no need to re init
    if nil != e.Fd {
        return
    }

    e.Modifiers = map[int]bool{
        K_LEFTCTRL:   false,
        K_RIGHTCTRL:  false,
        K_CTRL:       false,
        K_LEFTSHIFT:  false,
        K_RIGHTSHIFT: false,
        K_SHIFT:      false,
        K_LEFTALT:    false,
        K_RIGHTALT:   false,
        K_ALT:        false,
    }

    path, err := uinputPath()
    if nil != err {
        return
    }

    if e.Fd, err = os.OpenFile(path, os.O_WRONLY|syscall.O_NONBLOCK, os.ModeDevice); nil != err {
        return
    }

    if _, _, ec := syscall.Syscall(syscall.SYS_IOCTL, e.Fd.Fd(), UI_SET_EVBIT, uintptr(EV_KEY)); 0 != ec {
        return
    }

    if  _, _, ec := syscall.Syscall(syscall.SYS_IOCTL, e.Fd.Fd(), UI_SET_EVBIT, uintptr(EV_SYN)); 0 != ec {
        return
    }

    registerKeyBits(e)
    device := uinput_user_dev{}
    for i, c := range "VKB" {
        device.name[i] = C.char(c)
    }

    device.id.bustype = C.BUS_USB
    device.id.vendor  = 0x1
    device.id.product = 0x1
    device.id.version = 0x1

    if err = binary.Write(e.Fd, binary.LittleEndian, &device); nil != err {
        return
    }

    _, _, ec := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), _UI_DEV_CREATE, 0);

    return
}

func registerKeyBits(e *Event) (ec error) {
    for i := 0; i < 256; i++ {
        if _, _, ec = syscall.Syscall(syscall.SYS_IOCTL, e.Fd.Fd(), UI_SET_KEYBIT, uintptr(i)); ec != 0 {
            return
        }
    }

    return
}

// TODO: find a nice way of calling this before close
func closeDevice(e *Event) {
    syscall.Syscall(syscall.SYS_IOCTL, e.Fd.Fd(), UI_DEV_DESTROY, 0)
    e.Fd.Close()
}

// find and return the proper path for the uinput file descriptor
func uinputPath() (string, error) {
    for _, path := range []string{"/dev/uinput", "/dev/input/uinput"} {
        if _, err := os.Stat(path); err == nil {
            return path, nil
        }
    }

    return "", errors.New("No uinput file found.\n Try this cmd 'sudo modprobe uinput'")
}
