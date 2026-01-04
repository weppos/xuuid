#!/bin/bash

set -e

echo "Testing default (v7)..."
./xuuid > /dev/null || (echo "FAILED: default (v7)" && exit 1)

echo "Testing --v4..."
./xuuid --v4 > /dev/null || (echo "FAILED: --v4" && exit 1)

echo "Testing --v6..."
./xuuid --v6 > /dev/null || (echo "FAILED: --v6" && exit 1)

echo "Testing --v7..."
./xuuid --v7 > /dev/null || (echo "FAILED: --v7" && exit 1)

echo "Testing --md5..."
./xuuid --md5 > /dev/null || (echo "FAILED: --md5" && exit 1)

echo "Testing --sha1..."
./xuuid --sha1 > /dev/null || (echo "FAILED: --sha1" && exit 1)

echo "Testing --count 5..."
test $(./xuuid --count 5 | wc -l) -eq 5 || (echo "FAILED: --count 5" && exit 1)

echo "Testing --help..."
./xuuid --help > /dev/null 2>&1 || true

echo "All tests passed!"
