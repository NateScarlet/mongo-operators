package aggregation

// AddFields to documents. Similar to $project, $addFields reshapes
// each document in the stream; specifically, by adding new fields
// to output documents that contain both the existing fields from
// the input documents and the newly added fields.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/addFields
func AddFields(fields interface{}) M {
	return M{"$addFields": fields}
}

// BucketStage returned from Bucket
type BucketStage M

// Bucket categorizes incoming documents into groups, called buckets,
// based on a specified expression and bucket boundaries.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/bucket/
func Bucket(groupBy, boundaries interface{}) BucketStage {
	return BucketStage{"$bucket": M{
		"groupBy":    groupBy,
		"boundaries": boundaries,
	}}
}

// SetDefault option
func (stage BucketStage) SetDefault(literal interface{}) BucketStage {
	stage["$bucket"].(M)["default"] = literal
	return stage
}

// SetOutput option
func (stage BucketStage) SetOutput(document interface{}) BucketStage {
	stage["$bucket"].(M)["output"] = document
	return stage
}

// BucketAutoStage returned from BucketAuto
type BucketAutoStage M

// BucketAuto categorizes incoming documents into a specific number of groups,
// called buckets, based on a specified expression.
// Bucket boundaries are automatically determined in an attempt
// to evenly distribute the documents into the specified number of buckets.
// https://docs.mongodb.com/manual/reference/operator/aggregation/bucketAuto/
func BucketAuto(groupBy interface{}, buckets int) BucketAutoStage {
	return BucketAutoStage{"$bucketAuto": M{
		"groupBy": groupBy,
		"buckets": buckets,
	}}
}

// SetOutput option
func (stage BucketAutoStage) SetOutput(document interface{}) BucketAutoStage {
	stage["$bucketAuto"].(M)["output"] = document
	return stage
}

// SetGranularity option
func (stage BucketAutoStage) SetGranularity(v string) BucketAutoStage {
	stage["$bucketAuto"].(M)["granularity"] = v
	return stage
}

// CollStatsStage returned from CollStats
type CollStatsStage M

// CollStats returns statistics regarding a collection or view.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/collStats/
func CollStats() CollStatsStage {
	return CollStatsStage{"$collStats": M{}}
}

// SetLatencyStats option
func (stage CollStatsStage) SetLatencyStats(histograms bool) CollStatsStage {
	stage["$collStats"].(M)["latencyStats"] = M{"histograms": histograms}
	return stage
}

// SetStorageStats option
func (stage CollStatsStage) SetStorageStats(scale int64) CollStatsStage {
	stage["$collStats"].(M)["storageStats"] = M{"scale": scale}
	return stage
}

// SetCount option
func (stage CollStatsStage) SetCount() CollStatsStage {
	stage["$collStats"].(M)["count"] = M{}
	return stage
}

// Count returns a count of the number of documents
// at this stage of the aggregation pipeline.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/count/
func Count(outputField string) M {
	return M{"$count": outputField}
}

// CurrentOp returns a stream of documents containing information
// on active and/or dormant operations as well as inactive sessions
// that are holding locks as part of a transaction.
// The stage returns a document for each operation or session.
// To run $currentOp, use the db.aggregate() helper on the admin database.
// https://docs.mongodb.com/manual/reference/operator/aggregation/currentOp/
func CurrentOp(options interface{}) M {
	return M{"$currentOp": options}
}

// Facet Processes multiple aggregation pipelines within a single stage
// on the same set of input documents.
// Enables the creation of multi-faceted aggregations capable of
// characterizing data across multiple dimensions, or facets, in a single stage.
// https://docs.mongodb.com/manual/reference/operator/aggregation/facet/
func Facet(outputs interface{}) M {
	return M{"$facet": outputs}
}

// GeoNearStage returned from GeoNear
type GeoNearStage M

// GeoNear returns an ordered stream of documents based on the proximity to
// a geospatial point. Incorporates the functionality of $match, $sort,
// and $limit for geospatial data. The output documents include an additional
// distance field and can include a location identifier field.
// https://docs.mongodb.com/manual/reference/operator/aggregation/geoNear/
func GeoNear(near interface{}, distanceField string) GeoNearStage {
	return GeoNearStage{"$geoNear": M{
		"near":          near,
		"distanceField": distanceField,
	}}
}

// SetSpherical option
func (stage GeoNearStage) SetSpherical(v bool) GeoNearStage {
	stage["$geoNear"].(M)["spherical"] = v
	return stage
}

// SetMaxDistance option
func (stage GeoNearStage) SetMaxDistance(v float64) GeoNearStage {
	stage["$geoNear"].(M)["maxDistance"] = v
	return stage
}

// SetQuery option
func (stage GeoNearStage) SetQuery(v interface{}) GeoNearStage {
	stage["$geoNear"].(M)["query"] = v
	return stage
}

// SetDistanceMultiplier option
func (stage GeoNearStage) SetDistanceMultiplier(v float64) GeoNearStage {
	stage["$geoNear"].(M)["distanceMultiplier"] = v
	return stage
}

// SetIncludeLocs option
func (stage GeoNearStage) SetIncludeLocs(v string) GeoNearStage {
	stage["$geoNear"].(M)["includeLocs"] = v
	return stage
}

// SetUniqueDocs option
func (stage GeoNearStage) SetUniqueDocs(v bool) GeoNearStage {
	stage["$geoNear"].(M)["uniqueDocs"] = v
	return stage
}

// SetMinDistance option
func (stage GeoNearStage) SetMinDistance(v float64) GeoNearStage {
	stage["$geoNear"].(M)["minDistance"] = v
	return stage
}

// SetKey option
func (stage GeoNearStage) SetKey(v string) GeoNearStage {
	stage["$geoNear"].(M)["key"] = v
	return stage
}

// GraphLookupStage stage
type GraphLookupStage M

// GraphLookup performs a recursive search on a collection.
// To each output document, adds a new array field that contains
// the traversal results of the recursive search for that document.
// https://docs.mongodb.com/manual/reference/operator/aggregation/graphLookup/
func GraphLookup(from string, startWith interface{}, connectFromField string, connectToField string, as string) GraphLookupStage {
	return GraphLookupStage{"$graphLookup": M{
		"from":             from,
		"startWith":        startWith,
		"connectFromField": connectFromField,
		"connectToField":   connectToField,
		"as":               as,
	}}
}

// SetMaxDepth option
func (stage GraphLookupStage) SetMaxDepth(v int) GraphLookupStage {
	stage["$graphLookup"].(M)["maxDepth"] = v
	return stage
}

// SetDepthField option
func (stage GraphLookupStage) SetDepthField(v string) GraphLookupStage {
	stage["$graphLookup"].(M)["depthField"] = v
	return stage
}

// SetRestrictSearchWithMatch option
func (stage GraphLookupStage) SetRestrictSearchWithMatch(v interface{}) GraphLookupStage {
	stage["$graphLookup"].(M)["restrictSearchWithMatch"] = v
	return stage
}

// Group input documents by a specified identifier expression
// and applies the accumulator expression(s), if specified, to each group.
// Consumes all input documents and outputs one document per each distinct group.
// The output documents only contain the identifier field and, if specified,
// accumulated fields.
// https://docs.mongodb.com/manual/reference/operator/aggregation/group/
func Group(idAndFields interface{}) M {
	return M{"$group": idAndFields}
}

// IndexStats returns statistics regarding the use of each index
// for the collection.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/indexStats/
func IndexStats() M {
	return M{"$indexStats": M{}}
}

// Limit Passes the first n documents unmodified to the pipeline
// where n is the specified limit. For each input document, outputs
// either one document (for the first n documents) or zero documents
// (after the first n documents).
// https://docs.mongodb.com/manual/reference/operator/aggregation/limit/
func Limit(n int) M {
	return M{"$limit": n}
}

// ListLocalSessions lists the sessions cached in memory by the mongod or mongos instance.
// New in version 3.6.
// https://docs.mongodb.com/manual/reference/operator/aggregation/listLocalSessions/
func ListLocalSessions(args interface{}) M {
	return M{"$listLocalSessions": args}
}

// ListSessions that have been active long enough to propagate to the
// system.sessions collection.
// New in version 3.6.
// https://docs.mongodb.com/manual/reference/operator/aggregation/listSessions/
func ListSessions(args interface{}) M {
	return M{"$listSession": args}
}

// LookupF perform an equality match between a field from the input documents
// with a field from the documents of the “joined” collection.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/#equality-match
func LookupF(from, localField, foreignField, as string) M {
	return M{"$lookup": M{
		"from":         from,
		"localField":   localField,
		"foreignField": foreignField,
		"as":           as,
	}}
}

// LookupP perform uncorrelated subqueries between two collections
// as well as allow other join conditions besides a single equality match.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/#join-conditions-and-uncorrelated-sub-queries
func LookupP(from string, let M, pipeline A, as string) M {
	var opts = M{
		"from":     from,
		"pipeline": pipeline,
		"as":       as,
	}
	if let != nil {
		opts["let"] = let
	}
	return M{"$lookup": opts}
}

// Match filters the document stream to allow only matching documents
// to pass unmodified into the next pipeline stage.
// $match uses standard MongoDB queries. For each input document,
// outputs either one document (a match) or zero documents (no match).
// https://docs.mongodb.com/manual/reference/operator/aggregation/match/
// Match https://docs.mongodb.com/manual/reference/operator/aggregation/match
func Match(query interface{}) M {
	return M{"$match": query}
}

// MergeStage returned from Merge
type MergeStage M

// Merge writes the resulting documents of the aggregation pipeline to a collection.
// The stage can incorporate (insert new documents, merge documents,
// replace documents, keep existing documents, fail the operation,
// process documents with a custom update pipeline) the results
// into an output collection. To use the $merge stage,
// it must be the last stage in the pipeline.
// New in version 4.2.
func Merge(into interface{}) MergeStage {
	return MergeStage{"$merge": M{"into": into}}
}

// SetOn option
func (stage MergeStage) SetOn(v interface{}) MergeStage {
	stage["$merge"].(M)["on"] = v
	return stage
}

// SetWhenMatched option
func (stage MergeStage) SetWhenMatched(v interface{}) MergeStage {
	stage["$merge"].(M)["whenMatched"] = v
	return stage
}

// SetLet option
func (stage MergeStage) SetLet(v interface{}) MergeStage {
	stage["$merge"].(M)["let"] = v
	return stage
}

// SetWhenNotMatched option
func (stage MergeStage) SetWhenNotMatched(v interface{}) MergeStage {
	stage["$merge"].(M)["whenNotMatched"] = v
	return stage
}

// Out writes the resulting documents of the aggregation pipeline to a collection.
// To use the $out stage, it must be the last stage in the pipeline.
// https://docs.mongodb.com/manual/reference/operator/aggregation/out/
func Out(collection string) M {
	return M{"$out": collection}
}

// PlanCacheStats returns plan cache information for a collection.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/planCacheStats/
func PlanCacheStats() M {
	return M{"$planCacheStats": M{}}
}

// Project reshapes each document in the stream, such as by adding new fields
// or removing existing fields. For each input document, outputs one document.
// See also $unset for removing existing fields.
// https://docs.mongodb.com/manual/reference/operator/aggregation/project/
func Project(specifications interface{}) M {
	return M{"$project": specifications}
}

// Redact reshapes each document in the stream by restricting the content
// for each document based on information stored in the documents themselves.
// Incorporates the functionality of $project and $match.
// Can be used to implement field level redaction.
// For each input document, outputs either one or zero documents.
// https://docs.mongodb.com/manual/reference/operator/aggregation/redact/
func Redact(expr interface{}) M {
	return M{"$redact": expr}
}

// ReplaceRoot replaces a document with the specified embedded document.
// The operation replaces all existing fields in the input document,
// including the _id field. Specify a document embedded in the input document
// to promote the embedded document to the top level.
// $replaceWith is an alias for $replaceRoot stage.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/replaceRoot/
func ReplaceRoot(newRoot interface{}) M {
	return M{"$replaceRoot": M{"newRoot": newRoot}}
}

// ReplaceWith replaces a document with the specified embedded document.
// The operation replaces all existing fields in the input document, including the _id field.
// Specify a document embedded in the input document to promote
// the embedded document to the top level.
// $replaceWith is an alias for $replaceRoot stage.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/replaceWith/
func ReplaceWith(replacementDocument interface{}) M {
	return M{"$replaceWith": replacementDocument}
}

// Sample randomly selects the specified number of documents from its input.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sample/
func Sample(size int) M {
	return M{"$sample": M{"size": size}}
}

// Search aggregation pipleline stage performs a full-text search
// of the field or fields in an Atlas collection.
// The fields must be covered by an Atlas Search index.
// https://docs.mongodb.com/manual/reference/operator/aggregation/search/
// https://docs.atlas.mongodb.com/atlas-search/query-syntax/
func Search(args interface{}) M {
	return M{"$search": args}
}

// Set Adds new fields to documents.
// Similar to $project, $set reshapes each document in the stream;
// specifically, by adding new fields to output documents that contain
// both the existing fields from the input documents and the newly added fields.
// $set is an alias for $addFields stage.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/set/
func Set(fields interface{}) M {
	return M{"$set": fields}
}

type SetWindowFieldsStage M

type SetWindowFieldsStageOutput M

func SetWindowFieldsOutput(field string, operator M) SetWindowFieldsStageOutput {
	return SetWindowFieldsStageOutput{
		field: operator,
	}
}

func (o SetWindowFieldsStageOutput) operator() M {
	for _, v := range o {
		return v.(M)
	}
	panic("SetWindowFieldsOutput: output is empty")
}

func (o SetWindowFieldsStageOutput) window() M {
	var op = o.operator()
	if op["window"] == nil {
		op["window"] = M{}
	}
	return op["window"].(M)
}

// SetDocuments set a window where the lower and upper boundaries
// are specified relative to the position of the current document
// read from the collection.
//
// The window boundaries are specified using a two element array containing a lower and upper limit string or integer. Use:
//
// The "current" string for the current document position in the output.
// The "unbounded" string for the first or last document position in the partition.
// An integer for a position relative to the current document.
// Use a negative integer for a position before the current document.
// Use a positive integer for a position after the current document.
// 0 is the current document position.
func (o SetWindowFieldsStageOutput) SetDocuments(lower, upper interface{}) SetWindowFieldsStageOutput {
	o.window()["documents"] = A{lower, upper}
	return o
}

// SetRange set a window where the lower and upper boundaries are defined
// using a range of values based on the sortBy field in the current document.
//
// The window boundaries are specified using a two element array
// containing a lower and upper limit string or number. Use:
//
// The "current" string for the current document position in the output.
// The "unbounded" string for the first or last document position in the partition.
// A number to add to the value of the sortBy field for the current document.
// A document is in the window if the sortBy field value
// is inclusively within the lower and upper boundaries.
func (o SetWindowFieldsStageOutput) SetRange(lower, upper interface{}) SetWindowFieldsStageOutput {
	o.window()["range"] = A{lower, upper}
	return o
}

type SetWindowFieldsUnit string

const (
	SWFUYear        SetWindowFieldsUnit = "year"
	SWFUQuarter     SetWindowFieldsUnit = "quarter"
	SWFUMonth       SetWindowFieldsUnit = "month"
	SWFUWeek        SetWindowFieldsUnit = "week"
	SWFUDay         SetWindowFieldsUnit = "day"
	SWFUHour        SetWindowFieldsUnit = "hour"
	SWFUMinute      SetWindowFieldsUnit = "minute"
	SWFUSecond      SetWindowFieldsUnit = "second"
	SWFUMillisecond SetWindowFieldsUnit = "millisecond"
)

func (o SetWindowFieldsStageOutput) SetUnit(unit SetWindowFieldsUnit) SetWindowFieldsStageOutput {
	o.window()["unit"] = unit
	return o
}

// SetWindowFields groups documents into windows and applies one or more operators
// to the documents in each window.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setWindowFields/
func SetWindowFields(outputs ...SetWindowFieldsStageOutput) SetWindowFieldsStage {
	var output = M{}
	for _, i := range outputs {
		for k, v := range i {
			output[k] = v
		}
	}
	return SetWindowFieldsStage{
		"$setWindowFields": M{
			"output": output,
		},
	}
}

// SetPartitionBy specifies an expression to group the documents.
// In the $setWindowFields stage, the group of documents is known as a partition.
// Default is one partition for the entire collection.
func (stage SetWindowFieldsStage) SetPartitionBy(expression interface{}) SetWindowFieldsStage {
	stage["$setWindowFields"].(M)["partitionBy"] = expression
	return stage
}

// SetSortBy specifies the field(s) to sort the documents by in the partition.
// Uses the same syntax as the $sort stage. Default is no sorting.
func (stage SetWindowFieldsStage) SetSortBy(sort interface{}) SetWindowFieldsStage {
	stage["$setWindowFields"].(M)["sortBy"] = sort
	return stage
}

// Skip the first n documents where n is the specified skip number and passes
// the remaining documents unmodified to the pipeline.
// For each input document, outputs either zero documents (for the first n documents)
// or one document (if after the first n documents).
// https://docs.mongodb.com/manual/reference/operator/aggregation/skip/
func Skip(n int) M {
	return M{"$skip": n}
}

// Sort reorders the document stream by a specified sort key. Only the order changes;
// the documents remain unmodified. For each input document, outputs one document.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sort/
func Sort(order interface{}) M {
	return M{"$sort": order}
}

// SortByCount groups incoming documents based on the value of a specified expression,
// then computes the count of documents in each distinct group.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sortByCount/
func SortByCount(expr interface{}) M {
	return M{"$sortByCount": expr}
}

// UnionWith performs a union of two collections;
// i.e. $unionWith combines pipeline results
// from two collections into a single result set.
// The stage outputs the combined result set (including duplicates) to the next stage.
// New in version 4.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/unionWith/
func UnionWith(collection string, pipeline interface{}) M {
	if pipeline == nil {
		return M{"$unionWith": M{"coll": collection}}
	}
	return M{"$unionWith": M{
		"coll":     collection,
		"pipeline": pipeline,
	}}
}

// Unset removes/excludes fields from documents.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/unset/
func Unset(fields ...string) M {
	if len(fields) == 1 {
		return M{"$unset": fields[0]}
	}
	return M{"$unset": fields}
}

// UnwindStage returned from Unwind
type UnwindStage M

// Unwind deconstructs an array field from the input documents to output
// a document for each element. Each output document replaces the array
// with an element value. For each input document, outputs n documents
// where n is the number of array elements and can be zero for an empty array.
// https://docs.mongodb.com/manual/reference/operator/aggregation/unwind
func Unwind(path string) UnwindStage {
	if path[0] != '$' {
		path = "$" + path
	}
	return UnwindStage{"$unwind": path}
}

// SetIncludeArrayIndex option
func (o UnwindStage) SetIncludeArrayIndex(v string) UnwindStage {
	o["$unwind"] = wrap(o["$unwind"], "path")
	o["$unwind"].(M)["includeArrayIndex"] = v
	return o
}

// SetPreserveNullAndEmptyArrays option
// https://docs.mongodb.com/manual/reference/operator/aggregation/unwind/#document-operand-with-options
func (o UnwindStage) SetPreserveNullAndEmptyArrays(v bool) UnwindStage {
	o["$unwind"] = wrap(o["$unwind"], "path")
	o["$unwind"].(M)["preserveNullAndEmptyArrays"] = v
	return o
}
