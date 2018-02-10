package mymath

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	MaxInt int64 = 1<<63 - 1
	MinInt int64 = -1 << 63
)

// An absFunc is a function that returns the absolute value of an integer.
type absFunc func(int64) int64

var (
	testInputs  = []int64{MinInt + 1, MinInt + 2, -1, -0, 1, 2, MaxInt - 1, MaxInt}
	testOutputs = []int64{MaxInt, MaxInt - 1, 1, 0, 1, 2, MaxInt - 1, MaxInt}
	testFuncs   = []absFunc{
		AbsWithBranch,
		//AbsWithStdLib, // test failure expected on large numbers
		AbsWithTwosComplement,
		//WithASM,
	}
)

func funcName(v interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	return s[strings.LastIndex(s, ".")+1:]
}

func TestAbs(t *testing.T) {
	for _, f := range testFuncs {
		testName := funcName(f)
		t.Run(testName, func(t *testing.T) {
			for i := 0; i < len(testInputs); i++ {

				actual := f(testInputs[i])
				if actual != testOutputs[i] {
					t.Errorf("%s(%d)", testName, testInputs[i])
					t.Errorf("	input:		%064b (%d)", uint64(testInputs[i]), testInputs[i])
					t.Errorf("	expected:	%064b (%d)", uint64(testOutputs[i]), testOutputs[i])
					t.Errorf("	actual:		%064b (%d)", uint64(actual), actual)
				}
				t.Log(actual)
			}
		})
	}
}

func BenchmarkAbsWithBranch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AbsWithBranch(-1)
	}
}
func BenchmarkAbsWithStdLib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AbsWithStdLib(-1)
	}
}

func BenchmarkAbsWithTwosComplement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AbsWithTwosComplement(-1)
	}
}

/*
func BenchmarkWithASM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WithASM(-1)
	}
}*/
