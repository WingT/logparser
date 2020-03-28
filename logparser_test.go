package logparser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTextFormattedLog(t *testing.T) {
	testCases := []struct {
		Log   string
		Entry Entry
		Desc  string
	}{
		{
			Log: `time="2020-03-28T07:33:05+00:00" level=info msg="executing git command" command="gitaly-upload-pack tcp://10.10.10.10:8080 {\"repository\":{\"storage_name\":\"default\",\"relative_path\":\"test_ns/test_name.git\",\"git_object_directory\":\"\",\"git_alternate_object_directories\":[],\"gl_repository\":\"project-123456\"},\"gl_repository\":\"project-123456\",\"gl_id\":\"key-123456\",\"gl_username\":\"michael\",\"git_config_options\":[],\"git_protocol\":null}" pid=123456 user="user with id key-123456"`,
			Entry: Entry{
				"time":    `2020-03-28T07:33:05+00:00`,
				"level":   `info`,
				"msg":     `executing git command`,
				"command": `gitaly-upload-pack tcp://10.10.10.10:8080 {"repository":{"storage_name":"default","relative_path":"test_ns/test_name.git","git_object_directory":"","git_alternate_object_directories":[],"gl_repository":"project-123456"},"gl_repository":"project-123456","gl_id":"key-123456","gl_username":"michael","git_config_options":[],"git_protocol":null}`,
				"pid":     `123456`,
				"user":    `user with id key-123456`,
			},
			Desc: "complicated case",
		},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.Entry, ParseTextFormattedLog(tc.Log), tc.Desc)
	}
}
