/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package applications

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestApplicationPackage(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Application package suite")
}

