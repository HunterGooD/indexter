// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels(in *jlexer.Lexer, out *WordIdx) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = string(in.String())
		case "in_files":
			if in.IsNull() {
				in.Skip()
				out.InFiles = nil
			} else {
				in.Delim('[')
				if out.InFiles == nil {
					if !in.IsDelim(']') {
						out.InFiles = make([]WordFile, 0, 1)
					} else {
						out.InFiles = []WordFile{}
					}
				} else {
					out.InFiles = (out.InFiles)[:0]
				}
				for !in.IsDelim(']') {
					var v1 WordFile
					(v1).UnmarshalEasyJSON(in)
					out.InFiles = append(out.InFiles, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels(out *jwriter.Writer, in WordIdx) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"in_files\":"
		out.RawString(prefix)
		if in.InFiles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.InFiles {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WordIdx) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WordIdx) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WordIdx) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WordIdx) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels(l, v)
}
func easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(in *jlexer.Lexer, out *WordFile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "path":
			out.Path = string(in.String())
		case "start_word_idx":
			if in.IsNull() {
				in.Skip()
				out.StartWordIdx = nil
			} else {
				in.Delim('[')
				if out.StartWordIdx == nil {
					if !in.IsDelim(']') {
						out.StartWordIdx = make([]int, 0, 8)
					} else {
						out.StartWordIdx = []int{}
					}
				} else {
					out.StartWordIdx = (out.StartWordIdx)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int
					v4 = int(in.Int())
					out.StartWordIdx = append(out.StartWordIdx, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(out *jwriter.Writer, in WordFile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"path\":"
		out.RawString(prefix[1:])
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"start_word_idx\":"
		out.RawString(prefix)
		if in.StartWordIdx == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.StartWordIdx {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WordFile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WordFile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WordFile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WordFile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels1(l, v)
}
func easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(in *jlexer.Lexer, out *FileIndexer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "file":
			if in.IsNull() {
				in.Skip()
				out.File = nil
			} else {
				in.Delim('[')
				if out.File == nil {
					if !in.IsDelim(']') {
						out.File = make([]WordIdx, 0, 1)
					} else {
						out.File = []WordIdx{}
					}
				} else {
					out.File = (out.File)[:0]
				}
				for !in.IsDelim(']') {
					var v7 WordIdx
					(v7).UnmarshalEasyJSON(in)
					out.File = append(out.File, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "path":
			out.Path = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(out *jwriter.Writer, in FileIndexer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix[1:])
		if in.File == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.File {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileIndexer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileIndexer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileIndexer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileIndexer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComHunterGooDHWalgoCourseWorkInternalModels2(l, v)
}
