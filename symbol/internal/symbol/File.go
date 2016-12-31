// automatically generated by the FlatBuffers compiler, do not modify

package symbol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// File files of project.
type File struct {
	_tab flatbuffers.Table
}

func GetRootAsFile(buf []byte, offset flatbuffers.UOffsetT) *File {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &File{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *File) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *File) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *File) FileName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *File) TranslationUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *File) SymbolDatabase(obj *SymbolDatabase) *SymbolDatabase {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SymbolDatabase)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FileStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FileAddFileName(builder *flatbuffers.Builder, FileName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(FileName), 0)
}
func FileAddTranslationUnit(builder *flatbuffers.Builder, TranslationUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(TranslationUnit), 0)
}
func FileAddSymbolDatabase(builder *flatbuffers.Builder, SymbolDatabase flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(SymbolDatabase), 0)
}
func FileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}