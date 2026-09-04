package keys

// Руны клавиш.
const (
	Backspace            = "\b"
	Tab                  = "\t"
	Enter                = "\r"
	Escape               = "\u001b"
	Quote                = "'"
	Backslash            = "\\"
	Delete               = "\u007f"
	Alt                  = "\u0102"
	CapsLock             = "\u0104"
	Control              = "\u0105"
	Fn                   = "\u0106"
	FnLock               = "\u0107"
	Hyper                = "\u0108"
	Meta                 = "\u0109"
	NumLock              = "\u010a"
	ScrollLock           = "\u010c"
	Shift                = "\u010d"
	Super                = "\u010e"
	ArrowDown            = "\u0301"
	ArrowLeft            = "\u0302"
	ArrowRight           = "\u0303"
	ArrowUp              = "\u0304"
	End                  = "\u0305"
	Home                 = "\u0306"
	PageDown             = "\u0307"
	PageUp               = "\u0308"
	Clear                = "\u0401"
	Copy                 = "\u0402"
	Cut                  = "\u0404"
	Insert               = "\u0407"
	Paste                = "\u0408"
	Redo                 = "\u0409"
	Undo                 = "\u040a"
	Again                = "\u0502"
	Cancel               = "\u0504"
	ContextMenu          = "\u0505"
	Find                 = "\u0507"
	Help                 = "\u0508"
	Pause                = "\u0509"
	Props                = "\u050b"
	Select               = "\u050c"
	ZoomIn               = "\u050d"
	ZoomOut              = "\u050e"
	BrightnessDown       = "\u0601"
	BrightnessUp         = "\u0602"
	Eject                = "\u0604"
	LogOff               = "\u0605"
	Power                = "\u0606"
	PrintScreen          = "\u0608"
	WakeUp               = "\u060b"
	Convert              = "\u0705"
	ModeChange           = "\u070b"
	NonConvert           = "\u070d"
	HangulMode           = "\u0711"
	HanjaMode            = "\u0712"
	Hiragana             = "\u0716"
	KanaMode             = "\u0718"
	Katakana             = "\u071a"
	ZenkakuHankaku       = "\u071d"
	F1                   = "\u0801"
	F2                   = "\u0802"
	F3                   = "\u0803"
	F4                   = "\u0804"
	F5                   = "\u0805"
	F6                   = "\u0806"
	F7                   = "\u0807"
	F8                   = "\u0808"
	F9                   = "\u0809"
	F10                  = "\u080a"
	F11                  = "\u080b"
	F12                  = "\u080c"
	F13                  = "\u080d"
	F14                  = "\u080e"
	F15                  = "\u080f"
	F16                  = "\u0810"
	F17                  = "\u0811"
	F18                  = "\u0812"
	F19                  = "\u0813"
	F20                  = "\u0814"
	F21                  = "\u0815"
	F22                  = "\u0816"
	F23                  = "\u0817"
	F24                  = "\u0818"
	Close                = "\u0a01"
	MailForward          = "\u0a02"
	MailReply            = "\u0a03"
	MailSend             = "\u0a04"
	MediaPlayPause       = "\u0a05"
	MediaStop            = "\u0a07"
	MediaTrackNext       = "\u0a08"
	MediaTrackPrevious   = "\u0a09"
	New                  = "\u0a0a"
	Open                 = "\u0a0b"
	Print                = "\u0a0c"
	Save                 = "\u0a0d"
	SpellCheck           = "\u0a0e"
	AudioVolumeDown      = "\u0a0f"
	AudioVolumeUp        = "\u0a10"
	AudioVolumeMute      = "\u0a11"
	LaunchApplication2   = "\u0b01"
	LaunchCalendar       = "\u0b02"
	LaunchMail           = "\u0b03"
	LaunchMediaPlayer    = "\u0b04"
	LaunchMusicPlayer    = "\u0b05"
	LaunchApplication1   = "\u0b06"
	LaunchScreenSaver    = "\u0b07"
	LaunchSpreadsheet    = "\u0b08"
	LaunchWebBrowser     = "\u0b09"
	LaunchContacts       = "\u0b0c"
	LaunchPhone          = "\u0b0d"
	LaunchAssistant      = "\u0b0e"
	BrowserBack          = "\u0c01"
	BrowserFavorites     = "\u0c02"
	BrowserForward       = "\u0c03"
	BrowserHome          = "\u0c04"
	BrowserRefresh       = "\u0c05"
	BrowserSearch        = "\u0c06"
	BrowserStop          = "\u0c07"
	ChannelDown          = "\u0d0a"
	ChannelUp            = "\u0d0b"
	ClosedCaptionToggle  = "\u0d12"
	Exit                 = "\u0d15"
	Guide                = "\u0d22"
	Info                 = "\u0d25"
	MediaFastForward     = "\u0d2c"
	MediaLast            = "\u0d2d"
	MediaPause           = "\u0d2e"
	MediaPlay            = "\u0d2f"
	MediaRecord          = "\u0d30"
	MediaRewind          = "\u0d31"
	Settings             = "\u0d43"
	ZoomToggle           = "\u0d4e"
	AudioBassBoostToggle = "\u0e02"
	SpeechInputToggle    = "\u0f02"
	AppSwitch            = "\u1001"
)

// Key информация для генерации нажатия клавиши.
type Key struct {
	// Код клавиши.
	Code string
	// Значение клавиши.
	Key string
	// Текст клавиши.
	Text string
	// Неизменяемый текст клавиши
	Unmodified string
	// Исходный скан-код.
	Native int64
	// Код клавиши Windows.
	Windows int64
	// Отправлять модификатор Shift.
	Shift bool
	// Клавиша печатная.
	Print bool
}

// NewKey Конструктор.
func NewKey(
	code, key, text, unmodified string,
	native, windows int64,
	shift, print bool,
) *Key {
	return &Key{
		Code:       code,
		Key:        key,
		Text:       text,
		Unmodified: unmodified,
		Native:     native,
		Windows:    windows,
		Shift:      shift,
		Print:      print,
	}
}

// Keys таблица информации о клавишах.
var Keys = map[rune]*Key{
	'\b':     NewKey("Backspace", "Backspace", "", "", 8, 8, false, false),
	'\t':     NewKey("Tab", "Tab", "", "", 9, 9, false, false),
	'\r':     NewKey("Enter", "Enter", "\r", "\r", 13, 13, false, true),
	'\u001b': NewKey("Escape", "Escape", "", "", 27, 27, false, false),
	' ':      NewKey("Space", " ", " ", " ", 32, 32, false, true),
	'!':      NewKey("Digit1", "!", "!", "1", 49, 49, true, true),
	'"':      NewKey("Quote", "\"", "\"", "'", 222, 222, true, true),
	'#':      NewKey("Digit3", "#", "#", "3", 51, 51, true, true),
	'$':      NewKey("Digit4", "$", "$", "4", 52, 52, true, true),
	'%':      NewKey("Digit5", "%", "%", "5", 53, 53, true, true),
	'&':      NewKey("Digit7", "&", "&", "7", 55, 55, true, true),
	'\'':     NewKey("Quote", "'", "'", "'", 222, 222, false, true),
	'(':      NewKey("Digit9", "(", "(", "9", 57, 57, true, true),
	')':      NewKey("Digit0", ")", ")", "0", 48, 48, true, true),
	'*':      NewKey("Digit8", "*", "*", "8", 56, 56, true, true),
	'+':      NewKey("Equal", "+", "+", "=", 187, 187, true, true),
	',':      NewKey("Comma", ",", ",", ",", 188, 188, false, true),
	'-':      NewKey("Minus", "-", "-", "-", 189, 189, false, true),
	'.':      NewKey("Period", ".", ".", ".", 190, 190, false, true),
	'/':      NewKey("Slash", "/", "/", "/", 191, 191, false, true),
	'0':      NewKey("Digit0", "0", "0", "0", 48, 48, false, true),
	'1':      NewKey("Digit1", "1", "1", "1", 49, 49, false, true),
	'2':      NewKey("Digit2", "2", "2", "2", 50, 50, false, true),
	'3':      NewKey("Digit3", "3", "3", "3", 51, 51, false, true),
	'4':      NewKey("Digit4", "4", "4", "4", 52, 52, false, true),
	'5':      NewKey("Digit5", "5", "5", "5", 53, 53, false, true),
	'6':      NewKey("Digit6", "6", "6", "6", 54, 54, false, true),
	'7':      NewKey("Digit7", "7", "7", "7", 55, 55, false, true),
	'8':      NewKey("Digit8", "8", "8", "8", 56, 56, false, true),
	'9':      NewKey("Digit9", "9", "9", "9", 57, 57, false, true),
	':':      NewKey("Semicolon", ":", ":", ";", 186, 186, true, true),
	';':      NewKey("Semicolon", ";", ";", ";", 186, 186, false, true),
	'<':      NewKey("Comma", "<", "<", ",", 188, 188, true, true),
	'=':      NewKey("Equal", "=", "=", "=", 187, 187, false, true),
	'>':      NewKey("Period", ">", ">", ".", 190, 190, true, true),
	'?':      NewKey("Slash", "?", "?", "/", 191, 191, true, true),
	'@':      NewKey("Digit2", "@", "@", "2", 50, 50, true, true),
	'A':      NewKey("KeyA", "A", "A", "a", 65, 65, true, true),
	'B':      NewKey("KeyB", "B", "B", "b", 66, 66, true, true),
	'C':      NewKey("KeyC", "C", "C", "c", 67, 67, true, true),
	'D':      NewKey("KeyD", "D", "D", "d", 68, 68, true, true),
	'E':      NewKey("KeyE", "E", "E", "e", 69, 69, true, true),
	'F':      NewKey("KeyF", "F", "F", "f", 70, 70, true, true),
	'G':      NewKey("KeyG", "G", "G", "g", 71, 71, true, true),
	'H':      NewKey("KeyH", "H", "H", "h", 72, 72, true, true),
	'I':      NewKey("KeyI", "I", "I", "i", 73, 73, true, true),
	'J':      NewKey("KeyJ", "J", "J", "j", 74, 74, true, true),
	'K':      NewKey("KeyK", "K", "K", "k", 75, 75, true, true),
	'L':      NewKey("KeyL", "L", "L", "l", 76, 76, true, true),
	'M':      NewKey("KeyM", "M", "M", "m", 77, 77, true, true),
	'N':      NewKey("KeyN", "N", "N", "n", 78, 78, true, true),
	'O':      NewKey("KeyO", "O", "O", "o", 79, 79, true, true),
	'P':      NewKey("KeyP", "P", "P", "p", 80, 80, true, true),
	'Q':      NewKey("KeyQ", "Q", "Q", "q", 81, 81, true, true),
	'R':      NewKey("KeyR", "R", "R", "r", 82, 82, true, true),
	'S':      NewKey("KeyS", "S", "S", "s", 83, 83, true, true),
	'T':      NewKey("KeyT", "T", "T", "t", 84, 84, true, true),
	'U':      NewKey("KeyU", "U", "U", "u", 85, 85, true, true),
	'V':      NewKey("KeyV", "V", "V", "v", 86, 86, true, true),
	'W':      NewKey("KeyW", "W", "W", "w", 87, 87, true, true),
	'X':      NewKey("KeyX", "X", "X", "x", 88, 88, true, true),
	'Y':      NewKey("KeyY", "Y", "Y", "y", 89, 89, true, true),
	'Z':      NewKey("KeyZ", "Z", "Z", "z", 90, 90, true, true),
	'[':      NewKey("BracketLeft", "[", "[", "[", 219, 219, false, true),
	'\\':     NewKey("Backslash", "\\", "\\", "\\", 220, 220, false, true),
	']':      NewKey("BracketRight", "]", "]", "]", 221, 221, false, true),
	'^':      NewKey("Digit6", "^", "^", "6", 54, 54, true, true),
	'_':      NewKey("Minus", "_", "_", "-", 189, 189, true, true),
	'`':      NewKey("Backquote", "`", "`", "`", 192, 192, false, true),
	'a':      NewKey("KeyA", "a", "a", "a", 65, 65, false, true),
	'b':      NewKey("KeyB", "b", "b", "b", 66, 66, false, true),
	'c':      NewKey("KeyC", "c", "c", "c", 67, 67, false, true),
	'd':      NewKey("KeyD", "d", "d", "d", 68, 68, false, true),
	'e':      NewKey("KeyE", "e", "e", "e", 69, 69, false, true),
	'f':      NewKey("KeyF", "f", "f", "f", 70, 70, false, true),
	'g':      NewKey("KeyG", "g", "g", "g", 71, 71, false, true),
	'h':      NewKey("KeyH", "h", "h", "h", 72, 72, false, true),
	'i':      NewKey("KeyI", "i", "i", "i", 73, 73, false, true),
	'j':      NewKey("KeyJ", "j", "j", "j", 74, 74, false, true),
	'k':      NewKey("KeyK", "k", "k", "k", 75, 75, false, true),
	'l':      NewKey("KeyL", "l", "l", "l", 76, 76, false, true),
	'm':      NewKey("KeyM", "m", "m", "m", 77, 77, false, true),
	'n':      NewKey("KeyN", "n", "n", "n", 78, 78, false, true),
	'o':      NewKey("KeyO", "o", "o", "o", 79, 79, false, true),
	'p':      NewKey("KeyP", "p", "p", "p", 80, 80, false, true),
	'q':      NewKey("KeyQ", "q", "q", "q", 81, 81, false, true),
	'r':      NewKey("KeyR", "r", "r", "r", 82, 82, false, true),
	's':      NewKey("KeyS", "s", "s", "s", 83, 83, false, true),
	't':      NewKey("KeyT", "t", "t", "t", 84, 84, false, true),
	'u':      NewKey("KeyU", "u", "u", "u", 85, 85, false, true),
	'v':      NewKey("KeyV", "v", "v", "v", 86, 86, false, true),
	'w':      NewKey("KeyW", "w", "w", "w", 87, 87, false, true),
	'x':      NewKey("KeyX", "x", "x", "x", 88, 88, false, true),
	'y':      NewKey("KeyY", "y", "y", "y", 89, 89, false, true),
	'z':      NewKey("KeyZ", "z", "z", "z", 90, 90, false, true),
	'{':      NewKey("BracketLeft", "{", "{", "[", 219, 219, true, true),
	'|':      NewKey("Backslash", "|", "|", "\\", 220, 220, true, true),
	'}':      NewKey("BracketRight", "}", "}", "]", 221, 221, true, true),
	'~':      NewKey("Backquote", "~", "~", "`", 192, 192, true, true),
	'\u007f': NewKey("Delete", "Delete", "", "", 46, 46, false, false),
	'¥':      NewKey("IntlYen", "¥", "¥", "¥", 220, 220, false, true),
	'\u0102': NewKey("AltLeft", "Alt", "", "", 164, 164, false, false),
	'\u0104': NewKey("CapsLock", "CapsLock", "", "", 20, 20, false, false),
	'\u0105': NewKey("ControlLeft", "Control", "", "", 162, 162, false, false),
	'\u0106': NewKey("Fn", "Fn", "", "", 0, 0, false, false),
	'\u0107': NewKey("FnLock", "FnLock", "", "", 0, 0, false, false),
	'\u0108': NewKey("Hyper", "Hyper", "", "", 0, 0, false, false),
	'\u0109': NewKey("MetaLeft", "Meta", "", "", 91, 91, false, false),
	'\u010a': NewKey("NumLock", "NumLock", "", "", 144, 144, false, false),
	'\u010c': NewKey("ScrollLock", "ScrollLock", "", "", 145, 145, false, false),
	'\u010d': NewKey("ShiftLeft", "Shift", "", "", 160, 160, false, false),
	'\u010e': NewKey("Super", "Super", "", "", 0, 0, false, false),
	'\u0301': NewKey("ArrowDown", "ArrowDown", "", "", 40, 40, false, false),
	'\u0302': NewKey("ArrowLeft", "ArrowLeft", "", "", 37, 37, false, false),
	'\u0303': NewKey("ArrowRight", "ArrowRight", "", "", 39, 39, false, false),
	'\u0304': NewKey("ArrowUp", "ArrowUp", "", "", 38, 38, false, false),
	'\u0305': NewKey("End", "End", "", "", 35, 35, false, false),
	'\u0306': NewKey("Home", "Home", "", "", 36, 36, false, false),
	'\u0307': NewKey("PageDown", "PageDown", "", "", 34, 34, false, false),
	'\u0308': NewKey("PageUp", "PageUp", "", "", 33, 33, false, false),
	'\u0401': NewKey("NumpadClear", "Clear", "", "", 12, 12, false, false),
	'\u0402': NewKey("Copy", "Copy", "", "", 0, 0, false, false),
	'\u0404': NewKey("Cut", "Cut", "", "", 0, 0, false, false),
	'\u0407': NewKey("Insert", "Insert", "", "", 45, 45, false, false),
	'\u0408': NewKey("Paste", "Paste", "", "", 0, 0, false, false),
	'\u0409': NewKey("Redo", "Redo", "", "", 0, 0, false, false),
	'\u040a': NewKey("Undo", "Undo", "", "", 0, 0, false, false),
	'\u0502': NewKey("Again", "Again", "", "", 0, 0, false, false),
	'\u0504': NewKey("Abort", "Cancel", "", "", 3, 3, false, false),
	'\u0505': NewKey("ContextMenu", "ContextMenu", "", "", 93, 93, false, false),
	'\u0507': NewKey("Find", "Find", "", "", 0, 0, false, false),
	'\u0508': NewKey("Help", "Help", "", "", 47, 47, false, false),
	'\u0509': NewKey("Pause", "Pause", "", "", 19, 19, false, false),
	'\u050b': NewKey("Props", "Props", "", "", 0, 0, false, false),
	'\u050c': NewKey("Select", "Select", "", "", 41, 41, false, false),
	'\u050d': NewKey("ZoomIn", "ZoomIn", "", "", 0, 0, false, false),
	'\u050e': NewKey("ZoomOut", "ZoomOut", "", "", 0, 0, false, false),
	'\u0601': NewKey("BrightnessDown", "BrightnessDown", "", "", 216, 0, false, false),
	'\u0602': NewKey("BrightnessUp", "BrightnessUp", "", "", 217, 0, false, false),
	'\u0604': NewKey("Eject", "Eject", "", "", 0, 0, false, false),
	'\u0605': NewKey("LogOff", "LogOff", "", "", 0, 0, false, false),
	'\u0606': NewKey("Power", "Power", "", "", 152, 0, false, false),
	'\u0608': NewKey("PrintScreen", "PrintScreen", "", "", 44, 44, false, false),
	'\u060b': NewKey("WakeUp", "WakeUp", "", "", 0, 0, false, false),
	'\u0705': NewKey("Convert", "Convert", "", "", 28, 28, false, false),
	'\u070b': NewKey("KeyboardLayoutSelect", "ModeChange", "", "", 0, 0, false, false),
	'\u070d': NewKey("NonConvert", "NonConvert", "", "", 29, 29, false, false),
	'\u0711': NewKey("Lang1", "HangulMode", "", "", 21, 21, false, false),
	'\u0712': NewKey("Lang2", "HanjaMode", "", "", 25, 25, false, false),
	'\u0716': NewKey("Lang4", "Hiragana", "", "", 0, 0, false, false),
	'\u0718': NewKey("KanaMode", "KanaMode", "", "", 21, 21, false, false),
	'\u071a': NewKey("Lang3", "Katakana", "", "", 0, 0, false, false),
	'\u071d': NewKey("Lang5", "ZenkakuHankaku", "", "", 0, 0, false, false),
	'\u0801': NewKey("F1", "F1", "", "", 112, 112, false, false),
	'\u0802': NewKey("F2", "F2", "", "", 113, 113, false, false),
	'\u0803': NewKey("F3", "F3", "", "", 114, 114, false, false),
	'\u0804': NewKey("F4", "F4", "", "", 115, 115, false, false),
	'\u0805': NewKey("F5", "F5", "", "", 116, 116, false, false),
	'\u0806': NewKey("F6", "F6", "", "", 117, 117, false, false),
	'\u0807': NewKey("F7", "F7", "", "", 118, 118, false, false),
	'\u0808': NewKey("F8", "F8", "", "", 119, 119, false, false),
	'\u0809': NewKey("F9", "F9", "", "", 120, 120, false, false),
	'\u080a': NewKey("F10", "F10", "", "", 121, 121, false, false),
	'\u080b': NewKey("F11", "F11", "", "", 122, 122, false, false),
	'\u080c': NewKey("F12", "F12", "", "", 123, 123, false, false),
	'\u080d': NewKey("F13", "F13", "", "", 124, 124, false, false),
	'\u080e': NewKey("F14", "F14", "", "", 125, 125, false, false),
	'\u080f': NewKey("F15", "F15", "", "", 126, 126, false, false),
	'\u0810': NewKey("F16", "F16", "", "", 127, 127, false, false),
	'\u0811': NewKey("F17", "F17", "", "", 128, 128, false, false),
	'\u0812': NewKey("F18", "F18", "", "", 129, 129, false, false),
	'\u0813': NewKey("F19", "F19", "", "", 130, 130, false, false),
	'\u0814': NewKey("F20", "F20", "", "", 131, 131, false, false),
	'\u0815': NewKey("F21", "F21", "", "", 132, 132, false, false),
	'\u0816': NewKey("F22", "F22", "", "", 133, 133, false, false),
	'\u0817': NewKey("F23", "F23", "", "", 134, 134, false, false),
	'\u0818': NewKey("F24", "F24", "", "", 135, 135, false, false),
	'\u0a01': NewKey("Close", "Close", "", "", 0, 0, false, false),
	'\u0a02': NewKey("MailForward", "MailForward", "", "", 0, 0, false, false),
	'\u0a03': NewKey("MailReply", "MailReply", "", "", 0, 0, false, false),
	'\u0a04': NewKey("MailSend", "MailSend", "", "", 0, 0, false, false),
	'\u0a05': NewKey("MediaPlayPause", "MediaPlayPause", "", "", 179, 179, false, false),
	'\u0a07': NewKey("MediaStop", "MediaStop", "", "", 178, 178, false, false),
	'\u0a08': NewKey("MediaTrackNext", "MediaTrackNext", "", "", 176, 176, false, false),
	'\u0a09': NewKey("MediaTrackPrevious", "MediaTrackPrevious", "", "", 177, 177, false, false),
	'\u0a0a': NewKey("New", "New", "", "", 0, 0, false, false),
	'\u0a0b': NewKey("Open", "Open", "", "", 43, 43, false, false),
	'\u0a0c': NewKey("Print", "Print", "", "", 0, 0, false, false),
	'\u0a0d': NewKey("Save", "Save", "", "", 0, 0, false, false),
	'\u0a0e': NewKey("SpellCheck", "SpellCheck", "", "", 0, 0, false, false),
	'\u0a0f': NewKey("AudioVolumeDown", "AudioVolumeDown", "", "", 174, 174, false, false),
	'\u0a10': NewKey("AudioVolumeUp", "AudioVolumeUp", "", "", 175, 175, false, false),
	'\u0a11': NewKey("AudioVolumeMute", "AudioVolumeMute", "", "", 173, 173, false, false),
	'\u0b01': NewKey("LaunchApp2", "LaunchApplication2", "", "", 183, 183, false, false),
	'\u0b02': NewKey("LaunchCalendar", "LaunchCalendar", "", "", 0, 0, false, false),
	'\u0b03': NewKey("LaunchMail", "LaunchMail", "", "", 180, 180, false, false),
	'\u0b04': NewKey("MediaSelect", "LaunchMediaPlayer", "", "", 181, 181, false, false),
	'\u0b05': NewKey("LaunchMusicPlayer", "LaunchMusicPlayer", "", "", 0, 0, false, false),
	'\u0b06': NewKey("LaunchApp1", "LaunchApplication1", "", "", 182, 182, false, false),
	'\u0b07': NewKey("LaunchScreenSaver", "LaunchScreenSaver", "", "", 0, 0, false, false),
	'\u0b08': NewKey("LaunchSpreadsheet", "LaunchSpreadsheet", "", "", 0, 0, false, false),
	'\u0b09': NewKey("LaunchWebBrowser", "LaunchWebBrowser", "", "", 0, 0, false, false),
	'\u0b0c': NewKey("LaunchContacts", "LaunchContacts", "", "", 0, 0, false, false),
	'\u0b0d': NewKey("LaunchPhone", "LaunchPhone", "", "", 0, 0, false, false),
	'\u0b0e': NewKey("LaunchAssistant", "LaunchAssistant", "", "", 153, 0, false, false),
	'\u0c01': NewKey("BrowserBack", "BrowserBack", "", "", 166, 166, false, false),
	'\u0c02': NewKey("BrowserFavorites", "BrowserFavorites", "", "", 171, 171, false, false),
	'\u0c03': NewKey("BrowserForward", "BrowserForward", "", "", 167, 167, false, false),
	'\u0c04': NewKey("BrowserHome", "BrowserHome", "", "", 172, 172, false, false),
	'\u0c05': NewKey("BrowserRefresh", "BrowserRefresh", "", "", 168, 168, false, false),
	'\u0c06': NewKey("BrowserSearch", "BrowserSearch", "", "", 170, 170, false, false),
	'\u0c07': NewKey("BrowserStop", "BrowserStop", "", "", 169, 169, false, false),
	'\u0d0a': NewKey("ChannelDown", "ChannelDown", "", "", 0, 0, false, false),
	'\u0d0b': NewKey("ChannelUp", "ChannelUp", "", "", 0, 0, false, false),
	'\u0d12': NewKey("ClosedCaptionToggle", "ClosedCaptionToggle", "", "", 0, 0, false, false),
	'\u0d15': NewKey("Exit", "Exit", "", "", 0, 0, false, false),
	'\u0d22': NewKey("Guide", "Guide", "", "", 0, 0, false, false),
	'\u0d25': NewKey("Info", "Info", "", "", 0, 0, false, false),
	'\u0d2c': NewKey("MediaFastForward", "MediaFastForward", "", "", 0, 0, false, false),
	'\u0d2d': NewKey("MediaLast", "MediaLast", "", "", 0, 0, false, false),
	'\u0d2e': NewKey("MediaPause", "MediaPause", "", "", 0, 0, false, false),
	'\u0d2f': NewKey("MediaPlay", "MediaPlay", "", "", 0, 0, false, false),
	'\u0d30': NewKey("MediaRecord", "MediaRecord", "", "", 0, 0, false, false),
	'\u0d31': NewKey("MediaRewind", "MediaRewind", "", "", 0, 0, false, false),
	'\u0d43': NewKey("LaunchControlPanel", "Settings", "", "", 154, 0, false, false),
	'\u0d4e': NewKey("ZoomToggle", "ZoomToggle", "", "", 251, 251, false, false),
	'\u0e02': NewKey("AudioBassBoostToggle", "AudioBassBoostToggle", "", "", 0, 0, false, false),
	'\u0f02': NewKey("SpeechInputToggle", "SpeechInputToggle", "", "", 0, 0, false, false),
	'\u1001': NewKey("SelectTask", "AppSwitch", "", "", 0, 0, false, false),
}
