package usecase

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestFindUsers(t *testing.T) {
	t.Run("should return users", func(t *testing.T) {
		g := NewWithT(t)
		users, err := FindUsers()
		g.Expect(err).Should(BeNil())
		g.Expect(users).Should(HaveLen(1))
		g.Expect(users[0]).Should(HaveField("FirstName", "Heitor"))
	})
}
