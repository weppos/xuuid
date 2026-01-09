# xuuid

A simple command-line UUID generator supporting multiple UUID versions.

## Installation

```shell
go install github.com/weppos/xuuid@latest
```

or build from source:

```shell
git clone https://github.com/weppos/xuuid.git
cd xuuid
make build
```

The binary will be created in the current directory as `xuuid`.

## Usage

```shell
xuuid [flags]
```

Generate a single UUID v7 (default):

```shell
xuuid

# 018f1234-5678-7890-abcd-ef1234567890
```

Generate a UUID v4:

```bash
xuuid --v4

# 550e8400-e29b-41d4-a716-446655440000
```

Generate multiple UUIDs:

```bash
xuuid --count 2

# 018f1234-5678-7890-abcd-ef1234567890
# 018f1235-5678-7890-abcd-ef1234567891
```

Generate a UUID v7 in hex (no dashes):

```bash
xuuid --hex

# 018f123456787890abcdef1234567890
```

Generate a UUID v7 as a URN:

```bash
xuuid --urn

# urn:uuid:018f1234-5678-7890-abcd-ef1234567890
```

Show all available flags:

```shell
xuuid --help
```
