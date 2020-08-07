package pprof

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func TestPProf(t *testing.T) {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	src := RandomNumbers()
	uniq := UniqNumbers(src)
	fmt.Println(len(src), len(uniq), len(src)-len(uniq))
}

func RandomNumbers() []int64 {
	result := make([]int64, 0)
	for i := 0; i < 100000; i++ {
		result = append(result, int64(time.Now().Nanosecond()+i))
	}
	return result
}

func UniqNumbers(in []int64) []int64 {
	out := make([]int64, 0)
	for _, iv := range in {
		if IsInArray(iv, out) {
			continue
		}
		out = append(out, iv)
	}

	return out
}

func IsInArray(aim int64, res []int64) bool {
	for _, v := range res {
		if v == aim {
			fmt.Print(v)
			return true
		}
	}
	return false
}
