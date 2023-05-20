package user_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/domain/user/usecase"
)

var _ = Describe("User", func() {
	Describe("Find users", func() {
		Context("with one user", func() {
			It("should be ok", func() {
				users, err := usecase.FindUsers()
				Expect(err).Should(BeNil())
				Expect(users).Should(HaveLen(1))
			})
		})
	})
})
