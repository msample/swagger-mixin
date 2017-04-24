# swagger-mixin
--
swagger-mixin writes the mixed spec to stdout and exits with non-zero iff
expected collision count of mixed-in paths, definitions, parameters and
responses does not match the number given with the -c option (which defaults to
zero if unspecified)

The given Swagger 2.0 files can be YAML or JSON. YAML input requires a .yml or
.yaml filename suffix; everything else is considered to be in JSON format.
Always writes result in JSON.

This is a proof of concept for a PR to github.com/go-swagger/go-swagger

Install:

    go get -u github.com/msample/swagger-mixin

Run:

    swagger-mixin -c 12 test-data/s1.yml test-data/s2.yml test-data/s3.yml > mixed.json

Help, run:

    swagger-mixin
