/*
© 2019 Red Hat, Inc. and others.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crds

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getSubmarinerCRD", func() {
	When("When called", func() {
		It("Should parse the embedded yaml properly", func() {
			crd, err := getSubmarinerCRD()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(crd.Spec.Names.Kind).Should(Equal("Submariner"))
			Expect(crd.Spec.Versions[0].Name).Should(Equal("v1alpha1"))
		})
	})
})

func TestOperatorCRDs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operator CRDs handling")
}
