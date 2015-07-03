Tag
===

A project for automated semantic tagging.

## Usage

Commit tag locally

```bash
# Patch
$ semantic --patch
$ 0.0.1

# Minor
$ semantic --minor
$ 0.1.0

# Major
$ semantic --major
$ 1.0.0
```

Commit tag and push

```bash
$ TAG=$(semantic --patch)
$ git push origin $TAG
```

Tag against a different branch

```bash
$ semantic --patch --branch=releases
$ 0.0.1
```
