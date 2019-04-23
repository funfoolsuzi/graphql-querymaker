# GraphQL Query Maker

## Purpose

This package is mainly designed to reduce repitition of code in graphql.

This package will remove the need to have graphql query literal in the code or as text file next to the code. `MakeQuery` method uses Go `reflect` to generate graphql query string and query variable names based on an struct or pointer to struct. Assuming most of user query a graphql API and marshal the reponse data into Go objects(struct instance) with JSON decoder, `MakeQuery` method uses JSON struct field tag. User might not create struct or struct field tags specifically for graphql. However, user need to add struct field tags(`graphqlvar`) to fields where they are intended for query variables.

## How to use

1. We are given these structs

    ```go
    type Species struct {
        Origin OriginLocation `json:"origin"`
    }
    type Animal struct {
        Name string `json:"name"`
        Species string
    }
    type OriginLocation struct {
        Region string `json:"region"`
        Countries []string `json:"countries"`
    }
    ```

2. Make query struct. Let's pretend we are querying Animals in the zoo.

    ```go
    type AnimalQuery struct {
        Animals []Animal `json:"animals"`
    }
    ```

    Query can be generated just like this(without query variables)

    ```go
    q := &AnimalQuery{}
    tmpl, _ := qm.MakeQuery(q)
    ```

    A graphql query like this below will be generated

    ```graphql
    query AnimalQuery {
        animals {
            name
            Species
        }
    }
    ```

3. Add variable tags on original struct

    Edit Animal struct to be like this:

    ```go
    type Animal struct {
        Name string `json:"name",graphqlvar:"name,String"`
        Species string `json:"species"`
    }
    ```

    Run `MakeQuery` again.

    ```graphql
    query AnimalQuery (
        $name: String
    ) {
        animals(name: $name) {
            name
            species
        }
    }
    ```
