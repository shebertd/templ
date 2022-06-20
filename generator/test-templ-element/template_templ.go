// Code generated by templ@(devel) DO NOT EDIT.

package testtemplelement

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"

// GoExpression
import "fmt"

func wrapper(index int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		var_1 := ctx
		ctx = templ.ClearChildren(var_1)
		// Element (standard)
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = io.WriteString(w, " id=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(fmt.Sprint(index)))
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
		// Children
		err = templ.GetChildren(var_1).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func template() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		var_2 := ctx
		ctx = templ.ClearChildren(var_2)
		// TemplElement
		var_3 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			// Text
			var_4 := `child1`
			_, err = io.WriteString(w, var_4)
			if err != nil {
				return err
			}
			// Whitespace (normalised)
			_, err = io.WriteString(w, ` `)
			if err != nil {
				return err
			}
			// TemplElement
			var_5 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				// Text
				var_6 := `child2`
				_, err = io.WriteString(w, var_6)
				if err != nil {
					return err
				}
				// Whitespace (normalised)
				_, err = io.WriteString(w, ` `)
				if err != nil {
					return err
				}
				// TemplElement
				var_7 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
					// Text
					var_8 := `child3`
					_, err = io.WriteString(w, var_8)
					if err != nil {
						return err
					}
					// Whitespace (normalised)
					_, err = io.WriteString(w, ` `)
					if err != nil {
						return err
					}
					// TemplElement
					err = wrapper(4).Render(ctx, w)
					if err != nil {
						return err
					}
					return err
				})
				err = wrapper(3).Render(templ.WithChildren(ctx, var_7), w)
				if err != nil {
					return err
				}
				return err
			})
			err = wrapper(2).Render(templ.WithChildren(ctx, var_5), w)
			if err != nil {
				return err
			}
			return err
		})
		err = wrapper(1).Render(templ.WithChildren(ctx, var_3), w)
		if err != nil {
			return err
		}
		return err
	})
}

