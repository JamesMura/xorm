// Copyright 2013 The XORM Authors. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

/*
Package xorm is a simple and powerful ORM for Go.

Installation

Make sure you have installed Go 1.1+ and then:

    go get github.com/lunny/xorm

Create Engine

Firstly, we should new an engine for a database

    engine, err := xorm.NewEngine(driverName, dataSourceName)

Method NewEngine's parameters is the same as sql.Open. It depends
drivers' implementation.
Generally, one engine is enough. You can set it as package variable.

Raw Methods

Xorm also support raw sql execution:

1. query sql, the returned results is []map[string][]byte

    results, err := engine.Query("select * from user")

2. exec sql, the returned results

    affected, err := engine.Exec("update user set .... where ...")

ORM Methods

There are 7 major ORM methods and many helpful methods to use to operate database.

1. Insert one or multipe records to database

    affected, err := engine.Insert(&struct)
    // INSERT INTO struct () values ()
    affected, err := engine.Insert(&struct1, &struct2)
    // INSERT INTO struct1 () values ()
    // INSERT INTO struct2 () values ()
    affected, err := engine.Insert(&sliceOfStruct)
    // INSERT INTO struct () values (),(),()
    affected, err := engine.Insert(&struct1, &sliceOfStruct2)
    // INSERT INTO struct1 () values ()
    // INSERT INTO struct2 () values (),(),()

2. Query one record from database

    has, err := engine.Get(&user)
    // SELECT * FROM user LIMIT 1

3. Query multiple records from database

    err := engine.Find(...)
    // SELECT * FROM user

4. Query multiple records and record by record handle, there two methods, one is Iterate,
another is Raws

    raws, err := engine.Raws(...)
    // SELECT * FROM user
    for raws.Next() {
        raws.Scan(bean)
    }

    err := engine.Iterate(...)
    // SELECT * FROM user

5. Update one or more records

    affected, err := engine.Update(&user)
    // UPDATE user SET

6. Delete one or more records

    affected, err := engine.Delete(&user)
    // DELETE FROM user Where ...

7. Count records

    counts, err := engine.Count(&user)
    // SELECT count(*) AS total FROM user

Conditions

The above 7 methods could use with condition methods.

1. Id, In

    engine.Id(1).Get(&user)
    // SELECT * FROM user WHERE id = 1
    engine.In("id", 1, 2, 3).Find(&users)
    // SELECT * FROM user WHERE id IN (1, 2, 3)

2. Where, And, Or

    engine.Where().And().Or().Find()
    // SELECT * FROM user WHERE (.. AND ..) OR ...

3. OrderBy, Asc, Desc

    engine.Asc().Desc().Find()
    // SELECT * FROM user ORDER BY .. ASC, .. DESC
    engine.OrderBy().Find()
    // SELECT * FROM user ORDER BY ..

4. Limit, Top

    engine.Limit().Find()
    // SELECT * FROM user LIMIT .. OFFSET ..
    engine.Top().Find()
    // SELECT * FROM user LIMIT ..

5. Sql

    engine.Sql("select * from user").Find()

6. Cols, Omit, Distinct

    engine.Cols("col1, col2").Find()
    // SELECT col1, col2 FROM user
    engine.Omit("col1").Find()
    // SELECT col2, col3 FROM user
    engine.Distinct("col1").Find()
    // SELECT DISTINCT col1 FROM user

7. Join, GroupBy, Having

    engine.GroupBy("name").Having("name='xlw'").Find()
    //SELECT * FROM user GROUP BY name HAVING name='xlw'
    engine.Join("LEFT", "userdetail", "user.id=userdetail.id").Find()
    //SELECT * FROM user LEFT JOIN userdetail ON user.id=userdetail.id

More usage, please visit https://github.com/lunny/xorm/blob/master/docs/QuickStartEn.md
*/
package xorm
