# swagger-mixin
--
swagger-mixin writes the mixed spec to stdout and exits with non-zero iff
expected collision count of mixed-in paths, definitions, parameters and
responses does not match the number given with the -c option (which defaults to
zero if unspecified)

The given Swagger 2.0 files can be YAML or JSON. YAML input requires a .yml or
.yaml filename suffix; everything else is considered to be in JSON format.
Always writes result in JSON.

Install:

    go get -u github.com/msample/swagger-mixin

Run:

    swagger-mixin -c 3 primary.yml mixin1.yml mixin2.json > mixed.json
