package cursor

type CursorMode int

const (
	Insert CursorMode = iota
	Overwrite
)

type Cursor struct {
	mode            CursorMode
	currentRune     rune
	insertModeRune  rune
	overwriteModeRune rune
}

func NewCursor() Cursor {
	defaultCursorMode := Insert // Set default cursor mode to Insert
	return Cursor{
		mode:            defaultCursorMode,
		insertModeRune:  '|',
		overwriteModeRune: '█'
	}
}
