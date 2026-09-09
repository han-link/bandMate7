package service

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"bandMate7/internal/utils"
	"context"
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/google/uuid"
)

var ErrPerformancesNotFound = errors.New("performances not found")

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
)

type ValidationError struct {
	Message string
	Fields  map[string]string
	Notes   []string
}

func (e *ValidationError) Unwrap() error { return ErrInvalidInput }

func (e *ValidationError) add(field string, value string) {
	if e.Fields == nil {
		e.Fields = map[string]string{}
	}
	e.Fields[field] = value
}

func (e *ValidationError) note(msg string) {
	e.Notes = append(e.Notes, msg)
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Fields) > 0 || len(e.Notes) > 0
}

func (e *ValidationError) Error() string {
	if !e.HasErrors() {
		return e.Message
	}

	parts := append([]string(nil), e.Notes...)

	keys := make([]string, 0, len(e.Fields))
	for k := range e.Fields {
		keys = append(keys, k)
	}
	sort.Strings(keys) // stable output for logs/tests

	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %s", k, e.Fields[k]))
	}
	return fmt.Sprintf("%s (%s)", e.Message, strings.Join(parts, "; "))
}

type SetlistsService struct {
	store *store.Storage
}

type CreateSetlistRequest struct {
	Title          string
	PerformanceIds []uuid.UUID
}

func (s *SetlistsService) Create(ctx context.Context, payload CreateSetlistRequest) (*model.Setlist, error) {
	missingIds, err := s.store.Performances.CheckPerformancesExist(ctx, payload.PerformanceIds)
	if err != nil {
		return nil, err
	}

	ve := &ValidationError{Message: "invalid setlist order"}
	if len(missingIds) > 0 {
		ve.note(fmt.Sprintf("Missing Ids: %s", missingIds))
		return nil, ve
	}
	setlist := model.Setlist{
		Titel: payload.Title,
	}
	if err := s.store.SetLists.Create(ctx, &setlist, payload.PerformanceIds); err != nil {
		return nil, err
	}
	return &setlist, nil
}

type SelistWithoutPerformance struct {
	model.BaseModel
	Title string
}

func (s *SetlistsService) GetAll(ctx context.Context) ([]SelistWithoutPerformance, error) {
	setlists, err := s.store.SetLists.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	setlistsNoPerf := make([]SelistWithoutPerformance, len(setlists))
	for i, perf := range setlists {
		setlistsNoPerf[i] = SelistWithoutPerformance{
			BaseModel: perf.BaseModel,
			Title:     perf.Titel,
		}
	}

	return setlistsNoPerf, nil
}

type ChangeOrderItem struct {
	ID       uuid.UUID `json:"id"`
	Position int       `json:"position"`
}

type ChangeOrderPayload struct {
	Items []ChangeOrderItem `json:"items"`
}

func (s *SetlistsService) ChangeOrder(ctx context.Context, payload ChangeOrderPayload, setlist *model.Setlist) (*model.Setlist, error) {
	performanceIds := make([]uuid.UUID, len(payload.Items))
	positions := make([]int, len(payload.Items))
	for i, pid := range payload.Items {
		performanceIds[i] = pid.ID
		positions[i] = pid.Position
	}
	ve := &ValidationError{Message: "invalid setlist order"}
	// Check for duplicates
	if utils.ContainsDuplicates(performanceIds) {
		ve.note("contains duplicate performance ids")
	}
	missingPerfIds, notIncludedPerfIds, err := s.store.SetLists.CheckPerformancesExist(ctx, setlist, performanceIds)
	if err != nil {
		return nil, err
	}

	if len(missingPerfIds) > 0 {
		ve.add("missingIds", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(missingPerfIds)), ","), "[]"))
	}

	if len(notIncludedPerfIds) > 0 {
		ve.add("notIncludedIds", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(notIncludedPerfIds)), ","), "[]"))
	}

	for _, p := range positions {
		if p < 1 {
			ve.add("negativePosition", "Negative positions or 0 are not allowed")
		}
		if p > len(payload.Items)+1 {
			ve.add("outOfRangePosition", "Position is out of range, the setlist does not have that many performances")
		}
	}

	if len(slices.Compact(positions)) != len(positions) {
		ve.add("duplicatedPositions", "Duplicated positions found")
	}

	if ve.HasErrors() {
		return nil, ve
	}

	newOrderMap := make(map[uuid.UUID]int, len(payload.Items))

	for _, p := range payload.Items {
		newOrderMap[p.ID] = p.Position
	}

	if err := s.store.SetLists.UpdateOrder(ctx, setlist, newOrderMap); err != nil {
		return nil, err
	}

	return setlist, nil
}
