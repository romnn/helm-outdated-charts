## Helm outdated dependencies

Helm plugin to list and update outdated dependencies of a Helm chart.

## Install

```
helm plugin install https://github.com/romnn/helm-outdated --version=master
```

## Usage

```bash
Helm plugin to manage outdated dependencies of a Helm chart.

Examples:
  $ helm outdated list <pathToChart> 										- Checks if there's a newer version of any dependency available in the specified repository.
  $ helm outdated list <pathToChart> --repositories repo1.corp,repo2.corp 	- Checks if there's a newer version of any dependency available only using the given repositories.

  $ helm outdated update <pathToChart> 							- Updates all outdated dependencies to the latest version found in the repository.
  $ helm outdated update <pathToChart> --increment-chart-version	- Updates all outdated dependencies to the latest version found in the repository and increments the version of the Helm chart.
```

### Auto update

This plugin also provides a git integration to help contributing the updated version of the Helm chart generated by the `helm outdated update ...` command to an upstream github.com repository.
This feature is enabled via the `--auto-update` flag.
Minor changes are directly committed to the master branch. Major and potentially breaking changes are submitted via pull requests (PR).  
Using the flag `--only-pull-requests` prevents commits to master and will create a PR instead.

Requirements:  
[1] Git command line tools.  
[2] [Hub](https://github.com/github/hub) for the Github API.

Example:

```bash
helm outdated update <pathToChart> --auto-update --author-name=sapcc-bot --author-email=sapcc-bot@sap.com
```

## BUILD

To build the outdated plugin and test it locally:

```
make build
make remove
make install
```

## RELEASE

Update the version in the [plugin.yaml](plugin.yaml), export the `GORELEASER_GITHUB_TOKEN` (needs `repo` scope) and run `make release`.
