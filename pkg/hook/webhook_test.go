/*
Copyright 2019 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hook

import (
	"context"
	"testing"
	"time"

	"k8s.io/git-sync/pkg/logging"
)

func TestWebhookDo(t *testing.T) {
	t.Run("test invalid urls are handled", func(t *testing.T) {
		wh := NewWebhook(
			":http://localhost:601426/hooks/webhook",
			"POST",
			200,
			time.Second,
			logging.New("", ""),
		)
		err := wh.Do(context.Background(), "hash")
		if err == nil {
			t.Fatalf("expected error for invalid url but got none")
		}
	})
}
