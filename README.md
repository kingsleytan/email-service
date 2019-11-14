# Email Service

## To setup DB with Docker
- run command $`docker-compose up`

## Setup environment file
- Go to `.realize-example-yaml` file
- Change file name into `.realize.yaml`
- Update all relevant environment variables here, including Mailgun domain, key etc.
- Install `go get github.com/oxequa/realize`
- Start the server: `realize start`

## POSTMAN
- Documenter: [https://www.getpostman.com/collections/88c9e20f6baed6d70244](https://www.getpostman.com/collections/88c9e20f6baed6d70244)

## API
- BaseURL: localhost:2000/v1
### Healthcheck
- `GET /`
- This is to check that API is working fine.
- Response:
```
{
    "message": "API is working fine, ENV: development"
}
```
### Send Email with Template
- `POST /send/:template-name`
- Send email with specific template on Mailgun. Can use this to test email templates before publishing.

- Request:

| Parameter        | Type           | Description  |  Required | Example |
| ------------- |:-------------:| -----:| -----:| -----:|
| id     | string | Auto generated internal id for email | Yes | "12345" |
| to     | string | Recipient of email | Yes | "example@email.com" |
| from     | string | Sender email. If not provided, use a predefined email | No | "receive@email.com" |
| domain     | string | Domain to be used for when sending the email. If not provided, use a predefined domain | No | "abc.com" |
| template_data     | object | Variables to be injected to email templates | Yes | {"title": "test title", "body": "test body"} |
| template     | string | Name of email template | Yes | "forgotpassword-2019-11-13.135207" |
| reference_id     | string | Optional unique id for identifying request | No | "ref-123" |
| status     | string | Current status of email | No | "active" |
| events     | string | Array of webhook payload sent from mailgun in descending order, ie latest event first | No | "events" |

- Response:

| Parameter        | Type           | Description  | Example |
| ------------- |:-------------:| -----:| -----:|
| to     | string | Recipient of email  | "example@email.com" |
| from     | string | Sender email. If not provided, use a predefined email  | "receive@email.com" |
| domain     | string | Domain to be used for when sending the email. If not provided, use a predefined domain  | "abc.com" |
| template_data     | object | Variables to be injected to email templates  | {"title": "test title", "body": "test body"} |
| template     | string | Name of email template  | "forgotpassword-2019-11-13.135207" |
| reference_id     | string | Optional unique id for identifying request | "ref-123" |
| created     | string | ISO 8601 formatted datetime | "2019-11-14T15:16:16+07:00" |
| updated     | string | ISO 8601 formatted datetime | "2019-11-14T15:16:16+07:00" |
| result     | string | Response from Mailgun | "result: Queued: <id .mailgun.org>" |


