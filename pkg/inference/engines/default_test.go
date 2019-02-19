package engines_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/diego/pkg/inference/engines"
	"github.com/andygeiss/diego/pkg/inference/repositories"
	"testing"
)

func TestDefaultEngine_Run_Should_Fire_One_Action(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_1.json")
	e := engines.NewDefaultEngine("one", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(2))
	assert.That(t, facts[1], is.Equal("X"))
}

func TestDefaultEngine_Run_Should_Fire_Two_Actions(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_2.json")
	e := engines.NewDefaultEngine("two", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(4))
	assert.That(t, facts[3], is.Equal("Y"))
}

func TestDefaultEngine_Run_Should_Handle_One_Cycle(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_cycle_1.json")
	e := engines.NewDefaultEngine("cycle 1", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(7))
	assert.That(t, facts[5], is.Equal("X"))
}

func TestDefaultEngine_Run_Should_Handle_Two_Cycles(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_cycle_2.json")
	e := engines.NewDefaultEngine("cycle 2", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(8))
	assert.That(t, facts[7], is.Equal("Y"))
}

func TestDefaultEngine_Run_Should_Handle_Three_Cycles(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_cycle_3.json")
	e := engines.NewDefaultEngine("cycle 3", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(9))
	assert.That(t, facts[8], is.Equal("Z"))
}

func TestDefaultEngine_Run_Should_Handle_Conflict_Resolution(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_conflict.json")
	e := engines.NewDefaultEngine("conflict", r)
	facts, err := e.Run(nil)
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(2))
	assert.That(t, facts[1], is.Equal("Y"))
}

func TestDefaultEngine_Run_Should_Handle_Conditions(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_conditions.json")
	e := engines.NewDefaultEngine("conditions", r)
	facts, err := e.Run([]string{"A", "B", "C", "D", "E"})
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(9))
	assert.That(t, facts[8], is.Equal("Z"))
}

func TestDefaultEngine_Run_Should_Handle_Multiple_Conclusions(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_multi_conclusions.json")
	e := engines.NewDefaultEngine("multi conclusions", r)
	facts, err := e.Run([]string{"A"})
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(3))
	assert.That(t, facts[2], is.Equal("C"))
}
