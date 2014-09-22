package indexes

/* Information type*/
type Information struct {
	Phrase string
	Occurrencies int
	Sources map[string]struct{}
}

func NewInformation() *Information {
	return &Information{Sources: make(map[string]struct{})}
}

func (information *Information) Increase() {
	information.Occurrencies ++
}

func (information *Information) AddSource(source string) {
	_, ok := information.Sources[source]
	if !ok {
		information.Sources[source] =  struct{}{}
	}
}
