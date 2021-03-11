package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	jira "github.com/andygrunwald/go-jira"
)

func GetAllIssues(client *jira.Client, searchString string) ([]jira.Issue, error) {
	last := 0
	var issues []jira.Issue
	for {
		opt := &jira.SearchOptions{
			MaxResults: 1000,
			StartAt:    last,
		}
		chunk, resp, err := client.Issue.Search(searchString, opt)
		if err != nil {
			return nil, err
		}
		total := resp.Total
		if issues == nil {
			issues = make([]jira.Issue, 0, total)
		}
		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)
		if last >= total {
			return issues, nil
		}
	}
}

type Config struct {
	JQL JQLinfo
}
type JQLinfo struct {
	Jql_1 string
	Jql_2 string
}

func main() {
	base := "https://jira.qiniu.io"
	tp := jira.BasicAuthTransport{
		Username: "zengzhaoxia",
		Password: "2wsx3edc#",
		// Username: "qiniu-bot",
		// Password: "t9RJKLZb0t-n",
	}
	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}
	var jql Config
	if _, err := toml.DecodeFile("conf6.toml", &jql); err != nil {
		panic(err)
	}

	//jql := "project = link AND issuetype = 缺陷 ORDER BY created DESC"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql.JQL.Jql_2)
	issues, err := GetAllIssues(jiraClient, jql.JQL.Jql_2)
	if err != nil {
		panic(err)
	}
	fmt.Println(issues)

}
