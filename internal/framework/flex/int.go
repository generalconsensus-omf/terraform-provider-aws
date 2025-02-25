// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Int64FromFramework converts a Framework Int64 value to an int64 pointer.
// A null Int64 is converted to a nil int64 pointer.
func Int64FromFramework(ctx context.Context, v basetypes.Int64Valuable) *int64 {
	var output *int64

	must(Expand(ctx, v, &output))

	return output
}

func Int64ValueFromFramework(ctx context.Context, v basetypes.Int64Valuable) int64 {
	var output int64

	must(Expand(ctx, v, &output))

	return output
}

func Int64FromFrameworkLegacy(_ context.Context, v types.Int64) *int64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}

	i := v.ValueInt64()
	if i == 0 {
		return nil
	}

	return aws.Int64(i)
}

// Int64ToFramework converts an int64 pointer to a Framework Int64 value.
// A nil int64 pointer is converted to a null Int64.
func Int64ToFramework(ctx context.Context, v *int64) types.Int64 {
	var output types.Int64

	must(Flatten(ctx, v, &output))

	return output
}

// Int64ToFrameworkLegacy converts an int64 pointer to a Framework Int64 value.
// A nil int64 pointer is converted to a zero Int64.
func Int64ToFrameworkLegacy(_ context.Context, v *int64) types.Int64 {
	return types.Int64Value(aws.ToInt64(v))
}

func Int32ToFrameworkInt64(ctx context.Context, v *int32) types.Int64 {
	var output types.Int64

	must(Flatten(ctx, v, &output))

	return output
}

func Int32ValueToFrameworkInt64(ctx context.Context, v int32) types.Int64 {
	var output types.Int64

	must(Flatten(ctx, v, &output))

	return output
}

// Int32ToFrameworkInt64Legacy converts an int32 pointer to a Framework Int64 value.
// A nil int32 pointer is converted to a zero Int64.
func Int32ToFrameworkInt64Legacy(_ context.Context, v *int32) types.Int64 {
	return types.Int64Value(int64(aws.ToInt32(v)))
}

// Int32FromFrameworkInt64 coverts a Framework Int64 value to an int32 pointer.
// A null Int64 is converted to a nil int32 pointer.
func Int32FromFrameworkInt64(ctx context.Context, v basetypes.Int64Valuable) *int32 {
	var output *int32

	must(Expand(ctx, v, &output))

	return output
}

// Int32FromFramework coverts a Framework Int32 value to an int32 pointer.
// A null Int32 is converted to a nil int32 pointer.
func Int32FromFramework(ctx context.Context, v basetypes.Int32Valuable) *int32 {
	var output *int32

	must(Expand(ctx, v, &output))

	return output
}

func Int32FromFrameworkLegacy(_ context.Context, v types.Int32) *int32 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}

	i := v.ValueInt32()
	if i == 0 {
		return nil
	}

	return aws.Int32(i)
}

// Int32ValueFromFrameworkInt64 coverts a Framework Int64 value to an int32 value.
// A null Int64 is converted to a nil int32 pointer.
func Int32ValueFromFrameworkInt64(ctx context.Context, v basetypes.Int64Valuable) int32 {
	var output int32

	must(Expand(ctx, v, &output))

	return output
}

func ZeroInt32AsNull(v types.Int32) types.Int32 {
	if v.IsNull() || v.IsUnknown() {
		return v
	}

	if v.ValueInt32() == 0 {
		return types.Int32Null()
	}

	return v
}
