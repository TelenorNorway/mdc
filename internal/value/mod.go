package value

type Value struct {
	exists bool
	value  string
}

var None = &Value{exists: false}

func Some(value string) *Value {
	return &Value{exists: true, value: value}
}

func (v *Value) Exists() bool {
	return v.exists
}

func (v *Value) Value() string {
	return v.value
}

func From(exists bool, value string) *Value {
	return &Value{exists, value}
}

type Diff struct {
	Old *Value
	New *Value
}

func NewDiff(old, new *Value) *Diff {
	return &Diff{Old: old, New: new}
}
