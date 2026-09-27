# Project Tracker Backend

Backend API for a project tracking system built with Go.

## Technologies

- Go
- Fiber
- PostgreSQL
- sqlx
- Redis
- Docker
- REST API

## Features

- Project management
- Member management
- Task management
- Task status tracking
- Project progress calculation
- Member contribution calculation
- Redis caching
- Cache invalidation
- API validation
- Unit and HTTP tests

## Project Structure

```text
cmd/
  api/

internal/
  app/
  cache/
  domain/
  dtos/
  handlers/
  repos/
  routes/
  services/

migrations/