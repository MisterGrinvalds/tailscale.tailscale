// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package tailscale

import (
	"context"

	"tailscale.com/feature"
)

// HookResolveValueFromParameterStore resolves to [awsparamstore.ResolveValue] when
// the corresponding feature tag is enabled in the build process.
//
// It fetches a value from AWS Parameter Store given an ARN. If the provided
// value is not an Parameter Store ARN, it returns the value unchanged.
var HookResolveValueFromParameterStore feature.Hook[func(ctx context.Context, valueOrARN string) (string, error)]
