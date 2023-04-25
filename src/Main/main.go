package main

import (
	"log"
	"main/person"
)

func main() {
	{
		var pp person.IPerson = person.NewPerson("Ann", 21)
		var ss person.IStudent = person.NewStudent("Bob", 22, "cs")
		var ps person.IPerson = person.NewStudent("Carla", 23, "math")
		log.Printf("pp.Name() => %s\n", pp.Name())
		log.Printf("ss.Name() => %s\n", ss.Name())
		log.Printf("pp.String() => %s\n", pp)
		log.Printf("ss.String() => %s\n", ss)
		log.Printf("ps.Name() => %s\n", ps.Name())
		log.Printf("ps.String() => %s\n", ps)
	}
	{
		var pp *person.Person = person.NewPerson("Ann", 21)
		var ss *person.Student = person.NewStudent("Bob", 22, "cs")
		var ps person.IPerson = person.NewStudent("Carla", 23, "math")
		log.Printf("pp.Name() => %s\n", pp.Name())
		log.Printf("ss.Name() => %s\n", ss.Name())
		log.Printf("pp.String() => %s\n", pp)
		log.Printf("ss.String() => %s\n", ss)
		log.Printf("ps.Name() => %s\n", ps.Name())
		log.Printf("ps.String() => %s\n", ps)
	}
	{
		pp := person.NewPerson("Ann", 21)
		ss := person.NewStudent("Bob", 22, "cs")
		ps := person.NewStudent("Carla", 23, "math")
		log.Printf("pp.Name() => %s\n", pp.Name())
		log.Printf("ss.Name() => %s\n", ss.Name())
		log.Printf("pp.String() => %s\n", pp)
		log.Printf("ss.String() => %s\n", ss)
		log.Printf("ps.Name() => %s\n", ps.Name())
		log.Printf("ps.String() => %s\n", ps)
	}
}
