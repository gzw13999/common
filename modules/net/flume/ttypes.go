// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package flume

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Status int64

const (
	Status_OK      Status = 0
	Status_FAILED  Status = 1
	Status_ERROR   Status = 2
	Status_UNKNOWN Status = 3
)

func (p Status) String() string {
	switch p {
	case Status_OK:
		return "Status_OK"
	case Status_FAILED:
		return "Status_FAILED"
	case Status_ERROR:
		return "Status_ERROR"
	case Status_UNKNOWN:
		return "Status_UNKNOWN"
	}
	return "<UNSET>"
}

func StatusFromString(s string) (Status, error) {
	switch s {
	case "Status_OK":
		return Status_OK, nil
	case "Status_FAILED":
		return Status_FAILED, nil
	case "Status_ERROR":
		return Status_ERROR, nil
	case "Status_UNKNOWN":
		return Status_UNKNOWN, nil
	}
	return Status(0), fmt.Errorf("not a valid Status string")
}

func StatusPtr(v Status) *Status { return &v }

type ThriftFlumeEvent struct {
	Headers map[string]string `thrift:"headers,1,required" json:"headers"`
	Body    []byte            `thrift:"body,2,required" json:"body"`
}

func NewThriftFlumeEvent() *ThriftFlumeEvent {
	return &ThriftFlumeEvent{}
}

func (p *ThriftFlumeEvent) GetHeaders() map[string]string {
	return p.Headers
}

func (p *ThriftFlumeEvent) GetBody() []byte {
	return p.Body
}
func (p *ThriftFlumeEvent) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ThriftFlumeEvent) ReadField1(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Headers = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val1 = v
		}
		p.Headers[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *ThriftFlumeEvent) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Body = v
	}
	return nil
}

func (p *ThriftFlumeEvent) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ThriftFlumeEvent"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ThriftFlumeEvent) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("headers", thrift.MAP, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:headers: %s", p, err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Headers)); err != nil {
		return fmt.Errorf("error writing map begin: %s", err)
	}
	for k, v := range p.Headers {
		if err := oprot.WriteString(string(k)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return fmt.Errorf("error writing map end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:headers: %s", p, err)
	}
	return err
}

func (p *ThriftFlumeEvent) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("body", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:body: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Body); err != nil {
		return fmt.Errorf("%T.body (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:body: %s", p, err)
	}
	return err
}

func (p *ThriftFlumeEvent) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftFlumeEvent(%+v)", *p)
}