//go:build tools
// +build tools

package main

import (
	_ "github.com/onsi/ginkgo/v2/ginkgo"
	_ "go.uber.org/mock/mockgen"
)
