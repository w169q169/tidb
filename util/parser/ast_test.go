// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
	utilparser "github.com/pingcap/tidb/util/parser"
)

func TestSimpleCases(t *testing.T) {
	tests := []struct {
		sql string
		db  string
		ans string
	}{
		{
			sql: "insert into t values(1, 2)",
			db:  "test",
			ans: "insert into test.t values(1, 2)",
		},
		{
			sql: "insert into mydb.t values(1, 2)",
			db:  "test",
			ans: "insert into mydb.t values(1, 2)",
		},
		{
			sql: "insert into t(a, b) values(1, 2)",
			db:  "test",
			ans: "insert into test.t(a, b) values(1, 2)",
		},
		{
			sql: "insert into value value(2, 3)",
			db:  "test",
			ans: "insert into test.value value(2, 3)",
		},
	}

	for _, testCase := range tests {
		p := parser.New()

		stmt, err := p.ParseOneStmt(testCase.sql, "", "")
		require.Nil(t, err)

		ans, ok := utilparser.SimpleCases(stmt, testCase.db, testCase.sql)
		require.True(t, ok)
		require.Equal(t, testCase.ans, ans)
	}
}
