# Vali

`Vali` is a service for validating any yaml file. It can be some config or kubernetes resource, for example.

## Build and run locally

To build run the command:

```bash
$> make build
```

To run:

```bash
$> cat ./examples/kubernetes/limit_range.yaml ./vali --config.path=./examples/kubernetes/config.yaml
```

## Validation rules

The example of validation config:

```yaml
rules:
  # rules for metadata
  - field: "apiVersion"
    match: "v1"
  - field: "kind"
    match: "LimitRange"
  - field: "metadata.name"
    match_re: '^[a-zA-Z-]+$'

  # rules for spec
  - field: "spec.limits"
    items:
      - field: "max.cpu"
        match_re: '^\d{0,3}m$'
      - field: "max.memory"
        match_re: '^\dGi|\d{0,3}Mi$'
      - field: "min.cpu"
        match_re: '^\d{0,3}m$'
      - field: "min.memory"
        match_re: '^\dGi|\d{0,3}Mi$'
      - field: "type"
        anyOf: ["Pod", "Container"]
```

All rules are located in "rules" field.

#### Match specific string

```yaml
  - field: "apiVersion"
    match: "LimitRange"
```

#### Match by regular expression

```yaml
  - field: "max.cpu"
    match: '^\d{0,3}m$'
```

#### Any of string

```yaml
  - field: "environment"
    anyOf: ["production", "staging", "development"]
```

#### Required filed

```yaml
  - fields: ["metadata.name", "metadata.labels", "metadata.annotations"]
    required: true
```

#### Case insensitive

```yaml
  - field: "spec.environment"
    anyOf: ["PRODUCTION", "staging", "DEVelopment"]
    caseInsens: true
```

#### Tested fields

```yaml
  - field: "spec.limits"
    items:
      - field: "max.cpu"
        match_re: '^\d{0,3}m$'
      - field: "max.memory"
        match_re: '^\dGi|\d{0,3}Mi$'
```

#### Is a number

```yaml
  - field: "threshold"
    isNumber: true
```

#### Not empty

```yaml
  - filed: "threshold"
    notEmpty: true
```

#### Field should be absent

```yaml
  - filed: "limit"
    absent: true
```

#### Is a map

```yaml
  - filed: "thresholds"
    isMap: true
```

#### Is an array

```yaml
  - filed: "thresholds"
    isArray: true
```
