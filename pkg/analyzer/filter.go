package analyzer

import (
	"github.com/pirsch-analytics/pirsch/v6/pkg"
	"github.com/pirsch-analytics/pirsch/v6/pkg/util"
	"strings"
	"time"
)

// Filter are all fields that can be used to filter the result sets.
// Fields can be inverted by adding a "!" in front of the string.
// To compare to none/unknown/empty, set the value to "null" (case-insensitive).
type Filter struct {
	// ClientID is the optional.
	ClientID int64

	// Timezone sets the timezone used to interpret dates and times.
	// It will be set to UTC by default.
	Timezone *time.Location

	// From is the start date of the selected period.
	From time.Time

	// To is the end date of the selected period.
	To time.Time

	// Period sets the period to group results.
	// This is only used by Analyzer.ByPeriod, Analyzer.AvgSessionDuration, and Analyzer.AvgTimeOnPage.
	// Using it for other queries leads to wrong results and might return an error.
	// This can either be PeriodDay (default), PeriodWeek, or PeriodYear.
	Period pkg.Period

	// Path filters for the path.
	// Note that if this and PathPattern are both set, Path will be preferred.
	Path []string

	// AnyPath filters for any path in the list.
	AnyPath []string

	// EntryPath filters for the entry page.
	EntryPath []string

	// ExitPath filters for the exit page.
	ExitPath []string

	// PathPattern filters for the path using a (ClickHouse supported) regex pattern.
	// Note that if this and Path are both set, Path will be preferred.
	// Examples for useful patterns (all case-insensitive, * is used for every character but slashes, ** is used for all characters including slashes):
	//  (?i)^/path/[^/]+$ // matches /path/*
	//  (?i)^/path/[^/]+/.* // matches /path/*/**
	//  (?i)^/path/[^/]+/slashes$ // matches /path/*/slashes
	//  (?i)^/path/.+/slashes$ // matches /path/**/slashes
	PathPattern []string

	// Language filters for the ISO language code.
	Language []string

	// Country filters for the ISO country code.
	Country []string

	// City filters for the city name.
	City []string

	// Referrer filters for the full referrer.
	Referrer []string

	// ReferrerName filters for the referrer name.
	ReferrerName []string

	// OS filters for the operating system.
	OS []string

	// OSVersion filters for the operating system version.
	OSVersion []string

	// Browser filters for the browser.
	Browser []string

	// BrowserVersion filters for the browser version.
	BrowserVersion []string

	// Platform filters for the platform (desktop, mobile, unknown).
	Platform string

	// ScreenClass filters for the screen class.
	ScreenClass []string

	// UTMSource filters for the utm_source query parameter.
	UTMSource []string

	// UTMMedium filters for the utm_medium query parameter.
	UTMMedium []string

	// UTMCampaign filters for the utm_campaign query parameter.
	UTMCampaign []string

	// UTMContent filters for the utm_content query parameter.
	UTMContent []string

	// UTMTerm filters for the utm_term query parameter.
	UTMTerm []string

	// EventName filters for an event by its name.
	EventName []string

	// EventMetaKey filters for an event meta key.
	// This must be used together with an EventName.
	EventMetaKey []string

	// EventMeta filters for event metadata.
	EventMeta map[string]string

	// Search searches the results for given fields and inputs.
	Search []Search

	// Sort sorts the results.
	// This will overwrite the default order provided by the Analyzer.
	Sort []Sort

	// Offset limits the number of results. Less or equal to zero means no offset.
	Offset int

	// Limit limits the number of results. Less or equal to zero means no limit.
	Limit int

	// CustomMetricKey is used to calculate the average and total for an event metadata field.
	// This must be used together with EventName and CustomMetricType.
	CustomMetricKey string

	// CustomMetricType is used to calculate the average and total for an event metadata field.
	CustomMetricType pkg.CustomMetricType

	// IncludeTime sets whether the selected period should contain the time (hour, minute, second).
	IncludeTime bool

	// IncludeTitle indicates that the Analyzer.ByPath, Analyzer.Entry, and Analyzer.Exit should contain the page title.
	IncludeTitle bool

	// IncludeTimeOnPage indicates that the Analyzer.ByPath and Analyzer.Entry should contain the average time on page.
	IncludeTimeOnPage bool

	// IncludeCR indicates that Analyzer.Total and Analyzer.ByPeriod should contain the conversion rate.
	IncludeCR bool

	// MaxTimeOnPageSeconds is an optional maximum for the time spent on page.
	// Visitors who are idle artificially increase the average time spent on a page, this option can be used to limit the effect.
	// Set to 0 to disable this option (default).
	MaxTimeOnPageSeconds int
}

// Search filters results by searching for the given input for given field.
// The field needs to contain the search string and is performed case-insensitively.
type Search struct {
	Field Field
	Input string
}

// Sort sorts results by a field and direction.
type Sort struct {
	Field     Field
	Direction pkg.Direction
}

// NewFilter creates a new filter for given client ID.
func NewFilter(clientID int64) *Filter {
	return &Filter{
		ClientID: clientID,
	}
}

func (filter *Filter) validate() {
	if filter.Timezone == nil {
		filter.Timezone = time.UTC
	}

	if !filter.From.IsZero() {
		if filter.IncludeTime {
			filter.From = filter.From.In(time.UTC)
		} else {
			filter.From = filter.toDate(filter.From)
		}
	}

	if !filter.To.IsZero() {
		if filter.IncludeTime {
			filter.To = filter.To.In(time.UTC)
		} else {
			filter.To = filter.toDate(filter.To)
		}
	}

	if !filter.To.IsZero() && filter.From.After(filter.To) {
		filter.From, filter.To = filter.To, filter.From
	}

	// use tomorrow instead of limiting to "today", so that all timezones are included
	tomorrow := util.Today().Add(time.Hour * 24)

	if !filter.To.IsZero() && filter.To.After(tomorrow) {
		filter.To = tomorrow
	}

	if len(filter.Path) != 0 && len(filter.PathPattern) != 0 {
		filter.PathPattern = nil
	}

	for i := 0; i < len(filter.Search); i++ {
		filter.Search[i].Input = strings.TrimSpace(filter.Search[i].Input)
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	}

	if filter.Limit < 0 {
		filter.Limit = 0
	}

	if filter.CustomMetricType != "" &&
		filter.CustomMetricType != pkg.CustomMetricTypeInteger &&
		filter.CustomMetricType != pkg.CustomMetricTypeFloat {
		filter.CustomMetricType = ""
	}

	filter.Path = filter.removeDuplicates(filter.Path)
	filter.EntryPath = filter.removeDuplicates(filter.EntryPath)
	filter.ExitPath = filter.removeDuplicates(filter.ExitPath)
	filter.PathPattern = filter.removeDuplicates(filter.PathPattern)
	filter.Language = filter.removeDuplicates(filter.Language)
	filter.Country = filter.removeDuplicates(filter.Country)
	filter.City = filter.removeDuplicates(filter.City)
	filter.Referrer = filter.removeDuplicates(filter.Referrer)
	filter.ReferrerName = filter.removeDuplicates(filter.ReferrerName)
	filter.OS = filter.removeDuplicates(filter.OS)
	filter.OSVersion = filter.removeDuplicates(filter.OSVersion)
	filter.Browser = filter.removeDuplicates(filter.Browser)
	filter.BrowserVersion = filter.removeDuplicates(filter.BrowserVersion)
	filter.ScreenClass = filter.removeDuplicates(filter.ScreenClass)
	filter.UTMSource = filter.removeDuplicates(filter.UTMSource)
	filter.UTMMedium = filter.removeDuplicates(filter.UTMMedium)
	filter.UTMCampaign = filter.removeDuplicates(filter.UTMCampaign)
	filter.UTMContent = filter.removeDuplicates(filter.UTMContent)
	filter.UTMTerm = filter.removeDuplicates(filter.UTMTerm)
	filter.EventName = filter.removeDuplicates(filter.EventName)
	filter.EventMetaKey = filter.removeDuplicates(filter.EventMetaKey)
}

func (filter *Filter) removeDuplicates(in []string) []string {
	if len(in) == 0 {
		return nil
	}

	keys := make(map[string]struct{})
	list := make([]string, 0, len(in))

	for _, item := range in {
		if _, value := keys[item]; !value {
			keys[item] = struct{}{}
			list = append(list, item)
		}
	}

	return list
}

func (filter *Filter) toDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

func (filter *Filter) buildQuery(fields, groupBy, orderBy []Field) (string, []any) {
	q := queryBuilder{
		filter:  filter,
		from:    filter.table(fields),
		search:  filter.Search,
		groupBy: groupBy,
		orderBy: orderBy,
		offset:  filter.Offset,
		limit:   filter.Limit,
	}
	returnEventName := filter.fieldsContain(fields, FieldEventName)
	customMetric := filter.CustomMetricKey != "" || filter.CustomMetricType != ""

	if q.from == events && !returnEventName && !customMetric {
		q.from = sessions
		q.fields = filter.excludeFields(fields, FieldPath)
		q.includeEventFilter = true
		q.leftJoin = filter.leftJoinEvents(fields)
	} else if q.from == pageViews || returnEventName || customMetric {
		q.fields = fields
		q.join = filter.joinSessions(fields)

		if q.join != nil {
			q.join.parent = &q
		}

		if q.from != events {
			if filter.valuesContainPrefix(filter.EventName, "!") {
				q.includeEventFilter = true
				q.leftJoin = filter.leftJoinEvents(fields)
			} else {
				q.joinSecond = filter.joinEvents(fields)
			}
		}
	} else {
		q.fields = fields
		q.join = filter.joinPageViews(fields)

		if filter.valuesContainPrefix(filter.EventName, "!") {
			q.includeEventFilter = true
			q.leftJoin = filter.leftJoinEvents(fields)
		} else {
			q.joinSecond = filter.joinEvents(fields)
		}
	}

	q.leftJoinSecond = filter.lefJoinUniqueVisitorsByPeriod(fields)
	return q.query()
}

func (filter *Filter) buildTimeQuery() (string, []any) {
	q := queryBuilder{filter: filter}
	return q.whereTime(), q.args
}

func (filter *Filter) table(fields []Field) table {
	if !filter.fieldsContain(fields, FieldEntryPath) && !filter.fieldsContain(fields, FieldExitPath) {
		eventFilter := filter.fieldsContain(fields, FieldEventName) || filter.CustomMetricType != "" && filter.CustomMetricKey != ""

		if !eventFilter &&
			(len(filter.Path) != 0 ||
				len(filter.PathPattern) != 0 ||
				filter.fieldsContain(fields, FieldPath) ||
				filter.searchContains(FieldPath)) {
			return pageViews
		}

		if len(filter.EventName) != 0 || eventFilter {
			return events
		}
	} else if filter.fieldsContain(fields, FieldEntries) || filter.fieldsContain(fields, FieldExits) {
		return pageViews
	}

	return sessions
}

func (filter *Filter) joinSessions(fields []Field) *queryBuilder {
	if len(filter.EntryPath) != 0 ||
		len(filter.ExitPath) != 0 ||
		filter.fieldsContain(fields, FieldBounces) ||
		filter.fieldsContain(fields, FieldViews) ||
		filter.fieldsContain(fields, FieldEntryPath) ||
		filter.fieldsContain(fields, FieldExitPath) {
		sessionFields := []Field{FieldVisitorID, FieldSessionID}
		groupBy := []Field{FieldVisitorID, FieldSessionID}

		if len(filter.EntryPath) != 0 || filter.fieldsContain(fields, FieldEntryPath) || filter.searchContains(FieldEntryPath) {
			sessionFields = append(sessionFields, FieldEntryPath)
			groupBy = append(groupBy, FieldEntryPath)

			if filter.IncludeTitle {
				sessionFields = append(sessionFields, FieldEntryTitle)
				groupBy = append(groupBy, FieldEntryTitle)
			}
		}

		if len(filter.ExitPath) != 0 || filter.fieldsContain(fields, FieldExitPath) || filter.searchContains(FieldExitPath) {
			sessionFields = append(sessionFields, FieldExitPath)
			groupBy = append(groupBy, FieldExitPath)

			if filter.IncludeTitle {
				sessionFields = append(sessionFields, FieldExitTitle)
				groupBy = append(groupBy, FieldExitTitle)
			}
		}

		if filter.fieldsContain(fields, FieldBounces) {
			sessionFields = append(sessionFields, FieldBounces)
		}

		if filter.fieldsContain(fields, FieldViews) {
			sessionFields = append(sessionFields, FieldViews)
		}

		filterCopy := *filter
		filterCopy.Sort = nil
		return &queryBuilder{
			filter:  &filterCopy,
			fields:  sessionFields,
			from:    sessions,
			groupBy: groupBy,
		}
	}

	return nil
}

func (filter *Filter) joinPageViews(fields []Field) *queryBuilder {
	if len(filter.Path) != 0 || len(filter.PathPattern) != 0 || filter.searchContains(FieldPath) {
		pageViewFields := []Field{FieldVisitorID, FieldSessionID}

		if len(filter.PathPattern) != 0 {
			pageViewFields = append(pageViewFields, FieldPath)
		} else if len(filter.Path) != 0 || filter.fieldsContain(fields, FieldPath) || filter.searchContains(FieldPath) {
			pageViewFields = append(pageViewFields, FieldPath)
		}

		filterCopy := *filter
		filterCopy.Sort = nil
		return &queryBuilder{
			filter:  &filterCopy,
			fields:  pageViewFields,
			from:    pageViews,
			groupBy: pageViewFields,
		}
	}

	return nil
}

func (filter *Filter) joinEvents(fields []Field) *queryBuilder {
	if len(filter.EventName) != 0 {
		filterCopy := *filter
		filterCopy.Path = nil
		filterCopy.AnyPath = nil
		filterCopy.Sort = nil
		eventFields := []Field{FieldVisitorID, FieldSessionID}

		if filter.fieldsContain(fields, FieldEventPath) {
			eventFields = append(eventFields, FieldEventPath)
		}

		if filter.fieldsContain(fields, FieldEventTitle) {
			eventFields = append(eventFields, FieldEventTitle)
		}

		if filter.CustomMetricType != "" && filter.CustomMetricKey != "" {
			eventFields = append(eventFields, FieldEventMetaKeysRaw)
			eventFields = append(eventFields, FieldEventMetaValuesRaw)
		}

		return &queryBuilder{
			filter:  &filterCopy,
			fields:  eventFields,
			from:    events,
			groupBy: eventFields,
		}
	}

	return nil
}

func (filter *Filter) leftJoinEvents(fields []Field) *queryBuilder {
	filterCopy := *filter
	filterCopy.EventName = nil
	filterCopy.EventMetaKey = nil
	filterCopy.EventMeta = nil
	filterCopy.Sort = nil
	eventFields := []Field{FieldVisitorID, FieldSessionID, FieldEventName}

	if len(filter.EventMeta) != 0 || filter.fieldsContain(fields, FieldEventMeta) {
		eventFields = append(eventFields, FieldEventMetaKeysRaw, FieldEventMetaValuesRaw)
	} else if len(filter.EventMetaKey) != 0 || filter.fieldsContain(fields, FieldEventMetaKeys) {
		eventFields = append(eventFields, FieldEventMetaKeysRaw)
	}

	if filter.fieldsContain(fields, FieldEventPath) {
		eventFields = append(eventFields, FieldEventPath)
	}

	if filter.fieldsContain(fields, FieldEventTitle) {
		eventFields = append(eventFields, FieldEventTitle)
	}

	return &queryBuilder{
		filter:  &filterCopy,
		fields:  eventFields,
		from:    events,
		groupBy: eventFields,
	}
}

func (filter *Filter) lefJoinUniqueVisitorsByPeriod(fields []Field) *queryBuilder {
	if filter.fieldsContain(fields, FieldCRPeriod) {
		var groupBy Field

		if filter.fieldsContain(fields, FieldDay) {
			groupBy = FieldDay
		} else {
			groupBy = FieldHour
		}

		return &queryBuilder{
			filter: &Filter{
				ClientID:    filter.ClientID,
				Timezone:    filter.Timezone,
				From:        filter.From,
				To:          filter.To,
				Period:      filter.Period,
				IncludeTime: filter.IncludeTime,
			},
			fields:  []Field{groupBy, FieldVisitorsRaw},
			from:    sessions,
			groupBy: []Field{groupBy},
		}
	}

	return nil
}

func (filter *Filter) excludeFields(fields []Field, exclude ...Field) []Field {
	result := make([]Field, 0, len(fields))

	for _, field := range fields {
		if !filter.fieldsContain(exclude, field) {
			result = append(result, field)
		}
	}

	return result
}

func (filter *Filter) fieldsContain(haystack []Field, needle Field) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

func (filter *Filter) valuesContainPrefix(haystack []string, prefix string) bool {
	for i := range haystack {
		if strings.HasPrefix(haystack[i], prefix) {
			return true
		}
	}

	return false
}

func (filter *Filter) searchContains(needle Field) bool {
	for i := range filter.Search {
		if filter.Search[i].Field == needle {
			return true
		}
	}

	return false
}

func (filter *Filter) fieldsContainByQuerySession(haystack []Field, needle string) bool {
	for i := range haystack {
		if haystack[i].querySessions == needle {
			return true
		}
	}

	return false
}
