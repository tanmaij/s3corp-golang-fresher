// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Document is an object representing the database table.
type Document struct {
	DocumentId string      `boil:"documentid" json:"documentid" toml:"documentid" yaml:"documentid"`
	Subject    null.String `boil:"subject" json:"subject,omitempty" toml:"subject" yaml:"subject,omitempty"`
	Createdat  null.Time   `boil:"createdat" json:"createdat,omitempty" toml:"createdat" yaml:"createdat,omitempty"`

	R *documentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L documentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DocumentColumns = struct {
	DocumentId string
	Subject    string
	Createdat  string
}{
	DocumentId: "documentid",
	Subject:    "subject",
	Createdat:  "createdat",
}

var DocumentTableColumns = struct {
	DocumentId string
	Subject    string
	Createdat  string
}{
	DocumentId: "document.documentid",
	Subject:    "document.subject",
	Createdat:  "document.createdat",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DocumentWhere = struct {
	DocumentId whereHelperstring
	Subject    whereHelpernull_String
	Createdat  whereHelpernull_Time
}{
	DocumentId: whereHelperstring{field: "\"main\".\"document\".\"documentid\""},
	Subject:    whereHelpernull_String{field: "\"main\".\"document\".\"subject\""},
	Createdat:  whereHelpernull_Time{field: "\"main\".\"document\".\"createdat\""},
}

// DocumentRels is where relationship names are stored.
var DocumentRels = struct {
	DocumentidSubdocuments string
}{
	DocumentidSubdocuments: "DocumentidSubdocuments",
}

// documentR is where relationships are stored.
type documentR struct {
	DocumentidSubdocuments SubDocumentSlice `boil:"DocumentidSubdocuments" json:"DocumentidSubdocuments" toml:"DocumentidSubdocuments" yaml:"DocumentidSubdocuments"`
}

// NewStruct creates a new relationship struct
func (*documentR) NewStruct() *documentR {
	return &documentR{}
}

func (r *documentR) GetDocumentidSubdocuments() SubDocumentSlice {
	if r == nil {
		return nil
	}
	return r.DocumentidSubdocuments
}

// documentL is where Load methods for each relationship are stored.
type documentL struct{}

var (
	documentAllColumns            = []string{"documentid", "subject", "createdat"}
	documentColumnsWithoutDefault = []string{}
	documentColumnsWithDefault    = []string{"documentid", "subject", "createdat"}
	documentPrimaryKeyColumns     = []string{"documentid"}
	documentGeneratedColumns      = []string{}
)

type (
	// DocumentSlice is an alias for a slice of pointers to Document.
	// This should almost always be used instead of []Document.
	DocumentSlice []*Document
	// DocumentHook is the signature for custom Document hook methods
	DocumentHook func(context.Context, boil.ContextExecutor, *Document) error

	documentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	documentType                 = reflect.TypeOf(&Document{})
	documentMapping              = queries.MakeStructMapping(documentType)
	documentPrimaryKeyMapping, _ = queries.BindMapping(documentType, documentMapping, documentPrimaryKeyColumns)
	documentInsertCacheMut       sync.RWMutex
	documentInsertCache          = make(map[string]insertCache)
	documentUpdateCacheMut       sync.RWMutex
	documentUpdateCache          = make(map[string]updateCache)
	documentUpsertCacheMut       sync.RWMutex
	documentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var documentAfterSelectHooks []DocumentHook

var documentBeforeInsertHooks []DocumentHook
var documentAfterInsertHooks []DocumentHook

var documentBeforeUpdateHooks []DocumentHook
var documentAfterUpdateHooks []DocumentHook

var documentBeforeDeleteHooks []DocumentHook
var documentAfterDeleteHooks []DocumentHook

var documentBeforeUpsertHooks []DocumentHook
var documentAfterUpsertHooks []DocumentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Document) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Document) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Document) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Document) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Document) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Document) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Document) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Document) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Document) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range documentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDocumentHook registers your hook function for all future operations.
func AddDocumentHook(hookPoint boil.HookPoint, documentHook DocumentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		documentAfterSelectHooks = append(documentAfterSelectHooks, documentHook)
	case boil.BeforeInsertHook:
		documentBeforeInsertHooks = append(documentBeforeInsertHooks, documentHook)
	case boil.AfterInsertHook:
		documentAfterInsertHooks = append(documentAfterInsertHooks, documentHook)
	case boil.BeforeUpdateHook:
		documentBeforeUpdateHooks = append(documentBeforeUpdateHooks, documentHook)
	case boil.AfterUpdateHook:
		documentAfterUpdateHooks = append(documentAfterUpdateHooks, documentHook)
	case boil.BeforeDeleteHook:
		documentBeforeDeleteHooks = append(documentBeforeDeleteHooks, documentHook)
	case boil.AfterDeleteHook:
		documentAfterDeleteHooks = append(documentAfterDeleteHooks, documentHook)
	case boil.BeforeUpsertHook:
		documentBeforeUpsertHooks = append(documentBeforeUpsertHooks, documentHook)
	case boil.AfterUpsertHook:
		documentAfterUpsertHooks = append(documentAfterUpsertHooks, documentHook)
	}
}

// One returns a single document record from the query.
func (q documentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Document, error) {
	o := &Document{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for document")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Document records from the query.
func (q documentQuery) All(ctx context.Context, exec boil.ContextExecutor) (DocumentSlice, error) {
	var o []*Document

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Document slice")
	}

	if len(documentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Document records in the query.
func (q documentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count document rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q documentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if document exists")
	}

	return count > 0, nil
}

// DocumentidSubdocuments retrieves all the subdocument's SubDocuments with an executor via documentid column.
func (o *Document) DocumentidSubdocuments(mods ...qm.QueryMod) subDocumentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"main\".\"subdocument\".\"documentid\"=?", o.DocumentId),
	)

	return SubDocuments(queryMods...)
}

// LoadDocumentidSubdocuments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (documentL) LoadDocumentidSubdocuments(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDocument interface{}, mods queries.Applicator) error {
	var slice []*Document
	var object *Document

	if singular {
		object = maybeDocument.(*Document)
	} else {
		slice = *maybeDocument.(*[]*Document)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &documentR{}
		}
		args = append(args, object.DocumentId)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &documentR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DocumentId) {
					continue Outer
				}
			}

			args = append(args, obj.DocumentId)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`main.subdocument`),
		qm.WhereIn(`main.subdocument.documentid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subdocument")
	}

	var resultSlice []*SubDocument
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice subdocument")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subdocument")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subdocument")
	}

	if len(subDocumentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DocumentidSubdocuments = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subDocumentR{}
			}
			foreign.R.DocumentidDocument = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.DocumentId, foreign.Documentid) {
				local.R.DocumentidSubdocuments = append(local.R.DocumentidSubdocuments, foreign)
				if foreign.R == nil {
					foreign.R = &subDocumentR{}
				}
				foreign.R.DocumentidDocument = local
				break
			}
		}
	}

	return nil
}

// AddDocumentidSubdocuments adds the given related objects to the existing relationships
// of the document, optionally inserting them as new records.
// Appends related to o.R.DocumentidSubdocuments.
// Sets related.R.DocumentidDocument appropriately.
func (o *Document) AddDocumentidSubdocuments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SubDocument) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.Documentid, o.DocumentId)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"main\".\"subdocument\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"documentid"}),
				strmangle.WhereClause("\"", "\"", 2, subDocumentPrimaryKeyColumns),
			)
			values := []interface{}{o.DocumentId, rel.subDocumentId}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.Documentid, o.DocumentId)
		}
	}

	if o.R == nil {
		o.R = &documentR{
			DocumentidSubdocuments: related,
		}
	} else {
		o.R.DocumentidSubdocuments = append(o.R.DocumentidSubdocuments, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subDocumentR{
				DocumentidDocument: o,
			}
		} else {
			rel.R.DocumentidDocument = o
		}
	}
	return nil
}

// SetDocumentidSubdocuments removes all previously related items of the
// document replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.DocumentidDocument's DocumentidSubdocuments accordingly.
// Replaces o.R.DocumentidSubdocuments with related.
// Sets related.R.DocumentidDocument's DocumentidSubdocuments accordingly.
func (o *Document) SetDocumentidSubdocuments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SubDocument) error {
	query := "update \"main\".\"subdocument\" set \"documentid\" = null where \"documentid\" = $1"
	values := []interface{}{o.DocumentId}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.DocumentidSubdocuments {
			queries.SetScanner(&rel.Documentid, nil)
			if rel.R == nil {
				continue
			}

			rel.R.DocumentidDocument = nil
		}
		o.R.DocumentidSubdocuments = nil
	}

	return o.AddDocumentidSubdocuments(ctx, exec, insert, related...)
}

// RemoveDocumentidSubdocuments relationships from objects passed in.
// Removes related items from R.DocumentidSubdocuments (uses pointer comparison, removal does not keep order)
// Sets related.R.DocumentidDocument.
func (o *Document) RemoveDocumentidSubdocuments(ctx context.Context, exec boil.ContextExecutor, related ...*SubDocument) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.Documentid, nil)
		if rel.R != nil {
			rel.R.DocumentidDocument = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("documentid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.DocumentidSubdocuments {
			if rel != ri {
				continue
			}

			ln := len(o.R.DocumentidSubdocuments)
			if ln > 1 && i < ln-1 {
				o.R.DocumentidSubdocuments[i] = o.R.DocumentidSubdocuments[ln-1]
			}
			o.R.DocumentidSubdocuments = o.R.DocumentidSubdocuments[:ln-1]
			break
		}
	}

	return nil
}

// Documents retrieves all the records using an executor.
func Documents(mods ...qm.QueryMod) documentQuery {
	mods = append(mods, qm.From("\"main\".\"document\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"main\".\"document\".*"})
	}

	return documentQuery{q}
}

// FindDocument retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDocument(ctx context.Context, exec boil.ContextExecutor, documentId string, selectCols ...string) (*Document, error) {
	documentObj := &Document{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"main\".\"document\" where \"documentid\"=$1", sel,
	)

	q := queries.Raw(query, documentId)

	err := q.Bind(ctx, exec, documentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from document")
	}

	if err = documentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return documentObj, err
	}

	return documentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Document) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no document provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	documentInsertCacheMut.RLock()
	cache, cached := documentInsertCache[key]
	documentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			documentAllColumns,
			documentColumnsWithDefault,
			documentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(documentType, documentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(documentType, documentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"main\".\"document\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"main\".\"document\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into document")
	}

	if !cached {
		documentInsertCacheMut.Lock()
		documentInsertCache[key] = cache
		documentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Document.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Document) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	documentUpdateCacheMut.RLock()
	cache, cached := documentUpdateCache[key]
	documentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			documentAllColumns,
			documentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update document, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"main\".\"document\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, documentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(documentType, documentMapping, append(wl, documentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update document row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for document")
	}

	if !cached {
		documentUpdateCacheMut.Lock()
		documentUpdateCache[key] = cache
		documentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q documentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for document")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for document")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DocumentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"main\".\"document\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, documentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in document slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all document")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Document) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no document provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(documentColumnsWithDefault, o)

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

	documentUpsertCacheMut.RLock()
	cache, cached := documentUpsertCache[key]
	documentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			documentAllColumns,
			documentColumnsWithDefault,
			documentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			documentAllColumns,
			documentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert document, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(documentPrimaryKeyColumns))
			copy(conflict, documentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"main\".\"document\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(documentType, documentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(documentType, documentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert document")
	}

	if !cached {
		documentUpsertCacheMut.Lock()
		documentUpsertCache[key] = cache
		documentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Document record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Document) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Document provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), documentPrimaryKeyMapping)
	sql := "DELETE FROM \"main\".\"document\" WHERE \"documentid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from document")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for document")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q documentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no documentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from document")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for document")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DocumentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(documentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"main\".\"document\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, documentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from document slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for document")
	}

	if len(documentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Document) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDocument(ctx, exec, o.DocumentId)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DocumentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DocumentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), documentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"main\".\"document\".* FROM \"main\".\"document\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, documentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DocumentSlice")
	}

	*o = slice

	return nil
}

// DocumentExists checks if the Document row exists.
func DocumentExists(ctx context.Context, exec boil.ContextExecutor, documentId string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"main\".\"document\" where \"documentid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, documentId)
	}
	row := exec.QueryRowContext(ctx, sql, documentId)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if document exists")
	}

	return exists, nil
}