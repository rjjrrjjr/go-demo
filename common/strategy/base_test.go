package strategy

import (
	"context"
	"fmt"
	"testing"
)

// 策略模式

var IdpParseStrategyMap = make(map[uint32]interface{}, 0)

func doRealParse(ctx context.Context, parseKey string, parseType uint32) (string, error) {
	strategy, exist := IdpParseStrategyMap[parseType]
	if !exist {
		return "", nil
	}
	return strategy.(IdpParseStrategy).doParse(ctx, parseKey)
}

type IdpParseStrategy interface {
	register(ctx context.Context, parseType uint32) error

	doParse(ctx context.Context, parseKey string) (string, error)
}

type WellKnownParseStrategy struct {
}

func (w *WellKnownParseStrategy) register(ctx context.Context, parseType uint32) error {
	IdpParseStrategyMap[parseType] = w
	return nil
}

func (w *WellKnownParseStrategy) doParse(ctx context.Context, parseKey string) (string, error) {
	return "WellKnownParseStrategy" + parseKey, nil
}

type IssuersParseStrategy struct {
}

func (i *IssuersParseStrategy) register(ctx context.Context, parseType uint32) error {
	IdpParseStrategyMap[parseType] = i
	return nil
}

func (i *IssuersParseStrategy) doParse(ctx context.Context, parseKey string) (string, error) {
	return "IssuersParseStrategy" + parseKey, nil
}

func TestStrategy(t *testing.T) {

	ctx := context.Background()

	var wellKnownParseStrategy WellKnownParseStrategy
	_ = wellKnownParseStrategy.register(ctx, 1)

	var issuersParseStrategy IssuersParseStrategy
	_ = issuersParseStrategy.register(ctx, 2)

	fmt.Println(doRealParse(ctx, "jinrruan", 1))
	fmt.Println(doRealParse(ctx, "jinrruan", 2))
	fmt.Println(doRealParse(ctx, "jinrruan", 3))
}

// .varcheckerr
// bool.if/ bool.else
