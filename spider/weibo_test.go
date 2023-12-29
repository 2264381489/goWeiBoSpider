package main

import (
	"context"
	"fmt"
	"goSpider/config"
	"goSpider/svc"
	"testing"
)

const (
	testRedisHost = "127.0.0.1:6379"
	testMongoHost = "localhost:27017/"
	testDB        = "testSpider"
)

func TestSpider_Run(t *testing.T) {

	serviceContext := svc.NewServiceContext(config.Config{
		Mongo: config.MongoConfig{Url: fmt.Sprintf("mongodb://%s/%s", testMongoHost, testDB), DB: testDB},
	})
	ctx := context.Background()

	spider := NewSpider(ctx, serviceContext)

	spider.Run(5680343342)
}

func TestExec(t *testing.T) {
	type args struct {
		scriptPath string
		args       []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create test",
			args: args{
				scriptPath: "./test.sh",
				args:       []string{"1"},
			},
		}, {
			name: "delete test",
			args: args{
				scriptPath: "./delete.sh",
				args:       []string{"1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(tt.args.scriptPath, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
