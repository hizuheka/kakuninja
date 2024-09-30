package main

import "testing"

// containsStringInFile関数のテーブルドリブンテスト
func TestContainsStringInFile(t *testing.T) {
	// テストケースの構造体
	tests := []struct {
		name         string // テストの名前
		content      string // テスト対象のファイル内容
		searchString string // 検索する文字列
		expected     bool   // 期待する結果
	}{
		{
			name:         "String present in content",
			content:      "Hello, this is a test file.",
			searchString: "test",
			expected:     true,
		},
		{
			name:         "String not present in content",
			content:      "Hello, this is a sample file.",
			searchString: "example",
			expected:     false,
		},
		{
			name:         "Empty content",
			content:      "",
			searchString: "empty",
			expected:     false,
		},
		{
			name:         "Empty search string",
			content:      "Some content here.",
			searchString: "",
			expected:     true,
		},
		{
			name:         "Both content and search string are empty",
			content:      "",
			searchString: "",
			expected:     true,
		},
	}

	// 各テストケースをループして実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 実行結果を関数から取得
			result := containsStringInFile(tt.content, tt.searchString)

			// 結果が期待通りでない場合、エラーメッセージを出力
			if result != tt.expected {
				t.Errorf("Test %q failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
