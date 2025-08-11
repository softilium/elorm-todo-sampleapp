// THIS IS GENERATED CODE. DO NOT MODIFY IT
package main

import (
	"github.com/softilium/elorm"
	"time"
)

// Tag class
//////

type TagDefStruct struct {
	*elorm.EntityDef
	Ref         *elorm.FieldDef
	IsDeleted   *elorm.FieldDef
	DataVersion *elorm.FieldDef

	Name *elorm.FieldDef

	Color *elorm.FieldDef
}

func (T *TagDefStruct) SelectEntities(filters []*elorm.Filter, sorts []*elorm.SortItem, pageNo int, pageSize int) (result []*Tag, pages int, err error) {

	res, total, err := T.EntityDef.SelectEntities(filters, sorts, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}

	res2 := make([]*Tag, 0, len(res))

	for _, r := range res {
		if r == nil {
			continue
		}
		rt := T.Wrap(r)
		res2 = append(res2, rt.(*Tag))
	}

	return res2, total, nil

}

type Tag struct {
	*elorm.Entity

	field_Name  *elorm.FieldValueString
	field_Color *elorm.FieldValueString
}

func (T *Tag) Name() string {
	if T.field_Name == nil {
		T.field_Name = T.Values["Name"].(*elorm.FieldValueString)
	}
	return T.field_Name.Get()
}

func (T *Tag) SetName(newValue string) {
	if T.field_Name == nil {
		T.field_Name = T.Values["Name"].(*elorm.FieldValueString)
	}
	T.field_Name.Set(newValue)
}

func (T *Tag) Color() string {
	if T.field_Color == nil {
		T.field_Color = T.Values["Color"].(*elorm.FieldValueString)
	}
	return T.field_Color.Get()
}

func (T *Tag) SetColor(newValue string) {
	if T.field_Color == nil {
		T.field_Color = T.Values["Color"].(*elorm.FieldValueString)
	}
	T.field_Color.Set(newValue)
}

// TagInstance class
//////

type TagInstanceDefStruct struct {
	*elorm.EntityDef
	Ref         *elorm.FieldDef
	IsDeleted   *elorm.FieldDef
	DataVersion *elorm.FieldDef

	Tag *elorm.FieldDef

	TodoItem *elorm.FieldDef
}

func (T *TagInstanceDefStruct) SelectEntities(filters []*elorm.Filter, sorts []*elorm.SortItem, pageNo int, pageSize int) (result []*TagInstance, pages int, err error) {

	res, total, err := T.EntityDef.SelectEntities(filters, sorts, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}

	res2 := make([]*TagInstance, 0, len(res))

	for _, r := range res {
		if r == nil {
			continue
		}
		rt := T.Wrap(r)
		res2 = append(res2, rt.(*TagInstance))
	}

	return res2, total, nil

}

type TagInstance struct {
	*elorm.Entity

	field_Tag      *elorm.FieldValueRef
	field_TodoItem *elorm.FieldValueRef
}

func (T *TagInstance) Tag() *Tag {
	if T.field_Tag == nil {
		T.field_Tag = T.Values["Tag"].(*elorm.FieldValueRef)
	}
	r, err := T.field_Tag.Get()
	if err != nil {
		panic(err)
	}
	return r.(*Tag)
}

func (T *TagInstance) SetTag(newValue *Tag) {
	if T.field_Tag == nil {
		T.field_Tag = T.Values["Tag"].(*elorm.FieldValueRef)
	}
	err := T.field_Tag.Set(newValue)
	if err != nil {
		panic(err)
	}
}

func (T *TagInstance) TodoItem() *TodoItem {
	if T.field_TodoItem == nil {
		T.field_TodoItem = T.Values["TodoItem"].(*elorm.FieldValueRef)
	}
	r, err := T.field_TodoItem.Get()
	if err != nil {
		panic(err)
	}
	return r.(*TodoItem)
}

func (T *TagInstance) SetTodoItem(newValue *TodoItem) {
	if T.field_TodoItem == nil {
		T.field_TodoItem = T.Values["TodoItem"].(*elorm.FieldValueRef)
	}
	err := T.field_TodoItem.Set(newValue)
	if err != nil {
		panic(err)
	}
}

// TodoComment class
//////

type TodoCommentDefStruct struct {
	*elorm.EntityDef
	Ref         *elorm.FieldDef
	IsDeleted   *elorm.FieldDef
	DataVersion *elorm.FieldDef

	TodoItem *elorm.FieldDef

	Author *elorm.FieldDef

	Content *elorm.FieldDef

	CreatedDate *elorm.FieldDef
}

func (T *TodoCommentDefStruct) SelectEntities(filters []*elorm.Filter, sorts []*elorm.SortItem, pageNo int, pageSize int) (result []*TodoComment, pages int, err error) {

	res, total, err := T.EntityDef.SelectEntities(filters, sorts, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}

	res2 := make([]*TodoComment, 0, len(res))

	for _, r := range res {
		if r == nil {
			continue
		}
		rt := T.Wrap(r)
		res2 = append(res2, rt.(*TodoComment))
	}

	return res2, total, nil

}

type TodoComment struct {
	*elorm.Entity

	field_TodoItem    *elorm.FieldValueRef
	field_Author      *elorm.FieldValueRef
	field_Content     *elorm.FieldValueString
	field_CreatedDate *elorm.FieldValueDateTime
}

func (T *TodoComment) TodoItem() *TodoItem {
	if T.field_TodoItem == nil {
		T.field_TodoItem = T.Values["TodoItem"].(*elorm.FieldValueRef)
	}
	r, err := T.field_TodoItem.Get()
	if err != nil {
		panic(err)
	}
	return r.(*TodoItem)
}

func (T *TodoComment) SetTodoItem(newValue *TodoItem) {
	if T.field_TodoItem == nil {
		T.field_TodoItem = T.Values["TodoItem"].(*elorm.FieldValueRef)
	}
	err := T.field_TodoItem.Set(newValue)
	if err != nil {
		panic(err)
	}
}

func (T *TodoComment) Author() *User {
	if T.field_Author == nil {
		T.field_Author = T.Values["Author"].(*elorm.FieldValueRef)
	}
	r, err := T.field_Author.Get()
	if err != nil {
		panic(err)
	}
	return r.(*User)
}

func (T *TodoComment) SetAuthor(newValue *User) {
	if T.field_Author == nil {
		T.field_Author = T.Values["Author"].(*elorm.FieldValueRef)
	}
	err := T.field_Author.Set(newValue)
	if err != nil {
		panic(err)
	}
}

func (T *TodoComment) Content() string {
	if T.field_Content == nil {
		T.field_Content = T.Values["Content"].(*elorm.FieldValueString)
	}
	return T.field_Content.Get()
}

func (T *TodoComment) SetContent(newValue string) {
	if T.field_Content == nil {
		T.field_Content = T.Values["Content"].(*elorm.FieldValueString)
	}
	T.field_Content.Set(newValue)
}

func (T *TodoComment) CreatedDate() time.Time {
	if T.field_CreatedDate == nil {
		T.field_CreatedDate = T.Values["CreatedDate"].(*elorm.FieldValueDateTime)
	}
	return T.field_CreatedDate.Get()
}

func (T *TodoComment) SetCreatedDate(newValue time.Time) {
	if T.field_CreatedDate == nil {
		T.field_CreatedDate = T.Values["CreatedDate"].(*elorm.FieldValueDateTime)
	}
	T.field_CreatedDate.Set(newValue)
}

// TodoItem class
//////

type TodoItemDefStruct struct {
	*elorm.EntityDef
	Ref         *elorm.FieldDef
	IsDeleted   *elorm.FieldDef
	DataVersion *elorm.FieldDef

	Title *elorm.FieldDef

	Description *elorm.FieldDef

	Owner *elorm.FieldDef

	DueDate *elorm.FieldDef

	IsCompleted *elorm.FieldDef

	CompletedDate *elorm.FieldDef

	Priority *elorm.FieldDef
}

func (T *TodoItemDefStruct) SelectEntities(filters []*elorm.Filter, sorts []*elorm.SortItem, pageNo int, pageSize int) (result []*TodoItem, pages int, err error) {

	res, total, err := T.EntityDef.SelectEntities(filters, sorts, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}

	res2 := make([]*TodoItem, 0, len(res))

	for _, r := range res {
		if r == nil {
			continue
		}
		rt := T.Wrap(r)
		res2 = append(res2, rt.(*TodoItem))
	}

	return res2, total, nil

}

type TodoItem struct {
	*elorm.Entity

	field_Title         *elorm.FieldValueString
	field_Description   *elorm.FieldValueString
	field_Owner         *elorm.FieldValueRef
	field_DueDate       *elorm.FieldValueDateTime
	field_IsCompleted   *elorm.FieldValueBool
	field_CompletedDate *elorm.FieldValueDateTime
	field_Priority      *elorm.FieldValueInt
}

func (T *TodoItem) Title() string {
	if T.field_Title == nil {
		T.field_Title = T.Values["Title"].(*elorm.FieldValueString)
	}
	return T.field_Title.Get()
}

func (T *TodoItem) SetTitle(newValue string) {
	if T.field_Title == nil {
		T.field_Title = T.Values["Title"].(*elorm.FieldValueString)
	}
	T.field_Title.Set(newValue)
}

func (T *TodoItem) Description() string {
	if T.field_Description == nil {
		T.field_Description = T.Values["Description"].(*elorm.FieldValueString)
	}
	return T.field_Description.Get()
}

func (T *TodoItem) SetDescription(newValue string) {
	if T.field_Description == nil {
		T.field_Description = T.Values["Description"].(*elorm.FieldValueString)
	}
	T.field_Description.Set(newValue)
}

func (T *TodoItem) Owner() *User {
	if T.field_Owner == nil {
		T.field_Owner = T.Values["Owner"].(*elorm.FieldValueRef)
	}
	r, err := T.field_Owner.Get()
	if err != nil {
		panic(err)
	}
	return r.(*User)
}

func (T *TodoItem) SetOwner(newValue *User) {
	if T.field_Owner == nil {
		T.field_Owner = T.Values["Owner"].(*elorm.FieldValueRef)
	}
	err := T.field_Owner.Set(newValue)
	if err != nil {
		panic(err)
	}
}

func (T *TodoItem) DueDate() time.Time {
	if T.field_DueDate == nil {
		T.field_DueDate = T.Values["DueDate"].(*elorm.FieldValueDateTime)
	}
	return T.field_DueDate.Get()
}

func (T *TodoItem) SetDueDate(newValue time.Time) {
	if T.field_DueDate == nil {
		T.field_DueDate = T.Values["DueDate"].(*elorm.FieldValueDateTime)
	}
	T.field_DueDate.Set(newValue)
}

func (T *TodoItem) IsCompleted() bool {
	if T.field_IsCompleted == nil {
		T.field_IsCompleted = T.Values["IsCompleted"].(*elorm.FieldValueBool)
	}
	return T.field_IsCompleted.Get()
}

func (T *TodoItem) SetIsCompleted(newValue bool) {
	if T.field_IsCompleted == nil {
		T.field_IsCompleted = T.Values["IsCompleted"].(*elorm.FieldValueBool)
	}
	T.field_IsCompleted.Set(newValue)
}

func (T *TodoItem) CompletedDate() time.Time {
	if T.field_CompletedDate == nil {
		T.field_CompletedDate = T.Values["CompletedDate"].(*elorm.FieldValueDateTime)
	}
	return T.field_CompletedDate.Get()
}

func (T *TodoItem) SetCompletedDate(newValue time.Time) {
	if T.field_CompletedDate == nil {
		T.field_CompletedDate = T.Values["CompletedDate"].(*elorm.FieldValueDateTime)
	}
	T.field_CompletedDate.Set(newValue)
}

func (T *TodoItem) Priority() int64 {
	if T.field_Priority == nil {
		T.field_Priority = T.Values["Priority"].(*elorm.FieldValueInt)
	}
	return T.field_Priority.Get()
}

func (T *TodoItem) SetPriority(newValue int64) {
	if T.field_Priority == nil {
		T.field_Priority = T.Values["Priority"].(*elorm.FieldValueInt)
	}
	T.field_Priority.Set(newValue)
}

// User class
//////

type UserDefStruct struct {
	*elorm.EntityDef
	Ref         *elorm.FieldDef
	IsDeleted   *elorm.FieldDef
	DataVersion *elorm.FieldDef

	Username *elorm.FieldDef

	Email *elorm.FieldDef

	PasswordHash *elorm.FieldDef

	LastLoginDate *elorm.FieldDef

	IsActive *elorm.FieldDef
}

func (T *UserDefStruct) SelectEntities(filters []*elorm.Filter, sorts []*elorm.SortItem, pageNo int, pageSize int) (result []*User, pages int, err error) {

	res, total, err := T.EntityDef.SelectEntities(filters, sorts, pageNo, pageSize)
	if err != nil {
		return nil, 0, err
	}

	res2 := make([]*User, 0, len(res))

	for _, r := range res {
		if r == nil {
			continue
		}
		rt := T.Wrap(r)
		res2 = append(res2, rt.(*User))
	}

	return res2, total, nil

}

type User struct {
	*elorm.Entity

	field_Username      *elorm.FieldValueString
	field_Email         *elorm.FieldValueString
	field_PasswordHash  *elorm.FieldValueString
	field_LastLoginDate *elorm.FieldValueDateTime
	field_IsActive      *elorm.FieldValueBool
}

func (T *User) Username() string {
	if T.field_Username == nil {
		T.field_Username = T.Values["Username"].(*elorm.FieldValueString)
	}
	return T.field_Username.Get()
}

func (T *User) SetUsername(newValue string) {
	if T.field_Username == nil {
		T.field_Username = T.Values["Username"].(*elorm.FieldValueString)
	}
	T.field_Username.Set(newValue)
}

func (T *User) Email() string {
	if T.field_Email == nil {
		T.field_Email = T.Values["Email"].(*elorm.FieldValueString)
	}
	return T.field_Email.Get()
}

func (T *User) SetEmail(newValue string) {
	if T.field_Email == nil {
		T.field_Email = T.Values["Email"].(*elorm.FieldValueString)
	}
	T.field_Email.Set(newValue)
}

func (T *User) PasswordHash() string {
	if T.field_PasswordHash == nil {
		T.field_PasswordHash = T.Values["PasswordHash"].(*elorm.FieldValueString)
	}
	return T.field_PasswordHash.Get()
}

func (T *User) SetPasswordHash(newValue string) {
	if T.field_PasswordHash == nil {
		T.field_PasswordHash = T.Values["PasswordHash"].(*elorm.FieldValueString)
	}
	T.field_PasswordHash.Set(newValue)
}

func (T *User) LastLoginDate() time.Time {
	if T.field_LastLoginDate == nil {
		T.field_LastLoginDate = T.Values["LastLoginDate"].(*elorm.FieldValueDateTime)
	}
	return T.field_LastLoginDate.Get()
}

func (T *User) SetLastLoginDate(newValue time.Time) {
	if T.field_LastLoginDate == nil {
		T.field_LastLoginDate = T.Values["LastLoginDate"].(*elorm.FieldValueDateTime)
	}
	T.field_LastLoginDate.Set(newValue)
}

func (T *User) IsActive() bool {
	if T.field_IsActive == nil {
		T.field_IsActive = T.Values["IsActive"].(*elorm.FieldValueBool)
	}
	return T.field_IsActive.Get()
}

func (T *User) SetIsActive(newValue bool) {
	if T.field_IsActive == nil {
		T.field_IsActive = T.Values["IsActive"].(*elorm.FieldValueBool)
	}
	T.field_IsActive.Set(newValue)
}

// DbContext core
//////

type DbContext struct {
	*elorm.Factory
	TagDef         TagDefStruct
	TagInstanceDef TagInstanceDefStruct
	TodoCommentDef TodoCommentDefStruct
	TodoItemDef    TodoItemDefStruct
	UserDef        UserDefStruct
}

func CreateDbContext(dbDialect string, connectionString string) (*DbContext, error) {

	var err error
	frg := []string{}
	_ = frg // to avoid unused variable error if no fragments are defined

	r := &DbContext{}
	r.Factory, err = elorm.CreateFactory(dbDialect, connectionString)
	if err != nil {
		return nil, err
	}

	r.TagDef.EntityDef, err = r.CreateEntityDef("Tag", "Tags")
	if err != nil {
		return nil, err
	}

	r.TagDef.EntityDef.UseSoftDelete = false

	r.TagInstanceDef.EntityDef, err = r.CreateEntityDef("TagInstance", "TagInstances")
	if err != nil {
		return nil, err
	}

	r.TagInstanceDef.EntityDef.UseSoftDelete = false

	r.TodoCommentDef.EntityDef, err = r.CreateEntityDef("TodoComment", "TodoComments")
	if err != nil {
		return nil, err
	}

	r.TodoCommentDef.EntityDef.UseSoftDelete = false

	r.TodoItemDef.EntityDef, err = r.CreateEntityDef("TodoItem", "TodoItems")
	if err != nil {
		return nil, err
	}

	r.TodoItemDef.EntityDef.UseSoftDelete = false

	r.UserDef.EntityDef, err = r.CreateEntityDef("User", "Users")
	if err != nil {
		return nil, err
	}

	r.UserDef.EntityDef.UseSoftDelete = false

	// Tag
	//////

	r.TagDef.Ref = r.TagDef.FieldDefByName("Ref")
	r.TagDef.IsDeleted = r.TagDef.FieldDefByName("IsDeleted")
	r.TagDef.DataVersion = r.TagDef.FieldDefByName("DataVersion")

	r.TagDef.Name, _ = r.TagDef.AddStringFieldDef("Name", 15)
	r.TagDef.Color, _ = r.TagDef.AddStringFieldDef("Color", 20)

	r.TagDef.Wrap = func(source *elorm.Entity) any { return &Tag{Entity: source} }

	// TagInstance
	//////

	r.TagInstanceDef.Ref = r.TagInstanceDef.FieldDefByName("Ref")
	r.TagInstanceDef.IsDeleted = r.TagInstanceDef.FieldDefByName("IsDeleted")
	r.TagInstanceDef.DataVersion = r.TagInstanceDef.FieldDefByName("DataVersion")

	r.TagInstanceDef.Tag, _ = r.TagInstanceDef.AddRefFieldDef("Tag", r.TagDef.EntityDef)
	r.TagInstanceDef.TodoItem, _ = r.TagInstanceDef.AddRefFieldDef("TodoItem", r.TodoItemDef.EntityDef)

	r.TagInstanceDef.Wrap = func(source *elorm.Entity) any { return &TagInstance{Entity: source} }

	// TodoComment
	//////

	r.TodoCommentDef.Ref = r.TodoCommentDef.FieldDefByName("Ref")
	r.TodoCommentDef.IsDeleted = r.TodoCommentDef.FieldDefByName("IsDeleted")
	r.TodoCommentDef.DataVersion = r.TodoCommentDef.FieldDefByName("DataVersion")

	r.TodoCommentDef.TodoItem, _ = r.TodoCommentDef.AddRefFieldDef("TodoItem", r.TodoItemDef.EntityDef)
	r.TodoCommentDef.Author, _ = r.TodoCommentDef.AddRefFieldDef("Author", r.UserDef.EntityDef)
	r.TodoCommentDef.Content, _ = r.TodoCommentDef.AddStringFieldDef("Content", 500)
	r.TodoCommentDef.CreatedDate, _ = r.TodoCommentDef.AddDateTimeFieldDef("CreatedDate")

	r.TodoCommentDef.Wrap = func(source *elorm.Entity) any { return &TodoComment{Entity: source} }

	// TodoItem
	//////

	r.TodoItemDef.Ref = r.TodoItemDef.FieldDefByName("Ref")
	r.TodoItemDef.IsDeleted = r.TodoItemDef.FieldDefByName("IsDeleted")
	r.TodoItemDef.DataVersion = r.TodoItemDef.FieldDefByName("DataVersion")

	r.TodoItemDef.Title, _ = r.TodoItemDef.AddStringFieldDef("Title", 200)
	r.TodoItemDef.Description, _ = r.TodoItemDef.AddStringFieldDef("Description", 500)
	r.TodoItemDef.Owner, _ = r.TodoItemDef.AddRefFieldDef("Owner", r.UserDef.EntityDef)
	r.TodoItemDef.DueDate, _ = r.TodoItemDef.AddDateTimeFieldDef("DueDate")
	r.TodoItemDef.IsCompleted, _ = r.TodoItemDef.AddBoolFieldDef("IsCompleted")
	r.TodoItemDef.CompletedDate, _ = r.TodoItemDef.AddDateTimeFieldDef("CompletedDate")
	r.TodoItemDef.Priority, _ = r.TodoItemDef.AddIntFieldDef("Priority")

	r.TodoItemDef.Wrap = func(source *elorm.Entity) any { return &TodoItem{Entity: source} }

	// User
	//////

	r.UserDef.Ref = r.UserDef.FieldDefByName("Ref")
	r.UserDef.IsDeleted = r.UserDef.FieldDefByName("IsDeleted")
	r.UserDef.DataVersion = r.UserDef.FieldDefByName("DataVersion")

	r.UserDef.Username, _ = r.UserDef.AddStringFieldDef("Username", 50)
	r.UserDef.Email, _ = r.UserDef.AddStringFieldDef("Email", 100)
	r.UserDef.PasswordHash, _ = r.UserDef.AddStringFieldDef("PasswordHash", 256)
	r.UserDef.LastLoginDate, _ = r.UserDef.AddDateTimeFieldDef("LastLoginDate")
	r.UserDef.IsActive, _ = r.UserDef.AddBoolFieldDef("IsActive")

	r.UserDef.Wrap = func(source *elorm.Entity) any { return &User{Entity: source} }

	return r, nil
}

func (T *DbContext) CreateTag() (*Tag, error) {
	r, err := T.CreateEntityWrapped(T.TagDef.EntityDef)
	if err != nil {
		return nil, err
	}
	rt := r.(*Tag)
	return rt, nil
}

func (T *DbContext) LoadTag(Ref string) (*Tag, error) {
	r, err := T.LoadEntityWrapped(Ref)
	if err != nil {
		return nil, err
	}
	rt := r.(*Tag)
	return rt, nil
}

func (T *DbContext) CreateTagInstance() (*TagInstance, error) {
	r, err := T.CreateEntityWrapped(T.TagInstanceDef.EntityDef)
	if err != nil {
		return nil, err
	}
	rt := r.(*TagInstance)
	return rt, nil
}

func (T *DbContext) LoadTagInstance(Ref string) (*TagInstance, error) {
	r, err := T.LoadEntityWrapped(Ref)
	if err != nil {
		return nil, err
	}
	rt := r.(*TagInstance)
	return rt, nil
}

func (T *DbContext) CreateTodoComment() (*TodoComment, error) {
	r, err := T.CreateEntityWrapped(T.TodoCommentDef.EntityDef)
	if err != nil {
		return nil, err
	}
	rt := r.(*TodoComment)
	return rt, nil
}

func (T *DbContext) LoadTodoComment(Ref string) (*TodoComment, error) {
	r, err := T.LoadEntityWrapped(Ref)
	if err != nil {
		return nil, err
	}
	rt := r.(*TodoComment)
	return rt, nil
}

func (T *DbContext) CreateTodoItem() (*TodoItem, error) {
	r, err := T.CreateEntityWrapped(T.TodoItemDef.EntityDef)
	if err != nil {
		return nil, err
	}
	rt := r.(*TodoItem)
	return rt, nil
}

func (T *DbContext) LoadTodoItem(Ref string) (*TodoItem, error) {
	r, err := T.LoadEntityWrapped(Ref)
	if err != nil {
		return nil, err
	}
	rt := r.(*TodoItem)
	return rt, nil
}

func (T *DbContext) CreateUser() (*User, error) {
	r, err := T.CreateEntityWrapped(T.UserDef.EntityDef)
	if err != nil {
		return nil, err
	}
	rt := r.(*User)
	return rt, nil
}

func (T *DbContext) LoadUser(Ref string) (*User, error) {
	r, err := T.LoadEntityWrapped(Ref)
	if err != nil {
		return nil, err
	}
	rt := r.(*User)
	return rt, nil
}
