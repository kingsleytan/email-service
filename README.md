# Email Service

## To setup DB with Docker
- run command $`docker-compose up`

## Database Dependencies
go get -u cloud.google.com/go/datastore
go get -u github.com/go-sql-driver/mysql
go get -u github.com/si3nloong/goloquent

## Setup environment file
- Go to `.realize-example-yaml` file
- Change file name into `.realize.yaml`
- Update all relevant environment variables here, including Mailgun domain, key etc.
- Install `go get github.com/oxequa/realize`
- Start the server: `realize start`

## POSTMAN
- Documenter: [https://www.getpostman.com/collections/88c9e20f6baed6d70244](https://www.getpostman.com/collections/88c9e20f6baed6d70244)


