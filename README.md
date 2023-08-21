# CSV to Nginx Redirection Blocks

Dead-simple Go CLI that converts CSV data from stdin into Nginx redirection location blocks and outputs them to stdout. Useful for CI/CD pipelines where you need to generate Nginx redirection rules dynamically.

## Prerequisites

- Basic CSV data with two columns: e.g. `from,to`

## Usage Example

```sh
curl https://example.com/redirects.csv | csv-to-nginx > /etc/nginx/conf.d/redirects.conf

```

## Notes
Gracefully handles missing rows in the CSV data.

Inspiration for this tool came from this [Stack Overflow answer](https://stackoverflow.com/a/61107170). Handy for working with Google Sheets and other data sources that support csv exports.

Highly experimental. Do not use in production.