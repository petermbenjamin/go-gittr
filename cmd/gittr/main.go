package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	lang string
)

func init() {
	flag.StringVar(&lang, "l", "", "Search by Language")
	flag.Parse()
}

func main() {
	u, err := url.Parse("https://api.github.com/search/repositories")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("q", "language:"+lang)
	q.Set("sort", "stars")
	u.RawQuery = q.Encode()
	fmt.Println(u)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", data)

	var repos Repos
	err = json.Unmarshal(data, &repos)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos.Items {
		fmt.Println("Name:", repo.Name)
		fmt.Println("Desc:", repo.Description)
		fmt.Println("URL :", repo.URL)
		fmt.Println()
	}
}

// type Repos struct {
// 	IncompleteResults bool `json:"incomplete_results"`
// 	Items             []struct {
// 		ArchiveURL       string      `json:"archive_url"`
// 		AssigneesURL     string      `json:"assignees_url"`
// 		BlobsURL         string      `json:"blobs_url"`
// 		BranchesURL      string      `json:"branches_url"`
// 		CloneURL         string      `json:"clone_url"`
// 		CollaboratorsURL string      `json:"collaborators_url"`
// 		CommentsURL      string      `json:"comments_url"`
// 		CommitsURL       string      `json:"commits_url"`
// 		CompareURL       string      `json:"compare_url"`
// 		ContentsURL      string      `json:"contents_url"`
// 		ContributorsURL  string      `json:"contributors_url"`
// 		CreatedAt        string      `json:"created_at"`
// 		DefaultBranch    string      `json:"default_branch"`
// 		DeploymentsURL   string      `json:"deployments_url"`
// 		Description      string      `json:"description"`
// 		DownloadsURL     string      `json:"downloads_url"`
// 		EventsURL        string      `json:"events_url"`
// 		Fork             bool        `json:"fork"`
// 		Forks            int64       `json:"forks"`
// 		ForksCount       int64       `json:"forks_count"`
// 		ForksURL         string      `json:"forks_url"`
// 		FullName         string      `json:"full_name"`
// 		GitCommitsURL    string      `json:"git_commits_url"`
// 		GitRefsURL       string      `json:"git_refs_url"`
// 		GitTagsURL       string      `json:"git_tags_url"`
// 		GitURL           string      `json:"git_url"`
// 		HasDownloads     bool        `json:"has_downloads"`
// 		HasIssues        bool        `json:"has_issues"`
// 		HasPages         bool        `json:"has_pages"`
// 		HasWiki          bool        `json:"has_wiki"`
// 		Homepage         string      `json:"homepage"`
// 		HooksURL         string      `json:"hooks_url"`
// 		HTMLURL          string      `json:"html_url"`
// 		ID               int64       `json:"id"`
// 		IssueCommentURL  string      `json:"issue_comment_url"`
// 		IssueEventsURL   string      `json:"issue_events_url"`
// 		IssuesURL        string      `json:"issues_url"`
// 		KeysURL          string      `json:"keys_url"`
// 		LabelsURL        string      `json:"labels_url"`
// 		Language         string      `json:"language"`
// 		LanguagesURL     string      `json:"languages_url"`
// 		MergesURL        string      `json:"merges_url"`
// 		MilestonesURL    string      `json:"milestones_url"`
// 		MirrorURL        interface{} `json:"mirror_url"`
// 		Name             string      `json:"name"`
// 		NotificationsURL string      `json:"notifications_url"`
// 		OpenIssues       int64       `json:"open_issues"`
// 		OpenIssuesCount  int64       `json:"open_issues_count"`
// 		Owner            struct {
// 			AvatarURL         string `json:"avatar_url"`
// 			EventsURL         string `json:"events_url"`
// 			FollowersURL      string `json:"followers_url"`
// 			FollowingURL      string `json:"following_url"`
// 			GistsURL          string `json:"gists_url"`
// 			GravatarID        string `json:"gravatar_id"`
// 			HTMLURL           string `json:"html_url"`
// 			ID                int64  `json:"id"`
// 			Login             string `json:"login"`
// 			OrganizationsURL  string `json:"organizations_url"`
// 			ReceivedEventsURL string `json:"received_events_url"`
// 			ReposURL          string `json:"repos_url"`
// 			SiteAdmin         bool   `json:"site_admin"`
// 			StarredURL        string `json:"starred_url"`
// 			SubscriptionsURL  string `json:"subscriptions_url"`
// 			Type              string `json:"type"`
// 			URL               string `json:"url"`
// 		} `json:"owner"`
// 		Private         bool    `json:"private"`
// 		PullsURL        string  `json:"pulls_url"`
// 		PushedAt        string  `json:"pushed_at"`
// 		ReleasesURL     string  `json:"releases_url"`
// 		Score           float64 `json:"score"`
// 		Size            int64   `json:"size"`
// 		SSHURL          string  `json:"ssh_url"`
// 		StargazersCount int64   `json:"stargazers_count"`
// 		StargazersURL   string  `json:"stargazers_url"`
// 		StatusesURL     string  `json:"statuses_url"`
// 		SubscribersURL  string  `json:"subscribers_url"`
// 		SubscriptionURL string  `json:"subscription_url"`
// 		SvnURL          string  `json:"svn_url"`
// 		TagsURL         string  `json:"tags_url"`
// 		TeamsURL        string  `json:"teams_url"`
// 		TreesURL        string  `json:"trees_url"`
// 		UpdatedAt       string  `json:"updated_at"`
// 		URL             string  `json:"url"`
// 		Watchers        int64   `json:"watchers"`
// 		WatchersCount   int64   `json:"watchers_count"`
// 	} `json:"items"`
// 	TotalCount int64 `json:"total_count"`
// }

type Repos struct {
	Items []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"items"`
}
