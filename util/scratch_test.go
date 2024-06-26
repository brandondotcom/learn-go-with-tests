package util

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	t.Run("with zero value members", func(t *testing.T) {
		person := Person{}

		fmt.Printf("%#v\n\n", person)

		if person.address.streetNumber != 0 {
			t.Fail()
		}

		// panics
		if person.addressRef.streetNumber != 0 {
			t.Fail()
		}
	})
}

func BenchmarkAddress(b *testing.B) {
	b.Run("copy by value", func(b *testing.B) {
		b.ReportAllocs()

		original := Address{streetNumber: 10, streetName: "cardinal way"}
		copies := make([]Address, b.N)
		for i := 0; i < b.N; i++ {
			copies[i] = original
		}

		if copies[0].streetNumber != 10 {
			b.Fail()
		}
	})

	b.Run("copy by reference", func(b *testing.B) {
		b.ReportAllocs()

		original := &Address{streetNumber: 10, streetName: "cardinal way"}
		copies := make([]*Address, b.N)
		for i := 0; i < b.N; i++ {
			copies[i] = original
		}

		if copies[0].streetNumber != 10 {
			b.Fail()
		}
	})
}
