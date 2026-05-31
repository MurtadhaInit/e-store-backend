# Image coordinates for the supply-chain verification recipe.
# GHCR path must be lowercase (the registry is case-sensitive); the GitHub repo
# slug passed to --repo is not.
image := "ghcr.io/murtadhainit/e-store-backend"
repo := "MurtadhaInit/e-store-backend"

# List available recipes
default:
    just --list

# Generate Go function from SQL queries
sqlc-gen:
    sqlc generate

# Show schema migration status
goose-status:
    goose status

# Verify a published image's signed attestations (provenance + SBOM).
# Tag defaults to latest, e.g. `just verify-image v1.2.3`
verify-image tag="latest":
    # Provenance — proves who/how/where it was built (repo, commit, workflow, runner).
    # Stronger: pin the exact signer by replacing --repo {{ repo }} with
    #   --signer-workflow {{ repo }}/.github/workflows/ci-publish.yml
    gh attestation verify oci://{{ image }}:{{ tag }} --repo {{ repo }}

    # SBOM — proves what's inside (SPDX inventory of Go modules + OS packages).
    gh attestation verify oci://{{ image }}:{{ tag }} --repo {{ repo }} \
        --predicate-type https://spdx.dev/Document

    # List every attestation attached to the image in the registry.
    cosign tree {{ image }}:{{ tag }}
