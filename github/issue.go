package github

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/crosbymichael/octokat"
)

func (g GitHub) IssueInfoCheck(issueHook *octokat.IssueHook) error {
	title := strings.ToLower(issueHook.Issue.Title)

	// we don't care about proposals or features
	if strings.Contains(title, "proposal") || strings.Contains(title, "feature") {
		log.Debugf("Issue is talking about a proposal or feature so ignoring.")
		return nil
	}

	return nil
}

func (g GitHub) LabelIssueComment(issueHook *octokat.IssueHook) error {
	var labelmap map[string]string = map[string]string{
		"#dibs":    "status/claimed",
		"#claimed": "status/claimed",
		"#mine":    "status/claimed",
	}

	repo := getRepo(issueHook.Repo)

	for token, label := range labelmap {
		// if comment matches predefined actions AND author is not bot
		if strings.Contains(strings.ToLower(issueHook.Comment.Body), token) && g.User != issueHook.Sender.Login {
			log.Debugf("Adding label %#v to issue %d", label, issueHook.Issue.Number)
			if err := g.addLabel(repo, issueHook.Issue.Number, label); err != nil {
				return err
			}
			log.Infof("Added label %#v to issue %d", label, issueHook.Issue.Number)
		}
	}

	return nil
}
