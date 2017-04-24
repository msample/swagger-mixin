// swagger-mixin writes the mixed spec to stdout and exits with
// non-zero iff expected collision count of mixed-in paths,
// definitions, parameters and responses does not match the number
// given with the -c option (which defaults to zero if unspecified)
//
// The given Swagger 2.0 files can be YAML or JSON.  YAML input
// requires a .yml or .yaml filename suffix; everything else is
// considered to be in JSON format.  Always writes result in JSON.
//
// This is a proof of concept for a PR to github.com/go-swagger/go-swagger
//
// Install:
//    go get -u github.com/msample/swagger-mixin
//
// Run:
//    swagger-mixin -c 12 test-data/s1.yml test-data/s2.yml test-data/s3.yml > mixed.json
//
// Help, run:
//    swagger-mixin
//
package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/msample/swagger-mixin/mixer"
)

var opts struct {
	ExpectedCollisionCount uint `short:"c" description:"expected # of rejected mixin paths, defs, etc due to existing key. Non-zero exit if does match actual."`
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		// usage will have been printed already
		os.Exit(1)
	}

	if len(args) < 2 {
		log.Fatalln("Nothing to do. Need some swagger files to merge.\nUSAGE: swagger-mixin [-c <expected#Collisions>] <primary-swagger-file> <mixin-swagger-file>...")
	}

	collisions, err := mixer.MixinFiles(args[0], args[1:], os.Stdout)

	if err != nil {
		log.Fatalln(err)
	}

	if collisions != opts.ExpectedCollisionCount {
		if collisions != 0 {
			// use bash $? to get actual # collisions
			// (but has to be non-zero)
			os.Exit(int(collisions))
		}
		os.Exit(254)
	}
}
