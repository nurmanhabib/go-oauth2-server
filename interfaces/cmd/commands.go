package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/nurmanhabib/go-oauth2-server/config"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/hashing"
	"github.com/nurmanhabib/go-oauth2-server/util"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// Commands is list commands.
func Commands(conf *config.Config, db *gorm.DB) cli.Commands {
	return cli.Commands{
		{
			Name:  "db:migrate",
			Usage: "run database migration",
			Action: func(c *cli.Context) error {
				err := db.AutoMigrate(&entity.User{})
				if err != nil {
					return err
				}
				err = db.AutoMigrate(&entity.OauthClient{})
				if err != nil {
					return err
				}
				err = db.AutoMigrate(&entity.OauthAccessGrant{})
				if err != nil {
					return err
				}
				err = db.AutoMigrate(&entity.OauthAccessToken{})
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "create:user",
			Usage: "Create a New User",
			Action: func(ctx *cli.Context) error {
				f := faker.New()

				// the questions to ask
				questions := []*survey.Question{
					{
						Name: "name",
						Prompt: &survey.Input{
							Message: "Please type full name:",
							Default: f.Person().Name(),
						},
						Validate:  survey.Required,
						Transform: survey.Title,
					},
					{
						Name: "email",
						Prompt: &survey.Input{
							Message: "Please type email address:",
							Default: f.Internet().Email(),
						},
						Validate: survey.Required,
					},
					{
						Name: "password",
						Prompt: &survey.Password{
							Message: "Please type password:",
						},
					},
				}

				var err error

				var answers struct {
					Name     string
					Email    string
					Password string
				}

				err = survey.Ask(questions, &answers)
				if err != nil {
					return err
				}

				if answers.Password == "" {
					answers.Password = "password"
				}

				user := &entity.User{
					ID:    uuid.New().String(),
					Name:  answers.Name,
					Email: answers.Email,
				}

				hasher := &hashing.Bcrypt{}
				user.Password, err = hasher.Create(ctx.Context, answers.Password)
				if err != nil {
					return err
				}

				db.WithContext(ctx.Context).Create(user)

				fmt.Print("\nCongratulation!\n\n")
				fmt.Printf("Name: %s\n", user.Name)
				fmt.Printf("Email: %s\n", user.Email)
				fmt.Printf("Password: %s\n\n", answers.Password)

				return nil
			},
		},
		{
			Name:  "create:client",
			Usage: "Create OAuth Client",
			Action: func(ctx *cli.Context) error {
				var err error

				client := &entity.OauthClient{
					ID:     uuid.New().String(),
					Secret: util.RandomAlpha(12),
				}

				f := faker.New()

				// the questions to ask
				questions := []*survey.Question{
					{
						Name: "name",
						Prompt: &survey.Input{
							Message: "Client name:",
							Default: f.Company().Name(),
						},
						Validate:  survey.Required,
						Transform: survey.Title,
					},
					{
						Name: "redirect_uri",
						Prompt: &survey.Multiline{
							Message: "Redirect URIs (separate with newlines or spaces):",
							Default: "https://oauthdebugger.com/debug\nhttps://oauth.pstmn.io/v1/callback",
						},
						Validate: survey.Required,
					},
					{
						Name: "scopes",
						Prompt: &survey.Input{
							Message: "Scopes (separate with spaces):",
							Default: "public",
						},
						Validate: survey.Required,
					},
					{
						Name: "super_app",
						Prompt: &survey.Confirm{
							Message: "Set as Super App?",
							Default: false,
						},
					},
				}

				var answers struct {
					Name        string
					RedirectURI string `survey:"redirect_uri"`
					Scopes      string `survey:"scopes"`
					SuperApp    bool   `survey:"super_app"`
				}

				err = survey.Ask(questions, &answers)
				if err != nil {
					return err
				}

				client.Name = answers.Name
				client.RedirectURI = answers.RedirectURI
				client.Scopes = answers.Scopes
				client.SuperApp = answers.SuperApp

				db.WithContext(ctx.Context).Create(client)

				redirectURIs := strings.Fields(client.RedirectURI)
				query := url.Values{
					"client_id":     {client.ID},
					"redirect_uri":  {redirectURIs[0]},
					"scopes":        {client.Scopes},
					"response_type": {"code"},
				}

				authorize := url.URL{Path: "/oauth/authorize"}
				authorize.RawQuery = query.Encode()
				authorize.Host = fmt.Sprintf(":%d", conf.Port)
				authorize.Scheme = "http"

				fmt.Print("\nCongratulation!\n\n")
				fmt.Printf("Client ID: %s\n", client.ID)
				fmt.Printf("Client Secret: %s\n\n", client.Secret)
				fmt.Printf("Authorize URI Test: %s\n\n", authorize.String())

				return nil
			},
		},
	}
}
