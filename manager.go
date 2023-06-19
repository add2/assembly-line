package main

type manager struct {
	line       line
	mechanisms map[int]*mechanism
}

func newManager(l line) manager {
	return manager{
		line:       l,
		mechanisms: make(map[int]*mechanism, 10),
	}
}

func (m *manager) append(mech *mechanism) {
	m.mechanisms[mech.id] = mech
}

func (m manager) mechanism(id int) *mechanism {
	if mech, ok := m.mechanisms[id]; ok {
		return mech
	}
	panic("Mechanism not found")
}
