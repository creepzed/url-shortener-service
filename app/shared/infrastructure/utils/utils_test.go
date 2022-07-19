package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockStructure struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type mockStructureNew mockStructure

func Test_ConvertEntity(t *testing.T) {

	t.Parallel()

	t.Run("When an entity of a specific type is delivered and another of another type returns with the same structure ok", func(t *testing.T) {
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		e2 := new(mockStructureNew)

		err := ConvertEntity(e1, e2)

		require.NoError(t, err)
		assert.Equal(t, EntityToJson(e1), EntityToJson(e2))
	})

	t.Run("When an entity of a specific type is delivered and another of another type returns with a distillate structure and returns nil", func(t *testing.T) {
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		e2 := new(string)

		err := ConvertEntity(e1, e2)

		require.Error(t, err)
		assert.Equal(t, e2, "")
	})
}

func Test_EntityToJson(t *testing.T) {

	t.Parallel()

	t.Run("When an entity is delivered and json returns with values", func(t *testing.T) {

		jsonExpected := "{\"id\":1,\"name\":\"name\"}"
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		result := EntityToJson(e1)

		assert.Equal(t, jsonExpected, result)

	})

	t.Run("When an invalid entity is delivered and json returns empty", func(t *testing.T) {

		jsonExpected := "{}"

		result := EntityToJson(make(chan int))

		assert.Equal(t, jsonExpected, result)

	})
}

func Test_EntityToJsonEscape(t *testing.T) {

	t.Parallel()

	t.Run("When an entity is delivered and json returns with values", func(t *testing.T) {

		jsonExpected := "{\"id\":1,\"name\":\"name\"}"
		e1 := mockStructure{
			Id:   1,
			Name: "name",
		}

		result := EntityToJsonEscape(e1)

		assert.Equal(t, jsonExpected, result)

	})

	t.Run("When an invalid entity is delivered and json returns empty", func(t *testing.T) {

		jsonExpected := "{}"

		result := EntityToJsonEscape(make(chan int))

		assert.Equal(t, jsonExpected, result)

	})
}

func Test_JsonToEntity(t *testing.T) {
	t.Parallel()
	t.Run("When json is sent valid and return entity with data", func(t *testing.T) {

		json := "{\"id\":1,\"name\":\"name\"}"

		entity := new(mockStructure)
		JsonToEntity(json, entity)
		assert.Equal(t, 1, entity.Id)
		assert.Equal(t, "name", entity.Name)
	})

	t.Run("When json is sent valid and returns empty entity", func(t *testing.T) {
		json := "$%&/()/"
		entityExpected := new(mockStructure)
		entity := new(mockStructure)

		JsonToEntity(json, entity)
		assert.Equal(t, entityExpected, entity)
	})
}
