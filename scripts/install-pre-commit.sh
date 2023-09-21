#!/usr/bin/env bash

set -ex

brew install pre-commit

cat <<EOF > .pre-commit-config.yaml
exclude: CODEOWNERS
fail_fast: true

repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
    -   id: trailing-whitespace
    -   id: end-of-file-fixer
    -   id: check-yaml
    -   id: check-added-large-files
-   repo: https://github.com/golangci/golangci-lint
    rev: v1.53.3
    hooks:
    -   id: golangci-lint
-   repo: https://github.com/sqlfluff/sqlfluff
    rev: 2.1.1
    hooks:
    -   id: sqlfluff-lint
        args:
        - --dialect=postgres
        files: ".*.sql"
    -   id: sqlfluff-fix
        args:
        - --dialect=postgres
        files: ".*.sql"
-   repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.1
    hooks:
    -   id: go-fmt
    -   id: go-imports
    -   id: go-cyclo
        args:
        - -ignore=factory
        - -over=12
-   repo: https://github.com/Yelp/detect-secrets
    rev: v1.4.0
    hooks:
    -   id: detect-secrets
-   repo: https://github.com/daveshanley/vacuum
    rev: v0.1.7
    hooks:
    -   id: vacuum
        args:
        - -d
        files: "api/openapi-spec/openapi.yaml"
EOF

pre-commit autoupdate
pre-commit install --overwrite
