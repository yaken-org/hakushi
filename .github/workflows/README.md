# GitHub Actions CI/CD Workflows

## Backend Test Workflow

This workflow runs automatically on:
- Push to `main` or `develop` branches
- Pull requests to `main` or `develop` branches
- When backend files are changed

### Jobs

#### 1. Test
- Sets up MariaDB service container
- Runs database migrations
- Executes all tests with coverage
- Displays coverage report in CI logs

#### 2. Lint
- Runs golangci-lint
- Checks code formatting with `go fmt`
- Runs `go vet`

#### 3. Build
- Builds the backend binary
- Verifies the build is successful

### Local Testing

To run tests locally using Docker:

```bash
cd backend
make docker-test
```

To run tests without Docker:

```bash
cd backend
make test-coverage
```

### Environment Variables

The CI uses these environment variables for testing:
- `APP_ENV=test`
- `DB_HOST=localhost`
- `DB_PORT=3306`
- `DB_USER=hakushi`
- `DB_PASSWORD=hakushi`
- `DB_NAME=hakushi_test`

### Coverage Reports

Coverage reports are displayed in the CI logs.
To view detailed coverage locally, run:

```bash
cd backend
make test-coverage
# Opens coverage.html in your browser
```