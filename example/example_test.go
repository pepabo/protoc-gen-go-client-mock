package example

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pepabo/protoc-gen-go-client-mock/example/gen/go/myapp"
)

func TestClientUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := myapp.NewMockClient(ctrl)
	req := &myapp.CreateUserRequest{
		Name:  "alice",
		Email: "alice@example.com",
	}
	res := &myapp.UserResponse{
		Name:  "alice",
		Email: "alice@example.com",
	}
	client.EXPECT().UserService().CreateUser(ctx, req).Return(res, nil)
	got, err := client.UserService().CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	opts := []cmp.Option{
		cmp.AllowUnexported(myapp.UserResponse{}),
		cmpopts.IgnoreFields(myapp.UserResponse{}, "state", "sizeCache", "unknownFields"),
	}
	if diff := cmp.Diff(got, res, opts...); diff != "" {
		t.Errorf("%s", diff)
	}
}

func TestClientProject(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	client := myapp.NewMockClient(ctrl)
	t.Cleanup(func() {
		client.Close()
	})
	req := &myapp.ListProjectsRequest{
		Page: 1,
	}
	res := []*myapp.ProjectResponse{
		{Name: "staging"},
		{Name: "production"},
	}

	lc := myapp.NewMockProjectService_ListProjectsClient(ctrl)
	lc.EXPECT().Recv().Return(res[0], nil)
	lc.EXPECT().Recv().Return(res[1], nil)
	lc.EXPECT().Recv().Return(nil, io.EOF)
	client.EXPECT().ProjectService().ListProjects(ctx, req).Return(lc, nil)
	client.EXPECT().Close().Return(nil)

	rc, err := client.ProjectService().ListProjects(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for {
		got, err := rc.Recv()
		if err != nil {
			if errors.Is(io.EOF, err) {
				break
			}
			t.Fatal(err)
		}
		opts := []cmp.Option{
			cmp.AllowUnexported(myapp.ProjectResponse{}),
			cmpopts.IgnoreFields(myapp.ProjectResponse{}, "state", "sizeCache", "unknownFields"),
		}
		if diff := cmp.Diff(got, res[i], opts...); diff != "" {
			t.Errorf("%s", diff)
		}
		i++
	}
}
