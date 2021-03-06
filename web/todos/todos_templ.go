// Code generated by templ DO NOT EDIT.

package todos

import "github.com/a-h/templ"
import "context"
import "io"
import "fmt"

func TodosComponent(todos []Todo) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<table")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"table\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<thead>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<tr>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Name"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Date"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</tr>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</thead>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<tbody>")
		if err != nil {
			return err
		}
		for _, p := range todos {
			err = TodoComponent(p).Render(ctx, w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "</tbody>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</table>")
		if err != nil {
			return err
		}
		return err
	})
}

func TodoComponent(todo Todo) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<tr")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(fmt.Sprintf("post-%d", todo.ID)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(todo.Name))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(todo.Date))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</th>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</tr>")
		if err != nil {
			return err
		}
		return err
	})
}

