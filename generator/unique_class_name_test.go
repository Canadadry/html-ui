package generator

import (
	"reflect"
	"testing"
)

func TestSorted(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		{
			in:  []string{},
			out: []string{},
		},
		{
			in:  []string{"a"},
			out: []string{"a"},
		},
		{
			in:  []string{"b", "a"},
			out: []string{"a", "b"},
		},
		{
			in:  []string{"b", "a", "b"},
			out: []string{"a", "b"},
		},
	}

	for _, tt := range tests {
		u := UniqueClasses{}
		for _, v := range tt.in {
			u.Add(v)
		}

		result := u.Sorted()

		if !reflect.DeepEqual(result, tt.out) {
			t.Fatalf("[%s] failed \ngot \n-%v-\nexp \n-%v-\n", tt, result, tt.out)
		}
	}

}

func TestString(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{
			in:  []string{},
			out: "",
		},
		{
			in:  []string{"a"},
			out: "a",
		},
		{
			in:  []string{"b", "a"},
			out: "a b",
		},
		{
			in:  []string{"b", "a", "b"},
			out: "a b",
		},
	}

	for _, tt := range tests {
		u := UniqueClasses{}
		for _, v := range tt.in {
			u.Add(v)
		}

		result := u.String()

		if result != tt.out {
			t.Fatalf("[%s] failed \ngot \n-%v-\nexp \n-%v-\n", tt, result, tt.out)
		}
	}

}

func TestMerge(t *testing.T) {
	tests := []struct {
		in  [][]string
		out string
	}{
		{
			in:  [][]string{},
			out: "",
		},
		{
			in:  [][]string{[]string{"a"}},
			out: "a",
		},
		{
			in:  [][]string{[]string{"b", "a"}, []string{"c", "d"}},
			out: "a b c d",
		},
		{
			in:  [][]string{[]string{"b", "a", "b"}, []string{"b", "a"}},
			out: "a b",
		},
	}

	for _, tt := range tests {
		u := UniqueClasses{}
		for _, list := range tt.in {
			u2 := UniqueClasses{}
			for _, c := range list {
				u2.Add(c)
			}
			u.Merge(u2)
		}

		result := u.String()

		if result != tt.out {
			t.Fatalf("[%s] failed \ngot \n-%v-\nexp \n-%v-\n", tt, result, tt.out)
		}
	}

}

func TestHas(t *testing.T) {
	u := UniqueClassesFrom("a b")
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "",
			out: false,
		},
		{
			in:  "a",
			out: true,
		},
		{
			in:  "b",
			out: true,
		},
		{
			in:  "c",
			out: false,
		},
	}

	for i, tt := range tests {
		result := u.Has(tt.in)
		if result != tt.out {
			t.Fatalf("[%d] failed \ngot \n-%v-\nexp \n-%v-\n", i, result, tt.out)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		u   UniqueClasses
		in  string
		out string
	}{
		{
			u:   UniqueClassesFrom("a b"),
			in:  "",
			out: "a b",
		},
		{
			u:   UniqueClassesFrom("a b"),
			in:  "a",
			out: "b",
		},
		{
			u:   UniqueClassesFrom("a b"),
			in:  "c",
			out: "a b",
		},
	}

	for i, tt := range tests {
		tt.u.Remove(tt.in)
		result := tt.u.String()

		if result != tt.out {
			t.Fatalf("[%d] failed \ngot \n-%v-\nexp \n-%v-\n", i, result, tt.out)
		}
	}

}
