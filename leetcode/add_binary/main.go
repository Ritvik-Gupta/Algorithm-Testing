package add_binary

const ZERO = uint8('0')

type StrWrapper struct {
	store *string
}

func (wrapper *StrWrapper) get(index int) uint8 {
	if index < 0 {
		return 0
	}
	return uint8((*wrapper.store)[index]) - ZERO
}

func addBinary(a string, b string) string {
	x, y := StrWrapper{store: &a}, StrWrapper{store: &b}
	carry, result := uint8(0), ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry == 1; i, j = i-1, j-1 {
		sum := x.get(i) + y.get(j) + carry
		carry = sum / 2
		result = string(sum%2+ZERO) + result
	}
	return result
}
