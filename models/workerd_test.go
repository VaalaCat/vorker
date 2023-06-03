package models

import (
	"reflect"
	"testing"
	"voker/defs"
	"voker/entities"
)

func GenTestRecord() {
	testWorkers := []*Worker{
		{
			Worker: &entities.Worker{
				UID:          "a5eb7b7d602449f699cb7bda44fde846", //strings.Replace(uuid.New().String(), "-", "", -1),
				ExternalPath: "test",
				HostName:     "test",
				NodeName:     "test",
				Port:         8080,
				Entry:        "entry.js",
				Code:         []byte(defs.DefaultCode),
				Name:         "test",
			},
		},
		{
			Worker: &entities.Worker{
				UID:          "426fb59e086501262f806e89a9d20081", //strings.Replace(uuid.New().String(), "-", "", -1),
				ExternalPath: "test",
				HostName:     "test",
				NodeName:     "test",
				Port:         8081,
				Entry:        "entry.js",
				Code:         []byte(defs.DefaultCode),
				Name:         "test1",
			},
		},
	}

	for _, worker := range testWorkers {
		_, err := GetWorkerByUID(1, worker.UID)
		if err != nil {
			worker.Create()
		}
	}
}

func TestGetWorkerByUID(t *testing.T) {
	GenTestRecord()
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Worker
		wantErr bool
	}{
		{
			name: "test common",
			args: args{uid: "a5eb7b7d602449f699cb7bda44fde846"},
			want: &entities.Worker{
				UID:          "a5eb7b7d602449f699cb7bda44fde846", //strings.Replace(uuid.New().String(), "-", "", -1),
				ExternalPath: "test",
				HostName:     "test",
				NodeName:     "test",
				Port:         8080,
				Entry:        "entry.js",
				Code:         []byte(defs.DefaultCode),
				Name:         "test",
			},
			wantErr: false,
		},
		{
			name:    "test not found",
			args:    args{uid: "114514"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWorkerByUID(1, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Worker, tt.want) {
				t.Errorf("GetByUID() = %v, want %v", &got.Worker, tt.want)
			}
		})
	}
}

func TestGetWorkersByNames(t *testing.T) {
	GenTestRecord()
	type args struct {
		names []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Worker
		wantErr bool
	}{
		{
			name:    "test common",
			args:    args{names: []string{"test", "test1"}},
			want:    []*Worker{{Worker: &entities.Worker{}}, {Worker: &entities.Worker{}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWorkersByNames(1, tt.args.names)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetByNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
