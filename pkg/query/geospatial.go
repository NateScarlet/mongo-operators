package query

// https://docs.mongodb.com/manual/reference/operator/query-geospatial/

// GeoIntersects selects documents whose geospatial data
// intersects with a specified GeoJSON object;
// i.e. where the intersection of the data and the specified object is non-empty.
// https://docs.mongodb.com/manual/reference/operator/query/geoIntersects/
func GeoIntersects(geometry interface{}) M {
	return M{"$geoIntersects": geometry}
}

// GeoWithIn selects documents with geospatial data
// that exists entirely within a specified shape.
// https://docs.mongodb.com/manual/reference/operator/query/geoWithin/
func GeoWithIn(geometry interface{}) M {
	return M{"$geoWithin": geometry}
}

// NearOperator returned from Near.
type NearOperator M

// Near returns geospatial objects in proximity to a point.
// Requires a geospatial index.
// https://docs.mongodb.com/manual/reference/operator/query/near/
func Near(longitude, latitude) NearOperator {
	return NearOperator{"$near": Geometry("Point", []float64{longitude, latitude})}
}

// SetMinDistance option
func (op NearOperator) SetMinDistance(meters float64) NearOperator {
	op["$near"].(M)["$minDistance"] = meters
	return op
}

// SetMaxDistance option
func (op NearOperator) SetMaxDistance(meters float64) NearOperator {
	op["$near"].(M)["$maxDistance"] = meters
	return op
}

// NearLegacyOperator returned from NearLegacy.
type NearLegacyOperator M

// NearLegacy specify a point using legacy coordinates.
// https://docs.mongodb.com/manual/reference/operator/query/near/
func NearLegacy(x, y float64) NearLegacyOperator {
	return NearLegacyOperator{"$near": []float64{x, y}}
}

// SetMaxDistance option
func (op NearLegacyOperator) SetMaxDistance(meters float64) NearLegacyOperator {
	op["$near"].(M)["$maxDistance"] = meters
	return op
}

// NearSphereOperator returned from NearSphere.
type NearSphereOperator M

// NearSphere returns geospatial objects in proximity to a point on a sphere.
// Requires a geospatial index.
// https://docs.mongodb.com/manual/reference/operator/query/nearSphere/
func NearSphere(longitude, latitude float64) NearSphereOperator {
	return NearSphereOperator{"$nearSphere": Geometry("Point", []float64{longitude, latitude})}
}

// SetMaxDistance option
func (op NearSphereOperator) SetMaxDistance(meters int) NearSphereOperator {
	op["$nearSphere"].(M)["$maxDistance"] = meters
	return op
}

// SetMinDistance option
func (op NearSphereOperator) SetMinDistance(meters int) NearSphereOperator {
	op["$nearSphere"].(M)["$minDistance"] = meters
	return op
}

// NearSphereLegacyOperator returned from NearLegacy.
type NearSphereLegacyOperator M

// NearSphereLegacy specify a point using legacy coordinates.
// https://docs.mongodb.com/manual/reference/operator/query/nearSphere/
func NearSphereLegacy(x, y float64) NearSphereLegacyOperator {
	return NearSphereLegacyOperator{"$nearSphere": []float64{x, y}}
}

// SetMinDistance option
func (op NearSphereLegacyOperator) SetMinDistance(meters float64) NearSphereLegacyOperator {
	op["$near"].(M)["$minDistance"] = meters
	return op
}

// SetMaxDistance option
func (op NearSphereLegacyOperator) SetMaxDistance(meters float64) NearSphereLegacyOperator {
	op["$near"].(M)["$maxDistance"] = meters
	return op
}

// Box specifies a rectangular box using legacy coordinate pairs for $geoWithin queries.
// The 2d index supports $box.
// https://docs.mongodb.com/manual/reference/operator/query/box/
func Box(x1, y1, x2, y2 float64) M {
	return M{"$box": A{
		A{x1, y1},
		A{x2, y2},
	}}
}

// Center specifies a circle using legacy coordinate pairs
// to $geoWithin queries when using planar geometry. The 2d index supports $center.
// https://docs.mongodb.com/manual/reference/operator/query/center/
func Center(x, y, radius float64) M {
	return M{"$center": A{[2]float64{x, y}, radius}}
}

// CenterSphere specifies a circle using either legacy coordinate pairs
// or GeoJSON format for $geoWithin queries when using spherical geometry.
// The 2dsphere and 2d indexes support $centerSphere.
// https://docs.mongodb.com/manual/reference/operator/query/centerSphere/
func CenterSphere(x, y, radius float64) M {
	return M{"$centerSphere": A{[2]float64{x, y}, radius}}
}

// GeometryOperator returned from Geometry
type GeometryOperator M

// Geometry specifies a geometry in GeoJSON format to geospatial query operators.
// https://docs.mongodb.com/manual/reference/operator/query/geometry/
func Geometry(typ string, coordinates interface{}) GeometryOperator {
	return GeometryOperator{"$geometry": M{
		"type":        typ,
		"coordinates": coordinates,
	}}
}

// SetCRS option
func (op GeometryOperator) SetCRS(v interface{}) GeometryOperator {
	op["$geometry"].(M)["crs"] = v
	return op
}

// Polygon specifies a polygon to using legacy coordinate pairs
// for $geoWithin queries. The 2d index supports $center.
// https://docs.mongodb.com/manual/reference/operator/query/polygon/
func Polygon(point ...[2]float64{}) M {
	return M{"$polygon": points}
}
