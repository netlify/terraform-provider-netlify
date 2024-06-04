# Contributions

🎉 Thanks for considering contributing to this project! 🎉

These guidelines will help you send a pull request.

If you're submitting an issue instead, please skip this document.

If your pull request is related to a typo or the documentation being unclear, please click on the relevant page's `Edit`
button (pencil icon) and directly suggest a correction instead. Note that most documentation files are auto-generated from code files.

This project was made with ❤️. The simplest way to give back is by starring and sharing it online.

Everyone is welcome regardless of personal background. We enforce a [Code of conduct](CODE_OF_CONDUCT.md) in order to
promote a positive and inclusive environment.

## Local Development

### Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

### Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using `make build`

To generate or update documentation, run `make generate`.

If you updated the `openapi.json` file, you will need to run `make openapi_generate` to update the generated code.

### Testing the Provider Locally

To use the provider, you must [generate a Netlify Personal Access Token](https://docs.netlify.com/cli/get-started/#obtain-a-token-in-the-netlify-ui). The token can be provided to the provider using the `NETLIFY_TOKEN` environment variable, or by using the `token` argument in the provider configuration block.

```hcl
provider "netlify" {
  token = "YOUR_NETLIFY_TOKEN"
}
```

See the [Debugging](https://developer.hashicorp.com/terraform/plugin/debugging) page in the Terraform documentation on how to use the locally-built version of the provider. It is generally easiest to use the instructions under "Terraform CLI Development Overrides" for local testing.

## License

By contributing to Netlify Terraform provider, you agree that your contributions will be licensed under its
[MIT license](LICENSE).