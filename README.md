# Google Takeout Sucks Auth

This is the auth service for the CLI tool [Google Takeout Sucks](https://github.com/Fallenstedt/google-takeout-sucks)

## Running Locally

Use Google Cloud and create a Web application client using the Google Auth Platform. The client should have two Google drive scopes `/auth/drive.metadata.readonly` and `/auth/drive.readonly`.

The client should have an authorized JavaScript origin of `http://localhost:8080` and it should have an Authorized redirct URI of `http://localhost:8080/oauth2/callback`

Download the credentials and save to the root of directory as `dev-credentials.json`

## Building

Before you build, ensure you have a production Google Auth Client created and credentials are saved to the root of the directory as `credentials.json`. Refer to `Running Locally` for steps.

To build:

1. Build

   go build -o bin/server ./cmd/server

2. Run

   ./bin/server

The server listens on :8080 with several endpoints:

- GET /login -> initiates Google Oauth2 Workflow
- GET /oauth2/callback -> exchange code for auth token and complete oauth2 workflow
