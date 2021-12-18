/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dtmcli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yedf/dtm/dtmcli/dtmimp"
)

func TestTypes(t *testing.T) {
	err := dtmimp.CatchP(func() {
		MustGenGid("http://localhost:36789/api/no")
	})
	assert.Error(t, err)
	assert.Error(t, err)
	_, err = BarrierFromQuery(url.Values{})
	assert.Error(t, err)

}

func TestXaSqlTimeout(t *testing.T) {
	old := GetXaSqlTimeoutMs()
	SetXaSqlTimeoutMs(old)
}
