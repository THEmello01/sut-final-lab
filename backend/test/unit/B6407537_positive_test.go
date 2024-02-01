package entity

import (
	"testing"

	"github.com/THEmello01/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"
	 ."github.com/onsi/gomega"
)

func testSuccess(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`success`, func(t *testing.T) {
		employee := entity.Employee{
			Name:       "Nattachai",
			URL:        "github.com/THEmello01/sut-final-lab/entity",
			EmployeeID: "EM0123456789",
		}

		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}
