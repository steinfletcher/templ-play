// Code generated by templ DO NOT EDIT.

package layout

import "github.com/a-h/templ"
import "context"
import "io"
import "github.com/steinfletcher/templ-play/web/todos"

func IndexPage(todoItems []todos.Todo) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<html>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<link")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " rel=\"stylesheet\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " crossorigin=\"anonymous\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</link>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<body>")
		if err != nil {
			return err
		}
		err = todos.TodosComponent(todoItems).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</html>")
		if err != nil {
			return err
		}
		return err
	})
}
