package keys

var keyMap = map[byte]string {
    // mods
    MOD_ALT:    "Alt",
    MOD_CRTL:    "Ctrl",
    MOD_SHIFT:   "Shft",
    MOD_COMMAND: "Cmd",

    // number row
    K_0: "0",
    K_1: "1",
    K_2: "2",
    K_3: "3",
    K_4: "4",
    K_5: "5",
    K_6: "6",
    K_7: "7",
    K_8: "8",
    K_9: "9",

    // alphabet
    K_A: "A",
    K_B: "B",
    K_C: "C",
    K_D: "D",
    K_E: "E",
    K_F: "F",
    K_G: "G",
    K_H: "H",
    K_I: "I",
    K_J: "J",
    K_K: "K",
    K_L: "L",
    K_M: "M",
    K_N: "N",
    K_O: "O",
    K_P: "P",
    K_Q: "Q",
    K_R: "R",
    K_S: "S",
    K_T: "T",
    K_U: "U",
    K_V: "V",
    K_W: "W",
    K_X: "X",
    K_Y: "Y",
    K_Z: "Z",

    // Misc
    SPACE:     "Spce",
    TAB:       "Tab",
    ESC:       "Esc",
    ENTER:     "Rtn",
    BACKSPACE: "BkSp",
    DELETE:    "Del",
    HOME:      "Hme",
    END:       "End",
    PAGE_DOWN: "PgDn",
    PAGE_UP:   "PgUp",

    // Arrows
    UP:    "Up",
    DOWN:  "Down",
    LEFT:  "Left",
    RIGHT: "Right",

    // F keys
    F1 : "F1",
    F2 : "F2",
    F3 : "F3",
    F4 : "F4",
    F5 : "F5",
    F6 : "F6",
    F7 : "F7",
    F8 : "F8",
    F9 : "F9",
    F10: "F10",
    F11: "F11",
    F12: "F12",
    F13: "F13",
    F14: "F14",
    F15: "F15",
    F16: "F16",
    F17: "F17",
    F18: "F18",
    F19: "F19",
    F20: "F20",
    F21: "F21",
    F22: "F22",
    F23: "F23",
    F24: "F24",

    // Symbols
    SLASH      : "/",
    BACKTICK   : "`",
    COMMA      : ",",
    EQUALS     : "=",
    DASH       : "-",
    BRACE_OPEN : "[",
    BRACE_CLOSE: "]",
    QUOTE      : "'",
    PERIOD     : ".",
    BACKSLASH  : "\\",
    SEMICOLON  : ";",

    // Media
    MEDIA_PLAY    : "Play",
    MEDIA_STOP    : "Stop",
    MEDIA_NEXT    : "Next Track",
    MEDIA_PREV    : "Prev Track",
    MEDIA_VOL_UP  : "Vol+",
    MEDIA_VOL_DOWN: "Vol-",
    MEDIA_MUTE    : "Mute",
}


// convert a key code to its string representation
func ToString(key byte) string {
    if char, ok := keyMap[key]; ok {
        return char
    }

    return ""
}

