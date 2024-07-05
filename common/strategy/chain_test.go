package strategy

import (
	"fmt"
	"testing"
)

func TestI(t *testing.T) {
	a := new(CheckTokenRule)
	fmt.Println(a.Apply("haha"))

	var b RuleChain = new(CheckTokenRule)
	fmt.Println(b.Apply("haha"))

}

// 责任链模式

type RuleChain interface {
	Apply(req interface{}) error
	ApplyNext(req interface{}) error
}

type BaseRuleChain struct {
	next RuleChain
}

func (b *BaseRuleChain) Apply(req interface{}) error {
	fmt.Println("do Apply", req)
	return nil
}

func (b *BaseRuleChain) ApplyNext(req interface{}) error {
	if b.next != nil {
		return b.next.Apply(req)
	}
	fmt.Println("do ApplyNext end", req)
	return nil
}

func (b *BaseRuleChain) RegisterNext(r RuleChain) {
	b.next = r
}

type CheckTokenRule struct {
	BaseRuleChain
}

func (c *CheckTokenRule) Apply(req interface{}) error {
	fmt.Println("check token", req)
	return c.ApplyNext(req)
}

type CheckAgeRule struct {
	BaseRuleChain
}

func (c *CheckAgeRule) Apply(req interface{}) error {
	fmt.Println("check age", req)
	return c.ApplyNext(req)
}

type CheckAuthor struct {
	BaseRuleChain
}

func (c *CheckAuthor) Apply(req interface{}) error {
	fmt.Println("check author", req)
	return c.ApplyNext(req)
}

func TestChain(t *testing.T) {
	ruleChain := new(CheckTokenRule)
	ageRule := new(CheckAgeRule)
	ruleChain.RegisterNext(ageRule)
	ageRule.RegisterNext(new(CheckAuthor))
	err := ruleChain.Apply("jinrruan")
	fmt.Println(err)
}
