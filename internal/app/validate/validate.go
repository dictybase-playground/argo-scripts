package validate

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

// ValidateServerArgs validates that the necessary flags are not missing
func ValidateServerArgs(c *cli.Context) error {
	for _, p := range []string{
		"github-access-token",
	} {
		if len(c.String(p)) == 0 {
			return cli.NewExitError(
				fmt.Sprintf("argument %s is missing", p),
				2,
			)
		}
	}
	return nil
}
