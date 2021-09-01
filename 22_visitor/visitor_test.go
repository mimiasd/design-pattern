package visitor

import "testing"

func TestVisitor(t *testing.T) {
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Accept(&ServiceRequestVisitor{})

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("A company"))
	c2.Add(NewIndividualCustomer("bob"))
	c2.Add(NewEnterpriseCustomer("B company"))
	c2.Accept(&AnalysisVisitor{})
}
