package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

func Keys[K comparable, V any](input map[K]V) []K {
	keys := make([]K, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](input map[K]V) []V {
	values := make([]V, 0, len(input))
	for _, v := range input {
		values = append(values, v)
	}
	return values
}

func ContainsAny(s string, substrings ...string) bool {
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}

func MatchesAny(s string, substrings ...string) bool {
	for _, substring := range substrings {
		if ok, _ := regexp.MatchString(s, substring); ok {
			return true
		}
	}
	return false
}

func Title(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to of "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && index > 0 {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func Capitalize(input string) string {
	if input == "" {
		return ""
	}
	return strings.ToUpper(input[:1]) + input[1:]
}

func Json(obj any) template.HTML {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return template.HTML(`<span class="json">` + string(b) + `</span>`)
}

func If[T any](cond bool, v1, v2 T) T {
	if cond {
		return v1
	} else {
		return v2
	}
}

func Map[U, V any](list []U, mapper func(U) V) []V {
	var newList = make([]V, 0, len(list))
	for _, i := range list {
		newList = append(newList, mapper(i))
	}
	return newList
}

func FilterMap[K comparable, V any](input map[K]V, predicate func(V) bool, sorter func(V, V) bool) []V {
	var list []V
	for _, v := range input {
		if predicate(v) {
			list = append(list, v)
		}
	}
	sort.Slice(list, func(i, j int) bool { return sorter(list[i], list[j]) })
	return list
}

type Identifiable interface {
	Id() int
}

func Find[U any](list []*U, predicate func(*U) bool) (*U, bool) {
	for _, x := range list {
		if predicate(x) {
			return x, true
		}
	}
	return nil, false
}

func FindInMap[U any](list map[int]*U, predicate func(*U) bool) (int, *U, bool) {
	for id, x := range list {
		if predicate(x) {
			return id, x, true
		}
	}
	return -1, nil, false
}

func FirstInMap(a any, b string) bool {
	ks := reflect.ValueOf(a).MapKeys()
	sort.Slice(ks, func(i, j int) bool { return ks[i].String() < ks[j].String() })
	return ks[0].String() == b
}

func Strip(html template.HTML) string {
	r := regexp.MustCompile(`<.*?>`)
	return r.ReplaceAllString(string(html), "")
}

func String(html template.HTML) string {
	return string(html)
}

/*
	ArtFormShortDesc

	Gets a shorter description for the popovers by modifying the description data.

	Note there is no way yet to highlight the
  	author, as that info is not tracked by the XML, and would need a more
  	intensive patch to the backend (keeping a map of art form -> author in 
  	world context). Probably won't implement this for popovers unless I
  	get really really bored, as it's a massive pain for minimal gain.
*/
func ArtFormShortDesc(name string, desc string) string {
	if len(strings.TrimSpace(desc)) == 0 {
		return ""
	}

	// find where the title (plus "is") ends
	i := len(name) + 3

	j := strings.Index(desc, ".")
	if j == -1 {
		return desc	// no period found; text is one sentence
	}

	return strings.TrimSpace(desc[i:j+1])
}
