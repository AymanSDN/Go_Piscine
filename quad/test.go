@@ -2,55 +2,10 @@ package z01

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/bits"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	IntSize = bits.UintSize - 1 // 31 or 63

	//                                  32 bits           64 bits
	MinInt  = -1 << IntSize        // -2147483648  -9223372036854775808
	MaxInt  = 1<<IntSize - 1       //  2147483647   9223372036854775807
	MaxUint = 1<<bits.UintSize - 1 //  4294967295   18446744073709551615

	StrLen   = 13 // Default length of random strings
	SliceLen = 8  // Default length of slices
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))

	// charsets
	Digit = RuneRange('0', '9')         // Decimal digit characters
	Lower = RuneRange('a', 'z')         // Lowercase latin alphabet characters
	Upper = RuneRange('A', 'Z')         // Uppercase latin alphabet characters
	ASCII = RuneRange(' ', '~')         // ASCII printable characters
	Space = strings.Repeat(" ", StrLen) // Spaces characters
	Basic = Lower + Upper               // Lower and Upper characters
	Alnum = Basic + Digit               // Basic and Digit characters
	Words = Alnum + Space               // Alnum and Space characters
)

func init() {
	rand.Seed(nsSince1970)
}

// PrintRune prints a single rune (Unicode code point) and returns any error
// if the encoding or the writing fails.
func PrintRune(r rune) error {
@@ -63,372 +18,3 @@ func PrintRune(r rune) error {
	_, err := os.Stdout.Write(p)
	return err
}

// RuneRange returns a string containing all the valid runes from a to b.
func RuneRange(a, b rune) string {
	var s []rune
	for {
		if utf8.ValidRune(a) {
			s = append(s, a)
		}
		if a == b {
			return string(s)
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

// IntRange returns a slice containing all the int from a to b.
func IntRange(a, b int) (s []int) {
	for {
		s = append(s, a)
		if a == b {
			return
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

// RandIntBetween returns a random int between a and b included.
func RandIntBetween(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := big.NewInt(int64(b))      // b
	n.Sub(n, big.NewInt(int64(a))) // b-a
	n.Add(n, big.NewInt(1))        // b-a+1
	n.Rand(bigRand, n)             // 0 <= n <= b-a
	n.Add(n, big.NewInt(int64(a))) // a <= n <= b
	return int(n.Int64())
}

// RandPosZ returns a random int between 0 and MaxInt included.
func RandPosZ() int { return RandIntBetween(0, MaxInt) }

// RandPos returns a random int between 1 and MaxInt included.
func RandPos() int { return RandIntBetween(1, MaxInt) }

// RandInt returns a random int between MinInt and MaxInt included.
func RandInt() int { return RandIntBetween(MinInt, MaxInt) }

// RandNeg returns a random int between MinInt and 1 included.
func RandNeg() int { return RandIntBetween(MinInt, 1) }

// RandNegZ returns a random int between MinInt and 0 included.
func RandNegZ() int { return RandIntBetween(MinInt, 0) }

// MakeIntFunc returns a slice of ints created by f.
func MakeIntFunc(f func() int) (s []int) {
	i := 0
	for i < SliceLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandPosZ returns a slice of random ints between 0 and MaxInt included.
func MultRandPosZ() []int { return MakeIntFunc(RandPosZ) }

// MultRandPos returns a slice of random ints between 1 and MaxInt included.
func MultRandPos() []int { return MakeIntFunc(RandPos) }

// MultRandInt returns a slice of random ints between MinInt and MaxInt included.
func MultRandInt() []int { return MakeIntFunc(RandInt) }

// MultRandNeg returns a slice of random ints between MinInt and 1 included.
func MultRandNeg() []int { return MakeIntFunc(RandNeg) }

// MultRandNegZ returns a slice of random ints between MinInt and 0 included.
func MultRandNegZ() []int { return MakeIntFunc(RandNegZ) }

// MultRandIntBetween returns a slice of random ints between a and b included.
func MultRandIntBetween(a, b int) []int {
	return MakeIntFunc(func() int { return RandIntBetween(a, b) })
}

// RandRune returns a random printable rune
// (although you may not have the corresponding glyph).
// One-in-ten chance to get a rune higher than 0x10000 (1<<16).
func RandRune() rune {
	ranges := unicode.PrintRanges
	table := ranges[rand.Intn(len(ranges))]
	if rand.Intn(10) == 0 {
		r := table.R32[rand.Intn(len(table.R32))]
		n := uint32(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
		return rune(r.Lo + n*r.Stride)
	} else {
		r := table.R16[rand.Intn(len(table.R16))]
		n := uint16(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
		return rune(r.Lo + n*r.Stride)
	}
}

// RandStr returns a string with l random characters taken from chars.
// If chars is empty, the characters are random printable runes.
func RandStr(l int, chars string) string {
	if l <= 0 {
		return ""
	}
	dst := make([]rune, l)
	if chars == "" {
		for i := range dst {
			dst[i] = RandRune()
		}
	} else {
		src := []rune(chars)
		for i := range dst {
			r := rand.Intn(len(src))
			dst[i] = src[r]
		}
	}
	return string(dst)
}

// RandDigit returns a string containing random Decimal digit characters.
func RandDigit() string { return RandStr(StrLen, Digit) }

// RandLower returns a string containing random Lowercase latin alphabet
// characters.
func RandLower() string { return RandStr(StrLen, Lower) }

// RandUpper returns a string containing random Uppercase latin alphabet
// characters.
func RandUpper() string { return RandStr(StrLen, Upper) }

// RandASCII returns a string containing random ASCII printable characters.
func RandASCII() string { return RandStr(StrLen, ASCII) }

// RandSpace returns a string containing random Spaces characters.
func RandSpace() string { return RandStr(StrLen, Space) }

// RandBasic returns a string containing random Lower and Upper characters.
func RandBasic() string { return RandStr(StrLen, Basic) }

// RandAlnum returns a string containing random Basic and Digit characters.
func RandAlnum() string { return RandStr(StrLen, Alnum) }

// RandWords returns a string containing random Alnum and Space characters.
func RandWords() string { return RandStr(StrLen, Words) }

// MakeStrFunc returns a slice of strings created by f.
func MakeStrFunc(f func() string) (s []string) {
	i := 0
	for i < StrLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandDigit returns a slice of strings containing random Decimal digit
// characters.
func MultRandDigit() []string { return MakeStrFunc(RandDigit) }

// MultRandLower returns a slice of strings containing random Lowercase latin
// alphabet.
func MultRandLower() []string { return MakeStrFunc(RandLower) }

// MultRandUpper returns a slice of strings containing random Uppercase latin
// alphabet.
func MultRandUpper() []string { return MakeStrFunc(RandUpper) }

// MultRandASCII returns a slice of strings containing random ASCII printable
// characters.
func MultRandASCII() []string { return MakeStrFunc(RandASCII) }

// MultRandSpace returns a slice of strings containing random Spaces characters.
func MultRandSpace() []string { return MakeStrFunc(RandSpace) }

// MultRandBasic returns a slice of strings containing random Lower and Upper
// characters.
func MultRandBasic() []string { return MakeStrFunc(RandBasic) }

// MultRandAlnum returns a slice of strings containing random Basic and Digit
// characters.
func MultRandAlnum() []string { return MakeStrFunc(RandAlnum) }

// MultRandWords returns a slice of strings containing random Alnum and Space
// characters.
func MultRandWords() []string { return MakeStrFunc(RandWords) }

// Wrap returns an error composed of the string of err & s
// Append " : " if s is a single-line
func Wrap(err error, s string) error {
	if err == nil || s == "" {
		return err
	}
	if err.Error() == "" {
		return errors.New(s)
	}
	if !strings.Contains(s, "\n") {
		s += " : "
	} else if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	return errors.New(s + err.Error())
}

// ExecOut runs the command name with its args and returns its combined stdout
// and stderr as string.
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
func ExecOut(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()
	out := string(bytes)
	if err == nil {
		return out, nil
	}
	if out == "" {
		return "", Wrap(err, "Command failed")
	}
	return "", errors.New(out)
}

func Format(a ...interface{}) string {
	ss := make([]string, len(a))
	for i, v := range a {
		switch v.(type) {
		case nil:
			ss[i] = "nil" // instead of "<nil>"
		case
			string,
			byte, // uint8
			rune: // int32

			// string     : a double-quoted string safely escaped with Go syntax
			// byte, rune : a single-quoted character literal safely escaped with Go syntax
			ss[i] = fmt.Sprintf("%q", v)
		default:
			// a Go-syntax representation of the value
			ss[i] = fmt.Sprintf("%#v", v)
		}
	}
	return strings.Join(ss, ", ")
}

var valueOf = reflect.ValueOf

func NameOfFunc(fn interface{}) string {
	if rf := runtime.FuncForPC(valueOf(fn).Pointer()); rf != nil {
		name := rf.Name()
		tokens := strings.Split(name, ".")
		return tokens[len(tokens)-1]
	}
	return "unknownFunc"
}

func Call(fn interface{}, args []interface{}) []interface{} {
	// Convert args from []interface{} to []reflect.Value
	vals := make([]reflect.Value, len(args))
	for i, v := range args {
		if v != nil {
			vals[i] = valueOf(v)
		} else {
			vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
	}

	vals = valueOf(fn).Call(vals)

	// Convert the return values from []reflect.Value to []interface{}
	result := make([]interface{}, len(vals))
	for i, v := range vals {
		result[i] = v.Interface()
	}
	return result
}

type Output struct {
	Results []interface{}
	Stdout  string
}

func Monitor(fn interface{}, args []interface{}) (out Output) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalln("Cannot create pipe.")
	}
	os.Stdout = w
	out.Results = Call(fn, args)
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	w.Close()
	out.Stdout = <-outC
	return out
}

func Challenge(t *testing.T, fn1, fn2 interface{}, args ...interface{}) {
	st1 := Monitor(fn1, args)
	st2 := Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		t.Fatalf("%s(%s) == %s instead of %s\n",
			NameOfFunc(fn1),
			Format(args...),
			Format(st1.Results...),
			Format(st2.Results...),
		)
	} else if !reflect.DeepEqual(st1.Stdout, st2.Stdout) {
		t.Fatalf("%s(%s) prints:\n%s\ninstead of:\n%s\n",
			NameOfFunc(fn1),
			Format(args...),
			Format(st1.Stdout),
			Format(st2.Stdout),
		)
	}
}

var compiled = map[string]struct{}{}

func MainOut(pkg string, args ...string) (out string, err error) {
	main := path.Join(pkg, "a.out")
	if _, ok := compiled[main]; !ok {
		compiled[main] = struct{}{}
		if _, err = ExecOut("go", "build", "-o", main, pkg); err != nil {
			return
		}
	}
	out, err = ExecOut(main, args...)
	return
}

func ChallengeMain(t *testing.T, args ...string) {
	challengeProgram(t, "./student/", "./solutions/", args...)
}

func ChallengeMainExam(t *testing.T, args ...string) {
	challengeProgram(t, "../", "../../solutions/", args...)
}

func challengeProgram(t *testing.T, stuPath, solPath string, args ...string) {
	exercise := strings.ToLower(
		strings.TrimPrefix(t.Name(), "Test"))
	out, err := MainOut(stuPath+exercise, args...)
	if err != nil {
		t.Fatal(err)
	}
	correct, err := MainOut(solPath+exercise, args...)
	if err != nil {
		t.Fatal(err)
	}
	if out != correct {
		t.Fatalf("./%s \"%s\" prints %q instead of %q\n",
			exercise, strings.Join(args, "\" \""), out, correct)
	}
}
  32 changes: 1 addition & 31 deletions32  
z01_test.go
@@ -1,10 +1,6 @@
package z01_test

import (
	"fmt"

	"github.com/01-edu/z01"
)
import "github.com/01-edu/z01"

func ExamplePrintRune() {
	z01.PrintRune('0')
@@ -20,29 +16,3 @@ func ExamplePrintRune() {
	// 01
	// ♥
}

func ExampleExecOut() {
	output, err := z01.ExecOut("echo", "test")
	if err == nil {
		fmt.Print(output)
	} else {
		fmt.Print(err)
	}
	// Output: test
}

func ExampleFormat() {
	a := '♥'
	fmt.Println(fmt.Sprint(a))
	fmt.Println(z01.Format(a))

	b := []interface{}{' ', " ", 5, true, nil, []int{1}, [...]int{1}}
	fmt.Println(fmt.Sprint(b...))
	fmt.Println(z01.Format(b...))

	// Output:
	// 9829
	// '♥'
	// 32 5 true <nil> [1] [1]
	// ' ', " ", 5, true, nil, []int{1}, [1]int{1}
}
