package repeater

import "testing"

func assertEQ(t *testing.T, format string, actual, expect interface{}) {
	if actual != expect {
		t.Errorf(format, actual, expect)
	}
}

func TestRead15(t *testing.T) {
	r := NewReader([]byte("HELLOWORLD")) // 10 chars

	{
		// 15バイトのバッファにReadする。10バイト得られる。
		buf11 := make([]byte, 15)
		l, err := r.Read(buf11)
		assertEQ(t, "Read length %d. should be %d", l, 10)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf11[:10]), "HELLOWORLD")
	}
}

func TestRead4(t *testing.T) {

	r := NewReader([]byte("HELLOWORLD")) // 10 chars

	{
		// 4バイトのバッファにReadする。4バイトずつ得られる。
		buf4 := make([]byte, 4)

		l, err := r.Read(buf4)
		assertEQ(t, "Read length %d. should be %d", l, 4)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf4[:4]), "HELL")

		l, err = r.Read(buf4)
		assertEQ(t, "Read length %d. should be %d", l, 4)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf4[:4]), "OWOR")

		l, err = r.Read(buf4)
		assertEQ(t, "Read length %d. should be %d", l, 2)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf4[:2]), "LD")

		// 次の周回が始まる
		l, err = r.Read(buf4)
		assertEQ(t, "Read length %d. should be %d", l, 4)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf4[:4]), "HELL")
	}

}

func TestRead5(t *testing.T) {
	r := NewReader([]byte("HELLOWORLD")) // 10 chars

	{
		// 5バイトのバッファにReadする。5バイトずつ得られる。
		buf5 := make([]byte, 5)

		l, err := r.Read(buf5)
		assertEQ(t, "Read length %d. should be %d", l, 5)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf5[:5]), "HELLO")

		l, err = r.Read(buf5)
		assertEQ(t, "Read length %d. should be %d", l, 5)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf5[:5]), "WORLD")

		// 次の周回が始まる
		l, err = r.Read(buf5)
		assertEQ(t, "Read length %d. should be %d", l, 5)
		assertEQ(t, "Read error %s. should be %s", err, nil)
		assertEQ(t, "Read buffer %s. should be %s", string(buf5[:5]), "HELLO")
	}
}
