/*
 * Copyright 2022 Aspect Build Systems, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package shutdown

import (
	"context"

	"github.com/spf13/cobra"

	"aspect.build/cli/pkg/aspecterrors"
	"aspect.build/cli/pkg/bazel"
	"aspect.build/cli/pkg/ioutils"
)

type Shutdown struct {
	ioutils.Streams
}

func New(streams ioutils.Streams) *Shutdown {
	return &Shutdown{
		Streams: streams,
	}
}

func (v *Shutdown) Run(ctx context.Context, _ *cobra.Command, args []string) error {
	bazelCmd := []string{"shutdown"}
	bazelCmd = append(bazelCmd, args...)
	bzl, err := bazel.FindFromWd()
	if err != nil {
		return err
	}

	if exitCode, err := bzl.RunCommand(v.Streams, bazelCmd...); exitCode != 0 {
		err = &aspecterrors.ExitError{
			Err:      err,
			ExitCode: exitCode,
		}
		return err
	}

	return nil
}