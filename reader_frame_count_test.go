// Copyright 2018 The Cacophony Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cptv

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReaderFrameCount(t *testing.T) {
	frame := makeTestFrame()
	cptvBytes := new(bytes.Buffer)

	w := NewWriter(cptvBytes)
	require.NoError(t, w.WriteHeader("nz43"))
	require.NoError(t, w.WriteFrame(frame))
	require.NoError(t, w.WriteFrame(frame))
	require.NoError(t, w.WriteFrame(frame))
	require.NoError(t, w.Close())

	r, err := NewReader(cptvBytes)
	require.NoError(t, err)
	c, err := r.FrameCount()
	require.NoError(t, err)
	assert.Equal(t, 3, c)
}
