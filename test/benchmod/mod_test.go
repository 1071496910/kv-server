package benchmod

import (
	"testing"
)

func BenchmarkMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//fmt.Fprint(ioutil.Discard, i%4321)
		_ = i + 731
	}
}

/*func BenchmarkShift(b *testing.B) {
	k := (1 << 3) -1
	for i := 0; i < b.N; i++ {
		fmt.Fprint(ioutil.Discard, i&k)
	}
}*/
