package search

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	JQL JQLinfo
}
type JQLinfo struct {
	Jql_1 string
	Jql_2 string
}

func a() {
	var jql Config
	if _, err := toml.DecodeFile("conf6.toml", &jql); err != nil {
		panic(err)
	}
	fmt.Println(jql.JQL.Jql_1)
	fmt.Println(jql.JQL.Jql_2)
}
