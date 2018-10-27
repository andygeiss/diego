package inference_test

import (
	"github.com/andygeiss/diego/internal/inference"
	"github.com/andygeiss/diego/pkg/assert"
	"github.com/andygeiss/diego/pkg/assert/matchers"
	"testing"
)

func TestDefaultEngine_Run_Should_Fire_One_Action(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_1.json")
	e := inference.NewDefaultEngine("one", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(2))
	assert.That(t, facts[1], matchers.IsEqual("X"))
}

func TestDefaultEngine_Run_Should_Fire_Two_Actions(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_2.json")
	e := inference.NewDefaultEngine("two", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(4))
	assert.That(t, facts[3], matchers.IsEqual("Y"))
}

func TestDefaultEngine_Run_Should_Handle_One_Cycle(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_cycle_1.json")
	e := inference.NewDefaultEngine("cycle 1", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(7))
	assert.That(t, facts[5], matchers.IsEqual("X"))
}

func TestDefaultEngine_Run_Should_Handle_Two_Cycles(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_cycle_2.json")
	e := inference.NewDefaultEngine("cycle 2", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(8))
	assert.That(t, facts[7], matchers.IsEqual("Y"))
}

func TestDefaultEngine_Run_Should_Handle_Three_Cycles(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_cycle_3.json")
	e := inference.NewDefaultEngine("cycle 3", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(9))
	assert.That(t, facts[8], matchers.IsEqual("Z"))
}

func TestDefaultEngine_Run_Should_Handle_Conflict_Resolution(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_conflict.json")
	e := inference.NewDefaultEngine("conflict", r)
	facts, err := e.Run(nil)
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(2))
	assert.That(t, facts[1], matchers.IsEqual("Y"))
}

func TestDefaultEngine_Run_Should_Handle_Conditions(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_conditions.json")
	e := inference.NewDefaultEngine("conditions", r)
	facts, err := e.Run([]string{"A", "B", "C", "D", "E"})
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(9))
	assert.That(t, facts[8], matchers.IsEqual("Z"))
}

func TestDefaultEngine_Run_Should_Handle_Multiple_Conclusions(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_multi_conclusions.json")
	e := inference.NewDefaultEngine("multi conclusions", r)
	facts, err := e.Run([]string{"A"})
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(3))
	assert.That(t, facts[2], matchers.IsEqual("C"))
}