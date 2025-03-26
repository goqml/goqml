package goqml

const (
	QUrlParsingModeTolerant = 0
	QUrlParsingModeStrict   = 1
)

type QUrl struct {
	vptr DosQUrl
}

func (qurl *QUrl) Setup(url string, mode int) {
	qurl.vptr = dos.QUrlCreate(url, mode)
}

func (qurl *QUrl) Delete() {
	if qurl.vptr == nil {
		return
	}
	dos.QUrlDelete(qurl.vptr)
	qurl.vptr = nil
}

func (qurl *QUrl) ToString() string {
	ptr := dos.QUrlToString(qurl.vptr)
	defer dos.CharArrayDelete(ptr)
	return charPtrToString(ptr)
}
