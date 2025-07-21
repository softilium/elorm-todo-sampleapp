package main

import (
	"context"

	"github.com/softilium/elorm"
)

func (dbc *DbContext) setHandlers() {
	err := dbc.AddBeforeDeleteHandlerByRef(dbc.TodoItemDef.EntityDef, func(ctx context.Context, ref string) error {
		comments, _, err := dbc.TodoCommentDef.SelectEntities(
			[]*elorm.Filter{elorm.AddFilterEQ(dbc.TodoCommentDef.TodoItem, ref)},
			nil, 0, 0)
		if err != nil {
			return err
		}
		for _, comment := range comments {
			err = dbc.DeleteEntity(ctx, comment.RefString())
			if err != nil {
				return err
			}
		}
		return nil
	})
	logErr(err)

	err = dbc.AddBeforeDeleteHandlerByRef(dbc.UserDef.EntityDef, func(ctx context.Context, ref string) error {
		todos, _, err := dbc.TodoItemDef.SelectEntities(
			[]*elorm.Filter{elorm.AddFilterEQ(dbc.TodoItemDef.Owner, ref)},
			nil, 0, 0)
		if err != nil {
			return err
		}
		for _, t := range todos {
			err = dbc.DeleteEntity(ctx, t.RefString())
			if err != nil {
				return err
			}
		}
		return nil
	})
	logErr(err)
}
