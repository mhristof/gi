package util

import (
	"bytes"
	"os/exec"
	"sort"

	"github.com/pkg/errors"
)

func Bash(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), errors.Wrap(err, "cannot execute command")
	}

	return stdout.String(), stderr.String(), nil
}

func MapMerge(this, that map[string]int) map[string]int {
	for k, v := range that {
		this[k] += v
	}

	return this
}

func SortMap(this map[string]int) PairList {
	list := make(PairList, len(this))

	i := 0

	for k, v := range this {
		list[i] = Pair{Key: k, Value: v}
		i++
	}

	sort.Sort(sort.Reverse(list))

	return list
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
