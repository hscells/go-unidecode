package table

var x018 = [][]uint8{
	[]uint8(" @ "),   // 0x00
	[]uint8(" ... "), // 0x01
	[]uint8(", "),    // 0x02
	[]uint8(". "),    // 0x03
	[]uint8(": "),    // 0x04
	[]uint8(" // "),  // 0x05
	[]uint8(""),      // 0x06
	[]uint8("-"),     // 0x07
	[]uint8(", "),    // 0x08
	[]uint8(". "),    // 0x09
	[]uint8(""),      // 0x0a
	[]uint8(""),      // 0x0b
	[]uint8(""),      // 0x0c
	[]uint8(""),      // 0x0d
	[]uint8(""),      // 0x0e
	[]uint8("[?]"),   // 0x0f
	[]uint8("0"),     // 0x10
	[]uint8("1"),     // 0x11
	[]uint8("2"),     // 0x12
	[]uint8("3"),     // 0x13
	[]uint8("4"),     // 0x14
	[]uint8("5"),     // 0x15
	[]uint8("6"),     // 0x16
	[]uint8("7"),     // 0x17
	[]uint8("8"),     // 0x18
	[]uint8("9"),     // 0x19
	[]uint8("[?]"),   // 0x1a
	[]uint8("[?]"),   // 0x1b
	[]uint8("[?]"),   // 0x1c
	[]uint8("[?]"),   // 0x1d
	[]uint8("[?]"),   // 0x1e
	[]uint8("[?]"),   // 0x1f
	[]uint8("a"),     // 0x20
	[]uint8("e"),     // 0x21
	[]uint8("i"),     // 0x22
	[]uint8("o"),     // 0x23
	[]uint8("u"),     // 0x24
	[]uint8("O"),     // 0x25
	[]uint8("U"),     // 0x26
	[]uint8("ee"),    // 0x27
	[]uint8("n"),     // 0x28
	[]uint8("ng"),    // 0x29
	[]uint8("b"),     // 0x2a
	[]uint8("p"),     // 0x2b
	[]uint8("q"),     // 0x2c
	[]uint8("g"),     // 0x2d
	[]uint8("m"),     // 0x2e
	[]uint8("l"),     // 0x2f
	[]uint8("s"),     // 0x30
	[]uint8("sh"),    // 0x31
	[]uint8("t"),     // 0x32
	[]uint8("d"),     // 0x33
	[]uint8("ch"),    // 0x34
	[]uint8("j"),     // 0x35
	[]uint8("y"),     // 0x36
	[]uint8("r"),     // 0x37
	[]uint8("w"),     // 0x38
	[]uint8("f"),     // 0x39
	[]uint8("k"),     // 0x3a
	[]uint8("kha"),   // 0x3b
	[]uint8("ts"),    // 0x3c
	[]uint8("z"),     // 0x3d
	[]uint8("h"),     // 0x3e
	[]uint8("zr"),    // 0x3f
	[]uint8("lh"),    // 0x40
	[]uint8("zh"),    // 0x41
	[]uint8("ch"),    // 0x42
	[]uint8("-"),     // 0x43
	[]uint8("e"),     // 0x44
	[]uint8("i"),     // 0x45
	[]uint8("o"),     // 0x46
	[]uint8("u"),     // 0x47
	[]uint8("O"),     // 0x48
	[]uint8("U"),     // 0x49
	[]uint8("ng"),    // 0x4a
	[]uint8("b"),     // 0x4b
	[]uint8("p"),     // 0x4c
	[]uint8("q"),     // 0x4d
	[]uint8("g"),     // 0x4e
	[]uint8("m"),     // 0x4f
	[]uint8("t"),     // 0x50
	[]uint8("d"),     // 0x51
	[]uint8("ch"),    // 0x52
	[]uint8("j"),     // 0x53
	[]uint8("ts"),    // 0x54
	[]uint8("y"),     // 0x55
	[]uint8("w"),     // 0x56
	[]uint8("k"),     // 0x57
	[]uint8("g"),     // 0x58
	[]uint8("h"),     // 0x59
	[]uint8("jy"),    // 0x5a
	[]uint8("ny"),    // 0x5b
	[]uint8("dz"),    // 0x5c
	[]uint8("e"),     // 0x5d
	[]uint8("i"),     // 0x5e
	[]uint8("iy"),    // 0x5f
	[]uint8("U"),     // 0x60
	[]uint8("u"),     // 0x61
	[]uint8("ng"),    // 0x62
	[]uint8("k"),     // 0x63
	[]uint8("g"),     // 0x64
	[]uint8("h"),     // 0x65
	[]uint8("p"),     // 0x66
	[]uint8("sh"),    // 0x67
	[]uint8("t"),     // 0x68
	[]uint8("d"),     // 0x69
	[]uint8("j"),     // 0x6a
	[]uint8("f"),     // 0x6b
	[]uint8("g"),     // 0x6c
	[]uint8("h"),     // 0x6d
	[]uint8("ts"),    // 0x6e
	[]uint8("z"),     // 0x6f
	[]uint8("r"),     // 0x70
	[]uint8("ch"),    // 0x71
	[]uint8("zh"),    // 0x72
	[]uint8("i"),     // 0x73
	[]uint8("k"),     // 0x74
	[]uint8("r"),     // 0x75
	[]uint8("f"),     // 0x76
	[]uint8("zh"),    // 0x77
	[]uint8("[?]"),   // 0x78
	[]uint8("[?]"),   // 0x79
	[]uint8("[?]"),   // 0x7a
	[]uint8("[?]"),   // 0x7b
	[]uint8("[?]"),   // 0x7c
	[]uint8("[?]"),   // 0x7d
	[]uint8("[?]"),   // 0x7e
	[]uint8("[?]"),   // 0x7f
	[]uint8("[?]"),   // 0x80
	[]uint8("H"),     // 0x81
	[]uint8("X"),     // 0x82
	[]uint8("W"),     // 0x83
	[]uint8("M"),     // 0x84
	[]uint8(" 3 "),   // 0x85
	[]uint8(" 333 "), // 0x86
	[]uint8("a"),     // 0x87
	[]uint8("i"),     // 0x88
	[]uint8("k"),     // 0x89
	[]uint8("ng"),    // 0x8a
	[]uint8("c"),     // 0x8b
	[]uint8("tt"),    // 0x8c
	[]uint8("tth"),   // 0x8d
	[]uint8("dd"),    // 0x8e
	[]uint8("nn"),    // 0x8f
	[]uint8("t"),     // 0x90
	[]uint8("d"),     // 0x91
	[]uint8("p"),     // 0x92
	[]uint8("ph"),    // 0x93
	[]uint8("ss"),    // 0x94
	[]uint8("zh"),    // 0x95
	[]uint8("z"),     // 0x96
	[]uint8("a"),     // 0x97
	[]uint8("t"),     // 0x98
	[]uint8("zh"),    // 0x99
	[]uint8("gh"),    // 0x9a
	[]uint8("ng"),    // 0x9b
	[]uint8("c"),     // 0x9c
	[]uint8("jh"),    // 0x9d
	[]uint8("tta"),   // 0x9e
	[]uint8("ddh"),   // 0x9f
	[]uint8("t"),     // 0xa0
	[]uint8("dh"),    // 0xa1
	[]uint8("ss"),    // 0xa2
	[]uint8("cy"),    // 0xa3
	[]uint8("zh"),    // 0xa4
	[]uint8("z"),     // 0xa5
	[]uint8("u"),     // 0xa6
	[]uint8("y"),     // 0xa7
	[]uint8("bh"),    // 0xa8
	[]uint8("'"),     // 0xa9
	[]uint8("[?]"),   // 0xaa
	[]uint8("[?]"),   // 0xab
	[]uint8("[?]"),   // 0xac
	[]uint8("[?]"),   // 0xad
	[]uint8("[?]"),   // 0xae
	[]uint8("[?]"),   // 0xaf
	[]uint8("[?]"),   // 0xb0
	[]uint8("[?]"),   // 0xb1
	[]uint8("[?]"),   // 0xb2
	[]uint8("[?]"),   // 0xb3
	[]uint8("[?]"),   // 0xb4
	[]uint8("[?]"),   // 0xb5
	[]uint8("[?]"),   // 0xb6
	[]uint8("[?]"),   // 0xb7
	[]uint8("[?]"),   // 0xb8
	[]uint8("[?]"),   // 0xb9
	[]uint8("[?]"),   // 0xba
	[]uint8("[?]"),   // 0xbb
	[]uint8("[?]"),   // 0xbc
	[]uint8("[?]"),   // 0xbd
	[]uint8("[?]"),   // 0xbe
	[]uint8("[?]"),   // 0xbf
	[]uint8("[?]"),   // 0xc0
	[]uint8("[?]"),   // 0xc1
	[]uint8("[?]"),   // 0xc2
	[]uint8("[?]"),   // 0xc3
	[]uint8("[?]"),   // 0xc4
	[]uint8("[?]"),   // 0xc5
	[]uint8("[?]"),   // 0xc6
	[]uint8("[?]"),   // 0xc7
	[]uint8("[?]"),   // 0xc8
	[]uint8("[?]"),   // 0xc9
	[]uint8("[?]"),   // 0xca
	[]uint8("[?]"),   // 0xcb
	[]uint8("[?]"),   // 0xcc
	[]uint8("[?]"),   // 0xcd
	[]uint8("[?]"),   // 0xce
	[]uint8("[?]"),   // 0xcf
	[]uint8("[?]"),   // 0xd0
	[]uint8("[?]"),   // 0xd1
	[]uint8("[?]"),   // 0xd2
	[]uint8("[?]"),   // 0xd3
	[]uint8("[?]"),   // 0xd4
	[]uint8("[?]"),   // 0xd5
	[]uint8("[?]"),   // 0xd6
	[]uint8("[?]"),   // 0xd7
	[]uint8("[?]"),   // 0xd8
	[]uint8("[?]"),   // 0xd9
	[]uint8("[?]"),   // 0xda
	[]uint8("[?]"),   // 0xdb
	[]uint8("[?]"),   // 0xdc
	[]uint8("[?]"),   // 0xdd
	[]uint8("[?]"),   // 0xde
	[]uint8("[?]"),   // 0xdf
	[]uint8("[?]"),   // 0xe0
	[]uint8("[?]"),   // 0xe1
	[]uint8("[?]"),   // 0xe2
	[]uint8("[?]"),   // 0xe3
	[]uint8("[?]"),   // 0xe4
	[]uint8("[?]"),   // 0xe5
	[]uint8("[?]"),   // 0xe6
	[]uint8("[?]"),   // 0xe7
	[]uint8("[?]"),   // 0xe8
	[]uint8("[?]"),   // 0xe9
	[]uint8("[?]"),   // 0xea
	[]uint8("[?]"),   // 0xeb
	[]uint8("[?]"),   // 0xec
	[]uint8("[?]"),   // 0xed
	[]uint8("[?]"),   // 0xee
	[]uint8("[?]"),   // 0xef
	[]uint8("[?]"),   // 0xf0
	[]uint8("[?]"),   // 0xf1
	[]uint8("[?]"),   // 0xf2
	[]uint8("[?]"),   // 0xf3
	[]uint8("[?]"),   // 0xf4
	[]uint8("[?]"),   // 0xf5
	[]uint8("[?]"),   // 0xf6
	[]uint8("[?]"),   // 0xf7
	[]uint8("[?]"),   // 0xf8
	[]uint8("[?]"),   // 0xf9
	[]uint8("[?]"),   // 0xfa
	[]uint8("[?]"),   // 0xfb
	[]uint8("[?]"),   // 0xfc
	[]uint8("[?]"),   // 0xfd
	[]uint8("[?]"),   // 0xfe
}
