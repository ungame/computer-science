package roadmap

import "github.com/ungame/cs-roadmap/subjects"

func Print() {
	for _, subject := range _roadmap {
		subject.Print()
	}
}

var _roadmap = []subjects.Subject{
	{
		Name: "data-structures",
		Items: []subjects.Subject{
			{Name: "array"},
			{Name: "linked-list"},
			{Name: "stack"},
			{Name: "queue"},
			{Name: "hash-table"},
			{
				Name: "tree",
				Items: []subjects.Subject{
					{Name: "binary-tree"},
					{Name: "binary-search-tree"},
					{Name: "full-binary-tree"},
					{Name: "complete-binary-tree"},
					{Name: "balanced-tree"},
					{Name: "unbalanced-tree"},
				},
			},
			{
				Name: "graph",
				Items: []subjects.Subject{
					{Name: "directed-graph"},
					{Name: "undirected-graph"},
					{Name: "spanning-tree"},
					{Name: "adjacency-matrix"},
					{Name: "adjacency-list"},
				},
			},
			{Name: "heap"},
		},
	},
}
