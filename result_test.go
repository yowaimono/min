package min

import (
    "encoding/json"
    "testing"
)

func TestResultOf(t *testing.T) {
    data := map[string]string{"key": "value"}
    result := Of(data)

    if result.Msg != "成功" {
        t.Errorf("Expected Msg to be '成功', got '%s'", result.Msg)
    }
    if result.Code != CodeSuccess {
        t.Errorf("Expected Code to be %d, got %d", CodeSuccess, result.Code)
    }
    if result.Data["key"] != "value" {
        t.Errorf("Expected Data['key'] to be 'value', got '%s'", result.Data["key"])
    }
}

func TestError(t *testing.T) {
    result := Error("服务器错误")

    if result.Msg != "服务器错误" {
        t.Errorf("Expected Msg to be '服务器错误', got '%s'", result.Msg)
    }
    if result.Code != CodeError {
        t.Errorf("Expected Code to be %d, got %d", CodeError, result.Code)
    }
    if result.Data != nil {
        t.Errorf("Expected Data to be nil, got %v", result.Data)
    }
}

func TestSuccess(t *testing.T) {
    data := "成功的数据"
    result := Success(data)

    if result.Msg != "成功" {
        t.Errorf("Expected Msg to be '成功', got '%s'", result.Msg)
    }
    if result.Code != CodeSuccess {
        t.Errorf("Expected Code to be %d, got %d", CodeSuccess, result.Code)
    }
    if result.Data != data {
        t.Errorf("Expected Data to be '%s', got '%v'", data, result.Data)
    }
}

func TestFail(t *testing.T) {
    result := Fail("请求参数错误")

    if result.Msg != "请求参数错误" {
        t.Errorf("Expected Msg to be '请求参数错误', got '%s'", result.Msg)
    }
    if result.Code != CodeFail {
        t.Errorf("Expected Code to be %d, got %d", CodeFail, result.Code)
    }
    if result.Data != nil {
        t.Errorf("Expected Data to be nil, got %v", result.Data)
    }
}

func TestJSONSerialization(t *testing.T) {
    data := map[string]string{"key": "value"}
    result := Of(data)

    jsonBytes, err := json.Marshal(result)
    if err != nil {
        t.Fatalf("Failed to serialize Result to JSON: %v", err)
    }

    expectedJSON := `{"msg":"成功","code":200,"data":{"key":"value"}}`
    if string(jsonBytes) != expectedJSON {
        t.Errorf("Expected JSON to be '%s', got '%s'", expectedJSON, string(jsonBytes))
    }
}