/*
 * Copyright The Dragonfly Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"context"

	"github.com/dragonflyoss/Dragonfly/apis/types"
)

// PreheatInfo gets detailed information of a preheat task.
func (client *APIClient) PreheatInfo(ctx context.Context, id string) (*types.PreheatInfo, error) {
	resp, err := client.get(ctx, "/preheats/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	preheat := &types.PreheatInfo{}

	err = decodeBody(preheat, resp.Body)
	ensureCloseReader(resp)

	return preheat, err
}
