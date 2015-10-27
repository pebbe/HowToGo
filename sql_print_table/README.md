## database/sql: printing query results as table

interface{} as variable

```go
func makeFields(rows *sql.Rows) (fields []interface{}) {
    cols, err := rows.Columns()
    checkErr(err)
    for _, c := range cols {
        switch {
        case c[:2] == "b:":
            fields = append(fields, new(sql.NullBool))
        case c[:2] == "f:":
            fields = append(fields, new(sql.NullFloat64))
        case c[:2] == "i:":
            fields = append(fields, new(sql.NullInt64))
        case c[:2] == "s:":
            fields = append(fields, new(sql.NullString))
        case c[:2] == "t:":
            fields = append(fields, new(time.Time))
        default:
            fields = append(fields, new(sql.NullString))
        }
    }
    return
}

func pr(t interface{}) (r string) {
    r = "\\N"
    switch v := t.(type) {
    case *sql.NullBool:
        if v.Valid {
            r = fmt.Sprintf("%v", v.Bool)
        }
    case *sql.NullString:
        if v.Valid {
            r = v.String
        }
    case *sql.NullInt64:
        if v.Valid {
            r = fmt.Sprintf("%6d", v.Int64)
        }
    case *sql.NullFloat64:
        if v.Valid {
            r = fmt.Sprintf("%.2f", v.Float64)
        }
    case *time.Time:
        if v.Year() &gt; 1900 {
            r = v.Format("_2 Jan 2006")
        }
    default:
        r = fmt.Sprintf("%#v", t)
    }
    return
}

func printTable(rows *sql.Rows) {

    // print labels
    cols, err := rows.Columns()
    checkErr(err)
    for _, c := range cols {
        if len(c) &gt; 1 &amp;&amp; c[1] == ':' {
            fmt.Print(c[2:], "\t")
        } else {
            fmt.Print(c, "\t")
        }
    }
    fmt.Print("\n\n")

    // print data
    fields := makeFields(rows)
    for rows.Next() {
        checkErr(rows.Scan(fields...))
        for _, f := range fields {
            fmt.Print(pr(f), "\t")
        }
        fmt.Println()
    }
    fmt.Println()
}
```

Example usage:

```go
    rows, err = db.Query("select id as \"i:id\", name as \"s:name\", date as \"t:date\" from `items`;")
    if err == nil {
        printTable(rows)
    }
```
