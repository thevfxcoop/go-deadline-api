package schema

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Decoder struct {
	tag    string
	debug  bool
	fields map[reflect.Type]map[string]reflect.StructField
	hooks  map[reflect.Type]HookFunc
}

type HookFunc func(path string, in, out reflect.Value) error

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewDecoder(tag string, debug bool) *Decoder {
	this := new(Decoder)
	this.tag = tag
	this.debug = debug
	this.fields = make(map[reflect.Type]map[string]reflect.StructField)
	this.hooks = map[reflect.Type]HookFunc{
		reflect.TypeOf(""):               this.setString,
		reflect.TypeOf(false):            this.setBool,
		reflect.TypeOf(uint(0)):          this.setUint,
		reflect.TypeOf(time.Time{}):      this.setTime,
		reflect.TypeOf(time.Duration(0)): this.setDuration,
	}
	return this
}

///////////////////////////////////////////////////////////////////////////////
// DECODE

func (this *Decoder) Decode(in map[string]interface{}, out interface{}) error {
	// Cycle through the fields in the input
	if err := this.decode("", in, reflect.ValueOf(out)); err != nil {
		return err
	}
	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *Decoder) decode(root string, in map[string]interface{}, out reflect.Value) error {
	// Obtain the fields to encode
	fields, err := this.getFields(out)
	if err != nil {
		return err
	}

	for k, v := range in {
		path := filepath.Join(root, k)
		switch value := v.(type) {
		case map[string]interface{}:
			if err := this.decode(path, value, out); err != nil {
				return err
			}
		default:
			if field, exists := fields[path]; exists {
				if err := this.set(path, reflect.ValueOf(v), out.Elem().FieldByIndex(field.Index)); err != nil {
					return err
				}
			} else if this.debug {
				fmt.Printf("  %q => unused: %v\n", path, v)
			}
		}
	}
	// Retrn success
	return nil
}

func (this *Decoder) getFields(r reflect.Value) (map[string]reflect.StructField, error) {
	// Obtain type of destination
	if r.Kind() != reflect.Ptr {
		return nil, deadline.ErrInternalAppError.With("Decoder")
	} else if r = r.Elem(); r.Kind() != reflect.Struct {
		return nil, deadline.ErrInternalAppError.With("Decoder")
	}
	// Return cached data
	if fields, exists := this.fields[r.Type()]; exists {
		return fields, nil
	}
	// Extract information
	fields := make(map[string]reflect.StructField)
	for i := 0; i < r.NumField(); i++ {
		t := r.Type().Field(i)
		if path := this.getPath(t); path != "" {
			fields[path] = r.Type().Field(i)
		}
	}
	// Cache and return information
	this.fields[r.Type()] = fields
	return fields, nil
}

func (this *Decoder) getPath(f reflect.StructField) string {
	tag := strings.SplitN(f.Tag.Get(this.tag), ",", 2)
	if tag[0] == "" {
		tag[0] = f.Name
	}
	return tag[0]
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS - SET

func (this *Decoder) set(path string, src, dest reflect.Value) error {
	t := dest.Type()
	if hook, exists := this.hooks[t]; exists == false {
		// Check to see if a Marshal function exists for pointer to type
		obj := reflect.New(t)
		if method := obj.MethodByName("Unmarshal"); method.IsValid() {
			result := method.Call([]reflect.Value{src})
			err := result[0].Interface()
			if err != nil {
				return err.(error)
			} else {
				dest.Set(obj.Elem())
			}
		} else {
			return deadline.ErrNotFound.With("No marshaler for type: ", dest.Type())
		}
	} else if err := hook(path, src, dest); err != nil {
		return err
	}
	// Return success
	return nil
}

func (this *Decoder) setString(path string, src, dest reflect.Value) error {
	if src.IsValid() == false {
		dest.SetString("")
	} else if src.Kind() != reflect.String {
		return deadline.ErrBadParameter.With(path, ": ", src.Kind())
	} else {
		dest.Set(src)
	}
	return nil
}

func (this *Decoder) setBool(path string, src, dest reflect.Value) error {
	if src.Kind() != reflect.Bool {
		return deadline.ErrBadParameter.With(path, ": ", src.Kind())
	}
	dest.Set(src)
	return nil
}

func (this *Decoder) setUint(path string, src, dest reflect.Value) error {
	switch src.Kind() {
	case reflect.Uint:
		dest.Set(src)
	case reflect.Float64:
		f := src.Float()
		if f != -1 {
			dest.Set(reflect.ValueOf(uint(f)))
		}
	default:
		return deadline.ErrBadParameter.With(path, ": ", src.Kind())
	}
	return nil
}

func (this *Decoder) setTime(path string, src, dest reflect.Value) error {
	if src.Kind() != reflect.String {
		return deadline.ErrBadParameter.With(path, ": ", src.Kind())
	}
	if t, err := time.Parse(time.RFC3339, src.String()); err != nil {
		return err
	} else {
		dest.Set(reflect.ValueOf(t))
	}
	return nil
}

func (this *Decoder) setDuration(path string, src, dest reflect.Value) error {
	switch src.Kind() {
	case reflect.Float64:
		v := time.Duration(time.Second) * time.Duration(src.Float())
		dest.Set(reflect.ValueOf(v))
	default:
		return deadline.ErrBadParameter.With(path, ": ", src.Kind())
	}
	return nil
}
