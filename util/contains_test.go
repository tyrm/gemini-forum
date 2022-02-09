package util

import (
	"github.com/google/uuid"
	"testing"
)

func TestContainsString(t *testing.T) {
	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	test1 := "one"
	result := ContainsString(&stack, &test1)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test2 := "four"
	result = ContainsString(&stack, &test2)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test3 := "foo"
	result = ContainsString(&stack, &test3)
	if result != false {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestContainsOneOfStrings(t *testing.T) {
	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	test1 := []string{"one", "two", "three"}
	result := ContainsOneOfStrings(&stack, &test1)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test2 := []string{"foo", "five", "bar"}
	result = ContainsOneOfStrings(&stack, &test2)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test3 := []string{"foo", "bar", "fizz"}
	result = ContainsOneOfStrings(&stack, &test3)
	if result != false {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, false)
	}

}

func TestContainsUUID(t *testing.T) {
	stack := []uuid.UUID{
		uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
		uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
		uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
		uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f")),
		uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
	}

	test1 := uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138"))
	result := ContainsUUID(&stack, &test1)
	if result != true {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, true)
	}

	test2 := uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f"))
	result = ContainsUUID(&stack, &test2)
	if result != true {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, true)
	}

	test3 := uuid.Must(uuid.Parse("319df288-032e-4628-ac4d-be483f263c37"))
	result = ContainsUUID(&stack, &test3)
	if result != false {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, false)
	}
}

func TestContainsOneOfUUIDs(t *testing.T) {
	stack := []uuid.UUID{
		uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
		uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
		uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
		uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f")),
		uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
	}

	test1 := []uuid.UUID{
		uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
		uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
		uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
	}
	result := ContainsOneOfUUIDs(&stack, &test1)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test2 := []uuid.UUID{
		uuid.Must(uuid.Parse("e0f30d4a-250b-4200-9d0f-b057a820e58b")),
		uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
		uuid.Must(uuid.Parse("ac0c8444-eb3b-4a6c-97dc-219f07e66c6c")),
	}
	result = ContainsOneOfUUIDs(&stack, &test2)
	if result != true {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, true)
	}

	test3 := []uuid.UUID{
		uuid.Must(uuid.Parse("2d43417a-5da6-4123-80d6-2f2d42f7477a")),
		uuid.Must(uuid.Parse("4a90659a-423d-471e-baa1-b87a0a66c51a")),
		uuid.Must(uuid.Parse("03c6f272-bb38-486f-bde1-eba3ac6d6ff6")),
	}
	result = ContainsOneOfUUIDs(&stack, &test3)
	if result != false {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", result, false)
	}

}
