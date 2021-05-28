package client_test

import (
	"errors"
	"net/url"
	"os"
	"testing"

	// Modules
	deadline "github.com/thevfxcoop/go-deadline-api"
	deadlineclient "github.com/thevfxcoop/go-deadline-api/pkg/client"
)

func Test_Client_000(t *testing.T) {
	t.Log(t.Name())
}
func Test_Client_001(t *testing.T) {
	client := GetClient(t, false)
	t.Log(client)
}

func Test_Client_002(t *testing.T) {
	client := GetClient(t, false)
	if version, err := client.Ping(); err != nil {
		t.Error(err)
	} else {
		t.Logf("Version=%q", version)
	}
}

func Test_Client_003(t *testing.T) {
	client := GetClient(t, true)
	if jobs, err := client.GetJobs(); err != nil {
		t.Error(err)
	} else {
		for _, job := range jobs {
			t.Logf("%+v", job)
		}
	}
}
func Test_Client_004(t *testing.T) {
	client := GetClient(t, true)
	if jobs, err := client.GetJobs(deadlineclient.OptJobState(deadlineclient.JobStateCompleted)); err != nil {
		t.Error(err)
	} else {
		for _, job := range jobs {
			t.Logf("%+v", job)
		}
	}
}

func Test_Client_005(t *testing.T) {
	client := GetClient(t, true)
	if jobs, err := client.GetJobIds(); err != nil {
		t.Error(err)
	} else {
		for _, job := range jobs {
			t.Logf("%+v", job)
		}
	}
}

func Test_Client_006(t *testing.T) {
	client := GetClient(t, true)
	if job, err := client.GetJobWithId("xxx"); errors.Is(err, deadline.ErrNotFound) {
		// OK
	} else if err != nil {
		t.Error(err)
	} else if job == nil {
		t.Error("Unexpected response")
	} else {
		t.Error("Unexpected response")
	}
}
func Test_Client_007(t *testing.T) {
	client := GetClient(t, true)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	for _, id := range ids {
		if job, err := client.GetJobWithId(id); err != nil {
			t.Error(err)
		} else if job == nil {
			t.Error("Unexpected response")
		} else {
			t.Logf("%+v", job)
		}
	}
}

func Test_Client_008(t *testing.T) {
	client := GetClient(t, true)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	if len(ids) == 0 {
		t.Skip("No jobs to suspend")
	} else {
		t.Log("Suspending job", ids[0])
	}
	if err := client.SuspendJobWithId(ids[0]); err != nil {
		t.Error(err)
	} else if err := client.ResumeJobWithId(ids[0]); err != nil {
		t.Error(err)
	}
}

func Test_Client_009(t *testing.T) {
	client := GetClient(t, true)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	if len(ids) == 0 {
		t.Skip("No jobs")
	}
	if tasks, err := client.GetTaskIdsForJobId(ids[0]); err != nil {
		t.Error(err)
	} else {
		t.Log(tasks)
	}
}

func Test_Client_010(t *testing.T) {
	client := GetClient(t, true)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	if len(ids) == 0 {
		t.Skip("No jobs")
	}
	if tasks, err := client.GetTaskIdsForJobId(ids[0]); err != nil {
		t.Error(err)
	} else {
		for _, task := range tasks {
			if task, err := client.GetTaskWithId(ids[0], task); err != nil {
				t.Error(err)
			} else {
				t.Logf("%+v", task)
			}
		}
	}
}

func Test_Client_011(t *testing.T) {
	client := GetClient(t, false)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	if len(ids) == 0 {
		t.Skip("No jobs")
	}
	tasks, err := client.GetTasksForJobId(ids[0])
	if err != nil {
		t.Error(err)
	}
	for _, task := range tasks {
		t.Logf("%+v", task)
	}
}

func Test_Client_012(t *testing.T) {
	client := GetClient(t, true)
	ids, err := client.GetJobIds()
	if err != nil {
		t.Error(err)
	}
	if len(ids) == 0 {
		t.Skip("No jobs")
	}
	// This command can take a while - so it may timeout
	if err := client.RequeueTasksWithId(ids[0]); err != nil {
		t.Error(err)
	}
}

///////////////////////////////////////////////////////////////////////////////

// To run some tests, use environment variable DEADLINE_ENDPOINT

func GetClient(t *testing.T, verbose bool) *deadlineclient.Client {
	if endpoint := os.Getenv("DEADLINE_ENDPOINT"); endpoint == "" {
		t.Skip("Skipping test, set DEADLINE_ENDPOINT environment variable otherwise")
	} else if endpoint, err := url.Parse(endpoint); err != nil {
		t.Fatal(err)
	} else if client, err := deadlineclient.NewClient(endpoint, deadlineclient.OptTrace(os.Stderr, verbose)); err != nil {
		t.Fatal(err)
	} else {
		return client
	}

	// Skip tests
	return nil
}
