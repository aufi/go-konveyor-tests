package data

import (
	. "github.com/konveyor/go-konveyor-tests/config"
	"github.com/konveyor/tackle2-hub/api"
)

type JiraInstanceTC struct {
	Identity api.Identity
	JiraUrl  string
	JiraKind string
}

// Set of valid instances for tests and reuse.
// Important: Do not use it directly to not affect other tests.
var (
	JiraCloud = JiraInstanceTC{
		Identity: api.Identity{
			Kind:     "basic-auth",
			User:     Config.JIRA_CLOUD_USERNAME,
			Password: Config.JIRA_CLOUD_PASSWORD,
		},
		JiraUrl:  Config.JIRA_CLOUD_URL,
		JiraKind: "jira-cloud",
	}
	JiraServer = JiraInstanceTC{
		Identity: api.Identity{
			Kind:     "basic-auth",
			User:     Config.JIRA_SERVER_USERNAME,
			Password: Config.JIRA_SERVER_PASSWORD,
		},
		JiraUrl:  Config.JIRA_SERVER_URL,
		JiraKind: "jira-onprem",
	}
	JiraServerBearerToken = JiraInstanceTC{
		Identity: api.Identity{
			Kind: "bearer",
			Key:  Config.JIRA_SERVER_TOKEN,
		},
		JiraUrl:  Config.JIRA_SERVER_URL,
		JiraKind: "jira-onprem",
	}
)