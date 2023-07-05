# OpenEmote

> OpenEmote is an open source API for adding reactions to your blog, website, or app.

## Features

- **Self-hosted:** Keep control of your own data by hosting this on your own server. All you need is Docker and/or Golang + Redis.
- **Flexible reactions:** Users can respond with any emoji or text your define.
- **Real-time data:** Retrieve real-time reaction data for any website.

## Setup

### Hosting

You can deploy this anywhere you have Docker installed and an open port.
You can deploy this to DigitalOcean for $6/month.

Use the following `docker-compose.yml` file to get started:
Replace `YOUR_DOMAIN_HERE` with your domain - this can be a subdomain.

```yml
version: "3"
services:
  web:
    image: "ghcr.io/joshghent/openemote:latest"
    restart: always
    env_file:
      - .env
    depends_on:
      - redis
    networks:
      - traefik-net
    labels:
      - traefik.enable=true
      - traefik.http.routers.openemote.rule=Host(`YOUR_DOMAIN_HERE`)
      - traefik.http.routers.openemote.entrypoints=web
  redis:
    image: "redis:alpine"
    restart: always
    networks:
      - traefik-net
    command: redis-server --requirepass ${REDIS_PASSWORD}
    environment:
      REDIS_PASSWORD: ${REDIS_PASSWORD}
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
  traefik:
    image: traefik:v2.4
    command:
      - --api.insecure=true
      - --providers.docker=true
      - --providers.docker.exposedbydefault=false
      - --entrypoints.web.address=:80
    ports:
      - 80:80
      - 8080:8080
    networks:
      - traefik-net
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro

networks:
  traefik-net:

volumes:
  redis-data:
```

### Add to your website

#### React

```jsx
import React, { useState } from "react";
import axios from "axios";

const ReactionButton = () => {
  const [hasReacted, setHasReacted] = useState(false);

  const sendReaction = async () => {
    const reaction = "👍";
    const url = window.location.href; // current page url

    try {
      const response = await axios.post("YOUR_DOMAIN_HERE", { reaction, url });
      if (response.status === 201) {
        setHasReacted(true);
      }
    } catch (error) {
      console.error("Error sending reaction: ", error);
    }
  };

  return (
    <button onClick={sendReaction} disabled={hasReacted}>
      {hasReacted ? "Reacted" : "👍"}
    </button>
  );
};

export default ReactionButton;
```

#### HTML

```html
<button onclick="sendReaction()">👍</button>
<script>
  function sendReaction() {
    const reaction = "👍";
    const url = "joshghent.com/example-post";
    fetch("YOUR_DOMAIN_HERE", {
      method: "POST",
      body: JSON.stringify({ url, reaction }),
    });
  }
</script>
```

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
