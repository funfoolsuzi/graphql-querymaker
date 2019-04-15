# GraphQL Query Maker

## Purpose

This package is mainly designed to reduce repitition of code in graphql.

This package will remove the need to have graphql query literal in the code or as text file next to the code. `MakeQuery` method uses Go `reflect` to generate graphql query string and query variable names based on an struct or pointer to struct. Assuming most of user query a graphql API and marshal the reponse data into Go objects(struct instance) with JSON decoder, `MakeQuery` method uses JSON struct field tag. User might not create struct or struct field tags specifically for graphql. However, user need to add struct field tags(`graphqlvar`) to fields where they are intended for query variables.
