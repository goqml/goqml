package goqml

type QUrlParsingMode int

const (
	QUrlParsingModeTolerant QUrlParsingMode = 0
	QUrlParsingModeStrict   QUrlParsingMode = 1
)

type QUrl struct {
	vptr DosQUrl
}

func NewQUrl(url string) *QUrl {
	return NewQUrlWithMode(url, QUrlParsingModeTolerant)
}

func NewQUrlWithMode(url string, mode QUrlParsingMode) *QUrl {
	var qurl QUrl
	qurl.Setup(url, mode)
	return &qurl
}

func (qurl *QUrl) Setup(url string, mode QUrlParsingMode) {
	qurl.vptr = DosQUrlCreate(url, int(mode))
}

func (qurl *QUrl) Delete() {
	if qurl.vptr == nil {
		return
	}
	DosQUrlDelete(qurl.vptr)
	qurl.vptr = nil
}

func (qurl *QUrl) ToString() string {
	ptr := DosQUrlToString(qurl.vptr)
	defer DosCharArrayDelete(ptr)
	return charPtrToString(ptr)
}
