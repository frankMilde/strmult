package strmult

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/frankMilde/rol/testutils"
)

func Test_isMultiLined(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   `1234567890`,
			want: false,
		},
		{
			// test regular ASCII
			in:   "it is\nmulti-lined",
			want: true,
		},
		{
			// test regular ASCII
			in: `it is 
			multi-lined, too`,
			want: true,
		},
	}

	for _, test := range tests {
		got := IsMultiLined(test.in)
		if !reflect.DeepEqual(test.want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall isMultiLined(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.in, test.want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()
}

func Test_splitStringIntoLines(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{
			in:   `test`,
			want: []string{"test"},
		},
		{
			in:   "test\ntest",
			want: []string{"test", "test"},
		},
		{
			in: `test
test`,
			want: []string{"test", "test"},
		},
		{
			in: `test1
test2 test3
test4`,
			want: []string{"test1", "test2 test3", "test4"},
		},
		{
			in:   "",
			want: []string{},
		},
	}

	for _, test := range tests {
		got := SplitStringIntoLines(test.in)
		if !reflect.DeepEqual(test.want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall splitStringIntoLines(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.in, test.want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()

}

func Test_TotalLines(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   `test`,
			want: 1,
		},
		{
			in:   "test\ntest",
			want: 2,
		},
		{
			in: `test
test`,
			want: 2,
		},
		{
			in: `test1
test2 test3
test4`,
			want: 3,
		},
		{
			in:   "",
			want: 0,
		},
	}

	for _, test := range tests {
		got := TotalLines(test.in)
		if !reflect.DeepEqual(test.want, got) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\ncall TotalLines(%#v)\n\texp: %#v\n\n\tgot: %#v\n\n",
				filepath.Base(file), line, test.in, test.want, got)
			t.FailNow()
		}
	}
	testutils.Cleanup()
}
