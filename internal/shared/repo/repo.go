package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/GagulProject/go-whisky/internal/shared/errors"
	httperr "github.com/GagulProject/go-whisky/internal/shared/errors/http"
	"github.com/samber/lo"
)

// Repo는 기본적인 CRUD 작업을 위한 인터페이스입니다.
// D: Domain 모델, B: Boiler 모델
type Repo[D DomainModel, B BoilerModel] interface {
	DB() *sql.DB
	Create(context.Context, D) (D, error)
	Find(context.Context, string) (D, error)
	FindBy(context.Context, ...qm.QueryMod) (D, error)
	Fetch(context.Context, string) (D, error)
	FetchBy(context.Context, ...qm.QueryMod) (D, error)
	FindAllBy(context.Context, ...qm.QueryMod) ([]D, error)
	Delete(context.Context, string) error
	DeleteBy(context.Context, ...qm.QueryMod) error
	DeleteAllBy(context.Context, ...qm.QueryMod) error
}

type (
	// BoilerModel은 데이터베이스 작업을 위한 기본 인터페이스입니다.
	BoilerModel interface {
		comparable
		Insert(context.Context, boil.ContextExecutor, boil.Columns) error
		Update(context.Context, boil.ContextExecutor, boil.Columns) (int64, error)
		Upsert(context.Context, boil.ContextExecutor, bool, []string, boil.Columns, boil.Columns) error
		Delete(context.Context, boil.ContextExecutor) (int64, error)
	}

	BoilerSliceModel interface{}
	DomainModel      interface{ comparable }

	// 모델 변환을 위한 함수 타입들
	DomainToBoilerFunc[D DomainModel, B BoilerModel]         func(D) (B, error)
	BoilerToDomainFunc[D DomainModel, B BoilerModel]         func(B) (D, error)
	BoilerSliceConverter[BS BoilerSliceModel, B BoilerModel] func(BS) []B
)

// Query는 데이터베이스 쿼리 작업을 위한 인터페이스입니다.
type Query[B BoilerModel, BS BoilerSliceModel] interface {
	One(context.Context, boil.ContextExecutor) (B, error)
	All(context.Context, boil.ContextExecutor) (BS, error)
	DeleteAll(context.Context, boil.ContextExecutor) (int64, error)
}

type QueryStarter[B BoilerModel, BS BoilerSliceModel] func(...qm.QueryMod) Query[B, BS]

type repo[D DomainModel, B BoilerModel, BS BoilerSliceModel] struct {
	db             *sql.DB
	toDomainFunc   BoilerToDomainFunc[D, B]
	toBoilerFunc   DomainToBoilerFunc[D, B]
	sliceConverter BoilerSliceConverter[BS, B]
	queryStarter   QueryStarter[B, BS]
}

// New는 새로운 Repository 인스턴스를 생성합니다.
func New[D DomainModel, B BoilerModel, BS BoilerSliceModel](
	db *sql.DB,
	toBoilerFunc DomainToBoilerFunc[D, B],
	toDomainFunc BoilerToDomainFunc[D, B],
	sliceConverter BoilerSliceConverter[BS, B],
	starter QueryStarter[B, BS],
) Repo[D, B] {
	return &repo[D, B, BS]{
		db:             db,
		toDomainFunc:   toDomainFunc,
		toBoilerFunc:   toBoilerFunc,
		sliceConverter: sliceConverter,
		queryStarter:   starter,
	}
}

func (r *repo[D, B, BS]) DB() *sql.DB { return r.db }

func (r *repo[D, B, BS]) Create(ctx context.Context, d D) (D, error) {
	b, err := r.toBoilerFunc(d)
	if err != nil {
		return lo.Empty[D](), fmt.Errorf("failed to convert domain to boiler: %w", err)
	}

	if err := b.Insert(ctx, r.DB(), boil.Infer()); err != nil {
		return lo.Empty[D](), fmt.Errorf("failed to insert: %w", err)
	}

	return r.toDomainFunc(b)
}

// handleSQLError는 SQL 관련 에러를 처리합니다.
func (r *repo[D, B, BS]) handleSQLError(err error, operation string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return httperr.NotFound(fmt.Errorf("record not found: %w", err))
	}
	return fmt.Errorf("%s failed: %w", operation, err)
}

// Find는 ID로 레코드를 조회합니다. ID 컬럼명이 'id'인 경우에만 사용 가능합니다.
func (r *repo[D, B, BS]) Find(ctx context.Context, id string) (D, error) {
	b, err := r.queryStarter(qm.Where("\"id\" = ?", id)).One(ctx, r.DB())
	if errors.Is(err, sql.ErrNoRows) {
		return lo.Empty[D](), nil
	}
	if err != nil {
		return lo.Empty[D](), r.handleSQLError(err, "find")
	}
	return r.toDomainFunc(b)
}

func (r *repo[D, B, BS]) fetch(f func() (D, error)) (D, error) {
	d, err := f()
	if err != nil {
		return lo.Empty[D](), err
	}
	if lo.IsEmpty(d) {
		return lo.Empty[D](), httperr.NotFound(
			fmt.Errorf("%T not found", d),
		)
	}
	return d, nil
}

// Fetch는 ID로 레코드를 조회하며, 없을 경우 에러를 반환합니다.
func (r *repo[D, B, BS]) Fetch(ctx context.Context, id string) (D, error) {
	return r.fetch(func() (D, error) { return r.Find(ctx, id) })
}

func (r *repo[D, B, BS]) FindBy(ctx context.Context, mods ...qm.QueryMod) (D, error) {
	b, err := r.queryStarter(mods...).One(ctx, r.DB())
	if errors.Is(err, sql.ErrNoRows) {
		return lo.Empty[D](), nil
	}
	if err != nil {
		return lo.Empty[D](), r.handleSQLError(err, "findBy")
	}
	return r.toDomainFunc(b)
}

func (r *repo[D, B, BS]) FetchBy(ctx context.Context, mods ...qm.QueryMod) (D, error) {
	return r.fetch(func() (D, error) { return r.FindBy(ctx, mods...) })
}

func (r *repo[D, B, BS]) FindAllBy(ctx context.Context, mods ...qm.QueryMod) ([]D, error) {
	bs, err := r.queryStarter(mods...).All(ctx, r.DB())
	if err != nil {
		return nil, r.handleSQLError(err, "findAllBy")
	}

	ds := make([]D, 0)
	for _, b := range r.sliceConverter(bs) {
		d, err := r.toDomainFunc(b)
		if err != nil {
			return nil, fmt.Errorf("failed to convert boiler to domain: %w", err)
		}
		ds = append(ds, d)
	}
	return ds, nil
}

func (r *repo[D, B, BS]) Delete(ctx context.Context, id string) error {
	return r.DeleteBy(ctx, qm.Where("\"id\" = ?", id))
}

func (r *repo[D, B, BS]) DeleteBy(ctx context.Context, mods ...qm.QueryMod) error {
	b, err := r.queryStarter(mods...).One(ctx, r.DB())
	if err != nil {
		return r.handleSQLError(err, "deleteBy")
	}

	if _, err = b.Delete(ctx, r.DB()); err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}

func (r *repo[D, B, BS]) DeleteAllBy(ctx context.Context, mods ...qm.QueryMod) error {
	_, err := r.queryStarter(mods...).DeleteAll(ctx, r.DB())
	if err != nil {
		return r.handleSQLError(err, "deleteAllBy")
	}
	return nil
}
