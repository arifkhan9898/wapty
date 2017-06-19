package decode

import (
	"bytes"
	"encoding/hex"
)

const b16Alphabet = "0123456789abcdefABCDEF"

const b16name = "b16"

func init() {
	addCodecC(b16name, codecConstructor(NewB16CodecC))
}

type Base16 struct {
	dec   *decoder
	input string
}

// nolint: gocyclo
func NewB16CodecC(in string) CodecC {
	const (
		itemInvalid itemType = iota
		itemAlphabet
	)

	// emit should write into output what was read up until this point
	// and move l.start to l.pos
	emit := func(d *decoder, t itemType) {
		token := d.input[d.start:d.pos]

		var decodefunc func(string) []byte

		switch t {
		case itemAlphabet:
			decodefunc = func(in string) []byte {
				if len(in) < 2 {
					return []byte(genInvalid(len(in)))
				}

				odd := false
				if len(in)%2 != 0 {
					in = in[:len(in)-1]
					odd = true
				}

				buf, err := hex.DecodeString(in)
				if err != nil {
					return []byte(err.Error())
				}

				if odd {
					buf = append(buf, []byte(genInvalid(1))...)
				}
				return buf
			}

		case itemInvalid:
			decodefunc = func(in string) []byte {
				return []byte(genInvalid(len(in)))
			}
		}

		d.out.Write(decodefunc(token))
		d.start = d.pos
	}

	var (
		startState    stateFn
		invalidState  stateFn
		alphabetState stateFn
	)

	startState = func(d *decoder) stateFn {
		switch n := d.peek(); {
		case bytes.ContainsRune([]byte(b64Alphabet), n):
			return alphabetState
		case n == -1:
			return nil
		default:
			return invalidState
		}
	}

	invalidState = func(d *decoder) stateFn {
		for {
			switch n := d.next(); {
			case bytes.ContainsRune([]byte(b64Alphabet), n):
				d.backup()
				emit(d, itemInvalid)
				return alphabetState

			case n == -1:
				emit(d, itemInvalid)
				return nil
			}
		}
	}

	alphabetState = func(d *decoder) stateFn {
		for {
			switch n := d.next(); {
			case bytes.ContainsRune([]byte(b64Alphabet), n):
				d.acceptRun(b64Alphabet)
				continue

			case n == -1:
				emit(d, itemAlphabet)
				return nil

			default:
				d.backup()
				emit(d, itemAlphabet)
				return invalidState
			}
		}
	}

	return &Base16{
		dec:   newDecoder(in, startState),
		input: in,
	}
}

func (b *Base16) String() string {
	return b16name
}

func (b *Base16) Decode() (output string) {
	return string(b.dec.decode())
}

func (b *Base16) Encode() (output string) {
	return hex.EncodeToString([]byte(b.input))
}

func (b *Base16) Check() (acceptability float64) {
	var c int
	var tot int
	for _, r := range b.input {
		tot++
		if bytes.ContainsRune([]byte(b16Alphabet), r) {
			c++
		}
	}
	//Heuristic to consider uneven strings as less likely to be valid base16
	if delta := tot % 2; delta != 0 {
		tot += delta
	}
	return float64(c) / float64(tot)
}
