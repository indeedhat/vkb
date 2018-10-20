package keys

const (
    // mods
    MOD_ALT     = 0x3A
    MOD_CRTL    = 0x3B
    MOD_SHIFT   = 0x38
    MOD_COMMAND = 0x37

    // number row
    K_0 = 0x1D
    K_1 = 0x12
    K_2 = 0x13
    K_3 = 0x14
    K_4 = 0x15
    K_5 = 0x17
    K_6 = 0x16
    K_7 = 0x1A
    K_8 = 0x1C
    K_9 = 0x19

    // alphabet
    K_A = 0x00
    K_B = 0x0B
    K_C = 0x08
    K_D = 0x02
    K_E = 0x0E
    K_F = 0x03
    K_G = 0x05
    K_H = 0x04
    K_I = 0x22
    K_J = 0x26
    K_K = 0x28
    K_L = 0x25
    K_M = 0x2E
    K_N = 0x2D
    K_O = 0x1F
    K_P = 0x23
    K_Q = 0x0C
    K_R = 0x0F
    K_S = 0x01
    K_T = 0x11
    K_U = 0x20
    K_V = 0x09
    K_W = 0x0D
    K_X = 0x07
    K_Y = 0x10
    K_Z = 0x06

    // Misc
    SPACE     = 0x31
    TAB       = 0x30
    ESC       = 0x35
    ENTER     = 0x24
    BACKSPACE = 0x33
    DELETE    = 0x75
    HOME      = 0x73
    END       = 0x77
    PAGE_DOWN = 0x79
    PAGE_UP   = 0x74

    // Arrows
    UP    = 0x7E
    DOWN  = 0x7D
    LEFT  = 0x7B
    RIGHT = 0x7C

    // F keys
    F1  = 0x7A
    F2  = 0x78
    F3  = 0x63
    F4  = 0x76
    F5  = 0x60
    F6  = 0x61
    F7  = 0x62
    F8  = 0x64
    F9  = 0x65
    F10 = 0x6D
    F11 = 0x67
    F12 = 0x6F
    F13 = 0x69
    F14 = 0x6B
    F15 = 0x71
    F16 = 0x6A
    F17 = 0x40
    F18 = 0x4F
    F19 = 0x50
    F20 = 0x5A /* there are no f21+ so this is what we got */
    F21 = 0x5A /* there are no f21+ so this is what we got */
    F22 = 0x5A /* there are no f21+ so this is what we got */
    F23 = 0x5A /* there are no f21+ so this is what we got */
    F24 = 0x5A /* there are no f21+ so this is what we got */

    // Symbols
    SLASH       = 0x2C
    BACKTICK    = 0x32
    COMMA       = 0x2B
    EQUALS      = 0x18
    DASH        = 0x1B
    BRACE_OPEN  = 0x21
    BRACE_CLOSE = 0x1E
    QUOTE       = 0x27
    PERIOD      = 0x2F
    BACKSLASH   = 0x2A
    SEMICOLON   = 0x29

    // Media
    MEDIA_PLAY     = 0x77 /* mac doesn't have media keys so send a end key (seems least destructive) */
    MEDIA_STOP     = 0x77 /* mac doesn't have media keys so send a end key (seems least destructive) */
    MEDIA_NEXT     = 0x77 /* mac doesn't have media keys so send a end key (seems least destructive) */
    MEDIA_PREV     = 0x77 /* mac doesn't have media keys so send a end key (seems least destructive) */
    MEDIA_VOL_UP   = 0x48
    MEDIA_VOL_DOWN = 0x49
    MEDIA_MUTE     = 0x4A
)