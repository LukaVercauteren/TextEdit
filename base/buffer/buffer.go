package buffer

import "github.com/LukaVercauteren/TextEdit/base/text"

type Buffer struct {
	label   string
	content map[uint]text.Text
}
