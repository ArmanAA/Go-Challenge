package commandline

import (
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

var Flags []cli.Flag

func init() {
	Flags = []cli.Flag{

		cli.StringFlag{
			// Name
			Name:  "file,f",
			Usage: "file path",
		},
		cli.StringFlag{
			Name:  "workercount,w",
			Usage: "number of concurrent threads",
			Value: "1",
		},
		cli.StringFlag{
			Name:  "ratelimit,r",
			Usage: "rate limit (milliseconds)",
			Value: "1",
		},
	}

}

//ValidateArgs validates the user arguments
func ValidateArgs(c *cli.Context) (string, int, int) {

	_, err := os.Stat(c.String("file"))
	if err != nil {
		log.Fatal("Invalid file path")
	}
	workerCount, err := strconv.Atoi(c.String("workercount"))
	if err != nil {
		log.Fatal("Invalid worker count")
	}
	rateLimit, err := strconv.Atoi(c.String("ratelimit"))
	if err != nil || rateLimit == 0 {
		log.Fatal("Invalid rate limit")
	}
	return c.String("file"), workerCount, rateLimit

}
