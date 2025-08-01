package commands

import (
	"testing"
)

type mockService struct {
	GetGithubActivityCalled bool
}

func (s *mockService) GetGithubActivity(username string) error {
	s.GetGithubActivityCalled = true
	return nil
}

func TestManageCommands(t *testing.T) {

	t.Run("len args invalid", func(t *testing.T) {
		mockService := &mockService{}
		args := []string{"one", "two"}
		want := "command unknown. please see 'github-activity help' for more information"
		got := ManageCommands(args, mockService).Error()

		if got != want {
			t.Errorf("we got %q but we want %q", got, want)
		}
	})

	t.Run("command provided correctly", func(t *testing.T) {
		mockService := &mockService{}
		args := []string{"cmd", "github-activity", "username"}
		ManageCommands(args, mockService)
		if !mockService.GetGithubActivityCalled {
			t.Error("expected GetGithubActivity to be called")
		}
	})

	t.Run("command unknown", func(t *testing.T) {
		mockService := &mockService{}
		args := []string{"cmd", "github-activityas", "username"}
		want := "command unknown. please see 'github-activity help' for more information"
		got := ManageCommands(args, mockService).Error()

		if got != want {
			t.Errorf("we got %q but we want %q", got, want)
		}
	})

}
