// Code generated by github.com/deep-creek/go-jsonschema, DO NOT EDIT.

package test

type ObjectNested struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectNestedMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}

type ObjectNestedMyObject struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectNestedMyObjectMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}

type ObjectNestedMyObjectMyObject struct {
	// MyString corresponds to the JSON schema field "myString".
	MyString *string `json:"myString,omitempty" yaml:"myString,omitempty" mapstructure:"myString,omitempty"`
}
