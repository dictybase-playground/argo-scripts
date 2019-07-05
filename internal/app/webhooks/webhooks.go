package webhooks

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/dictybase-playground/argo-scripts/internal/logger"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
)

type WebhookInput struct {
	HookURL string   `yaml:"hookURL"`
	Secret  string   `yaml:"secret"`
	Owner   string   `yaml:"owner"`
	Repos   []string `yaml:"repos"`
}

type WebhookOutput struct {
	Hooks []Repo `yaml:"hooks"`
}

type Repo struct {
	Name string `yaml:"repo"`
	ID   string `yaml:"id"`
}

func InitializeGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

// RunCreateWebhooks creates GitHub webhooks based on an input YAML and provides
// an output YAML with its results.
func RunCreateWebhooks(c *cli.Context) error {
	l, err := logger.GetLogger(c)
	if err != nil {
		return fmt.Errorf("could not get logger %s", err)
	}
	i := &WebhookInput{}
	o := &WebhookOutput{}
	input, err := ioutil.ReadFile(c.String("input-file"))
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	err = yaml.Unmarshal(input, i)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	client := InitializeGitHubClient(c.String("github-access-token"))
	for _, name := range i.Repos {
		h := &github.Hook{
			Events: []string{"push"},
			Config: map[string]interface{}{
				"url":          i.HookURL + "/" + name,
				"content_type": "json",
				"insecure_ssl": "0",
				"secret":       i.Secret,
			},
		}
		hook, resp, err := client.Repositories.CreateHook(context.Background(), i.Owner, name, h)
		if err != nil {
			fmt.Errorf("test")
			return cli.NewExitError(err.Error(), 2)
		}
		if resp.StatusCode > 299 {
			l.Infof("request failed with %v", resp.Status)
			return cli.NewExitError(err.Error(), 2)
		}
		l.Infof("successfully created webhook for repo %s", name)
		o.Hooks = append(o.Hooks, Repo{
			ID:   strconv.Quote(strconv.FormatInt(hook.GetID(), 10)),
			Name: name,
		})
	}

	d, err := yaml.Marshal(&o)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	err = ioutil.WriteFile(c.String("output-file"), d, 0644)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	l.Infof("successfully created file %s", c.String("output-file"))
	return nil
}
