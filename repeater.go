package repeater

type Reader struct {
	token []byte
	pos   int
}

func NewReader(token []byte) *Reader {
	return &Reader{token, 0}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	left := len(r.token) - r.pos
	if len(p) >= left {
		// 終端まで出力
		copy(p, r.token[r.pos:])
		r.pos = 0
		return left, nil
	} else {
		copy(p, r.token[r.pos:r.pos+len(p)])
		r.pos += len(p)
		return len(p), nil
	}

}
