// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     schemas.v1.kafka.asvc
 */

package ethvm

type UnionNullString struct {
	Null      interface{}
	String    string
	UnionType UnionNullStringTypeEnum
}

type UnionNullStringTypeEnum int

const (
	UnionNullStringTypeEnumNull   UnionNullStringTypeEnum = 0
	UnionNullStringTypeEnumString UnionNullStringTypeEnum = 1
)