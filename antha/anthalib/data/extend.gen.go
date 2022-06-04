package data

// Code generated by gen.py. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// Float64 adds a float64 col using float64 inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to float64; no conversions are performed.
func (e *ExtendOn) Float64(f func(v ...float64) (float64, bool)) (*Table, error) {
	typ := typeFloat64
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterFloat64, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateFloat64(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendFloat64{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterFloat64 = (*extendFloat64)(nil)

type extendFloat64 struct {
	f      func(v ...float64) (float64, bool)
	source []iterFloat64
}

func (x *extendFloat64) Next() bool {
	return true
}

func (x *extendFloat64) Value() interface{} {
	v, ok := x.Float64()
	if !ok {
		return nil
	}
	return v
}

func (x *extendFloat64) Float64() (float64, bool) {
	args := make([]float64, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.Float64()
		if !ok {
			return float64(0), false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// Float64 adds a float64 col using float64 inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) Float64(f func(v ...float64) (float64, bool)) *Table {
	t, err := m.ExtendOn.Float64(f)
	handle(err)
	return t
}

// InterfaceFloat64 adds a float64 col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceFloat64(f func(v ...interface{}) (float64, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeFloat64,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceFloat64{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterFloat64 = (*extendInterfaceFloat64)(nil)

type extendInterfaceFloat64 struct {
	f      func(v ...interface{}) (float64, bool)
	source []iterator
}

func (x *extendInterfaceFloat64) Next() bool {
	return true
}

func (x *extendInterfaceFloat64) Value() interface{} {
	v, ok := x.Float64()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceFloat64) Float64() (float64, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceFloat64 adds a float64 col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceFloat64(f func(v ...interface{}) (float64, bool)) *Table {
	t, err := m.ExtendOn.InterfaceFloat64(f)
	handle(err)
	return t
}

// Int64 adds a int64 col using int64 inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to int64; no conversions are performed.
func (e *ExtendOn) Int64(f func(v ...int64) (int64, bool)) (*Table, error) {
	typ := typeInt64
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterInt64, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateInt64(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendInt64{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterInt64 = (*extendInt64)(nil)

type extendInt64 struct {
	f      func(v ...int64) (int64, bool)
	source []iterInt64
}

func (x *extendInt64) Next() bool {
	return true
}

func (x *extendInt64) Value() interface{} {
	v, ok := x.Int64()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInt64) Int64() (int64, bool) {
	args := make([]int64, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.Int64()
		if !ok {
			return int64(0), false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// Int64 adds a int64 col using int64 inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) Int64(f func(v ...int64) (int64, bool)) *Table {
	t, err := m.ExtendOn.Int64(f)
	handle(err)
	return t
}

// InterfaceInt64 adds a int64 col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceInt64(f func(v ...interface{}) (int64, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeInt64,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceInt64{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterInt64 = (*extendInterfaceInt64)(nil)

type extendInterfaceInt64 struct {
	f      func(v ...interface{}) (int64, bool)
	source []iterator
}

func (x *extendInterfaceInt64) Next() bool {
	return true
}

func (x *extendInterfaceInt64) Value() interface{} {
	v, ok := x.Int64()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceInt64) Int64() (int64, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceInt64 adds a int64 col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceInt64(f func(v ...interface{}) (int64, bool)) *Table {
	t, err := m.ExtendOn.InterfaceInt64(f)
	handle(err)
	return t
}

// Int adds a int col using int inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to int; no conversions are performed.
func (e *ExtendOn) Int(f func(v ...int) (int, bool)) (*Table, error) {
	typ := typeInt
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterInt, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateInt(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendInt{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterInt = (*extendInt)(nil)

type extendInt struct {
	f      func(v ...int) (int, bool)
	source []iterInt
}

func (x *extendInt) Next() bool {
	return true
}

func (x *extendInt) Value() interface{} {
	v, ok := x.Int()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInt) Int() (int, bool) {
	args := make([]int, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.Int()
		if !ok {
			return 0, false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// Int adds a int col using int inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) Int(f func(v ...int) (int, bool)) *Table {
	t, err := m.ExtendOn.Int(f)
	handle(err)
	return t
}

// InterfaceInt adds a int col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceInt(f func(v ...interface{}) (int, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeInt,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceInt{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterInt = (*extendInterfaceInt)(nil)

type extendInterfaceInt struct {
	f      func(v ...interface{}) (int, bool)
	source []iterator
}

func (x *extendInterfaceInt) Next() bool {
	return true
}

func (x *extendInterfaceInt) Value() interface{} {
	v, ok := x.Int()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceInt) Int() (int, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceInt adds a int col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceInt(f func(v ...interface{}) (int, bool)) *Table {
	t, err := m.ExtendOn.InterfaceInt(f)
	handle(err)
	return t
}

// String adds a string col using string inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to string; no conversions are performed.
func (e *ExtendOn) String(f func(v ...string) (string, bool)) (*Table, error) {
	typ := typeString
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterString, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateString(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendString{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterString = (*extendString)(nil)

type extendString struct {
	f      func(v ...string) (string, bool)
	source []iterString
}

func (x *extendString) Next() bool {
	return true
}

func (x *extendString) Value() interface{} {
	v, ok := x.String()
	if !ok {
		return nil
	}
	return v
}

func (x *extendString) String() (string, bool) {
	args := make([]string, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.String()
		if !ok {
			return "", false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// String adds a string col using string inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) String(f func(v ...string) (string, bool)) *Table {
	t, err := m.ExtendOn.String(f)
	handle(err)
	return t
}

// InterfaceString adds a string col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceString(f func(v ...interface{}) (string, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeString,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceString{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterString = (*extendInterfaceString)(nil)

type extendInterfaceString struct {
	f      func(v ...interface{}) (string, bool)
	source []iterator
}

func (x *extendInterfaceString) Next() bool {
	return true
}

func (x *extendInterfaceString) Value() interface{} {
	v, ok := x.String()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceString) String() (string, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceString adds a string col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceString(f func(v ...interface{}) (string, bool)) *Table {
	t, err := m.ExtendOn.InterfaceString(f)
	handle(err)
	return t
}

// Bool adds a bool col using bool inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to bool; no conversions are performed.
func (e *ExtendOn) Bool(f func(v ...bool) (bool, bool)) (*Table, error) {
	typ := typeBool
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterBool, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateBool(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendBool{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterBool = (*extendBool)(nil)

type extendBool struct {
	f      func(v ...bool) (bool, bool)
	source []iterBool
}

func (x *extendBool) Next() bool {
	return true
}

func (x *extendBool) Value() interface{} {
	v, ok := x.Bool()
	if !ok {
		return nil
	}
	return v
}

func (x *extendBool) Bool() (bool, bool) {
	args := make([]bool, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.Bool()
		if !ok {
			return false, false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// Bool adds a bool col using bool inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) Bool(f func(v ...bool) (bool, bool)) *Table {
	t, err := m.ExtendOn.Bool(f)
	handle(err)
	return t
}

// InterfaceBool adds a bool col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceBool(f func(v ...interface{}) (bool, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeBool,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceBool{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterBool = (*extendInterfaceBool)(nil)

type extendInterfaceBool struct {
	f      func(v ...interface{}) (bool, bool)
	source []iterator
}

func (x *extendInterfaceBool) Next() bool {
	return true
}

func (x *extendInterfaceBool) Value() interface{} {
	v, ok := x.Bool()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceBool) Bool() (bool, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceBool adds a bool col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceBool(f func(v ...interface{}) (bool, bool)) *Table {
	t, err := m.ExtendOn.InterfaceBool(f)
	handle(err)
	return t
}

// TimestampMillis adds a TimestampMillis col using TimestampMillis inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to TimestampMillis; no conversions are performed.
func (e *ExtendOn) TimestampMillis(f func(v ...TimestampMillis) (TimestampMillis, bool)) (*Table, error) {
	typ := typeTimestampMillis
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterTimestampMillis, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateTimestampMillis(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendTimestampMillis{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterTimestampMillis = (*extendTimestampMillis)(nil)

type extendTimestampMillis struct {
	f      func(v ...TimestampMillis) (TimestampMillis, bool)
	source []iterTimestampMillis
}

func (x *extendTimestampMillis) Next() bool {
	return true
}

func (x *extendTimestampMillis) Value() interface{} {
	v, ok := x.TimestampMillis()
	if !ok {
		return nil
	}
	return v
}

func (x *extendTimestampMillis) TimestampMillis() (TimestampMillis, bool) {
	args := make([]TimestampMillis, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.TimestampMillis()
		if !ok {
			return TimestampMillis(0), false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// TimestampMillis adds a TimestampMillis col using TimestampMillis inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) TimestampMillis(f func(v ...TimestampMillis) (TimestampMillis, bool)) *Table {
	t, err := m.ExtendOn.TimestampMillis(f)
	handle(err)
	return t
}

// InterfaceTimestampMillis adds a TimestampMillis col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceTimestampMillis(f func(v ...interface{}) (TimestampMillis, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeTimestampMillis,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceTimestampMillis{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterTimestampMillis = (*extendInterfaceTimestampMillis)(nil)

type extendInterfaceTimestampMillis struct {
	f      func(v ...interface{}) (TimestampMillis, bool)
	source []iterator
}

func (x *extendInterfaceTimestampMillis) Next() bool {
	return true
}

func (x *extendInterfaceTimestampMillis) Value() interface{} {
	v, ok := x.TimestampMillis()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceTimestampMillis) TimestampMillis() (TimestampMillis, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceTimestampMillis adds a TimestampMillis col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceTimestampMillis(f func(v ...interface{}) (TimestampMillis, bool)) *Table {
	t, err := m.ExtendOn.InterfaceTimestampMillis(f)
	handle(err)
	return t
}

// TimestampMicros adds a TimestampMicros col using TimestampMicros inputs.  Null on any null inputs.
// Returns error if any column cannot be assigned to TimestampMicros; no conversions are performed.
func (e *ExtendOn) TimestampMicros(f func(v ...TimestampMicros) (TimestampMicros, bool)) (*Table, error) {
	typ := typeTimestampMicros
	inputs, err := e.inputs(typ)
	if err != nil {
		return nil, err
	}
	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typ,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterTimestampMicros, len(inputs))
			var err error
			for i, ser := range inputs {
				iter := cache.Ensure(ser)
				colReader[i], err = ser.iterateTimestampMicros(iter) // note colReader[i] is not itself in the cache!
				if err != nil {
					panic(errors.Wrapf(err, "SHOULD NOT HAPPEN; when extending new column %q", e.extension.newCol))
				}
			}
			// end when table exhausted
			e.extension.extensionSource(cache)
			return &extendTimestampMicros{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterTimestampMicros = (*extendTimestampMicros)(nil)

type extendTimestampMicros struct {
	f      func(v ...TimestampMicros) (TimestampMicros, bool)
	source []iterTimestampMicros
}

func (x *extendTimestampMicros) Next() bool {
	return true
}

func (x *extendTimestampMicros) Value() interface{} {
	v, ok := x.TimestampMicros()
	if !ok {
		return nil
	}
	return v
}

func (x *extendTimestampMicros) TimestampMicros() (TimestampMicros, bool) {
	args := make([]TimestampMicros, len(x.source))
	var ok bool
	for i, s := range x.source {
		args[i], ok = s.TimestampMicros()
		if !ok {
			return TimestampMicros(0), false
		}
	}
	v, notNull := x.f(args...)
	return v, notNull
}

// TimestampMicros adds a TimestampMicros col using TimestampMicros inputs.  Null on any null inputs.
// Panics on error.
func (m *MustExtendOn) TimestampMicros(f func(v ...TimestampMicros) (TimestampMicros, bool)) *Table {
	t, err := m.ExtendOn.TimestampMicros(f)
	handle(err)
	return t
}

// InterfaceTimestampMicros adds a TimestampMicros col using arbitrary (interface{}) inputs.
func (e *ExtendOn) InterfaceTimestampMicros(f func(v ...interface{}) (TimestampMicros, bool)) (*Table, error) {
	projection, err := newProjection(e.extension.t.schema, e.inputCols...)
	if err != nil {
		return nil, err
	}

	return newFromSeries(append(append([]*Series(nil), e.extension.t.series...), &Series{
		col:  e.extension.newCol,
		typ:  typeTimestampMicros,
		meta: e.meta,
		read: func(cache *seriesIterCache) iterator {
			colReader := make([]iterator, len(projection.newToOld))
			for new, old := range projection.newToOld {
				colReader[new] = cache.Ensure(e.extension.t.series[old])
			}
			// end when table exhausted
			//e.extension.extensionSource(cache)
			return &extendInterfaceTimestampMicros{f: f, source: colReader}
		}},
	), e.extension.t.sortKey...), nil
}

var _ iterTimestampMicros = (*extendInterfaceTimestampMicros)(nil)

type extendInterfaceTimestampMicros struct {
	f      func(v ...interface{}) (TimestampMicros, bool)
	source []iterator
}

func (x *extendInterfaceTimestampMicros) Next() bool {
	return true
}

func (x *extendInterfaceTimestampMicros) Value() interface{} {
	v, ok := x.TimestampMicros()
	if !ok {
		return nil
	}
	return v
}

func (x *extendInterfaceTimestampMicros) TimestampMicros() (TimestampMicros, bool) {
	args := make([]interface{}, len(x.source))
	for i, s := range x.source {
		args[i] = s.Value()
	}
	return x.f(args...)
}

// InterfaceTimestampMicros adds a TimestampMicros col using arbitrary (interface{}) inputs.
// Panics on error.
func (m *MustExtendOn) InterfaceTimestampMicros(f func(v ...interface{}) (TimestampMicros, bool)) *Table {
	t, err := m.ExtendOn.InterfaceTimestampMicros(f)
	handle(err)
	return t
}