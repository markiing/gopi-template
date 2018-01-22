package connection

import (
	m "gopkg.in/mgo.v2"
)

var database = "go";

func openSession()(*m.Session,error){
	s,err := m.Dial("localhost")

	if err != nil {
		panic("Oops. Houve um problema ao estabelecer a conexão com a base de dados")
	}

	return s,err
}

func withCollection(collection string) (*m.Collection, *m.Session){
	c,_ := openSession()
	return c.DB(database).C(collection),c
}

func Find(collection string,param interface{})(q *m.Query){
	c, _ := withCollection(collection)
	query := c.Find(param)
	return query
}

func Insert(collection string, param interface{}){
	c, _ := withCollection(collection)
	c.Insert(param)
}


func CloseSession(session *m.Session){
	session.Close()
}