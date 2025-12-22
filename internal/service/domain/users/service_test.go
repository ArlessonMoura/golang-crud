package users

import (
	"context"
	"reflect"
	"testing"

	"meu-treino-golang/users-crud/internal/service"
)

type mockRepo struct {
    createID  uint
    createErr error
    listResp  []service.UserDTO
    listErr   error

    lastName  string
    lastEmail string
}

func (m *mockRepo) Create(ctx context.Context, name, email string) (uint, error) {
    m.lastName = name
    m.lastEmail = email
    return m.createID, m.createErr
}

func (m *mockRepo) List(ctx context.Context) ([]service.UserDTO, error) {
    return m.listResp, m.listErr
}

func (m *mockRepo) GetByID(ctx context.Context, id uint) (*service.UserDTO, error) {
    // simple stub: search in listResp
    for _, u := range m.listResp {
        if u.ID == id {
            copy := u
            return &copy, nil
        }
    }
    return nil, nil
}

func TestCreateUser_EmptyName(t *testing.T) {
    svc := NewService(&mockRepo{})
    if _, err := svc.CreateUser(context.Background(), "", "a@b.com"); err == nil {
        t.Fatalf("expected error for empty name")
    }
}

func TestCreateUser_Success(t *testing.T) {
    mr := &mockRepo{createID: 123}
    svc := NewService(mr)
    id, err := svc.CreateUser(context.Background(), "Alice", "alice@example.com")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if id != 123 {
        t.Fatalf("expected id 123 got %d", id)
    }
    if mr.lastName != "Alice" || mr.lastEmail != "alice@example.com" {
        t.Fatalf("repo received wrong args: %v %v", mr.lastName, mr.lastEmail)
    }
}

func TestListUsers_DelegatesToRepo(t *testing.T) {
    resp := []service.UserDTO{{ID: 1, Name: "A", Email: "a@a.com"}}
    mr := &mockRepo{listResp: resp}
    svc := NewService(mr)
    list, err := svc.ListUsers(context.Background())
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if !reflect.DeepEqual(list, resp) {
        t.Fatalf("expected %+v got %+v", resp, list)
    }
}
