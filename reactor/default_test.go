package reactor_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.zenithar.org/pkg/reactor"
)

func TestDefaultReactor_Do(t *testing.T) {

	testCases := []struct {
		name    string
		prepare func(reactor.Reactor)
		request interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "nil request",
			request: nil,
			wantErr: true,
		},
		{
			name:    "no handler registered",
			request: &struct{}{},
			wantErr: true,
		},
		{
			name:    "registered message",
			request: &struct{}{},
			prepare: func(r reactor.Reactor) {
				r.RegisterHandler(&struct{}{}, reactor.HandlerFunc(func(_ context.Context, req interface{}) (interface{}, error) {
					return req, nil
				}))
			},
			wantErr: false,
			want:    &struct{}{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Default instances
			ctx := context.Background()

			// Reactor
			underTest := reactor.New(tt.name)

			// Register handlers if needed
			if tt.prepare != nil {
				tt.prepare(underTest)
			}

			// Call operation
			got, err := underTest.Do(ctx, tt.request)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error mst be raised")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("error must not be raised, got %v", err)
			}
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("got %v, wanted %v", got, tt.want)
			}
		})
	}
}

func TestDefaultReactor_Send(t *testing.T) {

	testCases := []struct {
		name    string
		prepare func(reactor.Reactor)
		request interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "nil request",
			request: nil,
			wantErr: true,
		},
		{
			name:    "no handler registered",
			request: &struct{}{},
			wantErr: true,
		},
		{
			name:    "registered message",
			request: &struct{}{},
			prepare: func(r reactor.Reactor) {
				r.RegisterHandler(&struct{}{}, reactor.HandlerFunc(func(_ context.Context, req interface{}) (interface{}, error) {
					return req, nil
				}))
			},
			wantErr: false,
			want:    &struct{}{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Default instances
			ctx := context.Background()

			// Reactor
			underTest := reactor.New(tt.name)

			// Register handlers if needed
			if tt.prepare != nil {
				tt.prepare(underTest)
			}

			// Call operation
			err := underTest.Send(ctx, tt.request, func(_ context.Context, got interface{}, err error) {
				if !cmp.Equal(got, tt.want) {
					t.Fatalf("got %v, wanted %v", got, tt.want)
				}
				return
			})
			if tt.wantErr && err == nil {
				t.Fatalf("expected error mst be raised")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("error must not be raised, got %v", err)
			}

		})
	}
}
