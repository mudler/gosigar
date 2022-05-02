//go:build !386
// +build !386

package gosigar

import (
	"syscall"
	"time"
)

func (self *Uptime) Get() error {
	tv := syscall.Timeval32{}

	if err := sysctlbyname("kern.boottime", &tv); err != nil {
		return err
	}

	self.Length = time.Since(time.Unix(int64(tv.Sec), int64(tv.Usec)*1000)).Seconds()

	return nil
}

// generic Sysctl buffer unmarshalling
func sysctlbyname(name string, data interface{}) (err error) {
	val, err := syscall.Sysctl(name)
	if err != nil {
		return err
	}

	buf := []byte(val)

	switch v := data.(type) {
	case *uint64:
		*v = *(*uint64)(unsafe.Pointer(&buf[0]))
		return
	}

	bbuf := bytes.NewBuffer([]byte(val))
	return binary.Read(bbuf, binary.LittleEndian, data)
}

func (self *Mem) Get() error {
	return nil
}

func (self *HugeTLBPages) Get() error {
	return nil
}

func (self *FDUsage) Get() error {
	return nil
}

func (self *Swap) Get() error {
	return nil
}

func (l *LoadAverage) Get() error {
	return nil
}

func (self *Cpu) Get() error {
	return nil
}
