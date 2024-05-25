package shortestpath

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

func (g Graph[TValue]) ToDot(value func(TValue) string) (string, error) {
	baseTmpl := `digraph{
{{range .Nodes}}{{.}}
{{end}}}`
	nodeTmpl := `	"{{.From}}" -> "{{.To}}" [label="{{.Label}}"]`
	baseT := template.Must(template.New("base").Parse(baseTmpl))
	nodeT := template.Must(template.New("node").Parse(nodeTmpl))
	type nodeS struct {
		From  string
		To    string
		Label string
	}
	type baseS struct {
		Nodes []string
	}
	out := baseS{Nodes: []string{}}
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			node := nodeS{
				From:  value(n.Value),
				To:    value(e.To.Value),
				Label: fmt.Sprintf("%.2f", e.Distance),
			}
			buf := bytes.Buffer{}
			err := nodeT.Execute(io.Writer(&buf), node)
			if err != nil {
				return "", err
			}
			out.Nodes = append(out.Nodes, buf.String())
		}
	}
	buf := bytes.Buffer{}
	err := baseT.Execute(io.Writer(&buf), out)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
