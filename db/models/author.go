// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Author is an object representing the database table.
type Author struct {
	ID    int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email string `boil:"email" json:"email" toml:"email" yaml:"email"`
	Name  string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *authorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthorColumns = struct {
	ID    string
	Email string
	Name  string
}{
	ID:    "id",
	Email: "email",
	Name:  "name",
}

var AuthorTableColumns = struct {
	ID    string
	Email string
	Name  string
}{
	ID:    "author.id",
	Email: "author.email",
	Name:  "author.name",
}

// Generated where

var AuthorWhere = struct {
	ID    whereHelperint
	Email whereHelperstring
	Name  whereHelperstring
}{
	ID:    whereHelperint{field: "\"author\".\"id\""},
	Email: whereHelperstring{field: "\"author\".\"email\""},
	Name:  whereHelperstring{field: "\"author\".\"name\""},
}

// AuthorRels is where relationship names are stored.
var AuthorRels = struct {
	Articles string
}{
	Articles: "Articles",
}

// authorR is where relationships are stored.
type authorR struct {
	Articles ArticleSlice `boil:"Articles" json:"Articles" toml:"Articles" yaml:"Articles"`
}

// NewStruct creates a new relationship struct
func (*authorR) NewStruct() *authorR {
	return &authorR{}
}

func (r *authorR) GetArticles() ArticleSlice {
	if r == nil {
		return nil
	}
	return r.Articles
}

// authorL is where Load methods for each relationship are stored.
type authorL struct{}

var (
	authorAllColumns            = []string{"id", "email", "name"}
	authorColumnsWithoutDefault = []string{"email", "name"}
	authorColumnsWithDefault    = []string{"id"}
	authorPrimaryKeyColumns     = []string{"id"}
	authorGeneratedColumns      = []string{}
)

type (
	// AuthorSlice is an alias for a slice of pointers to Author.
	// This should almost always be used instead of []Author.
	AuthorSlice []*Author
	// AuthorHook is the signature for custom Author hook methods
	AuthorHook func(context.Context, boil.ContextExecutor, *Author) error

	authorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authorType                 = reflect.TypeOf(&Author{})
	authorMapping              = queries.MakeStructMapping(authorType)
	authorPrimaryKeyMapping, _ = queries.BindMapping(authorType, authorMapping, authorPrimaryKeyColumns)
	authorInsertCacheMut       sync.RWMutex
	authorInsertCache          = make(map[string]insertCache)
	authorUpdateCacheMut       sync.RWMutex
	authorUpdateCache          = make(map[string]updateCache)
	authorUpsertCacheMut       sync.RWMutex
	authorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authorAfterSelectHooks []AuthorHook

var authorBeforeInsertHooks []AuthorHook
var authorAfterInsertHooks []AuthorHook

var authorBeforeUpdateHooks []AuthorHook
var authorAfterUpdateHooks []AuthorHook

var authorBeforeDeleteHooks []AuthorHook
var authorAfterDeleteHooks []AuthorHook

var authorBeforeUpsertHooks []AuthorHook
var authorAfterUpsertHooks []AuthorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Author) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Author) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Author) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Author) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Author) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Author) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Author) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Author) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Author) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthorHook registers your hook function for all future operations.
func AddAuthorHook(hookPoint boil.HookPoint, authorHook AuthorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		authorAfterSelectHooks = append(authorAfterSelectHooks, authorHook)
	case boil.BeforeInsertHook:
		authorBeforeInsertHooks = append(authorBeforeInsertHooks, authorHook)
	case boil.AfterInsertHook:
		authorAfterInsertHooks = append(authorAfterInsertHooks, authorHook)
	case boil.BeforeUpdateHook:
		authorBeforeUpdateHooks = append(authorBeforeUpdateHooks, authorHook)
	case boil.AfterUpdateHook:
		authorAfterUpdateHooks = append(authorAfterUpdateHooks, authorHook)
	case boil.BeforeDeleteHook:
		authorBeforeDeleteHooks = append(authorBeforeDeleteHooks, authorHook)
	case boil.AfterDeleteHook:
		authorAfterDeleteHooks = append(authorAfterDeleteHooks, authorHook)
	case boil.BeforeUpsertHook:
		authorBeforeUpsertHooks = append(authorBeforeUpsertHooks, authorHook)
	case boil.AfterUpsertHook:
		authorAfterUpsertHooks = append(authorAfterUpsertHooks, authorHook)
	}
}

// OneG returns a single author record from the query using the global executor.
func (q authorQuery) OneG(ctx context.Context) (*Author, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single author record from the query.
func (q authorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Author, error) {
	o := &Author{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for author")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Author records from the query using the global executor.
func (q authorQuery) AllG(ctx context.Context) (AuthorSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Author records from the query.
func (q authorQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthorSlice, error) {
	var o []*Author

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Author slice")
	}

	if len(authorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Author records in the query using the global executor
func (q authorQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Author records in the query.
func (q authorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count author rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q authorQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q authorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if author exists")
	}

	return count > 0, nil
}

// Articles retrieves all the article's Articles with an executor.
func (o *Author) Articles(mods ...qm.QueryMod) articleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"article\".\"author_id\"=?", o.ID),
	)

	return Articles(queryMods...)
}

// LoadArticles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (authorL) LoadArticles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthor interface{}, mods queries.Applicator) error {
	var slice []*Author
	var object *Author

	if singular {
		var ok bool
		object, ok = maybeAuthor.(*Author)
		if !ok {
			object = new(Author)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthor))
			}
		}
	} else {
		s, ok := maybeAuthor.(*[]*Author)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthor))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authorR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authorR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`article`),
		qm.WhereIn(`article.author_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load article")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice article")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on article")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for article")
	}

	if len(articleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Articles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &articleR{}
			}
			foreign.R.Author = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.AuthorID {
				local.R.Articles = append(local.R.Articles, foreign)
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.Author = local
				break
			}
		}
	}

	return nil
}

// AddArticlesG adds the given related objects to the existing relationships
// of the author, optionally inserting them as new records.
// Appends related to o.R.Articles.
// Sets related.R.Author appropriately.
// Uses the global database handle.
func (o *Author) AddArticlesG(ctx context.Context, insert bool, related ...*Article) error {
	return o.AddArticles(ctx, boil.GetContextDB(), insert, related...)
}

// AddArticles adds the given related objects to the existing relationships
// of the author, optionally inserting them as new records.
// Appends related to o.R.Articles.
// Sets related.R.Author appropriately.
func (o *Author) AddArticles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Article) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AuthorID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"article\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"author_id"}),
				strmangle.WhereClause("\"", "\"", 2, articlePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.AuthorID = o.ID
		}
	}

	if o.R == nil {
		o.R = &authorR{
			Articles: related,
		}
	} else {
		o.R.Articles = append(o.R.Articles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &articleR{
				Author: o,
			}
		} else {
			rel.R.Author = o
		}
	}
	return nil
}

// Authors retrieves all the records using an executor.
func Authors(mods ...qm.QueryMod) authorQuery {
	mods = append(mods, qm.From("\"author\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"author\".*"})
	}

	return authorQuery{q}
}

// FindAuthorG retrieves a single record by ID.
func FindAuthorG(ctx context.Context, iD int, selectCols ...string) (*Author, error) {
	return FindAuthor(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAuthor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthor(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Author, error) {
	authorObj := &Author{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"author\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from author")
	}

	if err = authorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authorObj, err
	}

	return authorObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Author) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Author) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no author provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authorInsertCacheMut.RLock()
	cache, cached := authorInsertCache[key]
	authorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authorAllColumns,
			authorColumnsWithDefault,
			authorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authorType, authorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authorType, authorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"author\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"author\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into author")
	}

	if !cached {
		authorInsertCacheMut.Lock()
		authorInsertCache[key] = cache
		authorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Author record using the global executor.
// See Update for more documentation.
func (o *Author) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Author.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Author) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authorUpdateCacheMut.RLock()
	cache, cached := authorUpdateCache[key]
	authorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authorAllColumns,
			authorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update author, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"author\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authorType, authorMapping, append(wl, authorPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update author row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for author")
	}

	if !cached {
		authorUpdateCacheMut.Lock()
		authorUpdateCache[key] = cache
		authorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q authorQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q authorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for author")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for author")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthorSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"author\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in author slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all author")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Author) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Author) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no author provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authorColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	authorUpsertCacheMut.RLock()
	cache, cached := authorUpsertCache[key]
	authorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authorAllColumns,
			authorColumnsWithDefault,
			authorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			authorAllColumns,
			authorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert author, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authorPrimaryKeyColumns))
			copy(conflict, authorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"author\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(authorType, authorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authorType, authorMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert author")
	}

	if !cached {
		authorUpsertCacheMut.Lock()
		authorUpsertCache[key] = cache
		authorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Author record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Author) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Author record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Author) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Author provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authorPrimaryKeyMapping)
	sql := "DELETE FROM \"author\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from author")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for author")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q authorQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q authorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no authorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from author")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for author")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AuthorSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"author\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from author slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for author")
	}

	if len(authorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Author) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: no Author provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Author) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthor(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthorSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: empty AuthorSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"author\".* FROM \"author\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in AuthorSlice")
	}

	*o = slice

	return nil
}

// AuthorExistsG checks if the Author row exists.
func AuthorExistsG(ctx context.Context, iD int) (bool, error) {
	return AuthorExists(ctx, boil.GetContextDB(), iD)
}

// AuthorExists checks if the Author row exists.
func AuthorExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"author\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if author exists")
	}

	return exists, nil
}
