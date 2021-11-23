package rs485

import . "github.com/volkszaehler/mbmd/meters"

func init() {
	Register(NewSDM120Producer)
}

const (
	METERTYPE_SDM120 = "SDM120"
)

type SDM120Producer struct {
	Opcodes
}

func NewSDM120Producer() Producer {
	/**
	 * Opcodes as defined by Eastron SDM120.
	 * See https://bg-etech.de/download/manual/SDM120-register.pdf
	 */
	ops := Opcodes{
		Voltage:        0x0000,
		Current:        0x0006,
		Power:          0x000C,
		ApparentPower:  0x0012,
		ReactivePower:  0x0018,
		Cosphi:         0x001E,
		PhaseAngle:     0x0024,
		Frequency:      0x0046,
		Import:         0x0048,
		Export:         0x004A,
		ReactiveImport: 0x004C,
		ReactiveExport: 0x004E,
		Sum:            0x0156,
		ReactiveSum:    0x0158,
	}
	return &SDM120Producer{Opcodes: ops}
}

func (p *SDM120Producer) Type() string {
	return METERTYPE_SDM120
}

func (p *SDM120Producer) Description() string {
	return "Eastron SDM120"
}

func (p *SDM120Producer) snip(iec Measurement) Operation {
	operation := Operation{
		FuncCode:  ReadInputReg,
		OpCode:    p.Opcode(iec),
		ReadLen:   2,
		IEC61850:  iec,
		Transform: RTUIeee754ToFloat64,
	}
	return operation
}

func (p *SDM120Producer) Probe() Operation {
	return p.snip(Voltage)
}

func (p *SDM120Producer) Produce() (res []Operation) {
	for op := range p.Opcodes {
		res = append(res, p.snip(op))
	}

	return res
}
