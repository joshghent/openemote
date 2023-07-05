# OpenEmote

> OpenEmote is an open source API for adding reactions to your blog, website, or app.

## Features

- **Self-hosted:** Keep control of your own data by hosting this on your own server. All you need is Docker and/or Golang + Redis.
- **Flexible reactions:** Users can respond with any emoji or text your define.
- **Real-time data:** Retrieve real-time reaction data for any website.

## Setup

### Hosting

### Add to your website

## API Endpoints

- `GET /`: Fetches all reactions for a certain URL. Pass the URL as a query parameter.

- `POST /`: Create a new reaction for a URL. The request body should be a JSON object containing the `reaction` (the emoji) and the `URL`.

## Configurations

This application uses environment variables for configuration. The following variables are available:

- `REDIS_ADDR`: The address of your Redis instance.
- `REDIS_PASSWORD`: The password for your Redis instance.
- `ALLOWED_URLS`: A comma-separated list of URLs that are allowed to be used with this service.

These variables can be set in a `.env` file in the root directory of the project. For example:

## FAQ

### What happens is someone spams the API?

It saves the data!

### How do I delete a reaction?

To prevent having to add complex security to this app, simply delete it from Redis manually.

### Why not Webmentions?

This is way simpler and doesn't require any authentication.
