package runtime

// RemoteObjectId Уникальный id объекта.
type RemoteObjectId string

// ScriptId Unique script identifier.
type ScriptId string

// UniqueDebuggerId Unique identifier of current debugger.
type UniqueDebuggerId string

// RemoteObject Зеркальный объект, ссылающийся на исходный объект JavaScript.
type RemoteObject struct {
	// Тип объекта.
	// Возможные значения: object, function, undefined, string,
	// number, boolean, symbol, bigint
	Type string `json:"type"`
	// Подсказка по подтипу объекта. Указывается только для значений типа объекта.
	// ПРИМЕЧАНИЕ: Если вы внесете какие-либо изменения здесь, обязательно
	// обновите также подтип в ObjectPreview и PropertyPreview ниже.
	// Возможные значения: array, null, node, regexp, date, map, set, weakmap,
	// weakset, iterator, generator, error, proxy, promise, typedarray,
	// arraybuffer, dataview, webassemblymemory, wasmvalue, trustedtype.
	Subtype string `json:"subtype,omitempty"`
	// Имя класса объекта (конструктора). Указывается только для значений типа объекта.
	ClassName string `json:"className,omitempty"`
	// Удаленное значение объекта в случае примитивных значений или
	// значений JSON (если таковое было запрошено).
	Value any `json:"value,omitempty"`
	// Примитивное значение, которое нельзя преобразовать в JSON-строку, не имеет значения, но получает это свойство.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`
	// Строковое представление объекта.
	Description string `json:"description,omitempty"`
	// Глубоко сериализованное значение.
	DeepSerializedValue DeepSerializedValue `json:"deepSerializedValue,omitempty"`
	// Уникальный идентификатор объекта (для не примитивных значений).
	ObjectId RemoteObjectId `json:"objectId,omitempty"`
	// Предварительный просмотр, содержащий сокращенные значения свойств.
	// Указано только для значений типов объектов.
	Preview       ObjectPreview `json:"preview,omitempty"`
	CustomPreview CustomPreview `json:"customPreview,omitempty"`
}

// UnserializableValue Примитивное значение, которое нельзя преобразовать в
// JSON-строку. Включает значения -0, NaN, Infinity, -Infinity и литералы типа bigint.
type UnserializableValue string

// DeepSerializedValue Представляет собой глубоко сериализованное значение.
type DeepSerializedValue struct {
	// Возможные значения: undefined, null, string, number, boolean, bigint, regexp,
	// date, symbol, array, object, function, map, set, weakmap, weakset, error,
	// proxy, promise, typedarray, arraybuffer, node, window, generator.
	Type     string `json:"type"`
	Value    any    `json:"value,omitempty"`
	ObjectId string `json:"objectId,omitempty"`
	// Устанавливается, если ссылка на значение встречалась более одного раза
	// во время сериализации. В таком случае значение предоставляется только
	// одному из сериализованных значений. Уникален для каждого значения
	// в рамках одного вызова CDP.
	WeakLocalObjectReference int `json:"weakLocalObjectReference,omitempty"`
}

// ObjectPreview Объект, содержащий сокращенное значение удаленного объекта.
type ObjectPreview struct {
	// Тип объекта.
	// Возможные значения: object, function, undefined,
	// string, number, boolean, symbol, bigint.
	Type string `json:"type"`
	// Подсказка по подтипу объекта. Указывается только для значений типа объекта.
	// Возможные значения: array, null, node, regexp, date, map, set, weakmap,
	// weakset, iterator, generator, error, proxy, promise, typedarray,
	// arraybuffer, dataview, webassemblymemory, wasmvalue, trustedtype.
	Subtype string `json:"subtype,omitempty"`
	// Строковое представление объекта.
	Description string `json:"description,omitempty"`
	// Истина тогда и только тогда, когда некоторые свойства или записи
	// исходного объекта не соответствуют действительности.
	Overflow bool `json:"overflow"`
	// Список объектов недвижимости.
	Properties []PropertyPreview `json:"properties"`
	// Список записей. Указано только для значений подтипов map и set.
	Entries []EntryPreview `json:"entries,omitempty"`
}

// PropertyPreview
type PropertyPreview struct {
	// Название свойства.
	Name string `json:"name"`
	// Тип объекта. «Аксессор» означает, что само свойство является свойством-аксессором.
	// Возможные значения: object, function, undefined, string,
	// number, boolean, symbol, accessor, bigint.
	Type string `json:"type"`
	// Удобная для пользователя строка значений свойства.
	Value string `json:"value,omitempty"`
	// Предварительный просмотр вложенных значений.
	ValuePreview ObjectPreview `json:"valuePreview,omitempty"`
	// Подсказка по подтипу объекта. Указывается только для значений типа объекта.
	// Возможные значения: array, null, node, regexp, date, map, set, weakmap,
	// weakset, iterator, generator, error, proxy, promise, typedarray,
	// arraybuffer, dataview, webassemblymemory, wasmvalue, trustedtype.
	Subtype string `json:"subtype,omitempty"`
}

// EntryPreview
type EntryPreview struct {
	// Предварительный просмотр ключа. Предназначен для записей в виде карт.
	Key ObjectPreview `json:"key,omitempty"`
	// Предварительный просмотр значения.
	Value ObjectPreview `json:"value"`
}

// CustomPreview
type CustomPreview struct {
	// Результат вызова formatter.header(object, config), преобразованный в JSON-строку.
	// Он содержит массив JSON ML, представляющий RemoteObject.
	Header string `json:"header"`
	// Если метод formatter.hasBody возвращает true, то bodyGetterId
	// будет содержать RemoteObjectId для функции, которая возвращает
	// результат вызова formatter.body(object, config). Результирующее
	// значение представляет собой массив JSON ML.
	BodyGetterId RemoteObjectId `json:"bodyGetterId,omitempty"`
}

// ExecutionContextId Идентификатор контекста выполнения.
type ExecutionContextId int

// StackTrace Для получения утверждений или сообщений
// об ошибках используйте фреймы вызовов.
type StackTrace struct {
	
}
