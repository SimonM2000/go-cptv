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
	"io"
	"testing"
	"time"

	"github.com/TheCacophonyProject/lepton3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriterAndReader(t *testing.T) {
	frame0 := makeTestFrame()
	frame1 := makeOffsetFrame(frame0)
	frame2 := makeOffsetFrame(frame1)

	cptvBytes := new(bytes.Buffer)

	w := NewWriter(cptvBytes)
	require.NoError(t, w.WriteHeader("nz42"))
	require.NoError(t, w.WriteFrame(frame0))
	require.NoError(t, w.WriteFrame(frame1))
	require.NoError(t, w.WriteFrame(frame2))
	require.NoError(t, w.Close())

	r, err := NewReader(cptvBytes)
	require.NoError(t, err)
	assert.Equal(t, "nz42", r.DeviceName())
	assert.True(t, time.Since(r.Timestamp()) < time.Minute)

	frameD := new(lepton3.Frame)
	require.NoError(t, r.ReadFrame(frameD))
	assert.Equal(t, frame0, frameD)
	require.NoError(t, r.ReadFrame(frameD))
	assert.Equal(t, frame1, frameD)
	require.NoError(t, r.ReadFrame(frameD))
	assert.Equal(t, frame2, frameD)

	assert.Equal(t, io.EOF, r.ReadFrame(frameD))
}
