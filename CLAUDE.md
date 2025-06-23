# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Hakushi is a social image-sharing platform with annotation capabilities, similar to Pinterest but with the ability to add clickable annotations on images. The project uses a Go backend with Echo framework and a Next.js frontend with TypeScript.

## Essential Commands

### Development Environment
- `task up` - Start local development environment with Docker Compose
- `task down` - Stop and remove containers
- `task build` - Build all container images (frontend and backend)

### Frontend Development
Run these commands in the `frontend/` directory:
- `npm run dev` - Start Next.js development server
- `npm run build` - Build production bundle
- `npm run lint` - Run ESLint linting

### Database Migrations
- `task migrate:local` - Apply migrations to local database
- `task create-migrate NAME=migration_name` - Create new migration file with timestamp

## Architecture Overview

### Backend (Go)
The backend follows a layered architecture pattern:
- **Entry**: `main.go` → `internal/server/server.go`
- **Handlers**: `internal/server/handler/` - HTTP request handlers
- **Services**: `internal/service/` - Business logic layer
- **Models**: `internal/model/` - Data structures
- **Config**: `internal/config/` - Environment-based configuration

Key API endpoints:
- `/api/account` - User management
- `/api/post` - Post CRUD operations
- `/api/tag` - Tag management
- `/api/search` - Search functionality

### Frontend (Next.js)
- Uses Next.js 14 with App Router
- Authentication via NextAuth.js with Google OAuth
- API Gateway pattern: Frontend calls `/api/backend/*` which forwards to Go backend
- Image processing: Uploads are resized to 1024x1024 WebP format and stored in Cloudflare R2

### Database Schema
Core tables:
- `user_account` - User profiles with unique username
- `post` - Main content with image_id reference
- `annotation` - X/Y coordinates on images, links to products
- `tag` & `post_tag` - Many-to-many tag relationships
- `product` - Product information for annotations

### Infrastructure
- Local: Docker Compose with Traefik reverse proxy
- Production: Kubernetes deployment
- Database: MariaDB 10.5
- Image Storage: Cloudflare R2 (S3-compatible)

## Development Workflow

1. **Adding a new API endpoint**:
   - Create handler in `backend/internal/server/handler/`
   - Add service method in `backend/internal/service/`
   - Register route in `backend/internal/server/server.go`
   - Create corresponding Next.js API route in `frontend/src/app/api/backend/`

2. **Database changes**:
   - Create migration: `task create-migrate NAME=descriptive_name`
   - Write SQL in generated file under `migrations/`
   - Apply: `task migrate:local`

3. **Frontend-Backend Communication**:
   - Frontend components call Next.js API routes
   - API routes handle auth, image processing, and forward to backend
   - Backend returns JSON responses

## Important Notes

- Images are processed server-side before storage (Sharp library)
- Authentication tokens are managed by NextAuth.js
- Environment configs are in `backend/config/` (development.yaml, production.yaml)
- Frontend environment variables in `frontend/.env.local`
- The project uses Task (taskfile.dev) instead of Make for command orchestration