package scale

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/d3"
	"honnef.co/go/js/util"
)

var _d3 = js.Global.Get("d3")

// TODO document that you can't implement your own d3 scales, which is
// unfortunate.

type Scale interface {
	Underlying() js.Object
	private()
}

type scale struct{}

func (scale) private() {}

type blackbox struct {
	js.Object
	scale
}

func (s blackbox) Underlying() js.Object {
	return s.Object
}

func Wrap(o js.Object) Scale {
	return blackbox{o, scale{}}
}

type Linear struct {
	js.Object
	scale
}

type Identity struct {
	js.Object
	scale
}

type Power struct {
	js.Object
	scale
}

type Log struct {
	js.Object
	scale
}

type Quantize struct {
	js.Object
	scale
}

type Quantile struct {
	js.Object
	scale
}

type Threshold struct {
	js.Object
	scale
}

func (s Linear) Underlying() js.Object {
	return s.Object
}
func (s Identity) Underlying() js.Object {
	return s.Object
}
func (s Power) Underlying() js.Object {
	return s.Object
}
func (s Log) Underlying() js.Object {
	return s.Object
}
func (s Quantize) Underlying() js.Object {
	return s.Object
}
func (s Quantile) Underlying() js.Object {
	return s.Object
}
func (s Threshold) Underlying() js.Object {
	return s.Object
}

func NewLinear() Linear {
	return Linear{_d3.Get("scale").Call("linear"), scale{}}
}

func NewIdentity() Identity {
	return Identity{_d3.Get("scale").Call("identity"), scale{}}
}

func NewSqrt() Power {
	return Power{_d3.Get("scale").Call("sqrt"), scale{}}
}

func NewPower() Power {
	return Power{_d3.Get("scale").Call("pow"), scale{}}
}

func NewLog() Log {
	return Log{_d3.Get("scale").Call("log"), scale{}}
}

func NewQuantize() Quantize {
	return Quantize{_d3.Get("scale").Call("quantize"), scale{}}
}

func NewQuantile() Quantile {
	return Quantile{_d3.Get("scale").Call("quantile"), scale{}}
}

func NewThreshold() Threshold {
	return Threshold{_d3.Get("scale").Call("threshold"), scale{}}
}

func (s Linear) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Identity) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Power) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Log) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Quantize) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Quantile) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Threshold) Invoke(val interface{}) js.Object {
	return s.Invoke(val)
}

func (s Linear) Copy() Linear {
	return Linear{s.Call("copy"), scale{}}
}

func (s Identity) Copy() Identity {
	return Identity{s.Call("copy"), scale{}}
}

func (s Power) Copy() Power {
	return Power{s.Call("copy"), scale{}}
}

func (s Log) Copy() Log {
	return Log{s.Call("copy"), scale{}}
}

func (s Quantize) Copy() Quantize {
	return Quantize{s.Call("copy"), scale{}}
}

func (s Quantile) Copy() Quantile {
	return Quantile{s.Call("copy"), scale{}}
}

func (s Threshold) Copy() Threshold {
	return Threshold{s.Call("copy"), scale{}}
}

func (s Linear) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Identity) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Power) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Log) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Quantize) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Quantile) SetDomain(numbers []float64) {
	s.Call("domain", numbers)
}

func (s Linear) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Identity) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Power) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Log) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Quantize) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Quantile) Domain() []float64 {
	return util.Float64Slice(s.Call("domain"))
}

func (s Threshold) SetDomain(domain []interface{}) {
	s.Call("domain", domain)
}

func (s Threshold) Domain() []interface{} {
	return s.Call("domain").Interface().([]interface{})
}

func (s Linear) Invert(y float64) float64 {
	return s.Call("invert", y).Float()
}

func (s Identity) Invert(y float64) float64 {
	return s.Call("invert", y).Float()
}

func (s Power) Invert(y float64) float64 {
	return s.Call("invert", y).Float()
}

func (s Log) Invert(y float64) float64 {
	return s.Call("invert", y).Float()
}

func (s Linear) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Identity) SetRange(numbers []float64) {
	s.Call("range", numbers)
}

func (s Power) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Log) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Quantize) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Quantile) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Threshold) SetRange(values []interface{}) {
	s.Call("range", values)
}

func (s Linear) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Identity) Range() []float64 {
	return util.Float64Slice(s.Call("range"))
}

func (s Power) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Log) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Quantize) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Quantile) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Threshold) Range() []interface{} {
	return s.Call("range").Interface().([]interface{})
}

func (s Linear) Nice(count float64) {
	s.Call("nice", count)
}

func (s Power) Nice(m float64) {
	s.Call("nice", m)
}

func (s Log) Nice() {
	s.Call("nice")
}

func (s Quantile) Quantiles() []float64 {
	return util.Float64Slice(s.Call("quantiles"))
}

func (s Quantize) InvertExtent(y interface{}) (float64, float64) {
	vals := s.Call("invertExtent", y).Interface().([]interface{})
	return vals[0].(float64), vals[1].(float64)
}

func (s Quantile) InvertExtent(y interface{}) (float64, float64) {
	vals := s.Call("invertExtent", y).Interface().([]interface{})
	return vals[0].(float64), vals[1].(float64)
}

func (s Threshold) InvertExtent(y interface{}) (interface{}, interface{}) {
	vals := s.Call("invertExtent", y).Interface().([]interface{})
	return vals[0], vals[1]
}

func (s Power) SetExponent(exp float64) {
	s.Call("exponent", exp)
}

func (s Power) Exponent() float64 {
	return s.Call("exponent").Float()
}

func (s Linear) SetRangeRound(values []interface{}) {
	s.Call("rangeRound", values)
}

func (s Power) SetRangeRound(values []interface{}) {
	s.Call("rangeRound", values)
}

func (s Log) SetRangeRound(values []interface{}) {
	s.Call("rangeRound", values)
}

func (s Log) SetBase(base int) {
	s.Call("base", base)
}

func (s Log) Base() int {
	return s.Call("base").Int()
}

func (s Linear) SetClamp(clamp bool) {
	s.Call("clamp", clamp)
}

func (s Power) SetClamp(clamp bool) {
	s.Call("clamp", clamp)
}

func (s Log) SetClamp(clamp bool) {
	s.Call("clamp", clamp)
}

func (s Linear) Clamp() bool {
	return s.Call("clamp").Bool()
}

func (s Power) Clamp() bool {
	return s.Call("clamp").Bool()
}

func (s Log) Clamp() bool {
	return s.Call("clamp").Bool()
}

func (s Linear) Ticks(count int) []float64 {
	return util.Float64Slice(s.Call("ticks", count))
}

func (s Identity) Ticks(count int) []float64 {
	return util.Float64Slice(s.Call("ticks", count))
}

func (s Power) Ticks(count int) []float64 {
	return util.Float64Slice(s.Call("ticks", count))
}

func (s Log) Ticks() []float64 {
	return util.Float64Slice(s.Call("ticks"))
}

func (s Linear) TickFormat(count int, format string) d3.Formatter {
	args := []interface{}{count}
	if len(format) > 0 {
		args = append(args, format)
	}
	ret := s.Call("tickFormat", args...)
	return func(o js.Object) string { return ret.Invoke(o).Str() }
}

func (s Identity) TickFormat(count int, format string) d3.Formatter {
	args := []interface{}{count}
	if len(format) > 0 {
		args = append(args, format)
	}
	ret := s.Call("tickFormat", args...)
	return func(o js.Object) string { return ret.Invoke(o).Str() }
}

func (s Power) TickFormat(count int, format string) d3.Formatter {
	args := []interface{}{}
	if count > -1 {
		args = append(args, count)
	}
	if len(format) > 0 {
		args = append(args, format)
	}
	ret := s.Call("tickFormat", args...)
	return func(o js.Object) string { return ret.Invoke(o).Str() }
}

func (s Log) TickFormat(count int, format string) d3.Formatter {
	args := []interface{}{}
	if count > -1 {
		args = append(args, count)
	}
	if len(format) > 0 {
		args = append(args, format)
	}
	ret := s.Call("tickFormat", args...)
	return func(o js.Object) string { return ret.Invoke(o).Str() }
}
