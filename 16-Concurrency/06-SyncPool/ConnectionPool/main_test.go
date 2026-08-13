package main

import "testing"

func BenchmarkWithoutPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range 10000 {
			connection = &DbConnection{
				Host:     "127.0.0.1",
				DbName:   "test",
				User:     "root",
				Password: "root",
			}
			connection.DbName = "test"
		}
	}

}
func BenchmarkWithPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range 10000 {
			connection = connectionPool.Get().(*DbConnection)
			connection.DbName = "test"
			connectionPool.Put(connection)
		}
	}
}
