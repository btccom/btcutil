package base58

const b58digits_ordered string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func FastBase58Encoding(bin []byte) string {
	binsz := len(bin)
	var i, j, high, zcount, carry int

	for zcount < binsz && bin[zcount] == 0 {
		zcount++
	}

	size := (binsz-zcount)*138/100 + 1
	var buf = make([]byte, size)

	high = size - 1
	for i = zcount; i < binsz; i += 1 {
		j = size - 1
		for carry = int(bin[i]); j > high || carry != 0; j -= 1 {
			carry = carry + 256*int(buf[j])
			buf[j] = byte(carry % 58)
			carry /= 58
		}
		high = j
	}

	for j = 0; j < size && buf[j] == 0; j += 1 {
	}

	var b58 = make([]byte, size-j+zcount)

	if zcount != 0 {
		for i = 0; i < zcount; i++ {
			b58[i] = '1'
		}
	}

	for i = zcount; j < size; i += 1 {
		b58[i] = b58digits_ordered[buf[j]]
		j += 1
	}

	return string(b58)
}
