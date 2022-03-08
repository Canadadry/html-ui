package generator

import (
	"sort"
	"strings"
)

type UniqueClasses map[string]struct{}

func UniqueClassesFrom(v string) UniqueClasses {
	vs := strings.Split(v, " ")
	u := UniqueClasses{}
	for _, s := range vs {
		u.Add(s)
	}
	return u
}

func (u UniqueClasses) Add(c string) {
	u[c] = struct{}{}
}

func (u UniqueClasses) Has(c string) bool {
	_, ok := u[c]
	return ok
}

func (u UniqueClasses) Remove(c string) {
	delete(u, c)
}

func (u UniqueClasses) Merge(u2 UniqueClasses) {
	for c := range u2 {
		u.Add(c)
	}
}

func (u UniqueClasses) Sorted() []string {
	classes := make([]string, 0, len(u))
	for c := range u {
		classes = append(classes, c)
	}
	sort.Strings(classes)
	return classes
}

func (u UniqueClasses) String() string {
	return strings.Join(u.Sorted(), " ")
}
