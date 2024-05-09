# Updating swagger.yml

**This is a temporary measure. Feedback will be shared to fix upstream.**

This project uses a modified swagger.yml. Please maintain these instructions and follow them carefully.

1. Take the latest swagger.yml from [netlify/open-api](https://github.com/netlify/open-api/blob/master/swagger.yml).
1. Remove the billing_details property from the `accountMembership` object.
1. Change the response schema of `operation: getAccount` from an array to a single item.