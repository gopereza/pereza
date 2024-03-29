package boolean

import (
	"github.com/gopereza/pereza/fixtures/bson/boolean"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var (
	oneDataProvider = map[bool]string{
		false: "\x0d\x00\x00\x00\x08state\x00\x00\x00",
		true:  "\x0d\x00\x00\x00\x08state\x00\x01\x00",
	}
)

func TestOneMongoMarshalBSON(t *testing.T) {
	for state, expect := range oneDataProvider {
		source := boolean.BoolState{
			State: state,
		}

		actual, err := bson.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestOnePerezaMarshalBSON(t *testing.T) {
	for state, expect := range oneDataProvider {
		source := boolean.PerezaBoolState{
			State: state,
		}

		actual, err := source.MarshalBSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkOneMongoMarshalBSON(b *testing.B) {
	source := boolean.BoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = bson.Marshal(source)
	}
}

func BenchmarkOnePerezaMarshalBSON(b *testing.B) {
	source := boolean.PerezaBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalBSON()
	}
}
