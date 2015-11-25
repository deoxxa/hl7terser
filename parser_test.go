package hl7terser

import (
	"reflect"
	"testing"
)

type parserTestPair struct {
	s string
	q Query
}

var parserTestCases = []parserTestPair{
	parserTestPair{"MSH", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x0,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x0,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2(3)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2(3)-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2(3)-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2(3)-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		HasSegmentRepeat:   true,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x6,
		HasSubComponent:    true,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH(1)-2(3)-4(5)-6(7)", Query{
		Segment:               "MSH",
		SegmentRepeat:         0x1,
		HasSegmentRepeat:      true,
		Field:                 0x2,
		HasField:              true,
		FieldRepeat:           0x3,
		HasFieldRepeat:        true,
		Component:             0x4,
		HasComponent:          true,
		ComponentRepeat:       0x5,
		HasComponentRepeat:    true,
		SubComponent:          0x6,
		HasSubComponent:       true,
		SubComponentRepeat:    0x7,
		HasSubComponentRepeat: true,
	}},
	parserTestPair{"MSH-2", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2(3)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2(3)-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2(3)-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2(3)-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x3,
		HasFieldRepeat:     true,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x6,
		HasSubComponent:    true,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2(3)-4(5)-6(7)", Query{
		Segment:               "MSH",
		SegmentRepeat:         0x0,
		Field:                 0x2,
		HasField:              true,
		FieldRepeat:           0x3,
		HasFieldRepeat:        true,
		Component:             0x4,
		HasComponent:          true,
		ComponentRepeat:       0x5,
		HasComponentRepeat:    true,
		SubComponent:          0x6,
		HasSubComponent:       true,
		SubComponentRepeat:    0x7,
		HasSubComponentRepeat: true,
	}},
	parserTestPair{"MSH-2-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x5,
		HasComponentRepeat: true,
		SubComponent:       0x6,
		HasSubComponent:    true,
		SubComponentRepeat: 0x0,
	}},
	parserTestPair{"MSH-2-4(5)-6(7)", Query{
		Segment:               "MSH",
		SegmentRepeat:         0x0,
		Field:                 0x2,
		HasField:              true,
		FieldRepeat:           0x0,
		Component:             0x4,
		HasComponent:          true,
		ComponentRepeat:       0x5,
		HasComponentRepeat:    true,
		SubComponent:          0x6,
		HasSubComponent:       true,
		SubComponentRepeat:    0x7,
		HasSubComponentRepeat: true,
	}},
	parserTestPair{"MSH-2-4-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		HasField:           true,
		FieldRepeat:        0x0,
		Component:          0x4,
		HasComponent:       true,
		ComponentRepeat:    0x0,
		SubComponent:       0x6,
		HasSubComponent:    true,
		SubComponentRepeat: 0x0,
	}},
}

func TestParse(t *testing.T) {
	for _, c := range parserTestCases {
		q, err := Parse(c.s)

		if err != nil {
			t.Error(err)
		}

		if *q != c.q {
			av := reflect.ValueOf(*q)
			ev := reflect.ValueOf(c.q)

			for i := 0; i < av.NumField(); i++ {
				if !reflect.DeepEqual(av.Field(i).Interface(), ev.Field(i).Interface()) {
					t.Errorf("[%q] expected field %s to be %v; was %v", c.s, av.Type().Field(i).Name, ev.Field(i).Interface(), av.Field(i).Interface())
				}
			}
		}
	}
}