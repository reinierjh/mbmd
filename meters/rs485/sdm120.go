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
	 * Opcodes as defined by Eastron SDM230.
	 * See https://bg-etech.de/download/manual/SDM230-register.pdf
	 */
	ops := Opcodes{
		Voltage:        0x0000, // 220, 230
		Current:        0x0006, // 220, 230
		Power:          0x000C, //      230
		Import:         0x0048, // 220, 230
		Export:         0x004a, // 220, 230
		Cosphi:         0x001e, //      230
		Frequency:      0x0046, //      230
		ReactiveImport: 0x4C,   // 220, 230
		ReactiveExport: 0x4E,   // 220, 230
		ApparentPower:  0x0012, // 230
		ReactivePower:  0x0018, // 230
		Sum:            0x0156, // 230
		PhaseAngle:     0x0024, // 230
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
