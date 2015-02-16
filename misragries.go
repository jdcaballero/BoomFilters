package boom

type MisraGries struct {
	k         int                 // Number of counters
	len       int                 // Map len
	registers map[interface{}]int //Counter registers
}

func (m MisraGries) Len() int {
	return len(m.registers)
}

func NewMisraGries(k int) (*MisraGries, error) {
	return &MisraGries{
		k:         k,
		len:       0,
		registers: make(map[interface{}]int),
	}, nil
}

func (m *MisraGries) Add(data interface{}) *MisraGries {
	if _, ok := m.registers[data]; ok {
		m.registers[data]++
	} else if m.Len() <= m.k-1 {
		m.registers[data] = 1
		m.len++
	} else {
		for k, _ := range m.registers {
			if m.registers[k] -= 1; m.registers[k] == 0 {
				delete(m.registers, k)
				m.len--
			}
		}
	}
	return m
}

func (m *MisraGries) Counts(data interface{}) int {

	if m, ok := m.registers[data]; ok {
		return m
	}
	return 0
}

func (m *MisraGries) TopK() map[interface{}]int {
	return m.registers
}
