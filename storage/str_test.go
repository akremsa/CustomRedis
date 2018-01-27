package storage

import (
	"testing"
)

// func BenchmarkSetStrAndGetStr(b *testing.B) {
// 	st := Init()

// 	for i := 0; i < b.N; i++ {
// 		st.SetStr(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i), 1)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		st.GetStr(fmt.Sprintf("key-%d", i))
// 	}
// }

func Test_SetStr(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := "str1"

	strg.SetStr(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if valueToSet != v.Value {
			t.Fatalf("Expected: %s, actual: %s", valueToSet, v.Value)
		}
	} else {
		t.Fatalf("Expected to find value `%s` by key `%s`", valueToSet, key)
	}
}

func Test_SetStrNX(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := "str1"

	strg.SetStrNX(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if valueToSet != v.Value {
			t.Fatalf("Expected: %s, actual: %s", valueToSet, v.Value)
		}
	} else {
		t.Fatalf("Expected to find value `%s` by key `%s`", valueToSet, key)
	}
}

func Test_SetStrNX_GetKeyExistsError(t *testing.T) {
	strg := Init(0, 1)

	// set value beforehand
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	err := strg.SetStrNX(key, valueToSet, 0)
	if errCustom, ok := err.(ErrBusiness); ok {
		if errCustom.Error() != errKeyExists {
			t.Fatalf("Expected error: `%s`, actual: `%s`", errKeyExists, errCustom.Error())
		}
	} else {
		t.Fatal("Unexpected error type. Expected ErrBusiness")
	}
}

func Test_GetStr(t *testing.T) {
	strg := Init(0, 1)

	// set value beforehand
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	val, err := strg.GetStr(key)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if val != valueToSet {
		t.Fatalf("Expected: %s, actual: %s", valueToSet, val)
	}
}

func Test_GetStr_GetErrWronType(t *testing.T) {
	strg := Init(0, 1)

	// set value beforehand
	key := "key1"
	strg.shards[0].keyValues[key] = Item{Value: []string{"q", "w"}}

	val, err := strg.GetStr(key)
	if err == nil {
		t.Fatal("Expected errWronType but got nil")
	}

	if val != "" {
		t.Fatalf("Expected `val` to be empty, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}
}