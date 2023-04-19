package encryption

func Xor(input, key []byte) (output []byte) {
	klen := len(key)
	output = make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%klen]
	}

	return output
}
