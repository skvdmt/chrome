package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/input/dispatch_drag_event"
	"github.com/skvdmt/chrome/internal/devtools/options/input/dispatch_key_event"
	"github.com/skvdmt/chrome/internal/devtools/options/input/dispatch_mouse_event"
	"github.com/skvdmt/chrome/internal/devtools/options/input/dispatch_touch_event"
	"github.com/skvdmt/chrome/internal/devtools/options/input/emulate_touch_from_mouse_event"
	"github.com/skvdmt/chrome/internal/devtools/options/input/ime_set_composition"
	"github.com/skvdmt/chrome/internal/devtools/options/input/synthesize_pinch_gesture"
	"github.com/skvdmt/chrome/internal/devtools/options/input/synthesize_scroll_gesture"
	"github.com/skvdmt/chrome/internal/devtools/options/input/synthesize_tap_gesture"
	"github.com/skvdmt/chrome/internal/devtools/types/input"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
	"github.com/skvdmt/chrome/internal/model"
)

// Input Вврд.
type Input struct {
	client *model.Client
	debug  *model.Debug
	// Текущая сессия.
	CurrentSessionId *target.SessionId
}

// NewDriver Конструктор.
func NewInput(c *model.Client, d *model.Debug) *Input {
	d.Debug("input created")
	return &Input{
		client: c,
		debug:  d,
	}
}

// CancelDragging Отменяет любое активное перетаскивание по странице.
func (i *Input) CancelDragging() error {
	return i.client.Exec(
		input.CANCEL_DRAGGING,
		nil,
		model.WithSessionId(i.CurrentSessionId),
	)
}

// DispatchKeyEvent Отправляет ключевое событие на страницу.
func (i *Input) DispatchKeyEvent(typ string, options ...dispatch_key_event.Option) error {
	c := &dispatch_key_event.Config{
		Type: typ,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.DISPATCH_KEY_EVENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// DispatchMouseEvent Отправляет событие мыши на страницу.
func (i *Input) DispatchMouseEvent(typ string, x, y int, options ...dispatch_mouse_event.Option) error {
	c := &dispatch_mouse_event.Config{
		Type: typ,
		X:    x,
		Y:    y,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.DISPATCH_MOUSE_EVENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// DispatchTouchEvent Отправляет событие касания на страницу.
func (i *Input) DispatchTouchEvent(
	typ string,
	touchPoints []*input.TouchPoint,
	options ...dispatch_touch_event.Option,
) error {
	c := &dispatch_touch_event.Config{
		Type:        typ,
		TouchPoints: touchPoints,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.DISPATCH_TOUCH_EVENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// SetIgnoreInputEvents Игнорирует события ввода (полезно при
// проведении аудита страницы).
func (i *Input) SetIgnoreInputEvents(ignore bool) error {
	return i.client.Exec(
		input.SET_IGNORE_INPUT_EVENTS,
		model.ForceJSONMarshal(struct {
			// При установке значения true обработка входных событий игнорируется.
			Ignore bool `json:"ignore"`
		}{
			Ignore: ignore,
		}),
	)
}

// DispatchDragEvent Отправляет событие перетаскивания на страницу.
func (i *Input) DispatchDragEvent(
	typ string,
	x, y int,
	data *input.DragData,
	options ...dispatch_drag_event.Option,
) error {
	c := &dispatch_drag_event.Config{
		Type: typ,
		X:    x,
		Y:    y,
		Data: data,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.DISPATCH_DRAG_EVENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// EmulateTouchFromMouseEvent Имитирует сенсорное событие
// на основе параметров события мыши.
func (i *Input) EmulateTouchFromMouseEvent(
	typ string,
	x, y int,
	button *input.MouseButton,
	options ...emulate_touch_from_mouse_event.Option,
) error {
	c := &emulate_touch_from_mouse_event.Config{
		Type:   typ,
		X:      x,
		Y:      y,
		Button: button,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.EMULATE_TOUCH_FROM_MOUSE_EVENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// ImeSetComposition Этот метод устанавливает текущий текст-кандидат для IME.
// Используйте imeCommitComposition для подтверждения окончательного текста.
// Используйте imeSetComposition с пустой строкой в ​​качестве текста для отмены ввода.
func (i *Input) ImeSetComposition(
	text string,
	selectionStart, selectionEnd int,
	options ...ime_set_composition.Option,
) error {
	c := &ime_set_composition.Config{
		Text:           text,
		SelectionStart: selectionStart,
		SelectionEnd:   selectionEnd,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.IME_SET_COMPOSITION,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// InsertText Этот метод имитирует вставку текста, который не создается
// нажатием клавиши, например, с клавиатуры эмодзи или IME.
func (i *Input) InsertText(text string) error {
	return i.client.Exec(
		input.INSERT_TEXT,
		model.ForceJSONMarshal(struct {
			// Текст для вставки.
			Text string `json:"text"`
		}{
			Text: text,
		}),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// SetInterceptDrags Предотвращает стандартное поведение перетаскивания и
// вместо этого генерирует события Input.DragIntercepted. Поведение
// перетаскивания можно напрямую контролировать с помощью Input.DispatchDragEvent.
func (i *Input) SetInterceptDrags(enabled bool) error {
	return i.client.Exec(
		input.SET_INTERCEPT_DRAGS,
		model.ForceJSONMarshal(struct {
			Enabled bool `json:"enabled"`
		}{
			Enabled: enabled,
		}),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// SynthesizePinchGesture Обрабатывает жест масштабирования (щипок) в течение
// определенного периода времени, генерируя соответствующие события касания.
func (i *Input) SynthesizePinchGesture(
	x, y, scaleFactor int,
	options ...synthesize_pinch_gesture.Option,
) error {
	c := &synthesize_pinch_gesture.Config{
		X:           x,
		Y:           y,
		ScaleFactor: scaleFactor,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.SYNTHESIZE_PINCH_GESTURE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// SynthesizeScrollGesture Обрабатывает жест прокрутки в течение определенного
// периода времени, генерируя соответствующие события касания.
func (i *Input) SynthesizeScrollGesture(
	x, y int,
	options ...synthesize_scroll_gesture.Option,
) error {
	c := &synthesize_scroll_gesture.Config{
		X: x,
		Y: y,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.SYNTHESIZE_SCROLL_GESTURE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}

// SynthesizeTapGesture Обрабатывает жест касания в течение определенного
// периода времени, генерируя соответствующие события касания.
func (i *Input) SynthesizeTapGesture(
	x, y int,
	options ...synthesize_tap_gesture.Option,
) error {
	c := &synthesize_tap_gesture.Config{
		X: x,
		Y: y,
	}
	for _, o := range options {
		o(c)
	}
	return i.client.Exec(
		input.SYNTHESIZE_TAP_GESTURE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(i.CurrentSessionId),
	)
}
