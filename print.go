package querymaker

import "fmt"

func (q *query) generateTmplNVars() (tmpl string, vars []string) {
	tmpl = fmt.Sprintf("query %s%s {\n", q.name, q.getVarStr())
	for _, f := range q.subGqlFields {
		tmpl += printField(f, 1)
	}
	tmpl += "}\n"
	for key := range q.variables {
		vars = append(vars, key)
	}
	return
}

func (q *query) String() string {
	tmpl, _ := q.generateTmplNVars()
	return tmpl
}

func (q *query) getVarStr() string {
	if len(q.variables) == 0 {
		return ""
	}
	res := " (\n"
	for k, v := range q.variables {
		res += fmt.Sprintf("%s$%s: %s\n", tab, k, v)
	}
	res += ")"

	return res
}

// printField prints fields to a string recursively
func printField(f *gqlField, depth int) string {
	tabs := ""
	for i := 0; i < depth; i++ {
		tabs += tab
	}
	fn := tabs + f.name

	if len(f.subGqlFields) == 0 {
		return fmt.Sprintf("%s\n", fn)
	}

	vars := ""
	subfields := ""
	for _, sub := range f.subGqlFields {
		if sub.variableName != "" {
			if vars == "" { // first variable
				vars += fmt.Sprintf("%s: $%s", sub.name, sub.variableName)
			} else { // variables after the first one
				vars += fmt.Sprintf(", %s: $%s", sub.name, sub.variableName)
			}
		}
		subfields += printField(sub, depth+1)
	}
	if vars != "" {
		vars = fmt.Sprintf("(%s)", vars)
	}

	return fmt.Sprintf("%s%s {\n%s%s}\n", fn, vars, subfields, tabs)
}
