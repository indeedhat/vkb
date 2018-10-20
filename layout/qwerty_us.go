package layout

import "github.com/indeedhat/vkb/keys"

func QwertyUs() Layout {
    return Layout{
        Keys: map[byte]int{
            '.':  keys.PERIOD,
            ',':  keys.COMMA,
            '=':  keys.EQUALS,
            '-':  keys.DASH,
            ';':  keys.SEMICOLON,
            '/':  keys.SLASH,
            '`':  keys.BACKTICK,
            '[':  keys.BRACE_OPEN,
            '\\': keys.BACKSLASH,
            ']':  keys.BRACE_CLOSE,
            '\'': keys.QUOTE,

            'a':  keys.K_A,
            'b':  keys.K_B,
            'c':  keys.K_C,
            'd':  keys.K_D,
            'e':  keys.K_E,
            'f':  keys.K_F,
            'g':  keys.K_G,
            'h':  keys.K_H,
            'i':  keys.K_I,
            'j':  keys.K_J,
            'k':  keys.K_K,
            'l':  keys.K_L,
            'm':  keys.K_M,
            'n':  keys.K_N,
            'o':  keys.K_O,
            'p':  keys.K_P,
            'q':  keys.K_Q,
            'r':  keys.K_R,
            's':  keys.K_S,
            't':  keys.K_T,
            'u':  keys.K_U,
            'v':  keys.K_V,
            'w':  keys.K_W,
            'x':  keys.K_X,
            'y':  keys.K_Y,
            'z':  keys.K_Z,

            '1':  keys.K_1,
            '2':  keys.K_2,
            '3':  keys.K_3,
            '4':  keys.K_4,
            '5':  keys.K_5,
            '6':  keys.K_6,
            '7':  keys.K_7,
            '8':  keys.K_8,
            '9':  keys.K_9,
            '0':  keys.K_0,

            ' ':  keys.SPACE,
            '\t': keys.TAB,
        },

        ShiftKeys: map[byte]int{
            '>': keys.PERIOD,
            '<': keys.COMMA,
            '_': keys.EQUALS,
            '+': keys.DASH,
            ':': keys.SEMICOLON,
            '?': keys.SLASH,
            '~': keys.BACKTICK,
            '{': keys.BRACE_OPEN,
            '|': keys.BACKSLASH,
            '}': keys.BRACE_CLOSE,
            '"': keys.QUOTE,

            'A': keys.K_A,
            'B': keys.K_B,
            'C': keys.K_C,
            'D': keys.K_D,
            'E': keys.K_E,
            'F': keys.K_F,
            'G': keys.K_G,
            'H': keys.K_H,
            'I': keys.K_I,
            'J': keys.K_J,
            'K': keys.K_K,
            'L': keys.K_L,
            'M': keys.K_M,
            'N': keys.K_N,
            'O': keys.K_O,
            'P': keys.K_P,
            'Q': keys.K_Q,
            'R': keys.K_R,
            'S': keys.K_S,
            'T': keys.K_T,
            'U': keys.K_U,
            'V': keys.K_V,
            'W': keys.K_W,
            'X': keys.K_X,
            'Y': keys.K_Y,
            'Z': keys.K_Z,

            '!': keys.K_1,
            '@': keys.K_2,
            '#': keys.K_3,
            '$': keys.K_4,
            '%': keys.K_5,
            '^': keys.K_6,
            '&': keys.K_7,
            '*': keys.K_8,
            '(': keys.K_9,
            ')': keys.K_0,
        },
    }
}

