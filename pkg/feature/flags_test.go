// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package feature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
At runtime, there is a singleton instance of feature flags.
This test suite verifies the behavior of of such an instance.
To avoid test pollution, a fresh instance is created in each test.
*/

func TestMustBeFrozenBeforeFirstUse(t *testing.T) {
	assert.Panics(t, func() { newFlagSet().IsEnabled(Noop) })
}

func TestFeaturesAreDisabledByDefault(t *testing.T) {
	assert.False(t, newFlagSet().Freeze().IsEnabled(Noop))
}

func TestFeaturesCanBeEnabled(t *testing.T) {
	flagset := newFlagSet().Enable(Noop).Freeze()
	assert.True(t, flagset.IsEnabled(Noop))
}

func TestCannotBeModifiedAfterFreeze(t *testing.T) {
	flagset := newFlagSet().Freeze()
	assert.Panics(t, func() { flagset.Enable(Noop) })
}

func TestPanicsWhenFeatureIsUnknown(t *testing.T) {
	assert.Panics(t, func() { newFlagSet().Enable("not-a-feature") })
}
