package hawk

import (
	"testing"
)

func TestMac_String(t *testing.T) {
	m := &Mac{
		Type: Header,
		Credential: &Credential{
			ID:  "dh37fgj492je",
			Key: "werxhqb98rpaxn39848xrunpaw3489ruxnpa98w4rxn",
			Alg: SHA256,
		},
		Uri:    "http://example.com:8000/resource/1?b=1&a=2",
		Method: "GET",
		Option: &Option{
			TimeStamp: int64(1353832234),
			Nonce:     "j4h3g2",
			Ext:       "some-app-ext-data",
		},
	}

	act, err := m.String()
	if err != nil {
		t.Error("got an error", err.Error())
	}

	// expected value is reference from https://github.com/hueniverse/hawk#protocol-example
	expect := "6R4rV5iE+NPoym+WwjeHzjAGXUtLNIxmo1vpMofpLAE="

	if act != expect {
		t.Error("invalid mac.")
	}
}

func TestTsMac_String(t *testing.T) {
	tm := &TsMac{
		TimeStamp: 1365741469,
		Credential: &Credential{
			ID:  "123456",
			Key: "2983d45yun89q",
			Alg: SHA256,
		},
	}

	act := tm.String()
	expect := "h/Ff6XI1euObD78ZNflapvLKXGuaw1RiLI4Q6Q5sAbM="

	if act != expect {
		t.Error("Invalid TsMac result")
	}
}

func TestPayloadHash_String(t *testing.T) {
	h := &PayloadHash{
		ContentType: "text/plain",
		Payload:     "Thank you for flying Hawk",
		Alg:         SHA256,
	}

	// expected value is reference from https://github.com/hueniverse/hawk#payload-validation
	expect := "Yi9LfIIFRtBEPt74PVmbTF/xVAwPn7ub15ePICfgnuY="
	actual := h.String()

	if actual != expect {
		t.Error("invalid payload hash string.")
	}
}
