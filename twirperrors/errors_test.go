package twirperrors

import (
	"reflect"
	"testing"

	errors "github.com/segmentio/errors-go"
	"github.com/twitchtv/twirp"
)

func TestNew(t *testing.T) {
	tests := []struct {
		types []string
		code  twirp.ErrorCode
	}{
		{
			types: []string{"Canceled"},
			code:  twirp.Canceled,
		},

		{
			types: []string{"Unknown"},
			code:  twirp.Unknown,
		},

		{
			types: []string{"InvalidArgument"},
			code:  twirp.InvalidArgument,
		},

		{
			types: []string{"DeadlineExceeded"},
			code:  twirp.DeadlineExceeded,
		},

		{
			types: []string{"NotFound"},
			code:  twirp.NotFound,
		},

		{
			types: []string{"BadRoute"},
			code:  twirp.BadRoute,
		},

		{
			types: []string{"AlreadyExists"},
			code:  twirp.AlreadyExists,
		},

		{
			types: []string{"PermissionDenied"},
			code:  twirp.PermissionDenied,
		},

		{
			types: []string{"Unauthenticated"},
			code:  twirp.Unauthenticated,
		},

		{
			types: []string{"ResourceExhausted"},
			code:  twirp.ResourceExhausted,
		},

		{
			types: []string{"FailedPrecondition"},
			code:  twirp.FailedPrecondition,
		},

		{
			types: []string{"Aborted"},
			code:  twirp.Aborted,
		},

		{
			types: []string{"OutOfRange"},
			code:  twirp.OutOfRange,
		},

		{
			types: []string{"Unimplemented"},
			code:  twirp.Unimplemented,
		},

		{
			types: []string{"Internal"},
			code:  twirp.Internal,
		},

		{
			types: []string{"Unavailable"},
			code:  twirp.Unavailable,
		},

		{
			types: []string{"DataLoss"},
			code:  twirp.DataLoss,
		},

		{
			types: []string{"Whatever"},
			code:  twirp.Unknown,
		},

		{
			types: []string{},
			code:  twirp.Unknown,
		},
	}

	t.Run("<nil>", func(t *testing.T) {
		if twerr := New(nil); twerr != nil {
			t.Error("calling New on a nil error did not return a nil error")
		}
	})

	for _, test := range tests {
		t.Run(string(test.code), func(t *testing.T) {
			twerr := New(
				errors.WithTags(
					errors.WithTypes(errors.New("oops"), test.types...),
					errors.T("hello", "world"),
					errors.T("twitch", "tv"),
				),
			)

			if msg := twerr.Msg(); msg != "oops" {
				t.Error("wrong error message:", msg)
			}

			if code := twerr.Code(); code != test.code {
				t.Error("wrong error code:", code)
			}

			if meta := twerr.MetaMap(); !reflect.DeepEqual(meta, map[string]string{
				"hello":  "world",
				"twitch": "tv",
			}) {
				t.Error("wrong meta map:", meta)
			}
		})
	}
}
