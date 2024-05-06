package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/netlify/terraform-provider-netlify/internal/provider"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name netlify

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger flatten swagger.yml -o swagger_flat.json
//go:generate sh -c "cat swagger_flat.json | jq '[., (.paths | map_values(.[] |= del(.tags?)) | {paths: .})] | add' > swagger_go.json"
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -A netlify -f swagger_go.json -t internal -c plumbing --default-scheme=https --with-flatten=full

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/netlify/netlify",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
