// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"context"
	"fmt"
)

type EventService interface {
	Publish(ctx context.Context, topic string, delay int, event *NitricEvent) error
	ListTopics(ctx context.Context) ([]string, error)
}

type UnimplementedeventsPlugin struct {
	EventService
}

func (*UnimplementedeventsPlugin) Publish(ctx context.Context, topic string, delay int, event *NitricEvent) error {
	return fmt.Errorf("UNIMPLEMENTED")
}

func (*UnimplementedeventsPlugin) ListTopics(ctx context.Context) ([]string, error) {
	return nil, fmt.Errorf("UNIMPLEMENTED")
}
