package main;

import "bytes"
func backspaceCompare(a string, b string) bool {
	return bytes.Compare(apply(a), apply(b)) == 0;
}

func apply(s string) []byte{
	aBuff := bytes.NewBufferString("")
	aSkip := 0;

	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '#' {
			aSkip++;
		} else {
			if aSkip > 0{
				aSkip --;
			}else{
				aBuff.WriteByte(s[i]);
			}
		}
	}

	return aBuff.Bytes();

}