package table

var x030 = [][]uint8{
	[]uint8(" "),     // 0x00
	[]uint8(", "),    // 0x01
	[]uint8(". "),    // 0x02
	[]uint8("\""),    // 0x03
	[]uint8("[JIS]"), // 0x04
	[]uint8("\""),    // 0x05
	[]uint8("/"),     // 0x06
	[]uint8("0"),     // 0x07
	[]uint8("<"),     // 0x08
	[]uint8("> "),    // 0x09
	[]uint8("<<"),    // 0x0a
	[]uint8(">> "),   // 0x0b
	[]uint8("["),     // 0x0c
	[]uint8("] "),    // 0x0d
	[]uint8("{"),     // 0x0e
	[]uint8("} "),    // 0x0f
	[]uint8("[("),    // 0x10
	[]uint8(")] "),   // 0x11
	[]uint8("@"),     // 0x12
	[]uint8("X "),    // 0x13
	[]uint8("["),     // 0x14
	[]uint8("] "),    // 0x15
	[]uint8("[["),    // 0x16
	[]uint8("]] "),   // 0x17
	[]uint8("(("),    // 0x18
	[]uint8(")) "),   // 0x19
	[]uint8("[["),    // 0x1a
	[]uint8("]] "),   // 0x1b
	[]uint8("~ "),    // 0x1c
	[]uint8("``"),    // 0x1d
	[]uint8("''"),    // 0x1e
	[]uint8(",,"),    // 0x1f
	[]uint8("@"),     // 0x20
	[]uint8("1"),     // 0x21
	[]uint8("2"),     // 0x22
	[]uint8("3"),     // 0x23
	[]uint8("4"),     // 0x24
	[]uint8("5"),     // 0x25
	[]uint8("6"),     // 0x26
	[]uint8("7"),     // 0x27
	[]uint8("8"),     // 0x28
	[]uint8("9"),     // 0x29
	[]uint8(""),      // 0x2a
	[]uint8(""),      // 0x2b
	[]uint8(""),      // 0x2c
	[]uint8(""),      // 0x2d
	[]uint8(""),      // 0x2e
	[]uint8(""),      // 0x2f
	[]uint8("~"),     // 0x30
	[]uint8("+"),     // 0x31
	[]uint8("+"),     // 0x32
	[]uint8("+"),     // 0x33
	[]uint8("+"),     // 0x34
	[]uint8(""),      // 0x35
	[]uint8("@"),     // 0x36
	[]uint8(" // "),  // 0x37
	[]uint8("+10+"),  // 0x38
	[]uint8("+20+"),  // 0x39
	[]uint8("+30+"),  // 0x3a
	[]uint8("[?]"),   // 0x3b
	[]uint8("[?]"),   // 0x3c
	[]uint8("[?]"),   // 0x3d
	[]uint8(""),      // 0x3e
	[]uint8(""),      // 0x3f
	[]uint8("[?]"),   // 0x40
	[]uint8("a"),     // 0x41
	[]uint8("a"),     // 0x42
	[]uint8("i"),     // 0x43
	[]uint8("i"),     // 0x44
	[]uint8("u"),     // 0x45
	[]uint8("u"),     // 0x46
	[]uint8("e"),     // 0x47
	[]uint8("e"),     // 0x48
	[]uint8("o"),     // 0x49
	[]uint8("o"),     // 0x4a
	[]uint8("ka"),    // 0x4b
	[]uint8("ga"),    // 0x4c
	[]uint8("ki"),    // 0x4d
	[]uint8("gi"),    // 0x4e
	[]uint8("ku"),    // 0x4f
	[]uint8("gu"),    // 0x50
	[]uint8("ke"),    // 0x51
	[]uint8("ge"),    // 0x52
	[]uint8("ko"),    // 0x53
	[]uint8("go"),    // 0x54
	[]uint8("sa"),    // 0x55
	[]uint8("za"),    // 0x56
	[]uint8("shi"),   // 0x57
	[]uint8("zi"),    // 0x58
	[]uint8("su"),    // 0x59
	[]uint8("zu"),    // 0x5a
	[]uint8("se"),    // 0x5b
	[]uint8("ze"),    // 0x5c
	[]uint8("so"),    // 0x5d
	[]uint8("zo"),    // 0x5e
	[]uint8("ta"),    // 0x5f
	[]uint8("da"),    // 0x60
	[]uint8("chi"),   // 0x61
	[]uint8("di"),    // 0x62
	[]uint8("tsu"),   // 0x63
	[]uint8("tsu"),   // 0x64
	[]uint8("du"),    // 0x65
	[]uint8("te"),    // 0x66
	[]uint8("de"),    // 0x67
	[]uint8("to"),    // 0x68
	[]uint8("do"),    // 0x69
	[]uint8("na"),    // 0x6a
	[]uint8("ni"),    // 0x6b
	[]uint8("nu"),    // 0x6c
	[]uint8("ne"),    // 0x6d
	[]uint8("no"),    // 0x6e
	[]uint8("ha"),    // 0x6f
	[]uint8("ba"),    // 0x70
	[]uint8("pa"),    // 0x71
	[]uint8("hi"),    // 0x72
	[]uint8("bi"),    // 0x73
	[]uint8("pi"),    // 0x74
	[]uint8("hu"),    // 0x75
	[]uint8("bu"),    // 0x76
	[]uint8("pu"),    // 0x77
	[]uint8("he"),    // 0x78
	[]uint8("be"),    // 0x79
	[]uint8("pe"),    // 0x7a
	[]uint8("ho"),    // 0x7b
	[]uint8("bo"),    // 0x7c
	[]uint8("po"),    // 0x7d
	[]uint8("ma"),    // 0x7e
	[]uint8("mi"),    // 0x7f
	[]uint8("mu"),    // 0x80
	[]uint8("me"),    // 0x81
	[]uint8("mo"),    // 0x82
	[]uint8("ya"),    // 0x83
	[]uint8("ya"),    // 0x84
	[]uint8("yu"),    // 0x85
	[]uint8("yu"),    // 0x86
	[]uint8("yo"),    // 0x87
	[]uint8("yo"),    // 0x88
	[]uint8("ra"),    // 0x89
	[]uint8("ri"),    // 0x8a
	[]uint8("ru"),    // 0x8b
	[]uint8("re"),    // 0x8c
	[]uint8("ro"),    // 0x8d
	[]uint8("wa"),    // 0x8e
	[]uint8("wa"),    // 0x8f
	[]uint8("wi"),    // 0x90
	[]uint8("we"),    // 0x91
	[]uint8("wo"),    // 0x92
	[]uint8("n"),     // 0x93
	[]uint8("vu"),    // 0x94
	[]uint8("[?]"),   // 0x95
	[]uint8("[?]"),   // 0x96
	[]uint8("[?]"),   // 0x97
	[]uint8("[?]"),   // 0x98
	[]uint8(""),      // 0x99
	[]uint8(""),      // 0x9a
	[]uint8(""),      // 0x9b
	[]uint8(""),      // 0x9c
	[]uint8("\""),    // 0x9d
	[]uint8("\""),    // 0x9e
	[]uint8("[?]"),   // 0x9f
	[]uint8("[?]"),   // 0xa0
	[]uint8("a"),     // 0xa1
	[]uint8("a"),     // 0xa2
	[]uint8("i"),     // 0xa3
	[]uint8("i"),     // 0xa4
	[]uint8("u"),     // 0xa5
	[]uint8("u"),     // 0xa6
	[]uint8("e"),     // 0xa7
	[]uint8("e"),     // 0xa8
	[]uint8("o"),     // 0xa9
	[]uint8("o"),     // 0xaa
	[]uint8("ka"),    // 0xab
	[]uint8("ga"),    // 0xac
	[]uint8("ki"),    // 0xad
	[]uint8("gi"),    // 0xae
	[]uint8("ku"),    // 0xaf
	[]uint8("gu"),    // 0xb0
	[]uint8("ke"),    // 0xb1
	[]uint8("ge"),    // 0xb2
	[]uint8("ko"),    // 0xb3
	[]uint8("go"),    // 0xb4
	[]uint8("sa"),    // 0xb5
	[]uint8("za"),    // 0xb6
	[]uint8("shi"),   // 0xb7
	[]uint8("zi"),    // 0xb8
	[]uint8("su"),    // 0xb9
	[]uint8("zu"),    // 0xba
	[]uint8("se"),    // 0xbb
	[]uint8("ze"),    // 0xbc
	[]uint8("so"),    // 0xbd
	[]uint8("zo"),    // 0xbe
	[]uint8("ta"),    // 0xbf
	[]uint8("da"),    // 0xc0
	[]uint8("chi"),   // 0xc1
	[]uint8("di"),    // 0xc2
	[]uint8("tsu"),   // 0xc3
	[]uint8("tsu"),   // 0xc4
	[]uint8("du"),    // 0xc5
	[]uint8("te"),    // 0xc6
	[]uint8("de"),    // 0xc7
	[]uint8("to"),    // 0xc8
	[]uint8("do"),    // 0xc9
	[]uint8("na"),    // 0xca
	[]uint8("ni"),    // 0xcb
	[]uint8("nu"),    // 0xcc
	[]uint8("ne"),    // 0xcd
	[]uint8("no"),    // 0xce
	[]uint8("ha"),    // 0xcf
	[]uint8("ba"),    // 0xd0
	[]uint8("pa"),    // 0xd1
	[]uint8("hi"),    // 0xd2
	[]uint8("bi"),    // 0xd3
	[]uint8("pi"),    // 0xd4
	[]uint8("hu"),    // 0xd5
	[]uint8("bu"),    // 0xd6
	[]uint8("pu"),    // 0xd7
	[]uint8("he"),    // 0xd8
	[]uint8("be"),    // 0xd9
	[]uint8("pe"),    // 0xda
	[]uint8("ho"),    // 0xdb
	[]uint8("bo"),    // 0xdc
	[]uint8("po"),    // 0xdd
	[]uint8("ma"),    // 0xde
	[]uint8("mi"),    // 0xdf
	[]uint8("mu"),    // 0xe0
	[]uint8("me"),    // 0xe1
	[]uint8("mo"),    // 0xe2
	[]uint8("ya"),    // 0xe3
	[]uint8("ya"),    // 0xe4
	[]uint8("yu"),    // 0xe5
	[]uint8("yu"),    // 0xe6
	[]uint8("yo"),    // 0xe7
	[]uint8("yo"),    // 0xe8
	[]uint8("ra"),    // 0xe9
	[]uint8("ri"),    // 0xea
	[]uint8("ru"),    // 0xeb
	[]uint8("re"),    // 0xec
	[]uint8("ro"),    // 0xed
	[]uint8("wa"),    // 0xee
	[]uint8("wa"),    // 0xef
	[]uint8("wi"),    // 0xf0
	[]uint8("we"),    // 0xf1
	[]uint8("wo"),    // 0xf2
	[]uint8("n"),     // 0xf3
	[]uint8("vu"),    // 0xf4
	[]uint8("ka"),    // 0xf5
	[]uint8("ke"),    // 0xf6
	[]uint8("va"),    // 0xf7
	[]uint8("vi"),    // 0xf8
	[]uint8("ve"),    // 0xf9
	[]uint8("vo"),    // 0xfa
	[]uint8(""),      // 0xfb
	[]uint8(""),      // 0xfc
	[]uint8("\""),    // 0xfd
	[]uint8("\""),    // 0xfe
}
