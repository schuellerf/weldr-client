// Copyright 2021 by Red Hat, Inc. All rights reserved.
// Use of this source is goverend by the Apache License
// that can be found in the LICENSE file.

// +build integration

package weldr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListProjects(t *testing.T) {
	projects, r, err := testState.client.ListProjects()
	require.Nil(t, err)
	require.Nil(t, r)
	require.NotNil(t, projects)
	assert.GreaterOrEqual(t, len(projects), 2)
}
