package manager

import (
	"errors"
	"fmt"
	"task-tracker/model"
	"testing"
)

type mockTaskService struct {
	AddCalled                bool
	UpdateCalled             bool
	DeleteCalled             bool
	ChangeStatusCalled       bool
	ListAllCalled            bool
	ListAllWithFiltersCalled bool
}

func TestManageAdd(t *testing.T) {
	t.Run("added task successfully", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"add", "test task"}
		err := manageAdd(args, mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.AddCalled {
			t.Error("expected Add to be called")
		}
	})

	t.Run("description not provided", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"add"}
		got := manageAdd(args, mock).Error()
		want := errors.New("description not provided").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.AddCalled {
			t.Error("expected Add to NOT be called")
		}
	})
}

func TestManageUpdate(t *testing.T) {
	t.Run("updated task successfully", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"update", "1", "test task"}
		err := manageUpdate(args, mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.UpdateCalled {
			t.Error("expected Update to be called")
		}
	})

	t.Run("id -96 does not exist", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"update", "-96", "test task"}

		id := args[1]
		got := manageUpdate(args, mock).Error()
		want := errors.New("does not exist task with ID:" + fmt.Sprint(id)).Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if !mock.UpdateCalled {
			t.Error("expected Update to be called")
		}
	})
	t.Run("id is not numeric", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"update", "blabla", "test task"}
		got := manageUpdate(args, mock).Error()
		want := errors.New("the id must to be numeric").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.UpdateCalled {
			t.Error("expected Update NOT to be called")
		}
	})

	t.Run("updated without description", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"update", "2"}
		got := manageUpdate(args, mock).Error()
		want := errors.New("update operation must to have {update {id} 'new description'}").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.UpdateCalled {
			t.Error("expected Update NOT to be called")
		}
	})

}

func TestManageDelete(t *testing.T) {
	t.Run("deleted task successfully", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"delete", "1"}
		err := manageDelete(args, mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.DeleteCalled {
			t.Error("expected Delete to be called")
		}
	})

	t.Run("delete without id", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"delete"}
		got := manageDelete(args, mock).Error()
		want := errors.New("delete operation must to have {delete {id}}").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.DeleteCalled {
			t.Error("expected Delete NOT to be called")
		}
	})

	t.Run("id is not numeric", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"delete", "abc"}
		got := manageDelete(args, mock).Error()
		want := errors.New("the id must to be numeric").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.UpdateCalled {
			t.Error("expected Delete NOT to be called")
		}
	})

	t.Run("id does not exist", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"delete", "-96"}
		got := manageDelete(args, mock).Error()
		want := errors.New("does not exist task with ID:" + fmt.Sprint(args[1])).Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.UpdateCalled {
			t.Error("expected Delete NOT to be called")
		}
	})
}

func TestManageChangeStatus(t *testing.T) {
	t.Run("change status successfully", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"mark-done", "1"}
		err := manageChangeStatus(args, "done", mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.ChangeStatusCalled {
			t.Error("expected ChangeStatus to be called")
		}
	})

	t.Run("id not provided", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"mark-in-progress"}
		got := manageChangeStatus(args, "done", mock).Error()
		want := errors.New("id not provided").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.ChangeStatusCalled {
			t.Error("expected ChangeStatus NOT to be called")
		}
	})

	t.Run("id is not numeric", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"mark-in-progress", "abc"}
		got := manageChangeStatus(args, "done", mock).Error()
		want := errors.New("the id must to be numeric").Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if mock.ChangeStatusCalled {
			t.Error("expected ChangeStatus NOT to be called")
		}
	})

	t.Run("id does not exist", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"mark-done", "-96"}
		got := manageChangeStatus(args, "done", mock).Error()
		want := errors.New("does not exist task with ID:" + fmt.Sprint(args[1])).Error()
		if got != want {
			t.Errorf("we got %s but we want %s", got, want)
		}
		if !mock.ChangeStatusCalled {
			t.Error("expected ChangeStatus to be called")
		}
	})
}

func TestManageList(t *testing.T) {
	t.Run("too many arguments", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"list", "done", "extra"}
		got := manageList(args, mock).Error()
		want := errors.New("we only support 'list', 'list done', 'list todo', 'list in-progress'").Error()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("list all tasks", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"list"}
		err := manageList(args, mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.ListAllCalled {
			t.Error("expected ListAll to be called")
		}
	})

	t.Run("list with valid filter", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"list", "done"}
		err := manageList(args, mock)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if !mock.ListAllWithFiltersCalled {
			t.Error("expected ListAllWithFilters to be called")
		}
	})

	t.Run("list with invalid filter", func(t *testing.T) {
		mock := &mockTaskService{}
		args := []string{"list", "invalid"}
		got := manageList(args, mock).Error()
		want := errors.New("we only support 'done', 'todo', 'in-progress'").Error()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		if mock.ListAllWithFiltersCalled {
			t.Error("expected ListAllWithFilters NOT to be called")
		}
	})
}

func (m *mockTaskService) Add(description string) int {
	m.AddCalled = true
	return 1
}

func (m *mockTaskService) Update(newDescription string, id int) error {
	m.UpdateCalled = true
	if id == -96 {
		return errors.New("")
	}
	return nil
}

func (m *mockTaskService) Delete(id int) error {
	m.DeleteCalled = true
	if id == -96 {
		return errors.New("")
	}
	return nil
}

func (m *mockTaskService) ListAll() []model.Task {
	m.ListAllCalled = true
	return nil
}

func (m *mockTaskService) ListAllWithFilters(filter string) []model.Task {
	m.ListAllWithFiltersCalled = true
	return nil
}

func (m *mockTaskService) ChangeStatus(status model.TaskStatus, id int) error {
	m.ChangeStatusCalled = true
	if id == -96 {
		return errors.New("")
	}
	return nil
}
