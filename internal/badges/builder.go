package badges

const (
	Blue   string = "blue"
	Red    string = "red"
	Green  string = "green"
	Yellow string = "yellow"
)

type BadgeBuilder struct {
	Label        string
	Message      string
	MessageColor string
	Style        string
}

func NewBadgeBuilder() *BadgeBuilder {
	return &BadgeBuilder{
		MessageColor: Green,
	}
}

func (b *BadgeBuilder) SetLabel(label string) *BadgeBuilder {
	b.Label = label
	return b
}

func (b *BadgeBuilder) SetMessage(message string) *BadgeBuilder {
	b.Message = message
	return b
}

func (b *BadgeBuilder) SetMessageColor(messageColor string) *BadgeBuilder {
	b.MessageColor = messageColor
	return b
}

func (b *BadgeBuilder) SetStyle(style string) *BadgeBuilder {
	b.Style = style
	return b
}

func (b *BadgeBuilder) Build() *BadgeBuilder {
	return b
}
