package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type QuickSortSuite struct {
	suite.Suite
	name               string
	arr                []int
	want               []int
	asyncLimit         string
	asyncLimitStartVal string
	asyncExpected      bool
	asyncExecuted      bool
}

func (s *QuickSortSuite) SetupTest() {
	s.asyncLimitStartVal = os.Getenv(AsyncLimitEnvName)
	_ = os.Setenv(AsyncLimitEnvName, s.asyncLimit)
	v = func() {
		s.asyncExecuted = true
	}
}

func (s *QuickSortSuite) TearDownTest() {
	_ = os.Setenv(AsyncLimitEnvName, s.asyncLimitStartVal)
}

func (s *QuickSortSuite) TestSort() {
	s.Equal(s.want, Sort(s.arr), s.name)
	if s.asyncExpected != s.asyncExecuted {
		s.T().Fail()
	}
}

func TestSortBySuite(t *testing.T) {
	tests := []struct {
		name          string
		arr           []int
		want          []int
		asyncLimit    string
		asyncExpected bool
	}{
		{
			name: "Sorts good",
			arr:  []int{12, 34, 23, 45, 547, 286},
			want: []int{12, 23, 34, 45, 286, 547},
		},
		{
			name: "Not initialized slice",
		},
		{
			name: "Empty slice",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "One element slice",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name:          "Sorts good async",
			arr:           []int{12, 34, 23, 45, 547, 286, 1},
			want:          []int{1, 12, 23, 34, 45, 286, 547},
			asyncLimit:    "1",
			asyncExpected: true,
		},
		{
			name:          "Sorts good with wrong env variable",
			arr:           []int{12, 34, 23, 45, 547, 286, 1},
			want:          []int{1, 12, 23, 34, 45, 286, 547},
			asyncLimit:    "q",
			asyncExpected: false,
		},
	}
	for _, tt := range tests {
		suite.Run(t, &QuickSortSuite{
			name:          tt.name,
			arr:           tt.arr,
			want:          tt.want,
			asyncLimit:    tt.asyncLimit,
			asyncExpected: tt.asyncExpected,
		})
	}
}
