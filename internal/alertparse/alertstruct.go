// base on https://github.com/prometheus/alertmanager/blob/main/template/template.go

package alertparse

import (
	"strings"
	"sort"
	"fmt"
        "time"
//	"encoding/json"
)

// Pair is a key/value string pair.
type Pair struct {
	Name, Value string
}

// Pairs is a list of key/value string pairs.
type Pairs []Pair

// Names returns a list of names of the pairs.
func (ps Pairs) Names() []string {
	ns := make([]string, 0, len(ps))
	for _, p := range ps {
		ns = append(ns, p.Name)
	}
	return ns
}

// Values returns a list of values of the pairs.
func (ps Pairs) Values() []string {
	vs := make([]string, 0, len(ps))
	for _, p := range ps {
		vs = append(vs, p.Value)
	}
	return vs
}

func (ps Pairs) String() string {
	b := strings.Builder{}
	for i, p := range ps {
		b.WriteString(p.Name)
		b.WriteRune('=')
		b.WriteString(p.Value)
		if i < len(ps)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}

// KV is a set of key/value string pairs.
type KV map[string]string

// SortedPairs returns a sorted list of key/value pairs.
func (kv KV) SortedPairs() Pairs {
	var (
		pairs     = make([]Pair, 0, len(kv))
		keys      = make([]string, 0, len(kv))
		sortStart = 0
	)
	for k := range kv {
		fmt.Println(k)
		/*if k == string(model.AlertNameLabel) {
			keys = append([]string{k}, keys...)
			sortStart = 1
		} else {
			keys = append(keys, k)
		}*/
	}
	sort.Strings(keys[sortStart:])

	for _, k := range keys {
		pairs = append(pairs, Pair{k, kv[k]})
	}
	return pairs
}


// Names returns the names of the label names in the LabelSet.
func (kv KV) Names() []string {
	return kv.SortedPairs().Names()
}

// Values returns a list of the values in the LabelSet.
func (kv KV) Values() []string {
	return kv.SortedPairs().Values()
}

func (kv KV) String() string {
	return kv.SortedPairs().String()
}

// Data is the data passed to notification templates and webhook pushes.
//
// End-users should not be exposed to Go's type system, as this will confuse them and prevent
// simple things like simple equality checks to fail. Map everything to float64/string.
type Data struct {
	Status   string   `json:"status"`
	Data     *Alerts  `json:"data"`
        //Data json.RawMessage `json:"data"`
}

// Alerts is a list of Alert objects.
type Alerts struct {
	Alerts   []Alert   `json:"alerts"`
}

// Alert holds one alert for notification templates.
type Alert struct {
	State        string    `json:"state"`
	Labels       KV        `json:"labels"`
	Annotations  KV        `json:"annotations"`
	ActiveAt     time.Time `json:"activeAt"`
	value	     float64   `json:"value"`
}


