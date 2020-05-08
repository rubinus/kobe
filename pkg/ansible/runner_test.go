package ansible

import (
	"kobe/pkg/models"
	"testing"
	"time"
)

func TestPlaybookRunner_Run(t *testing.T) {
	result := models.Result{
		StartTime: time.Now(),
		EndTime:   time.Time{},
		Message:   "",
		Success:   false,
		Content:   nil,
	}
	runner := PlaybookRunner{Project: models.Project{
		Name:      "test",
		Playbooks: nil,
	}}
	runner.Run("fd37fdc0-eaed-4178-a8c2-b8087dbe01b6", "main.yml", &result)
	t.Log(result)
}
