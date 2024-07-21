package buffer

import "github.com/LukaVercauteren/TextEdit/base/text"

type Buffer struct {
	Label   string
	Content map[uint]text.Text
}
