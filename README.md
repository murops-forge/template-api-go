# template-api-go

Go HTTP service template used by [Forge](https://github.com/murops/forge).

## Running

```bash
go test ./...
go run ./cmd/server
curl 'localhost:8080/hello?name=forge'
```

Or with Docker:

```bash
docker build -t app .
docker run --rm -p 8080:8080 app
```

## CI

`.github/workflows/ci.yml` runs on pushes to `main`, `homol` and `desen` and on every pull request:

- `test`: gofmt, go vet and tests with coverage
- `trivy`: scans the Dockerfile and the built image, fails on fixable HIGH/CRITICAL findings and uploads the report to the Security tab
- `sonar`: SonarCloud analysis with quality gate, skipped when the organization has no `SONAR_TOKEN`

These are the checks Forge requires on `main` and `homol`.
